package entry

import (
	"bytes"
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
