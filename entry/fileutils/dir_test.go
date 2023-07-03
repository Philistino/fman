package fileutils

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/spf13/afero"
)

func TestCopyDir(t *testing.T) {
	fs := afero.NewOsFs()

	// Create a temporary source directory.
	srcDir, err := os.MkdirTemp("", "test-src")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(srcDir)

	// Create some files in the source directory.
	err = fs.Mkdir(filepath.Join(srcDir, "subdir"), 0755)
	if err != nil {
		t.Fatal(err)
	}
	err = afero.WriteFile(fs, filepath.Join(srcDir, "subdir", "file1.txt"), []byte("hello"), 0644)
	if err != nil {
		t.Fatal(err)
	}
	err = afero.WriteFile(fs, filepath.Join(srcDir, "file2.txt"), []byte("world"), 0644)
	if err != nil {
		t.Fatal(err)
	}

	// Create a temporary destination directory.
	destDir, err := os.MkdirTemp("", "test-dest")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(destDir)

	// Copy the source directory to the destination directory.
	err = CopyDir(fs, srcDir, destDir)
	if err != nil {
		t.Fatal(err)
	}

	// Check that the files were copied correctly.
	file1, err := afero.ReadFile(fs, filepath.Join(destDir, "subdir", "file1.txt"))
	if err != nil {
		t.Fatal(err)
	}
	if string(file1) != "hello" {
		t.Errorf("expected file1.txt to contain 'hello', but got '%s'", string(file1))
	}

	file2, err := afero.ReadFile(fs, filepath.Join(destDir, "file2.txt"))
	if err != nil {
		t.Fatal(err)
	}
	if string(file2) != "world" {
		t.Errorf("expected file2.txt to contain 'world', but got '%s'", string(file2))
	}
}

// TestCopyDirErrorSrc tests that CopyDir returns an error if the source directory
// does not exist.
func TestCopyDirErrorSrc(t *testing.T) {
	fs := afero.NewOsFs()

	// Create a temporary destination directory.
	destDir, err := os.MkdirTemp("", "test-dest")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(destDir)

	// Copy the source directory to the destination directory.
	err = CopyDir(fs, "does-not-exist", destDir)
	if err == nil {
		t.Fatal("expected CopyDir to return an error, but it did not")
	}
}
