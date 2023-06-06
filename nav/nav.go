package nav

import (
	"errors"
	"os"
	"path/filepath"
	"strings"

	"github.com/Philistino/fman/entry"
	"github.com/Philistino/fman/nav/history"
)

// on startup, create a filesystem, read the cwd and display it. Walk the filetree up to root while
// caching directories and setting up watchers. Also, walk the filetree down 3(?) levels from current
// and setup file watchers. On directory change, shift the window of file watchers.
// "github.com/gohugoio/hugo/watcher"

// Need to do something with symlinks

type Nav struct {
	hist        history.History[string] // history of paths visited. Set to record max. 5000 entries
	currentPath string                  // current path
	entries     []entry.Entry           // current entries
	showHidden  bool                    // if true, show hidden files and directories
	dirsMixed   bool                    // if true, directories are mixed in with files
	cursorHist  map[string]string       // path -> cursor. This can grow unchecked but should not be a problem
}

func NewNav(showHidden bool, dirsMixed bool, startPath string) *Nav {
	navi := &Nav{
		hist:        history.NewHistory[string](5000),
		showHidden:  showHidden,
		dirsMixed:   dirsMixed,
		currentPath: startPath,
		cursorHist:  make(map[string]string),
	}
	return navi
}

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

	entries, err := n.getEntries(path)
	if err != nil {
		return n.newDirState(n.entries, currState, err)
	}

	var newState NavState

	newState.cursor = n.handleCursor(path)
	newState.path = path

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
	cursor = strings.TrimPrefix(cursor, "/")
	cursor = strings.Split(cursor, "/")[0]
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
	entries, err := n.getEntries(newPath)
	if err != nil {
		return n.newDirState(n.entries, currState, err)
	}
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

	entries, err := n.getEntries(newPath)
	if err != nil {
		return n.newDirState(n.entries, currState, err)
	}
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
	entries, err := n.getEntries(n.currentPath)
	n.entries = entries
	return n.newDirState(entries, state, err)
}

func (n *Nav) CurrentPath() string {
	return n.currentPath
}

func (n *Nav) Path() string {
	return n.currentPath
}

func (n *Nav) Entries() []entry.Entry {
	return n.entries
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
	return entry.GetEntries(path, n.showHidden, n.dirsMixed)
}

// returns map of slice to use a set
func mapStruct[T comparable](list []T) map[T]struct{} {
	mapped := make(map[T]struct{}, len(list))
	for _, elem := range list {
		mapped[elem] = struct{}{}
	}
	return mapped
}
