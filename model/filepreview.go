package model

import (
	"context"
	"fmt"
	"path/filepath"
	"strings"
	"time"

	"github.com/Philistino/fman/entry"
	"github.com/Philistino/fman/model/keys"
	"github.com/Philistino/fman/model/message"
	"github.com/Philistino/fman/theme"
	"github.com/Philistino/fman/theme/colors"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/viewport"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/termenv"
)

const margin = 2

var previewStyle = lipgloss.NewStyle()

type state uint8

const (
	previewing state = iota
	loadingFile
	loadingDir
)

type filePreview struct {
	theme    colors.Theme
	viewPort viewport.Model // viewport for scrolling preview

	width         int
	height        int
	previewHeight int

	previewCancel   context.CancelFunc
	dirPath         string
	entry           entry.Entry
	viewPortUpdated bool
	loadingDelay    time.Duration

	spinner spinner.Model

	state state
}

func NewFilePreviewer(theme colors.Theme, previewDelay int) *filePreview {

	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(theme.SelectedItemBgColor)

	// set loading delay to 1.5x the preview delay or 500ms, whichever is less
	previewDelay = previewDelay * 3 / 2
	if previewDelay < 500 || previewDelay > 750 {
		previewDelay = 500
	}
	loadingDelay := time.Duration(time.Millisecond * time.Duration(previewDelay))
	f := &filePreview{
		height:       10,
		width:        10,
		theme:        theme,
		viewPort:     viewport.New(1_000_000, 10), // set width super wide to avoid text wrapping
		state:        previewing,
		loadingDelay: loadingDelay,
		spinner:      s,
	}
	return f
}

func (fp *filePreview) Init() tea.Cmd {
	return fp.spinner.Tick
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
		fp.viewPortUpdated = true
		fp.state = previewing
		return nil
	}

	if entry.Size() == 0 {
		fp.viewPort.SetContent(fp.renderNoPreview("Empty file"))
		fp.viewPortUpdated = true
		fp.state = previewing
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
		fp.viewPortUpdated = true
		fp.state = previewing
		return nil
	}

	ctx, cancel := context.WithCancel(context.Background())
	fp.previewCancel = cancel

	cmd := message.GetPreviewCmd(ctx, fp.getFullPath())

	fp.viewPortUpdated = false

	return tea.Batch(cmd, previewLoadingCmd(fp.getFullPath(), fp.loadingDelay))
}

type fileType uint8

const (
	unknown fileType = iota
	file
	directory
)

type previewLoadingMsg struct {
	Path     string
	FileType fileType
}

func previewLoadingCmd(path string, delay time.Duration) tea.Cmd {
	return func() tea.Msg {
		time.Sleep(delay)
		return previewLoadingMsg{Path: path}
	}
}

func dirLoadingCmd(path string, delay time.Duration) tea.Cmd {
	return func() tea.Msg {
		time.Sleep(delay)
		return dirLoadingMsg{Path: path}
	}
}

type previewReadyMsg struct {
	Path    string
	Preview string
	Err     error
}

type dirLoadingMsg struct {
	Path string
}

func (fp *filePreview) handleDirLoadingMsg(msg dirLoadingMsg) {
	if msg.Path != fp.dirPath {
		return
	}
	fp.state = loadingDir
}

func (fp *filePreview) handlePreviewLoadingMsg(msg previewLoadingMsg) {
	if msg.Path != fp.getFullPath() {
		return
	}
	if fp.viewPortUpdated {
		return
	}
	fp.state = loadingFile
}

func (fp *filePreview) Update(msg tea.Msg) (*filePreview, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case message.DirChangedMsg: // Can this can be merged with message.EntryMsg?
		fp.dirPath = msg.Path()
		cmd = dirLoadingCmd(msg.Path(), fp.loadingDelay)
		cmds = append(cmds, cmd)
		fp.state = loadingDir
	case previewLoadingMsg:
		fp.handlePreviewLoadingMsg(msg)
	case message.NewEntryMsg:
		cmd = fp.setNewEntry(msg.Entry)
		cmds = append(cmds, cmd)
	case previewReadyMsg:
		fp.state = previewing
		fp.handlePreviewMsg(msg)
	case message.EmptyDirMsg:
		fp.state = previewing
		fp.viewPort.SetContent(fp.renderNoPreview("Select a file to preview"))
	case tea.KeyMsg:
		if key.Matches(msg, keys.Map.ScrollPreviewDown) {
			fp.viewPort.LineDown(1)
		}
		if key.Matches(msg, keys.Map.ScrollPreviewUp) {
			fp.viewPort.LineUp(1)
		}
	case spinner.TickMsg:
		fp.spinner, cmd = fp.spinner.Update(msg)
		cmds = append(cmds, cmd)
	}
	return fp, tea.Batch(cmds...)
}

func (fp *filePreview) handlePreviewMsg(msg previewReadyMsg) {
	// check that the path matches so we don't set the current preview based on the previous file
	if msg.Path != fp.getFullPath() {
		return
	}

	fp.viewPortUpdated = true
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

	var mainView string
	if fp.state == previewing {
		mainView = fp.viewPort.View()
	} else {
		mainView = fp.renderLoadingPreview()
	}
	return theme.EntryInfoStyle.Render(
		lipgloss.JoinVertical(
			lipgloss.Left,
			previewStyle.
				MaxHeight(fp.previewHeight).
				Height(fp.previewHeight).
				MaxWidth(fp.width-margin).
				Render(mainView),
			fileInfo,
		),
	)
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

func (fp *filePreview) renderLoadingPreview() string {
	loadingContent := "Loading preview..."
	if fp.state == loadingDir {
		loadingContent = "Loading directory..."
	}
	return lipgloss.Place(
		fp.width-2,
		fp.viewPort.Height,
		lipgloss.Center,
		lipgloss.Center,
		fmt.Sprintf("%s %s", fp.spinner.View(), loadingContent),
		lipgloss.WithWhitespaceChars("."),
		lipgloss.WithWhitespaceForeground(theme.EvenItemStyle.GetBackground()),
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
