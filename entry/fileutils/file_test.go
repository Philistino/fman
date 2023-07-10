package fileutils

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"testing"

	"github.com/spf13/afero"
)

func TestCommonPrefix(t *testing.T) {
	testCases := map[string]struct {
		paths []string
		want  string
	}{
		"same lvl": {
			paths: []string{
				"/home/user/file1",
				"/home/user/file2",
			},
			want: "/home/user",
		},
		"sub folder": {
			paths: []string{
				"/home/user/folder",
				"/home/user/folder/file",
			},
			want: "/home/user/folder",
		},
		"relative path": {
			paths: []string{
				"/home/user/folder",
				"/home/user/folder/../folder2",
			},
			want: "/home/user",
		},
		"no common path": {
			paths: []string{
				"/home/user/folder",
				"/etc/file",
			},
			want: "",
		},
		"empty path": {
			paths: []string{},
			want:  "",
		},
		"single path": {
			paths: []string{"file"},
			want:  "file",
		},
	}
	for name, tt := range testCases {
		t.Run(name, func(t *testing.T) {
			if got := CommonPrefix('/', tt.paths...); got != tt.want {
				t.Errorf("CommonPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMoveFileMany(t *testing.T) {
	fs := afero.NewMemMapFs()

	// Create source files
	err := afero.WriteFile(fs, "file1.txt", []byte("file1"), 0644)
	if err != nil {
		t.Fatal(err)
	}
	err = afero.WriteFile(fs, "file2.txt", []byte("file2"), 0644)
	if err != nil {
		t.Fatal(err)
	}
	// Move files to destination
	errs := MoveFileMany(context.Background(), fs, []string{"file1.txt", "file2.txt"}, "destination")
	for _, err := range errs {
		if err != nil {
			t.Fatal(err)
		}
	}

	// Check if files were moved
	_, err = fs.Stat("destination/file1.txt")
	if err != nil {
		t.Errorf("file1.txt was not moved: %v", err)
	}
	_, err = fs.Stat("destination/file2.txt")
	if err != nil {
		t.Errorf("file2.txt was not moved: %v", err)
	}

	// Move non-existent file
	errs = MoveFileMany(context.Background(), fs, []string{"file3.txt"}, "destination")
	if len(errs) != 1 {
		t.Errorf("expected 1 error, got %d", len(errs))
	}
	if errs[0] == nil {
		t.Errorf("expected error, got nil")
	}
}

func TestMkFileIfNotExist(t *testing.T) {
	// Create a new in-memory filesystem for testing
	fs := afero.NewMemMapFs()

	// Test creating a new file
	err := MkFileIfNotExist(fs, "/path/to/new/file.txt")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Test creating a file that already exists
	err = MkFileIfNotExist(fs, "/path/to/new/file.txt")
	if !errors.Is(err, PathAlreadyExistsError) {
		t.Errorf("Expected error: %v, but got: %v", PathAlreadyExistsError, err)
	}
}

func TestMoveFile(t *testing.T) {
	dir, err := os.MkdirTemp("", "test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(dir)

	fsys := afero.NewOsFs()
	src := filepath.Join(dir, "test")
	file, err := fsys.Create(src)
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()
	_, err = file.WriteString("Bingo Bango Bongo")
	if err != nil {
		t.Fatal(err)
	}
	err = file.Close()
	if err != nil {
		t.Fatal(err)
	}

	dst := filepath.Join(dir, "test2")
	err = MoveFile(afero.NewOsFs(), src, dst)
	if err != nil {
		t.Fatal(err)
	}

	_, err = fsys.Stat(src)
	if !errors.Is(err, os.ErrNotExist) {
		t.Errorf("expected error: %v, got: %v", os.ErrNotExist, err)
	}

	file, err = fsys.Open(dst)
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			t.Errorf("expected error: %v, got: %v", os.ErrNotExist, err)
		}
		t.Fatal(err)
	}
	defer file.Close()

	contents, err := afero.ReadAll(file)
	if err != nil {
		t.Fatal(err)
	}
	if string(contents) != "Bingo Bango Bongo" {
		t.Errorf("expected contents: %s, got: %s", "Bingo Bango Bongo", string(contents))
	}

}
