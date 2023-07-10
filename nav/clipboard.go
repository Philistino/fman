package nav

import (
	"path/filepath"
)

type clipBoard struct {
	paths []string
	cut   bool
}

// Empty returns true if the clipboard is empty
func (cb *clipBoard) Empty() bool {
	return len(cb.paths) == 0
}

// IsCut returns true if the clipboard is in cut mode
func (cb *clipBoard) IsCut() bool {
	return cb.cut
}

// ClipboardCopy sets the internal clipboard to the selected entries
func (n *Nav) ClipboardCopy(selectedNames map[string]struct{}, cut bool) {
	paths := make([]string, 0, len(selectedNames))
	dir := n.CurrentPath()
	for name := range selectedNames {
		paths = append(paths, filepath.Join(dir, name))
	}
	clip := clipBoard{paths: paths, cut: cut}
	n.clipboard = clip
}

// func (n *Nav) ClipboardPaste() []error {
// 	if n.clipboard.Empty() {
// 		return []error{errors.New("Nothing to paste")}
// 	}
// 	var errs []error

// 	if n.clipboard.IsCut() {
// 		// move the files to the current directory
// 		for _, path := range n.clipboard.paths {
// 			err := fileutils.MoveOrCopy(n.fsys, path, n.CurrentPath())
// 			if err != nil {
// 				errs = append(errs, err)
// 			}
// 		}
// 	} else {
// 		// copy the files to the current directory
// 		for _, path := range n.clipboard.paths {
// 			err := n.Copy(path, n.CurrentPath())
// 			if err != nil {
// 				errs = append(errs, err)
// 			}
// 		}
// 	}
// }

// // TODO: make this real
// func (app *App) handlePaste() tea.Cmd {
// 	if app.clipboard.Empty() {
// 		return message.NewNotificationCmd("Nothing to paste")
// 	}

// }
