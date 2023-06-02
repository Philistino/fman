package nav

import (
	"os"

	"github.com/nore-dev/fman/entry"
	"github.com/nore-dev/fman/history"
)

// on startup, create a filesystem, read the cwd and display it. Walk the filetree up to root while
// caching directories and setting up watchers. Also, walk the filetree down 3(?) levels from current
// and setup file watchers. On directory change, shift the window of file watchers.
// "github.com/gohugoio/hugo/watcher"

// Need to do something with symlinks

type NavState struct {
	path     string
	selected map[string]struct{} // TODO: should this be a map[string]struct{} instead?
}

// Path returns the path to the directory of the NavState
func (n NavState) Path() string {
	return n.path
}

// Selected returns the selected items in the directory
func (n NavState) Selected() map[string]struct{} {
	return n.selected
}

type Nav struct {
	hist        history.History[NavState]
	currentPath string
	entries     []entry.Entry
	showHidden  bool
	dirsMixed   bool
}

func NewNav(path string, showHidden bool, dirsMixed bool) (*Nav, error) {
	navi := &Nav{
		hist:        history.NewHistory[NavState](1000),
		showHidden:  showHidden,
		dirsMixed:   dirsMixed,
		currentPath: path,
	}
	entries, err := navi.getEntries(path)
	if err != nil {
		return navi, err
	}
	navi.entries = entries
	return navi, nil
}

func (n *Nav) Go(path string, currSelected []string) ([]entry.Entry, NavState, error) {
	var err error
	if path == "~" {
		path, err = os.UserHomeDir()
		if err != nil {
			return nil, NavState{}, err
		}
	}

	var emptyNav NavState // TODO: track navigation traversing up and down
	entries, err := n.getEntries(path)
	if err != nil {
		return nil, emptyNav, err
	}
	n.hist.Visit(NavState{path: n.currentPath, selected: mapStruct(currSelected)})
	n.currentPath = path
	return entries, emptyNav, nil
}

func (n *Nav) Back(currSelected []string) ([]entry.Entry, NavState, error) {
	state, commit, err := n.hist.Back(NavState{path: n.currentPath, selected: mapStruct(currSelected)})
	if err != nil {
		return nil, state, err
	}
	entries, err := n.getEntries(state.Path())
	if err != nil {
		return nil, state, err
	}
	commit()
	n.currentPath = state.Path()
	return entries, state, err
}

func (n *Nav) Forward(currSelected []string) ([]entry.Entry, NavState, error) {
	state, commit, err := n.hist.Foreward(NavState{path: n.currentPath, selected: mapStruct(currSelected)})
	if err != nil {
		return nil, state, err
	}
	entries, err := n.getEntries(state.Path())
	if err != nil {
		return nil, state, err
	}
	commit()
	n.currentPath = state.Path()
	return entries, state, err
}

// Reload reads and returns the current directory contents
func (n *Nav) Reload(currSelected []string) ([]entry.Entry, NavState, error) {
	state := NavState{path: n.currentPath, selected: mapStruct(currSelected)}
	entries, err := n.getEntries(n.currentPath)
	return entries, state, err
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
