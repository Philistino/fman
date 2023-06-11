package entryinfo

import (
	"bufio"
	"context"
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/Philistino/fman/entry"
	"golang.org/x/text/unicode/norm"
)

type Previewer interface {
	Preview() (string, error)
	ScrollUp() error
	ScrollDown() error
}

type filePreviewer struct {
	path         string   // SHOULD THIS BE AN FS.FILE INSTEAD?
	wholeText    []string //
	start        int      // start line of preview
	height       int      // number of lines to return in preview
	nLines       int      // number of additional lines to read from the file at each update
	validStr     bool     // is text a valid string?
	validChecked bool     // has the text been checked for validity?
	eofReached   bool
}

func NewFilePreviewer(path string, height int, readLines int) *filePreviewer {
	return &filePreviewer{
		path:         path,
		wholeText:    nil,
		start:        0,
		height:       height,
		nLines:       readLines,
		validStr:     false,
		validChecked: false,
		eofReached:   false,
	}
}

func (fp *filePreviewer) updatePreview(ctx context.Context) error {
	if fp.validChecked && !fp.validStr {
		return nil
	}
	file, err := os.Open(fp.path)
	if err != nil {
		fp.validChecked = true
		fp.validStr = false
		return err
	}
	defer file.Close()
	if !fp.validChecked {
		fp.validStr = checkMimeType(file)
		fp.validChecked = true
	}
	if !fp.validStr {
		return nil
	}
	// read lines
	preview, eof, err := createPreview(ctx, filepath.Base(fp.path), file, len(fp.wholeText)+fp.nLines)
	if err != nil {
		return err
	}
	fp.eofReached = eof
	fp.wholeText = strings.Split(preview, "\n")
	return nil
}

// checkMimeType returns true if the mime type is text/*
func checkMimeType(seeker io.ReadSeeker) bool {
	mime, err := entry.GetMimeType(seeker)
	if err != nil {
		return false
	}
	return strings.HasPrefix(mime, "text/")
}

func readLines(ctx context.Context, reader io.Reader, nLines int, maxBytes int) (string, bool, error) {
	// _, err := seeker.Seek(0, 0)
	// if err != nil {
	// 	return "", false, err
	// }
	strBuilder := strings.Builder{}
	strBuilder.Grow(maxBytes)
	scanner := bufio.NewScanner(reader)
	scanner.Buffer(make([]byte, maxBytes), maxBytes)
	eofReached := false

	// read from the start to the end of nLines or EOF
	for i := 0; i < nLines; i++ {

		if ctx.Err() != nil {
			return "", eofReached, ctx.Err()
		}

		ok := scanner.Scan()
		// ok returns false if EOF is reached or an error occurs
		if !ok {
			// err will be nil if EOF is reached
			if scanner.Err() == nil {
				eofReached = true
			}
			break
		}

		if i != 0 {
			strBuilder.WriteByte('\n')
		}

		// get the size of bytes read by scanner
		// if the size is greater than maxBytes, add up to maxBytes from the scanner
		// to strBuilder and break else add all the bytes to strBuilder
		scannedBytes := scanner.Bytes()
		maxAdd := maxBytes - strBuilder.Len()
		if len(scannedBytes) > maxAdd {
			strBuilder.Write(scannedBytes[:maxAdd])
			break
		}
		strBuilder.Write(scannedBytes)
	}
	text := norm.NFC.String(strBuilder.String())
	return text, eofReached, scanner.Err()
}

func readBytes(ctx context.Context, reader io.Reader, maxBytes int) (string, error) {
	buf := make([]byte, maxBytes)
	nRead, err := io.ReadAtLeast(reader, buf, maxBytes)
	if errors.Is(err, io.ErrUnexpectedEOF) || errors.Is(err, io.EOF) {
		err = nil
	}
	return string(buf[:nRead]), err
}

func createPreview(ctx context.Context, fileName string, reader io.Reader, nLines int) (string, bool, error) {
	preview, eofReached, err := readLines(ctx, reader, nLines, 1_000_000)
	if err != nil {
		return "", eofReached, err
	}

	if preview == "" {
		return "", eofReached, nil
	}

	preview, err = highlightSyntax(fileName, preview)
	if err != nil {
		return "", eofReached, HighlightErr
	}

	// tabs are rendered with different widths based on terminal and font settings
	// so we replace the tab with four spaces so we can reliably truncate each line
	preview = strings.ReplaceAll(preview, "\t", "    ")
	return preview, eofReached, nil
}
