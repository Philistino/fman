package model

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/Philistino/fman/entry"
	"github.com/Philistino/fman/model/keys"
	"github.com/Philistino/fman/model/message"
	"github.com/Philistino/fman/theme"
	"github.com/Philistino/fman/theme/colors"
	"github.com/alecthomas/chroma/formatters"
	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/styles"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/termenv"
)

const margin = 2

var UnreadableErr = errors.New("unreadable content")
var HighlightErr = errors.New("failed to highlight syntax")
var previewStyle = lipgloss.NewStyle()

type FilePreview struct {
	theme        colors.Theme
	viewPort     viewport.Model // viewport for scrolling preview
	previewDelay int            //  delay in ms before openning a file to create a preview

	width         int
	height        int
	previewHeight int

	previewCancel context.CancelFunc
	path          string
	entry         entry.Entry
}

func NewFilePreviewer(theme colors.Theme, previewDelay int) FilePreview {
	return FilePreview{
		height:       10,
		width:        10,
		theme:        theme,
		previewDelay: previewDelay,
		viewPort:     viewport.New(1_000_000, 10), // set width super wide to avoid text wrapping
	}
}

func (fp *FilePreview) Init() tea.Cmd {
	return nil
}

func (fp *FilePreview) setNewEntry(entry entry.Entry) tea.Cmd {
	fp.entry = entry
	// handle preview context cancellation
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

	// set default preview content
	fp.viewPort.SetContent(fp.renderNoPreview("Loading preview..."))
	return fp.getPreview(true)
}

func (fp *FilePreview) getPreview(delay bool) tea.Cmd {
	ctx, cancel := context.WithCancel(context.Background())
	fp.previewCancel = cancel

	pDelay := 0
	if delay {
		pDelay = fp.previewDelay
	}
	return getPreviewCmd(
		ctx,
		fp.getFullPath(),
		pDelay,
	)
}

type previewReadyMsg struct {
	Path    string
	Preview string
	Err     error
}

func (fp *FilePreview) Update(msg tea.Msg) (FilePreview, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case message.DirChangedMsg: // Can this can be merged with message.EntryMsg?
		fp.path = msg.Path()
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
	return *fp, cmd
}

func (fp *FilePreview) handlePreviewMsg(msg previewReadyMsg) {
	// check that the path matches so we don't set the current preview based on the previous file
	if msg.Path != fp.getFullPath() {
		return
	}
	if msg.Err != nil {
		fp.viewPort.SetContent(fp.renderNoPreview("Failed to load preview"))
	}
	if msg.Preview != "" {
		fp.viewPort.SetContent(msg.Preview)
	}
}

func (fp *FilePreview) getFullPath() string {
	if fp.entry.SymLinkPath != "" {
		return fp.entry.SymLinkPath
	}
	return filepath.Join(fp.path, fp.entry.Name())
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
	return theme.EntryInfoStyle.Render(
		lipgloss.JoinVertical(
			lipgloss.Left,
			previewStyle.
				MaxHeight(fp.previewHeight).
				Height(fp.previewHeight). // could set Width(entryInfo.width-margin) here to to wrap lines.
				MaxWidth(fp.width-margin).
				Render(fp.viewPort.View()),
			fileInfo,
		),
	)
}

func (fp *FilePreview) SetWidth(width int) {
	fp.width = width
}

func (fp *FilePreview) Height() int {
	return fp.height
}

func (fp *FilePreview) SetHeight(height int) {
	fp.viewPort.Height = height - lipgloss.Height(fp.fileInfoView()) - margin
	fp.height = height
}

func (fp *FilePreview) renderNoPreview(text string) string {
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

// readDelay is how long in milliseconds to wait before reading the file.
// If negative, it will not wait.
// This is meant to avoid unnecessary disk io when the user is navigating quickly.
func getPreviewCmd(ctx context.Context, path string, readDelay int) tea.Cmd {
	return func() tea.Msg {
		previewChan := make(chan previewReadyMsg)
		errc := make(chan error, 1)
		go func() {
			defer close(previewChan)
			defer close(errc)

			if readDelay >= 0 {
				time.Sleep(time.Millisecond * time.Duration(readDelay))
			}

			// return early if context is cancelled
			if ctx.Err() != nil {
				errc <- ctx.Err()
			}
			f, err := os.Open(path)
			if err != nil {
				errc <- err
				return
			}
			defer f.Close()

			// return early if context is cancelled
			if ctx.Err() != nil {
				errc <- ctx.Err()
			}
			isText := checkMimeType(f)
			if !isText {
				errc <- fmt.Errorf("not a text file")
				return
			}

			preview, err := createPreview(ctx, filepath.Base(path), f)
			p := previewReadyMsg{
				Preview: preview,
				Err:     err,
				Path:    path,
			}
			previewChan <- p
		}()
		select {
		case <-ctx.Done():
			return previewReadyMsg{Err: ctx.Err(), Path: path}
		case err := <-errc:
			return previewReadyMsg{Err: err, Path: path}
		case p := <-previewChan:
			return p
		}
	}
}

func checkMimeType(seeker io.ReadSeeker) bool {
	mime, err := entry.GetMimeType(seeker)
	if err != nil {
		return false
	}
	return strings.HasPrefix(mime, "text/")
}

func readBytes(ctx context.Context, reader io.Reader, maxBytes int) (string, error) {
	buf := make([]byte, maxBytes)
	nRead, err := io.ReadAtLeast(reader, buf, maxBytes)
	if errors.Is(err, io.ErrUnexpectedEOF) || errors.Is(err, io.EOF) {
		err = nil
	}
	return string(buf[:nRead]), err
}

func createPreview(ctx context.Context, fileName string, reader io.Reader) (string, error) {
	preview, err := readBytes(ctx, reader, 1_000_000)
	if err != nil || preview == "" {
		return preview, err
	}

	if ctx.Err() != nil {
		return "", ctx.Err()
	}

	if filepath.Ext(fileName) == ".md" {
		preview, _ = renderMarkdown(preview)
	} else {
		preview, _ = highlightSyntax(fileName, preview)
	}

	// tabs are rendered with different widths based on terminal and font settings
	// so we replace the tab with four spaces so we can reliably truncate each line
	preview = strings.ReplaceAll(preview, "\t", "    ")
	return preview, nil
}

func highlightSyntax(name string, preview string) (string, error) {
	lexer := lexers.Match(name)
	if lexer == nil {
		lexer = lexers.Fallback
	}
	style := styles.Get("monokai")
	formatter := formatters.Get("terminal")

	iterator, err := lexer.Tokenise(nil, preview)
	if err != nil {
		return preview, err
	}

	var buffer bytes.Buffer
	err = formatter.Format(&buffer, style, iterator)
	if err != nil {
		return preview, err
	}

	return buffer.String(), nil
}

func renderMarkdown(content string) (string, error) {
	str, err := glamour.Render(content, "dracula")
	if err != nil {
		return content, err
	}
	return str, nil
}
