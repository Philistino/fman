package model

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
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/termenv"
)

const margin = 2

var previewStyle = lipgloss.NewStyle()

type filePreview struct {
	theme        colors.Theme
	viewPort     viewport.Model // viewport for scrolling preview
	previewDelay int            //  delay in ms before openning a file to create a preview

	width         int
	height        int
	previewHeight int

	previewCancel context.CancelFunc
	dirPath       string
	entry         entry.Entry
}

func NewFilePreviewer(theme colors.Theme, previewDelay int) *filePreview {
	return &filePreview{
		height:       10,
		width:        10,
		theme:        theme,
		previewDelay: previewDelay,
		viewPort:     viewport.New(1_000_000, 10), // set width super wide to avoid text wrapping
	}
}

func (fp *filePreview) Init() tea.Cmd {
	return nil
}

func (fp *filePreview) setNewEntry(entry entry.Entry) tea.Cmd {
	fp.entry = entry
	// handle preview context cancellation for previous file
	if fp.previewCancel != nil {
		fp.previewCancel()
		fp.previewCancel = nil
	}
	// reset read position to start of file
	fp.viewPort.SetYOffset(0)

	if entry.IsDir() {
		fp.viewPort.SetContent(fp.renderNoPreview("Directory"))
		return nil
	}

	if entry.Size() == 0 {
		fp.viewPort.SetContent(fp.renderNoPreview("Directory"))
		return nil
	}

	isText := false
	types := []string{"application/xml", "application/json", "application/javascript", "application/x-sh", "text/"}
	for _, t := range types {
		text := strings.Contains(entry.MimeType, t)
		if text {
			isText = true
			break
		}
	}
	// if file is not immediately identifiable as text and is larger than 1MB, don't preview
	if !isText && entry.Size() > 1_000_000 {
		fp.viewPort.SetContent(fp.renderNoPreview("No preview available"))
		return nil
	}

	ctx, cancel := context.WithCancel(context.Background())
	fp.previewCancel = cancel

	cmd := message.GetPreviewCmd(ctx, fp.getFullPath())
	return tea.Batch(cmd)
}

type previewReadyMsg struct {
	Path    string
	Preview string
	Err     error
}

func (fp *filePreview) Update(msg tea.Msg) (*filePreview, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case message.DirChangedMsg: // Can this can be merged with message.EntryMsg?
		fp.dirPath = msg.Path()
	case message.NewEntryMsg:
		cmd = fp.setNewEntry(msg.Entry)
	case previewReadyMsg:
		fp.handlePreviewMsg(msg)
	case tea.KeyMsg:
		if key.Matches(msg, keys.Map.ScrollPreviewDown) {
			fp.viewPort.LineDown(1)
		}
		if key.Matches(msg, keys.Map.ScrollPreviewUp) {
			fp.viewPort.LineUp(1)
		}
	}
	return fp, cmd
}

func (fp *filePreview) handlePreviewMsg(msg previewReadyMsg) {
	// check that the path matches so we don't set the current preview based on the previous file
	if msg.Path != fp.getFullPath() {
		return
	}
	if msg.Err != nil {
		fp.viewPort.SetContent(fp.renderNoPreview("No preview available"))
	}
	if msg.Preview != "" {
		fp.viewPort.SetContent(msg.Preview)
	}
}

func (fp *filePreview) getFullPath() string {
	if fp.entry.SymLinkPath != "" {
		return fp.entry.SymLinkPath
	}
	return filepath.Join(fp.dirPath, fp.entry.Name())
}

func (fp *filePreview) fileInfoView() string {
	str := strings.Builder{}
	str.WriteString(termenv.String(strings.Repeat("-", fp.width-margin)).Foreground(termenv.RGBColor(fp.theme.InfobarBgColor)).String())
	str.WriteByte('\n')
	str.WriteString(termenv.String("Modified ").Italic().String())
	str.WriteString(fp.entry.ModifyTime)
	return str.String()
}

func (fp *filePreview) View() string {
	fileInfo := fp.fileInfoView()
	fp.previewHeight = fp.height - lipgloss.Height(fileInfo) - margin
	return theme.EntryInfoStyle.Render(
		lipgloss.JoinVertical(
			lipgloss.Left,
			previewStyle.
				MaxHeight(fp.previewHeight).
				Height(fp.previewHeight).
				MaxWidth(fp.width-margin).
				Render(fp.viewPort.View()),
			fileInfo,
		),
	)
}

func (fp *filePreview) SetWidth(width int) {
	fp.width = width
}

func (fp *filePreview) Height() int {
	return fp.height
}

func (fp *filePreview) SetHeight(height int) {
	fp.viewPort.Height = height - lipgloss.Height(fp.fileInfoView()) - margin
	fp.height = height
}

func (fp *filePreview) renderNoPreview(text string) string {
	return lipgloss.Place(
		fp.width-2,
		fp.viewPort.Height,
		lipgloss.Center,
		lipgloss.Center,
		text,
		lipgloss.WithWhitespaceChars("."),
		lipgloss.WithWhitespaceForeground(theme.EvenItemStyle.GetBackground()),
	)
}
