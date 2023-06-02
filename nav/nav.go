package nav

import (
	"os"
	"path/filepath"

	"github.com/nore-dev/fman/entry"
	"github.com/nore-dev/fman/history"
)

// on startup, create a filesystem, read the cwd and display it. Walk the filetree up to root while
// caching directories and setting up watchers. Also, walk the filetree down 3(?) levels from current
// and setup file watchers. On directory change, shift the window of file watchers.
// "github.com/gohugoio/hugo/watcher"

// Need to do something with symlinks

type Nav struct {
	hist        history.History[NavState]
	currentPath string
	entries     []entry.Entry
	showHidden  bool
	dirsMixed   bool
}

func NewNav(showHidden bool, dirsMixed bool, startPath string) *Nav {
	navi := &Nav{
		hist:        history.NewHistory[NavState](1000),
		showHidden:  showHidden,
		dirsMixed:   dirsMixed,
		currentPath: startPath,
	}
	return navi
}

func (n *Nav) Go(path string, currSelected []string) DirState {

	var err error
	var state NavState
	if path == "~" {
		path, err = os.UserHomeDir()
		if err != nil {
			return n.newDirState(nil, state, err)
		}
	}
	if path == n.currentPath {
		state.selected = mapStruct(currSelected)
		state.path = path
		return n.newDirState(n.entries, state, err)
	}

	entries, err := n.getEntries(path)
	if err != nil {
		return n.newDirState(nil, state, err)
	}

	// if the new path is the parent of the current path set the cursor to the current path
	if path == filepath.Dir(n.currentPath) {
		state.selected = mapStruct([]string{filepath.Base(n.currentPath)})
	}
	state.path = path

	n.hist.Go(NavState{path: n.currentPath, selected: mapStruct(currSelected)})
	n.currentPath = path
	n.entries = entries
	return n.newDirState(n.entries, state, err)
}

func (n *Nav) Back(currSelected []string) DirState {
	state, commit, err := n.hist.Back(NavState{path: n.currentPath, selected: mapStruct(currSelected)})
	if err != nil {
		return n.newDirState(nil, state, err)
	}
	entries, err := n.getEntries(state.Path())
	if err != nil {
		return n.newDirState(nil, state, err)
	}
	commit()
	n.currentPath = state.Path()
	n.entries = entries
	return n.newDirState(entries, state, err)
}

func (n *Nav) Forward(currSelected []string) DirState {
	state, commit, err := n.hist.Foreward(NavState{path: n.currentPath, selected: mapStruct(currSelected)})
	if err != nil {
		return n.newDirState(nil, state, err)
	}

	entries, err := n.getEntries(state.Path())
	if err != nil {
		return n.newDirState(nil, state, err)
	}
	commit()
	n.currentPath = state.Path()
	n.entries = entries
	return n.newDirState(entries, state, err)
}

// Reload reads and returns the current directory contents
func (n *Nav) Reload(currSelected []string) DirState {
	state := NavState{path: n.currentPath, selected: mapStruct(currSelected)}
	entries, err := n.getEntries(n.currentPath)
	n.entries = entries
	return n.newDirState(entries, state, err)
}

func (n *Nav) CurrentPath() string {
	return n.currentPath
}

// returns map of slice to use a set
func mapStruct[T comparable](list []T) map[T]struct{} {
	mapped := make(map[T]struct{}, len(list))
	for _, elem := range list {
		mapped[elem] = struct{}{}
	}
	return mapped
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

func (n *Nav) SetDirsMixed(dirsMixed bool) {
	n.dirsMixed = dirsMixed
}

func (n *Nav) getEntries(path string) ([]entry.Entry, error) {
	return entry.GetEntries(path, n.showHidden, n.dirsMixed)
}

type NavState struct {
	path     string
	selected map[string]struct{}
}

// Path returns the path to the directory of the NavState
func (n NavState) Path() string {
	return n.path
}

// Selected returns the selected items in the directory
func (n NavState) Selected() map[string]struct{} {
	return n.selected
}

type DirState struct {
	NavState
	entries    []entry.Entry
	backActive bool
	fwdActive  bool
	upActive   bool
	err        error
}

// Error returns the error or nil
func (d DirState) Error() error {
	return d.err
}

// Entries returns the directory entries
func (d DirState) Entries() []entry.Entry {
	return d.entries
}

// BackActive returns true if there is a back history
func (d DirState) BackActive() bool {
	return d.backActive
}

// ForwardActive returns true if there is a forward history
func (d DirState) ForwardActive() bool {
	return d.fwdActive
}

// UpActive returns true if the current directory is not the root
func (d DirState) UpActive() bool {
	return d.upActive
}

func isRoot(name string) bool { return filepath.Dir(name) == name }

func (n *Nav) newDirState(entries []entry.Entry, nState NavState, err error) DirState {
	return DirState{
		NavState:   nState,
		entries:    entries,
		backActive: !n.hist.BackEmpty(),
		fwdActive:  !n.hist.ForewardEmpty(),
		upActive:   !isRoot(nState.Path()),
		err:        err,
	}
}
