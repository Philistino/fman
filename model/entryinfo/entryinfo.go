package entryinfo

import (
	"context"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/reflow/truncate"
	"github.com/muesli/termenv"
	"github.com/nore-dev/fman/entry"
	"github.com/nore-dev/fman/keymap"
	"github.com/nore-dev/fman/message"
	"github.com/nore-dev/fman/theme"
)

type EntryInfo struct {
	entry entry.Entry

	width  int
	height int

	path    string
	preview string

	previewHeight int
	previewOffset int

	theme         *theme.Theme
	eofReached    bool // is set to true when the end of the file is reached in the preview
	previewCancel context.CancelFunc
}

const margin = 2

var previewStyle = lipgloss.NewStyle()

func New(theme *theme.Theme, firstEntry entry.Entry) EntryInfo {
	return EntryInfo{
		entry:         firstEntry,
		previewHeight: 10,
		theme:         theme,
		width:         10,
	}
}

func (entryInfo *EntryInfo) Init() tea.Cmd {
	return nil
}

func (entryInfo *EntryInfo) setNewEntry(entry entry.Entry) tea.Cmd {
	entryInfo.entry = entry
	// handle preview context cancellation
	if entryInfo.previewCancel != nil {
		entryInfo.previewCancel()
		entryInfo.previewCancel = nil
	}
	// reset read position to start of file
	entryInfo.previewOffset = 0
	entryInfo.eofReached = false

	if entry.IsDir() {
		entryInfo.preview = entryInfo.renderNoPreview("Directory")
		return nil
	}
	// set default preview content
	entryInfo.preview = entryInfo.renderNoPreview("Loading preview...")
	return entryInfo.getPreview()
}

func (entryInfo *EntryInfo) getPreview() tea.Cmd {
	ctx, cancel := context.WithCancel(context.Background())
	entryInfo.previewCancel = cancel
	return getPreviewCmd(ctx, entryInfo.getFullPath(), entryInfo.previewHeight, entryInfo.previewOffset)
}

func (entryInfo *EntryInfo) handlePreviewMsg(msg previewReadyMsg) {
	// check that the path matches so we don't set the current preview based on the previous file
	if msg.Path != entryInfo.getFullPath() {
		return
	}
	if msg.Err != nil {
		entryInfo.preview = entryInfo.renderNoPreview("Failed to load preview")
		return
	}
	if msg.Preview != "" {
		entryInfo.preview = msg.Preview
	}
	if msg.EndReached {
		entryInfo.eofReached = true
	}
}

func (entryInfo *EntryInfo) Update(msg tea.Msg) (EntryInfo, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case message.PathMsg:
		entryInfo.path = msg.Path
		entryInfo.eofReached = false
	case tea.KeyMsg:
		if key.Matches(msg, keymap.Default.ScrollPreviewDown) {
			if !entryInfo.eofReached {
				entryInfo.previewOffset++
			}
			cmd = entryInfo.getPreview()
		}
		if key.Matches(msg, keymap.Default.ScrollPreviewUp) && entryInfo.previewOffset > 0 {
			entryInfo.previewOffset--
			entryInfo.eofReached = false
			cmd = entryInfo.getPreview()
		}
	case message.EntryMsg:
		cmd = entryInfo.setNewEntry(msg.Entry)
	case previewReadyMsg:
		entryInfo.handlePreviewMsg(msg)
	}
	return *entryInfo, cmd
}

func (entryInfo *EntryInfo) getFileInfo() string {

	str := strings.Builder{}
	str.WriteByte('\n')
	name := termenv.String(entryInfo.entry.Name()).Bold().Underline().String()
	str.WriteString(truncate.StringWithTail(name, uint(entryInfo.width-margin-1), "..."))
	str.WriteByte('\n')
	typeStr := entryInfo.entry.Type

	if typeStr == "" {
		typeStr = "Unknown type"
	}

	if entryInfo.entry.IsDir() {
		typeStr = "Folder"
	}

	{
		padding := 1
		style := lipgloss.NewStyle().
			Padding(0, padding).
			Width(lipgloss.Width(typeStr) + 2*padding + 2).
			Foreground(entryInfo.theme.BackgroundColor)

		icon := theme.GetActiveIconTheme().FileIcon

		if entryInfo.entry.IsDir() {
			style.Background(entryInfo.theme.FolderColor)
			icon = theme.GetActiveIconTheme().FolderIcon
		} else {
			style.Background(entryInfo.theme.HiddenFileColor)
		}

		str.WriteString(truncate.StringWithTail(style.Render(string(icon)+" "+typeStr), uint(entryInfo.width-margin), ".."))
		str.WriteByte('\n')

		str.WriteString(termenv.String(strings.Repeat("-", entryInfo.width-margin)).Foreground(termenv.RGBColor(entryInfo.theme.InfobarBgColor)).String())
		str.WriteByte('\n')
	}

	str.WriteString(termenv.String("Modified ").Italic().String())
	str.WriteString(entryInfo.entry.ModifyTime)
	str.WriteByte('\n')
	str.WriteString(termenv.String("Changed ").Italic().String())
	str.WriteString(entryInfo.entry.ChangeTime)
	str.WriteByte('\n')
	return str.String()
}

func (entryInfo *EntryInfo) View() string {
	fileInfo := entryInfo.getFileInfo()
	entryInfo.previewHeight = entryInfo.height - lipgloss.Height(fileInfo)

	return theme.EntryInfoStyle.Render(
		lipgloss.JoinVertical(
			lipgloss.Left,
			previewStyle.
				MaxHeight(entryInfo.previewHeight).
				Height(entryInfo.previewHeight-margin).
				// Width(entryInfo.width-margin). // uncomment this line to wrap lines.
				MaxWidth(entryInfo.width-margin).
				Render(entryInfo.preview),
			fileInfo,
		),
	)
}

func (entryInfo *EntryInfo) Width() int {
	return entryInfo.width
}

func (entryInfo *EntryInfo) SetWidth(width int) {
	entryInfo.width = width
}

func (entryInfo *EntryInfo) Height() int {
	return entryInfo.height
}

func (entryInfo *EntryInfo) SetHeight(height int) {
	entryInfo.height = height
}

func (entryInfo *EntryInfo) renderNoPreview(text string) string {
	return lipgloss.Place(
		entryInfo.width-2,
		entryInfo.previewHeight,
		lipgloss.Center,
		lipgloss.Center,
		text,
		lipgloss.WithWhitespaceChars("."),
		lipgloss.WithWhitespaceForeground(theme.EvenItemStyle.GetBackground()),
	)
}

func (entryInfo *EntryInfo) getFullPath() string {
	if entryInfo.entry.SymLinkPath != "" {
		return entryInfo.entry.SymLinkPath
	}
	return filepath.Join(entryInfo.path, entryInfo.entry.Name())
}
