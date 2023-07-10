package bookmark_ui

import tea "github.com/charmbracelet/bubbletea"

type BookmarkMsg struct {
	paths []string
}

func (m BookmarkMsg) Paths() []string {
	return m.paths
}

func BookmarkCmd(paths []string) tea.Cmd {
	return func() tea.Msg {
		return BookmarkMsg{paths: paths}
	}
}

type UnbookmarkMsg struct {
	paths []string
}

func (m UnbookmarkMsg) Paths() []string {
	return m.paths
}

func UnbookmarkCmd(paths []string) tea.Cmd {
	return func() tea.Msg {
		return UnbookmarkMsg{paths: paths}
	}
}
