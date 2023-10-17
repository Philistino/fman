package infobar

import (
	"github.com/Philistino/fman/ui/theme"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/mattn/go-runewidth"
)

type Infobar struct {
	width     int // the width of the infobar, which should be the window width
	prompt    prompt
	notis     notifications
	logo      string
	selected  itemTracker
	logoWidth int
}

func New() Infobar {
	logo := theme.LogoStyle.Render(string(theme.GetActiveIconTheme().GopherIcon) + "FMAN")
	return Infobar{
		width:     0,
		prompt:    newPrompt(),
		notis:     newNotifications(),
		logo:      logo,
		logoWidth: lipgloss.Width(logo),
		selected:  itemTracker{},
	}
}

// Init initializes the model. It must be called.
func (m *Infobar) Init() tea.Cmd {
	m.notis.Init()
	m.prompt.Init()
	return nil
}

// Update handles messages from the program.
func (m Infobar) Update(msg tea.Msg) (Infobar, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.prompt.width = m.width - m.logoWidth
		m.prompt.textInput.Width = m.prompt.width - 4 // 4 is the width of the prompt prefix and cursor
		m.prompt.textInput.CharLimit = m.prompt.textInput.Width
	}
	var promptCmd, notiCmd, itemCmd tea.Cmd
	m.prompt, promptCmd = m.prompt.Update(msg)
	m.notis, notiCmd = m.notis.Update(msg)
	m.selected, itemCmd = m.selected.Update(msg)
	return m, tea.Batch(promptCmd, notiCmd, itemCmd)
}

func (m Infobar) View() string {
	style := theme.InfobarStyle.Copy()
	var mainContent string
	switch {
	case m.prompt.textInput.Focused():
		mainContent = style.Width(m.width - m.logoWidth).Render(" " + m.prompt.View())
	default:
		noti := " " + m.notis.View()
		selected := style.Render(m.selected.View())
		width := m.width - m.logoWidth - lipgloss.Width(selected)
		noti = runewidth.Truncate(noti, width, "...")
		mainContent = style.Width(width).Render(noti) + selected
	}
	return lipgloss.JoinHorizontal(
		lipgloss.Center,
		m.logo,
		mainContent,
	)
}
