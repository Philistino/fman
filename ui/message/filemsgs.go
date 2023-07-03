package message

import tea "github.com/charmbracelet/bubbletea"

// RenameMsg is used to communicate to the main program
// that a rename operation is requested.
type RenameMsg struct{}

// RenameCmd is used to create a command that will
// communicate to the main program that a rename
// operation is requested.
func RenameCmd() tea.Cmd {
	return func() tea.Msg {
		return RenameMsg{}
	}
}

// NewFileMsg is used to communicate to the main program
// that a new file operation is requested.
type NewFileMsg struct{}

// NewFileCmd is used to create a command that will
// communicate to the main program that a new file
// operation is requested.
func NewFileCmd() tea.Cmd {
	return func() tea.Msg {
		return NewFileMsg{}
	}
}

// MkDirMsg is used to communicate to the main program
// that a new directory operation is requested.
type MkDirMsg struct{}

// MkDirCmd is used to create a command that will
// communicate to the main program that a new directory
// operation is requested.
func MkDirCmd() tea.Cmd {
	return func() tea.Msg {
		return MkDirMsg{}
	}
}

// DeleteMsg is used to communicate to the main program
// that a delete operation is requested.
type DeleteMsg struct{}

// DeleteCmd is used to create a command that will
// communicate to the main program that a delete
// operation is requested.
func DeleteCmd() tea.Cmd {
	return func() tea.Msg {
		return DeleteMsg{}
	}
}
