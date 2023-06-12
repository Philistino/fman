package nav

import (
	"path/filepath"

	"github.com/Philistino/fman/entry"
)

// NavState represents the state of the navigation within a directory
type NavState struct {
	path     string              // path to the directory
	selected map[string]struct{} // map of selected items in the directory
	cursor   string              // name of the item in the directory for the cursor
}

// Path returns the path to the directory of the NavState
func (n NavState) Path() string {
	return n.path
}

// Selected returns the selected items in the directory
func (n NavState) Selected() map[string]struct{} {
	return n.selected
}

// Cursor returns the cursor in the directory
func (n NavState) Cursor() string {
	return n.cursor
}

// DirState represents the state of the directory
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
