package fileutils

import (
	"context"

	"github.com/spf13/afero"
	"golang.org/x/sync/errgroup"
)

// Remove removes a file or directory
// If the file is a directory, this removes the directory and any children it contains.
func Remove(ctx context.Context, fs afero.Fs, file string) error {
	stat, err := fs.Stat(file)
	if err != nil {
		return err
	}
	if ctx.Err() != nil {
		return ctx.Err()
	}
	if stat.IsDir() {
		return fs.RemoveAll(file)
	}
	return fs.Remove(file)
}

// RemoveMany removes many files and/or directories
// If any path points to a directory, it and all children will be removed.
func RemoveMany(ctx context.Context, fs afero.Fs, files []string) []error {
	g := errgroup.Group{}
	g.SetLimit(10)
	errs := make([]error, len(files))
	for i, file := range files {
		file := file
		i := i
		if ctx.Err() != nil {
			errs[i] = ctx.Err()
			continue
		}
		g.Go(func() error {
			err := Remove(ctx, fs, file)
			errs[i] = err
			return nil
		})
	}
	g.Wait()
	return errs
}
