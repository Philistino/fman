package infobar

import (
	"log"
	"strings"

	"github.com/Philistino/fman/entry/storage"
	"github.com/Philistino/fman/ui/theme"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/dustin/go-humanize"
)

type Infobar struct {
	width             int // the width of the infobar, which should be the window width
	prompt            prompt
	notis             notifications
	storageInfo       storage.StorageInfo // the disk storage info
	logo              string
	diskFree          string
	diskFreeWidth     int // the width of the progress bar
	diskFreeIndicator string
	middleWidth       int
}

func New() Infobar {
	info, err := storage.GetStorageInfo()
	if err != nil {
		log.Println("An error occurred while getting storage info", err.Error())
	}
	notis := newNotifications()
	prompt := newPrompt()
	diskFreeWidth := 20
	return Infobar{
		width:             0,
		diskFreeWidth:     diskFreeWidth,
		storageInfo:       info,
		prompt:            prompt,
		notis:             notis,
		logo:              theme.LogoStyle.Render(string(theme.GetActiveIconTheme().GopherIcon) + "FMAN"),
		diskFree:          lipgloss.JoinHorizontal(lipgloss.Center, " ", humanize.Bytes(info.AvailableSpace), "/", humanize.Bytes(info.TotalSpace), " free"),
		diskFreeIndicator: renderProgress(diskFreeWidth, info.AvailableSpace, info.TotalSpace),
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
		m.middleWidth = m.width - (lipgloss.Width(m.diskFree) + lipgloss.Width(m.diskFreeIndicator) + lipgloss.Width(m.logo) + 1)
		m.notis.width = m.middleWidth
		m.prompt.width = m.middleWidth
		m.prompt.textInput.Width = m.middleWidth - 4 // 4 is the width of the prompt prefix and cursor
	}
	var promptCmd, notiCmd tea.Cmd
	m.prompt, promptCmd = m.prompt.Update(msg)
	m.notis, notiCmd = m.notis.Update(msg)
	return m, tea.Batch(promptCmd, notiCmd)
}

// renderProgress returns a string representing a progress bar with the given width, based on the used and total disk space.
// The progress bar is rendered using the theme.ProgressStyle style.
func renderProgress(width int, usedSpace uint64, totalSpace uint64) string {
	usedWidth := (int(usedSpace) * width / int(totalSpace))
	usedStr := strings.Repeat("█", int(width-usedWidth))
	return theme.ProgressStyle.Width(width).Render(usedStr)
}

func (m Infobar) View() string {
	style := theme.InfobarStyle
	var middleSection string
	if m.prompt.textInput.Focused() {
		middleSection = m.prompt.View()
	} else {
		middleSection = m.notis.View()
	}

	return lipgloss.JoinHorizontal(
		lipgloss.Center,
		m.logo,
		style.Width(m.middleWidth).Render(" "+middleSection),
		m.diskFreeIndicator,
		style.Render(m.diskFree),
	)
}
