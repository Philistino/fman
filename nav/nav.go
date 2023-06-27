package nav

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/Philistino/fman/entry"
	"github.com/Philistino/fman/nav/history"
	"github.com/spf13/afero"
)

// on startup, create a filesystem, read the cwd and display it. Walk the filetree up to root while
// caching directories and setting up watchers. Also, walk the filetree down 3(?) levels from current
// and setup file watchers. On directory change, shift the window of file watchers.
// "github.com/gohugoio/hugo/watcher"

// Need to do something with symlinks

const pathSeparator = string(filepath.Separator)

type Nav struct {
	mu             sync.Mutex              // mutex for Nav
	hist           history.History[string] // history of paths visited. Set to record max. 5000 entries
	currentPath    string                  // current path
	entries        []entry.Entry           // current entries
	showHidden     bool                    // if true, show hidden files and directories
	dirsMixed      bool                    // if true, directories are mixed in with files
	cursorHist     map[string]string       // path -> cursor. This can grow unchecked but should not be a problem
	fsys           afero.Fs                // filesystem
	previewer      *PreviewHandler         // previewer
	idleWalkCancel context.CancelFunc
}

func NewNav(showHidden bool, dirsMixed bool, startPath string, fsys afero.Fs, previewDelay int) *Nav {

	navi := &Nav{
		hist:        history.NewHistory[string](5000),
		showHidden:  showHidden,
		dirsMixed:   dirsMixed,
		currentPath: startPath,
		cursorHist:  make(map[string]string),
		fsys:        fsys,
		previewer: NewPreviewHandler(
			context.Background(),
			previewDelay,
			50_000, // 50 kB
			100,
			time.Second*time.Duration(30),
		),
	}

	return navi
}

// Go changes the current directory to the given path and returns a Dirstate struct. If the path is "~", the home directory is used.
func (n *Nav) Go(path string, currCursor string, currSelected []string) DirState {
	currState := NavState{path: n.currentPath, cursor: currCursor, selected: mapStruct(currSelected)}

	var err error
	if path == "~" {
		path, err = os.UserHomeDir()
		if err != nil {
			return n.newDirState(n.entries, currState, err)
		}
	}
	if path == n.currentPath {
		return n.newDirState(n.entries, currState, err)
	}

	// if n.idleWalkCancel != nil {
	// 	n.idleWalkCancel()
	// }

	entries, err := n.getEntries(path)
	if err != nil {
		return n.newDirState(n.entries, currState, err)
	}

	var newState NavState

	newState.cursor = n.handleCursor(path)
	newState.path = path

	n.mu.Lock()
	defer n.mu.Unlock()
	n.idleWalk()

	n.hist.Go(n.currentPath)

	n.cursorHist[n.currentPath] = currCursor // make sure this is set before the path is set to the new one
	n.currentPath = path
	n.entries = entries
	return n.newDirState(n.entries, newState, nil)
}

func (n *Nav) handleCursor(dst string) string {
	var cursor string
	src := n.currentPath

	// if the destination is the parent of the source, set the cursor to the source
	if dst == filepath.Dir(src) {
		return filepath.Base(src)
	}

	// if the destination is in the history, set the cursor based on last visit
	cursor, ok := n.cursorHist[dst]
	if ok {
		return cursor
	}

	// if the destination is a [great-]grandparent of the source, set the cursor
	// to be in the tree of the source
	if !strings.HasPrefix(src, dst) {
		return ""
	}
	cursor = strings.Replace(src, dst, "", 1)
	cursor = strings.TrimPrefix(cursor, pathSeparator)
	cursor = strings.Split(cursor, pathSeparator)[0]
	return cursor
}

func (n *Nav) Back(currSelected []string, currCursor string) DirState {
	currState := NavState{path: n.currentPath, cursor: currCursor, selected: mapStruct(currSelected)}
	newPath, commit, err := n.hist.Back(n.currentPath)
	if err != nil {
		if errors.Is(err, history.ErrStackEmpty) {
			return n.newDirState(n.entries, currState, nil)
		}
		return n.newDirState(n.entries, currState, err)
	}

	// if n.idleWalkCancel != nil {
	// 	n.idleWalkCancel()
	// }

	entries, err := n.getEntries(newPath)
	if err != nil {
		return n.newDirState(n.entries, currState, err)
	}

	n.mu.Lock()
	defer n.mu.Unlock()
	n.idleWalk()

	commit()
	n.cursorHist[n.currentPath] = currCursor // save the cursor for the path we are leaving
	n.currentPath = newPath
	n.entries = entries
	cursor := n.cursorHist[newPath] // note this may return an empty string
	state := NavState{path: newPath, cursor: cursor}
	return n.newDirState(entries, state, err)
}

func (n *Nav) Forward(currSelected []string, currCursor string) DirState {
	currState := NavState{path: n.currentPath, cursor: currCursor, selected: mapStruct(currSelected)}
	newPath, commit, err := n.hist.Foreward(n.currentPath)
	if err != nil {
		if errors.Is(err, history.ErrStackEmpty) {
			return n.newDirState(n.entries, currState, nil)
		}
		return n.newDirState(n.entries, currState, err)
	}

	// if n.idleWalkCancel != nil {
	// 	n.idleWalkCancel()
	// }

	entries, err := n.getEntries(newPath)
	if err != nil {
		return n.newDirState(n.entries, currState, err)
	}

	n.mu.Lock()
	defer n.mu.Unlock()
	n.idleWalk()

	commit()
	n.cursorHist[n.currentPath] = currCursor // save the cursor for the path we are leaving
	n.currentPath = newPath
	n.entries = entries
	cursor := n.cursorHist[newPath] // note this may return an empty string
	state := NavState{path: newPath, cursor: cursor}
	return n.newDirState(entries, state, err)
}

// Reload reads and returns the current directory contents
func (n *Nav) Reload(currSelected []string, currCursor string) DirState {
	state := NavState{path: n.currentPath, cursor: currCursor, selected: mapStruct(currSelected)}

	// if n.idleWalkCancel != nil {
	// 	n.idleWalkCancel()
	// }

	entries, err := n.getEntries(n.currentPath)

	n.mu.Lock()
	defer n.mu.Unlock()
	n.idleWalk()

	n.entries = entries
	return n.newDirState(entries, state, err)
}

// idleWalk walks the current directory in the background so they are cached by the os
func (n *Nav) idleWalk() {
	ctx, cancel := context.WithCancel(context.Background())
	n.idleWalkCancel = cancel
	go entry.WalkDown(ctx, n.fsys, n.currentPath, 3, 8, false)
	go entry.WalkUp(ctx, n.fsys, n.currentPath, 3, 3, false)
}

func (n *Nav) CurrentPath() string {
	return n.currentPath
}

func (n *Nav) SetShowHidden(showHidden bool) {
	n.showHidden = showHidden
}

func (n *Nav) ShowHidden() bool {
	return n.showHidden
}

func (n *Nav) SetDirsMixed(dirsMixed bool) {
	n.dirsMixed = dirsMixed
}

func (n *Nav) getEntries(path string) ([]entry.Entry, error) {
	entries, _, err := entry.GetEntries(n.fsys, path, n.showHidden, n.dirsMixed)
	return entries, err
}

// returns map of slice to use a set
func mapStruct[T comparable](list []T) map[T]struct{} {
	mapped := make(map[T]struct{}, len(list))
	for _, elem := range list {
		mapped[elem] = struct{}{}
	}
	return mapped
}

func (n *Nav) GetPreview(ctx context.Context, path string) entry.Preview {
	return n.previewer.GetPreview(ctx, n.fsys, path)
}
