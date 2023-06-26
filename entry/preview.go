package entry

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"github.com/alecthomas/chroma/formatters"
	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/styles"
	"github.com/charmbracelet/glamour"
	"github.com/spf13/afero"
)

type Preview struct {
	Content  string
	Err      error
	Path     string
	ReadTime time.Time
}

func CreatePreview(ctx context.Context, fsys afero.Fs, preview Preview, maxBytes int) Preview {
	previewChan := make(chan Preview)
	errc := make(chan error, 1)
	go func(prev Preview) {
		defer close(previewChan)
		defer close(errc)

		stat, err := fsys.Stat(prev.Path)
		if err != nil {
			errc <- err
			return
		}
		if !prev.ReadTime.IsZero() {
			if prev.ReadTime.After(stat.ModTime()) {
				prev.ReadTime = time.Now()
				previewChan <- prev
				return
			}
		}

		// return early if context is cancelled
		if ctx.Err() != nil {
			errc <- ctx.Err()
		}

		file, err := fsys.Open(prev.Path)
		if err != nil {
			errc <- err
			return
		}
		defer file.Close()

		// return early if context is cancelled
		if ctx.Err() != nil {
			errc <- ctx.Err()
		}

		mime, err := GetMimeTypeByRead(file)
		if err != nil {
			errc <- err
			return
		}

		isText := strings.HasPrefix(mime, "text/")
		if !isText {
			errc <- fmt.Errorf("not a text file")
			return
		}

		// return early if context is cancelled
		if ctx.Err() != nil {
			errc <- ctx.Err()
		}

		content, err := createPreview(ctx, filepath.Base(prev.Path), file, maxBytes)
		p := Preview{
			Content:  content,
			Err:      err,
			Path:     prev.Path,
			ReadTime: time.Now(),
		}
		previewChan <- p
	}(preview)

	select {
	case <-ctx.Done():
		return preview
	case err := <-errc:
		preview.Err = err
		return preview
	case p := <-previewChan:
		return p
	}
}

func readNBytes(ctx context.Context, reader io.Reader, nBytes int) (string, error) {
	buf := make([]byte, nBytes)
	nRead, err := io.ReadAtLeast(reader, buf, nBytes)
	if errors.Is(err, io.ErrUnexpectedEOF) || errors.Is(err, io.EOF) {
		err = nil
	}
	return string(buf[:nRead]), err
}

func createPreview(ctx context.Context, fileName string, reader io.Reader, maxBytes int) (string, error) {
	preview, err := readNBytes(ctx, reader, maxBytes)
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

// GetMimeTypeByRead returns the mime type of a file by reading up to 512 bytes of its content.
func GetMimeTypeByRead(seeker io.ReadSeeker) (string, error) {
	// At most the first 512 bytes of data are used:
	// https://golang.org/src/net/http/sniff.go?s=646:688#L11
	// Without this buffer, http.DetectContentType will not correctly identify text files if they are smaller than 512 bytes
	buff := make([]byte, 512)

	_, err := seeker.Seek(0, io.SeekStart)
	if err != nil {
		return "", err
	}

	bytesRead, err := seeker.Read(buff)
	if err != nil && err != io.EOF {
		return "", err
	}
	seeker.Seek(0, io.SeekStart)
	// Slice to remove fill-up zero values which cause a wrong content type detection in the next step
	buff = buff[:bytesRead]

	return http.DetectContentType(buff), nil
}
