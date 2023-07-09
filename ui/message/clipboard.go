package message

import tea "github.com/charmbracelet/bubbletea"

// InternalCopyMsg is used to communicate to the main program
// that a "clipboard" copy operation is requested.
type InternalCopyMsg struct{}

// InternalCopyCmd is used to create a command that will
// communicate to the main program that a "clipboard" copy
// operation is requested.
func InternalCopyCmd() tea.Cmd {
	return func() tea.Msg {
		return InternalCopyMsg{}
	}
}

// CutMsg is used to communicate to the main program
// that a "clipboard" cut operation is requested.
type CutMsg struct{}

// CutCmd is used to create a command that will
// communicate to the main program that a "clipboard" cut
// operation is requested.
func CutCmd() tea.Cmd {
	return func() tea.Msg {
		return CutMsg{}
	}
}

type InternalPasteMsg struct{}

func InternalPasteCmd() tea.Cmd {
	return func() tea.Msg {
		return InternalPasteMsg{}
	}
}
