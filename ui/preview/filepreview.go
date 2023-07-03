package preview

import (
	"context"
	"fmt"
	"path/filepath"
	"strings"
	"time"

	"github.com/Philistino/fman/entry"
	"github.com/Philistino/fman/ui/keys"
	"github.com/Philistino/fman/ui/message"
	"github.com/Philistino/fman/ui/theme"
	"github.com/Philistino/fman/ui/theme/colors"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/viewport"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/termenv"
)

const margin = 2

var previewStyle = lipgloss.NewStyle()

type PreviewReadyMsg struct {
	Path    string
	Preview string
	Err     error
}

type previewState uint8

const (
	previewStatePreviewing      previewState = iota // show preview
	previewStateLoadingDirPre                       // show old preview
	previewStateLoadingDirPost                      // show spinner
	previewStateLoadingFilePre                      // show old preview
	previewStateLoadingFilePost                     // show spinner
)

type FilePreview struct {
	theme    colors.Theme
	viewPort viewport.Model // viewport for scrolling preview

	width         int
	height        int
	previewHeight int

	previewCancel context.CancelFunc
	dirPath       string
	entry         entry.Entry
	loadingDelay  time.Duration

	spinner spinner.Model

	state previewState
}

func NewFilePreviewer(theme colors.Theme, previewDelay int) *FilePreview {

	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(theme.SelectedItemBgColor)

	// set loading delay to 1.5x the preview delay or 500ms, whichever is less
	previewDelay = previewDelay * 3 / 2
	if previewDelay < 250 || previewDelay > 750 {
		previewDelay = 500
	}
	f := &FilePreview{
		height:       10,
		width:        10,
		theme:        theme,
		viewPort:     viewport.New(1_000_000, 10), // set width super wide to avoid text wrapping
		state:        previewStateLoadingDirPre,
		loadingDelay: time.Duration(time.Millisecond * time.Duration(previewDelay)),
		spinner:      s,
	}
	return f
}

func (fp *FilePreview) Init() tea.Cmd {
	return fp.spinner.Tick
}

func (fp *FilePreview) setNewEntry(entry entry.Entry) tea.Cmd {
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
		fp.state = previewStatePreviewing
		return nil
	}
	if entry.Size() == 0 {
		fp.viewPort.SetContent(fp.renderNoPreview("Empty file"))
		fp.state = previewStatePreviewing
		return nil
	}

	// if file is not immediately identifiable as text and is larger than 1MB, don't preview
	if !fp.isText() && entry.Size() > 1_000_000 {
		fp.viewPort.SetContent(fp.renderNoPreview("No preview available"))
		fp.state = previewStatePreviewing
		return nil
	}

	ctx, cancel := context.WithCancel(context.Background())
	fp.previewCancel = cancel

	cmd := message.GetPreviewCmd(ctx, fp.getFullPath())

	fp.state = previewStateLoadingFilePre

	return tea.Batch(cmd, fileLoadingCmd(fp.getFullPath(), fp.loadingDelay))
}

// showSpinFileMsg is sent after a delay to indicate that the spinner
// should be shown
type showSpinFileMsg struct {
	Path string
}

type showSpinDirMsg struct {
	Path string
}

// fileLoadingCmd is called when a new preview is needed.
// it returns a message indicating that loading spinner should be shown
// if the preview has not yet been loaded
func fileLoadingCmd(path string, delay time.Duration) tea.Cmd {
	return func() tea.Msg {
		time.Sleep(delay)
		return showSpinFileMsg{Path: path}
	}
}

// dirLoadingCmd is called when a new directory has been selected.
// it returns a message indicating that loading spinner should be shown
// if the preview has not yet been loaded
func dirLoadingCmd(path string, delay time.Duration) tea.Cmd {
	return func() tea.Msg {
		time.Sleep(delay)
		return showSpinDirMsg{Path: path}
	}
}

func (fp *FilePreview) handleShowSpinDirMsg(msg showSpinDirMsg) {
	if msg.Path != fp.dirPath {
		return
	}
	if fp.state == previewStateLoadingDirPre {
		fp.state = previewStateLoadingDirPost
	}
}

