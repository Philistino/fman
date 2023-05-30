package entryinfo

import (
	"bufio"
	"bytes"
	"context"
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"
	"unicode/utf8"

	"github.com/alecthomas/chroma/formatters"
	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/styles"
	tea "github.com/charmbracelet/bubbletea"
)

var UnreadableErr = errors.New("unreadable content")
var HighlightErr = errors.New("failed to highlight syntax")

type previewReadyMsg struct {
	Path       string
	Preview    string
	Err        error
	EndReached bool
}

func (p previewReadyMsg) Error() error {
	return p.Err
}

func getPreviewCmdOLDKEEP(ctx context.Context, path string, height int, offset int) tea.Cmd {
	return func() tea.Msg {
		f, err := os.Open(path)
		if err != nil {
			return err
		}
		defer f.Close()

		preview, endReached, err := handlePreviewFunc(ctx, f, path, height, offset)
		return previewReadyMsg{
			Path:       path,
			Preview:    preview,
			Err:        err,
			EndReached: endReached,
		}
	}
}

func getPreviewCmd(ctx context.Context, path string, height int, offset int) tea.Cmd {
	return func() tea.Msg {
		previewChan := make(chan previewReadyMsg)
		errc := make(chan error, 1)
		go func() {
			f, err := os.Open(path)
			if err != nil {
				errc <- err
				return
			}
			defer f.Close()
			preview, endReached, err := handlePreviewFunc(ctx, f, path, height, offset)
			p := previewReadyMsg{
				Preview:    preview,
				Err:        err,
				EndReached: endReached,
				Path:       path,
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

// this implementation opens and closes the file for reading on every scroll, which is probably a little slow
func getFilePreviewFunc(ctx context.Context, reader io.Reader, height int, offset int) (string, bool, error) {
	strBuilder := strings.Builder{}
	scanner := bufio.NewScanner(reader)
	eofReached := false
	for i := 0; i < height+offset; i++ {
		ok := scanner.Scan()
		if !ok {
			eofReached = true
			break
		}
		if i < offset {
			continue
		}
		strBuilder.Write(scanner.Bytes())
		strBuilder.WriteByte('\n')
	}

	if !utf8.ValidString(strBuilder.String()) {
		return "", eofReached, errors.New("unable to show preview")
	}
	if err := scanner.Err(); err != nil {
		return "", eofReached, err
	}
	return strBuilder.String(), eofReached, nil
}

func handlePreviewFunc(ctx context.Context, reader io.Reader, path string, height int, offset int) (string, bool, error) {
	preview, eofReached, err := getFilePreviewFunc(ctx, reader, height, offset)
	if err != nil {
		return "", eofReached, err
	}

	if preview == "" {
		return "", eofReached, nil
	}

	preview, err = highlightSyntax(filepath.Base(path), preview)
	if err != nil {
		return "", eofReached, HighlightErr
	}

	// tabs are rendered with different widths based on terminal and font settings
	// so we replace the tab with four spaces so we can reliably truncate each line
	preview = strings.ReplaceAll(preview, "\t", "    ")
	return preview, eofReached, nil
}

func highlightSyntax(name string, preview string) (string, error) {
	var buffer bytes.Buffer

	lexer := lexers.Match(name)
	if lexer == nil {
		lexer = lexers.Fallback
	}
	style := styles.Get("monokai")
	formatter := formatters.Get("terminal")

	iterator, err := lexer.Tokenise(nil, preview)
	if err != nil {
		return "", err
	}

	if err = formatter.Format(&buffer, style, iterator); err != nil {
		return "", err
	}

	return buffer.String(), nil
}
