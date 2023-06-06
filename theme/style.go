package theme

import (
	"github.com/Philistino/fman/theme/colors"
	"github.com/charmbracelet/lipgloss"
)

var (
	EntryInfoStyle      = lipgloss.NewStyle().Border(lipgloss.RoundedBorder())
	ListStyle           = lipgloss.NewStyle().Padding(1)
	AppStyle            = lipgloss.NewStyle().Align(lipgloss.Center)
	EvenItemStyle       = lipgloss.NewStyle().Height(1)
	PathStyle           = lipgloss.NewStyle().Padding(0, 1)
	SelectedItemStyle   = lipgloss.NewStyle().Height(1)
	LogoStyle           = lipgloss.NewStyle().Padding(0, 1)
	ProgressStyle       = lipgloss.NewStyle().Border(lipgloss.NormalBorder(), false, true)
	InfobarStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("#000"))
	ArrowStyle          = lipgloss.NewStyle().Align(lipgloss.Center)
	EmptyFolderStyle    = lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).Padding(2)
	ButtonStyle         = lipgloss.NewStyle().Padding(0, 1).Border(lipgloss.NormalBorder(), false, true)
	InactiveButtonStyle = lipgloss.NewStyle().Padding(0, 1).
				Border(lipgloss.NormalBorder(), false, true).
				Foreground(lipgloss.Color("#707070")) // TODO: make this a theme color
)

type StyleSet struct {
	EntryInfoStyle      lipgloss.Style
	ListStyle           lipgloss.Style
	AppStyle            lipgloss.Style
	EvenItemStyle       lipgloss.Style
	PathStyle           lipgloss.Style
	SelectedItemStyle   lipgloss.Style
	LogoStyle           lipgloss.Style
	ProgressStyle       lipgloss.Style
	InfobarStyle        lipgloss.Style
	ArrowStyle          lipgloss.Style
	EmptyFolderStyle    lipgloss.Style
	ButtonStyle         lipgloss.Style
	InactiveButtonStyle lipgloss.Style
}

// func NewStyleSet() StyleSet {

func SetTheme(theme colors.Theme) {
	EvenItemStyle.Background(theme.EvenItemBgColor)

	SelectedItemStyle.Background(theme.SelectedItemBgColor)
	SelectedItemStyle.Foreground(theme.SelectedItemFgColor)

	ButtonStyle.BorderForeground(theme.ButtonBorderFgColor)
	ButtonStyle.Background(theme.ButtonBgColor)

	InactiveButtonStyle.BorderForeground(theme.ButtonBorderFgColor)
	InactiveButtonStyle.Background(theme.ButtonBgColor)

	PathStyle.Background(theme.PathElementBgColor)
	PathStyle.BorderForeground(theme.PathElementBorderFgColor)

	AppStyle.Background(theme.ListBgColor)
	AppStyle.Foreground(theme.ListFgColor)

	LogoStyle.Background(theme.LogoBgColor)
	LogoStyle.Foreground(theme.LogoFgColor)

	ProgressStyle.Background(theme.ProgressBarBgColor)
	ProgressStyle.Foreground(theme.ProgressBarFgColor)
	ProgressStyle.BorderForeground(theme.ProgressBarFgColor)

	InfobarStyle.Background(theme.InfobarBgColor)
	InfobarStyle.Foreground(theme.InfobarFgColor)

	EntryInfoStyle.BorderForeground(theme.SeparatorColor)

	ArrowStyle.Foreground(theme.ArrowColor)
}
