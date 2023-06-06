package entry

import (
	"bytes"
	"io"
	"net/http"
	"os"
)

// IsZipFile checks if file is zip or not.
// Play: https://go.dev/play/p/9M0g2j_uF_e
func IsZipFile(filepath string) (bool, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return false, err
	}
	defer f.Close()

	buf := make([]byte, 4)
	if n, err := f.Read(buf); err != nil || n < 4 {
		return false, err
	}

	isZip := bytes.Equal(buf, []byte("PK\x03\x04"))
	return isZip, nil
}

func GetMimeType(seeker io.ReadSeeker) (string, error) {
	// At most the first 512 bytes of data are used:
	// https://golang.org/src/net/http/sniff.go?s=646:688#L11
	buff := make([]byte, 512)

	_, err := seeker.Seek(0, io.SeekStart)
	if err != nil {
		return "", err
	}

	bytesRead, err := seeker.Read(buff)
	if err != nil && err != io.EOF {
		return "", err
	}

	// Slice to remove fill-up zero values which cause a wrong content type detection in the next step
	buff = buff[:bytesRead]

	return http.DetectContentType(buff), nil
}
