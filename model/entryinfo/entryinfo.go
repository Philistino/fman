package entryinfo

import (
	"context"
	"path/filepath"
	"strings"

	"github.com/Philistino/fman/entry"
	"github.com/Philistino/fman/model/keys"
	"github.com/Philistino/fman/model/message"
	"github.com/Philistino/fman/theme"
	"github.com/Philistino/fman/theme/colors"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/termenv"
)

// Use a viewport for a file preview. On scroll up/down, update the viewport. If viewport.AtBottom(), read more lines from the file.
// When reading from the file, show a spinner at the bottom?

type EntryInfo struct {
	entry entry.Entry

	width  int
	height int

	path    string
	preview string

	previewHeight int
	previewOffset int

	theme         colors.Theme
	eofReached    bool // is set to true when the end of the file is reached in the preview
	previewCancel context.CancelFunc

	previewDelay int //  delay in ms before reading a preview
}

const margin = 2

var previewStyle = lipgloss.NewStyle()

// previewDelay is how long to wait on a file before reading it to create a preview.
// This is meant to avoid unnecessary disk io when the user is navigating quickly.
func New(theme colors.Theme, previewDelay int) EntryInfo {
	return EntryInfo{
		previewHeight: 10,
		theme:         theme,
		width:         10,
		previewDelay:  previewDelay,
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

	if entry.Size() == 0 {
		entryInfo.preview = entryInfo.renderNoPreview("Empty file")
		entryInfo.eofReached = true
		return nil
	}

	// set default preview content
	entryInfo.preview = entryInfo.renderNoPreview("Loading preview...")
	return entryInfo.getPreview(true)
}

func (entryInfo *EntryInfo) getPreview(delay bool) tea.Cmd {
	ctx, cancel := context.WithCancel(context.Background())
	entryInfo.previewCancel = cancel
	var pDelay int
	if delay {
		pDelay = entryInfo.previewDelay
	} else {
		pDelay = 0
	}
	return getPreviewCmd(
		ctx,
		entryInfo.getFullPath(),
		entryInfo.previewHeight,
		entryInfo.previewOffset,
		pDelay,
	)
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
	case message.DirChangedMsg: // Can this can be merged with message.EntryMsg?
		entryInfo.path = msg.Path()
		entryInfo.eofReached = false
	case tea.KeyMsg:
		if key.Matches(msg, keys.Map.ScrollPreviewDown) {
			if entryInfo.eofReached {
				break
			}
			entryInfo.previewOffset++
			cmd = entryInfo.getPreview(false)
		}
		if key.Matches(msg, keys.Map.ScrollPreviewUp) {
			if entryInfo.previewOffset < 1 {
				break
			}
			entryInfo.previewOffset--
			entryInfo.eofReached = false
			cmd = entryInfo.getPreview(false)
		}
	case message.NewEntryMsg:
		cmd = entryInfo.setNewEntry(msg.Entry)
	case previewReadyMsg:
		entryInfo.handlePreviewMsg(msg)
	}
	return *entryInfo, cmd
}

func (entryInfo *EntryInfo) getFileInfo() string {
	str := strings.Builder{}
	str.WriteString(termenv.String(strings.Repeat("-", entryInfo.width-margin)).Foreground(termenv.RGBColor(entryInfo.theme.InfobarBgColor)).String())
	str.WriteByte('\n')
	str.WriteString(termenv.String("Modified ").Italic().String())
	str.WriteString(entryInfo.entry.ModifyTime)
	return str.String()
}

func (entryInfo *EntryInfo) View() string {
	fileInfo := entryInfo.getFileInfo()
	entryInfo.previewHeight = entryInfo.height - lipgloss.Height(fileInfo) - margin
	return theme.EntryInfoStyle.Render(
		lipgloss.JoinVertical(
			lipgloss.Left,
			previewStyle.
				MaxHeight(entryInfo.previewHeight).
				Height(entryInfo.previewHeight). // could set Width(entryInfo.width-margin) here to to wrap lines.
				MaxWidth(entryInfo.width-margin).
				Render(entryInfo.preview),
			fileInfo,
		),
	)
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