func (fp *FilePreview) handleShowSpinFileMsg(msg showSpinFileMsg) {
	if msg.Path != fp.getFullPath() {
		return
	}
	if fp.state == previewStateLoadingFilePre {
		fp.state = previewStateLoadingFilePost
	}
}

// handleDirChangedMsg is called when the directory changes
func (fp *FilePreview) handleDirChangedMsg(msg message.DirChangedMsg) tea.Cmd {
	// if there was an error loading the directory, its a no-op
	if msg.Error() != nil {
		return nil
	}

	fp.dirPath = msg.Path()

	if len(msg.Entries()) == 0 {
		fp.state = previewStatePreviewing
		fp.viewPort.SetContent(fp.renderNoPreview("Select a file to preview"))
		return nil
	}

	cmd := dirLoadingCmd(msg.Path(), fp.loadingDelay)
	fp.state = previewStateLoadingDirPre
	return cmd
}

func (fp *FilePreview) handlePreviewReadyMsg(msg PreviewReadyMsg) {
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
	fp.state = previewStatePreviewing
}

func (fp *FilePreview) Update(msg tea.Msg) (*FilePreview, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case message.DirChangedMsg: // Can this can be merged with message.EntryMsg?
		cmd = fp.handleDirChangedMsg(msg)
		cmds = append(cmds, cmd)
	case message.NewEntryMsg:
		cmd = fp.setNewEntry(msg.Entry)
		cmds = append(cmds, cmd)
	case showSpinFileMsg:
		fp.handleShowSpinFileMsg(msg)
	case showSpinDirMsg:
		fp.handleShowSpinDirMsg(msg)
	case PreviewReadyMsg:
		fp.handlePreviewReadyMsg(msg)
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

func (fp *FilePreview) fileInfoView() string {
	str := strings.Builder{}
	str.WriteString(termenv.String(strings.Repeat("-", fp.width-margin)).Foreground(termenv.RGBColor(fp.theme.InfobarBgColor)).String())
	str.WriteByte('\n')
	str.WriteString(termenv.String("Modified ").Italic().String())
	str.WriteString(fp.entry.ModifyTime)
	return str.String()
}

func (fp *FilePreview) View() string {
	fileInfo := fp.fileInfoView()
	fp.previewHeight = fp.height - lipgloss.Height(fileInfo) - margin

	var mainView string
	if fp.state == previewStateLoadingFilePost || fp.state == previewStateLoadingDirPost {
		mainView = fp.renderLoadingPreview()
	} else {
		mainView = fp.viewPort.View()
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

func (fp *FilePreview) renderNoPreview(text string) string {
	return fp.renderText(text)
}

func (fp *FilePreview) renderLoadingPreview() string {
	loadingContent := "Loading preview..."
	if fp.state == previewStateLoadingDirPost {
		loadingContent = "Loading directory..."
	}
	return fp.renderText(fmt.Sprintf("%s %s", fp.spinner.View(), loadingContent))
}

func (fp *FilePreview) renderText(text string) string {
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

func (fp *FilePreview) getFullPath() string {
	if fp.entry.SymLinkPath != "" {
		return fp.entry.SymLinkPath
	}
	return filepath.Join(fp.dirPath, fp.entry.Name())
}

// isText returns true if the file is likely a text file
func (fp *FilePreview) isText() bool {
	types := []string{"application/xml", "application/json", "application/javascript", "application/x-sh", "text/"}
	for _, t := range types {
		if strings.Contains(fp.entry.MimeType, t) {
			return true
		}
	}
	return false
}

// SetWidth sets the width of the preview
func (fp *FilePreview) SetWidth(width int) {
	fp.width = width
}

// Height returns the height of the preview
func (fp *FilePreview) Height() int {
	return fp.height
}

// SetHeight sets the height of the preview
func (fp *FilePreview) SetHeight(height int) {
	fp.viewPort.Height = height - lipgloss.Height(fp.fileInfoView()) - margin
	fp.height = height
}
