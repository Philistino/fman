package fileutils

import (
	"errors"
	"os"
	"path/filepath"
	"testing"

	"github.com/spf13/afero"
)

func TestMoveOrCopyFile(t *testing.T) {
	tt := []struct {
		name string
		src  string
		dst  string
	}{
		{
			name: "same folder",
			src:  "/path/to/source",
			dst:  "/path/to/dest",
		},
		{
			name: "different folder",
			src:  "/path/to/source",
			dst:  "/path/to/destDir/destFile",
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			// Create a new memory-based file system for testing
			fs := afero.NewMemMapFs()

			// Create the source file
			err := afero.WriteFile(fs, tc.src, []byte("source file contents"), 0644)
			if err != nil {
				t.Fatalf("failed to create source file: %v", err)
			}

			// Test moving the source file to the destination
			err = MoveOrCopy(fs, tc.src, tc.dst)
			if err != nil {
				t.Fatalf("failed to move file: %v", err)
			}

			// Ensure that the dst file exists
			_, err = fs.Stat(tc.dst)
			if err != nil {
				t.Fatalf("failed to find destination file: %v", err)
			}

			// Ensure that the source file was deleted
			_, err = fs.Stat(tc.src)
			if !errors.Is(err, os.ErrNotExist) {
				t.Fatalf("source file still exists after move: %v", err)
			}

			// Ensure that the destination parent directory exists
			info, err := fs.Stat(filepath.Dir(tc.dst))
			if err != nil {
				t.Fatalf("failed to find destination parent: %v", err)
			}
			if !info.IsDir() {
				t.Fatalf("destination parent is not a directory")
			}
		})
	}
}

func TestMoveOrCopyDir(t *testing.T) {
	tt := []struct {
		name string
		src  string
		dst  string
	}{
		{
			name: "same folder",
			src:  "/path/to/srcDir",
			dst:  "/path/to/destDir",
		},
		{
			name: "different folder",
			src:  "/path/to/srcDir",
			dst:  "/path/to/parent/destDir",
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			// Create a new memory-based file system for testing
			fs := afero.NewMemMapFs()
			fileName := "Bingo"
			// Create the source file
			err := afero.WriteFile(fs, filepath.Join(tc.src, fileName), []byte("source file contents"), 0644)
			if err != nil {
				t.Fatalf("failed to create source file: %v", err)
			}

			// Test moving the source folder to the destination
			err = MoveOrCopy(fs, tc.src, tc.dst)
			if err != nil {
				t.Fatalf("failed to move file: %v", err)
			}

			// Ensure that the dst folder exists
			info, err := fs.Stat(tc.dst)
			if err != nil {
				t.Fatalf("failed to find destination file: %v", err)
			}
			if !info.IsDir() {
				t.Fatalf("destination is not a directory")
			}

			// Ensure that the file in the source dir was moved
			_, err = fs.Stat(filepath.Join(tc.dst, fileName))
			if err != nil {
				t.Fatalf("failed to find destination file: %v", err)
			}

			// Ensure that the source dir was deleted
			_, err = fs.Stat(tc.src)
			if !errors.Is(err, os.ErrNotExist) {
				t.Fatalf("source file still exists after move: %v", err)
			}

			// Ensure that the destination parent directory exists
			info, err = fs.Stat(filepath.Dir(tc.dst))
			if err != nil {
				t.Fatalf("failed to find destination parent: %v", err)
			}
			if !info.IsDir() {
				t.Fatalf("destination parent is not a directory")
			}
		})
	}
}
