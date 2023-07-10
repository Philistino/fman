package fileutils

import (
	"context"
	"io"
	"os"
	"path"
	"path/filepath"

	"github.com/spf13/afero"
	"golang.org/x/sync/errgroup"
)

// MoveFile moves file from src to dst.
// By default the rename filesystem system call is used. If src and dst point to different volumes
// the file copy is used as a fallback
func MoveFile(fs afero.Fs, src, dst string) error {
	err := RenameOrCopy(fs, src, dst)
	return err
}

// MoveFileMany moves files from src into dst.
// By default the rename filesystem system call is used. If src and dst point to different volumes
// the file copy is used as a fallback
func MoveFileMany(ctx context.Context, fs afero.Fs, src []string, dst string) []error {
	g := errgroup.Group{}
	g.SetLimit(10)
	errs := make([]error, len(src))
	for i, f := range src {
		i, f := i, f
		if ctx.Err() != nil {
			errs[i] = ctx.Err()
			continue
		}
		g.Go(func() error {
			err := MoveFile(fs, f, filepath.Join(dst, filepath.Base(f)))
			if err != nil {
				errs[i] = err
			}
			return nil
		})
	}
	g.Wait()
	return errs
}

// CopyFile copies a file from source to dest and returns
// an error if any.
func CopyFile(fs afero.Fs, source, dest string) error {
	// Open the source file.
	src, err := fs.Open(source)
	if err != nil {
		return err
	}
	defer src.Close()

	// Makes the directory needed to create the dst
	// file.
	err = fs.MkdirAll(filepath.Dir(dest), 0666) //nolint:gomnd
	if err != nil {
		return err
	}

	// Create the destination file.
	dst, err := fs.OpenFile(dest, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0775) //nolint:gomnd
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy the contents of the file.
	_, err = io.Copy(dst, src)
	if err != nil {
		return err
	}

	// Copy the mode
	info, err := fs.Stat(source)
	if err != nil {
		return err
	}
	// this ignores an error if the file system does not support chmod
	fs.Chmod(dest, info.Mode())

	return nil
}

// CommonPrefix returns common directory path of provided files
func CommonPrefix(sep byte, paths ...string) string {
	// Handle special cases.
	switch len(paths) {
	case 0:
		return ""
	case 1:
		return path.Clean(paths[0])
	}

	// Note, we treat string as []byte, not []rune as is often
	// done in Go. (And sep as byte, not rune). This is because
	// most/all supported OS' treat paths as string of non-zero
	// bytes. A filename may be displayed as a sequence of Unicode
	// runes (typically encoded as UTF-8) but paths are
	// not required to be valid UTF-8 or in any normalized form
	// (e.g. "é" (U+00C9) and "é" (U+0065,U+0301) are different
	// file names.
	c := []byte(path.Clean(paths[0]))

	// We add a trailing sep to handle the case where the
	// common prefix directory is included in the path list
	// (e.g. /home/user1, /home/user1/foo, /home/user1/bar).
	// path.Clean will have cleaned off trailing / separators with
	// the exception of the root directory, "/" (in which case we
	// make it "//", but this will get fixed up to "/" bellow).
	c = append(c, sep)

	// Ignore the first path since it's already in c
	for _, v := range paths[1:] {
		// Clean up each path before testing it
		v = path.Clean(v) + string(sep)

		// Find the first non-common byte and truncate c
		if len(v) < len(c) {
			c = c[:len(v)]
		}
		for i := 0; i < len(c); i++ {
			if v[i] != c[i] {
				c = c[:i]
				break
			}
		}
	}

	// Remove trailing non-separator characters and the final separator
	for i := len(c) - 1; i >= 0; i-- {
		if c[i] == sep {
			c = c[:i]
			break
		}
	}

	return string(c)
}

var PathAlreadyExistsError = &os.PathError{Op: "Create", Err: os.ErrExist}

// MkFileIfNotExist creates a file if it does not exist
// returns an error if the file already exists
func MkFileIfNotExist(fs afero.Fs, path string) error {
	exists, err := afero.Exists(fs, path)
	if err != nil {
		return err
	}
	if exists {
		return PathAlreadyExistsError
	}
	file, err := fs.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	return nil
}
