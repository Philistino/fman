//go:build windows

package entry

import (
	"os"
	"syscall"
	"testing"
)

// hide hides a
// https://github.com/tobychui/goHidden/blob/main/hide.go
func hide(filename string) error {
	filenameW, err := syscall.UTF16PtrFromString(filename)
	if err != nil {
		return err
	}
	err = syscall.SetFileAttributes(filenameW, syscall.FILE_ATTRIBUTE_HIDDEN)
	if err != nil {
		return err
	}
	return nil
}

// This test is a little bit self-fullfiling.
// I manually checked that the hide function works on Windows 11.
// Hopefully Windows does not change its behavior in the future.
func TestIsHiddenWindows(t *testing.T) {
	t.Parallel()
	file, err := os.CreateTemp(os.TempDir(), "test.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(file.Name())

	info, err := file.Stat()
	if err != nil {
		t.Fatal(err)
	}
	if isHidden(info, os.TempDir(), nil) {
		t.Error("file should not be hidden")
	}

	err = hide(file.Name())
	if err != nil {
		t.Fatal(err)
	}

	if !isHidden(info, os.TempDir(), nil) {
		t.Error("file should be hidden")
	}
}
