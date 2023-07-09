package message

import tea "github.com/charmbracelet/bubbletea"

// NavBackMsg is used to communicate to the main program
// that a navigation back operation is requested.
type NavBackMsg struct{}

// NavBackCmd is used to create a command that will
// communicate to the main program that a navigation back
// operation is requested.
func NavBackCmd() tea.Cmd {
	return func() tea.Msg {
		return NavBackMsg{}
	}
}

// NavFwdMsg is used to communicate to the main program
// that a navigation forward operation is requested.
type NavFwdMsg struct{}

// NavFwdCmd is used to create a command that will
// communicate to the main program that a navigation forward
// operation is requested.
func NavFwdCmd() tea.Cmd {
	return func() tea.Msg {
		return NavFwdMsg{}
	}
}

// NavUpMsg is used to communicate to the main program
// that a navigation to parent directory operation is requested.
type NavUpMsg struct{}

// NavUpCmd is used to create a command that will
// communicate to the main program that a navigation to the parent directory
// operation is requested.
func NavUpCmd() tea.Cmd {
	return func() tea.Msg {
		return NavUpMsg{}
	}
}

// NavHomeMsg is used to communicate to the main program
// that a navigation to the home directory operation is requested.
type NavHomeMsg struct{}

// NavHomeCmd is used to create a command that will
// communicate to the main program that a navigation to the home directory
// operation is requested.
func NavHomeCmd() tea.Cmd {
	return func() tea.Msg {
		return NavHomeMsg{}
	}
}

// NavDownMsg is used to communicate to the main program
// that a navigation to the selected directory is requested.
type NavDownMsg struct {
	Name string
}

// NavDownCmd is used to create a command that will
// communicate to the main program that a navigation to the selected directory
// is requested.
func NavDownCmd(name string) tea.Cmd {
	return func() tea.Msg {
		return NavDownMsg{
			Name: name,
		}
	}
}

// NavOtherMsg is used to communicate to the main program
// that a navigation to a directory other than the current one
type NavOtherMsg struct {
	Path string
}

// NavOtherCmd is used to create a command that will
// communicate to the main program that a navigation to a directory other than
// the current one is requested.
func NavOtherCmd(path string) tea.Cmd {
	return func() tea.Msg {
		return NavOtherMsg{
			Path: path,
		}
	}
}
