package fileutils

import (
	"errors"
	"io"
	"path/filepath"
	"runtime"
	"strings"

	"io/fs"

	"github.com/spf13/afero"
)

// RenameOrCopy renames a file, leaving source file intact in case of failure.
// Tries hard to succeed on various systems by temporarily tweaking directory
// permissions and removing the destination file when necessary.
func RenameOrCopy(fsys afero.Fs, from, to string) error {

	return withPreparedTarget(fsys, from, to, func() error {
		if fsys.Rename(from, to) == nil {
			return nil
		}

		// Everything is sad, do a copy and delete.
		if _, err := fsys.Stat(to); !errors.Is(err, fs.ErrNotExist) {
			err := fsys.Remove(to)
			if err != nil {
				return err
			}
		}

		err := copyFileContents(fsys, from, to)
		if err != nil {
			_ = fsys.Remove(to)
			return err
		}

		return withPreparedTarget(fsys, from, from, func() error {
			return fsys.Remove(from)
		})
	})
}

// Copy copies the file content from source to destination.
// Tries hard to succeed on various systems by temporarily tweaking directory
// permissions and removing the destination file when necessary.
func Copy(fsys afero.Fs, from, to string) (err error) {
	return withPreparedTarget(fsys, from, to, func() error {
		return copyFileContents(fsys, from, to)
	})
}

// Tries hard to succeed on various systems by temporarily tweaking directory
// permissions and removing the destination file when necessary.
func withPreparedTarget(filesystem afero.Fs, from, to string, f func() error) error {
	// Make sure the destination directory is writeable
	toDir := filepath.Dir(to)
	if info, err := filesystem.Stat(toDir); err == nil && info.IsDir() && info.Mode()&0200 == 0 {
		filesystem.Chmod(toDir, 0755)
		defer filesystem.Chmod(toDir, info.Mode())
	}

	// On Windows, make sure the destination file is writeable (or we can't delete it)
	if runtime.GOOS == "windows" {
		filesystem.Chmod(to, 0666)
		if !strings.EqualFold(from, to) {
			err := filesystem.Remove(to)
			if err != nil && !errors.Is(err, fs.ErrNotExist) {
				return err
			}
		}
	}
	return f()
}

// copyFileContents copies the contents of the file named src to the file named
// by dst. The file will be created if it does not already exist. If the
// destination file exists, all its contents will be replaced by the contents
// of the source file.
func copyFileContents(fsys afero.Fs, from, to string) error {
	in, err := fsys.Open(from)
	if err != nil {
		return err
	}
	defer in.Close()
	out, err := fsys.Create(to)
	if err != nil {
		return err
	}
	defer func() {
		cerr := out.Close()
		if err == nil {
			err = cerr
		}
	}()
	_, err = io.Copy(out, in)
	return err
}
