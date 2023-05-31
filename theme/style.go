package theme

import "github.com/charmbracelet/lipgloss"

var (
	EntryInfoStyle = lipgloss.NewStyle().Border(lipgloss.RoundedBorder())
	ListStyle      = lipgloss.NewStyle().Padding(1)

	AppStyle = lipgloss.NewStyle().Align(lipgloss.Center)

	EvenItemStyle = lipgloss.NewStyle().
			Height(1)

	PathStyle = lipgloss.NewStyle().Padding(0, 1)

	SelectedItemStyle = lipgloss.NewStyle().Height(1)

	ButtonStyle = lipgloss.NewStyle().
			Padding(0, 1).
			Border(lipgloss.NormalBorder(), false, true)

	InactiveButtonStyle = lipgloss.NewStyle().
				Padding(0, 1).
				Border(lipgloss.NormalBorder(), false, true).
				Foreground(lipgloss.Color("#707070"))

	LogoStyle = lipgloss.NewStyle().Padding(0, 1)

	ProgressStyle = lipgloss.NewStyle().Border(lipgloss.NormalBorder(), false, true)

	InfobarStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#000"))

	ArrowStyle = lipgloss.NewStyle().Align(lipgloss.Center)

	EmptyFolderStyle = lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).Padding(2)
)
