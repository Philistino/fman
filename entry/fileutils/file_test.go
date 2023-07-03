package fileutils

import (
	"context"
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
