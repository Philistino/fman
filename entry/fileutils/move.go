package fileutils

import (
	"path/filepath"

	"github.com/spf13/afero"
)

// MoveOrCopy moves or copies a file or directory from src to dst.
// Both src and dst must be absolute paths.
// A rename is attempted but if it fails a copy operation is performed.
func MoveOrCopy(fs afero.Fs, src, dst string) error {
	src = filepath.Clean(src)
	dst = filepath.Clean(dst)

	err := fs.Rename(src, dst)
	if err == nil {
		return nil
	}
	// fallback
	stat, err := fs.Stat(src)
	if err != nil {
		return err
	}
	if stat.IsDir() {
		err := CopyDir(fs, src, dst)
		if err != nil {
			return err
		}
		return fs.RemoveAll(src)
	}
	return CopyFile(fs, src, dst)
}
