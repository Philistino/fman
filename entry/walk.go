package entry

import (
	"context"
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/spf13/afero"
	"golang.org/x/sync/errgroup"
)

const Separator = string(filepath.Separator)

type DirEntry struct {
	fs.FileInfo
	path string
}

func (d DirEntry) Path() string {
	return d.path
}

// use absolute paths. Do not pass "." or ".." as startDir
func WalkDown(ctx context.Context, fsys afero.Fs, startDir string, depth int, nRoutines int, passVals bool) (<-chan DirEntry, <-chan error, error) {

	maxDepth := depth
	if depth >= 0 && !isRoot(startDir) {
		maxDepth = strings.Count(startDir, Separator) + depth // this assumes
	}

	// set up channels to pass values back to caller
	var entriesCh chan DirEntry // default nil
	var walkErrors chan error   // default nil
	if passVals {
		entriesCh = make(chan DirEntry)
		walkErrors = make(chan error)
	}

	g, ctx := errgroup.WithContext(ctx)
	if nRoutines <= 0 || nRoutines > 100 {
		nRoutines = 10
	}
	g.SetLimit(nRoutines + 1) // +1 to launch the goroutine that waits for the group to finish

	switch depth {
	case 0:
		g.Go(func() error {
			readDir(ctx, fsys, entriesCh, walkErrors, startDir)
			return nil
		})
	case 1:
		g.Go(func() error {
			readDirAndSubDirs(ctx, fsys, entriesCh, walkErrors, startDir)
			return nil
		})
	default:
		g.Go(func() error {
			return walkDown(ctx, g, fsys, maxDepth, startDir, entriesCh, walkErrors)
		})
	}

	go func() {
		g.Wait()
		if passVals {
			close(entriesCh)
			close(walkErrors)
		}
	}()
	return entriesCh, walkErrors, nil
}

// THIS DOES NOT WORK FOR READING ONE LEVEL DEEP WHEN READING AT ROOT
func walkDown(
	ctx context.Context,
	g *errgroup.Group,
	fsys afero.Fs,
	maxDepth int,
	root string,
	pathsCh chan<- DirEntry,
	walkErrors chan<- error,
) error {
	return afero.Walk(fsys, root, func(path string, info fs.FileInfo, err error) error {
		// if context is cancelled, return the error
		if ctx.Err() != nil {
			return ctx.Err()
		}

		// if there is an error, send it to the walkErrors channel and move on
		if err != nil {
			if walkErrors != nil {
				select {
				case <-ctx.Done():
					return ctx.Err()
				case walkErrors <- err:
				}
			}
			return nil
		}

		maxDepthReached := maxDepth > 0 && strings.Count(path, Separator) == maxDepth

		if !info.IsDir() || maxDepthReached || root == path {

			if pathsCh != nil {
				select {
				case <-ctx.Done():
					return ctx.Err()
				case pathsCh <- DirEntry{info, path}:
				}
			}

			if maxDepthReached && info.IsDir() {
				return fs.SkipDir
			}
			return nil
		}

		// if there are idle go routines, kick off a new routine
		// to walk the current subdirectory. This will return the fs.SkipDir
		// error so the current fs.WalkDir does not also walk down the same subdirectory.
		// IMPORTANT: this will does not put the current path in the pathsCh because
		// the new go routine will start reading at the given root, which is the current
		// path, and put this path in the channel
		started := g.TryGo(func() error {
			return walkDown(ctx, g, fsys, maxDepth, path, pathsCh, walkErrors)
		})
		if started {
			return fs.SkipDir
		}

		// a new go routine was not started, so put the object in the channel
		if pathsCh != nil {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case pathsCh <- DirEntry{info, path}:
			}
		}
		return nil
	})
}

// WalkUp walks up the file system from the given startDir, sending DirEntry objects
// to the entriesCh channel. If passVals is true, the entriesCh and walkErrors channels
// are created and returned to the caller. If passVals is false, the channels are nil.
// If passVals is true, the caller must read from the channels until they are closed.
// If passVals is false, reading from the channels will cause a panic.
// Also, if passVals is false, the
func WalkUp(ctx context.Context, fsys afero.Fs, startDir string, depth int, nRoutines int, passVals bool) (<-chan DirEntry, <-chan error, error) {

	// set up channels to pass values back to caller
	var entriesCh chan DirEntry // default nil
	var walkErrors chan error   // default nil
	if passVals {
		entriesCh = make(chan DirEntry)
		walkErrors = make(chan error)
	}

	g, ctx := errgroup.WithContext(ctx)
	if nRoutines <= 0 || nRoutines > 100 {
		nRoutines = 10
	}
	g.SetLimit(nRoutines + 1) // +1 to launch the goroutine that waits for the group to finish

	// g.Go(func() error {
	// 	// +1 to depth to get same behavior as WalkDown (n levels + current level)
	// 	walkUp(ctx, g, fsys, depth+1, startDir, entriesCh, walkErrors)
	// 	return nil
	// })

	go func() {
		walkUp(ctx, g, fsys, depth+1, startDir, entriesCh, walkErrors)
		g.Wait()
		if passVals {
			close(entriesCh)
			close(walkErrors)
		}
	}()
	return entriesCh, walkErrors, nil
}

func walkUp(
	ctx context.Context,
	g *errgroup.Group,
	fsys afero.Fs,
	maxDepth int,
	startDir string,
	pathCh chan<- DirEntry,
	errCh chan<- error,
) error {
	maxSteps := maxDepth
	for {
		path := startDir
		if maxDepth > 0 {
			if maxSteps == 0 {
				return nil
			}
		}
		if ctx.Err() != nil {
			return ctx.Err()
		}

		g.Go(func() error {
			readDir(ctx, fsys, pathCh, errCh, path)
			return nil
		})

		if isRoot(startDir) {
			return nil
		}
		startDir = filepath.Dir(startDir)
		maxSteps--
	}
}

func readDir(ctx context.Context, fsys afero.Fs, pathsCh chan<- DirEntry, walkErrors chan<- error, path string) {
	info, err := fsys.Stat(path)

	// if the channels are nil, no need to pass values back to caller
	if pathsCh == nil && walkErrors == nil {
		return
	}
	// if an error occurs, send it to the walkErrors channel and return
	if err != nil {
		select {
		case <-ctx.Done():
			return
		case walkErrors <- err:
			return
		}
	}
	select {
	case <-ctx.Done():
		return
	case pathsCh <- DirEntry{info, path}:
		return
	}
}

func readDirAndSubDirs(ctx context.Context, fsys afero.Fs, pathsCh chan<- DirEntry, walkErrors chan<- error, path string) {

	readDir(ctx, fsys, pathsCh, walkErrors, path)

	entries, err := afero.ReadDir(fsys, path)

	// if the channels are nil, no need to pass values back to caller
	if pathsCh == nil && walkErrors == nil {
		return
	}

	// if an error occurs, send it to the walkErrors channel and return
	if err != nil {
		select {
		case <-ctx.Done():
			return
		case walkErrors <- err:
			return
		}
	}

	for _, entry := range entries {
		select {
		case <-ctx.Done():
			return
		case pathsCh <- DirEntry{entry, filepath.Join(path, entry.Name())}:
		}
	}
}
