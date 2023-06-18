package entry

import (
	"testing"

	"github.com/spf13/afero"
)

func TestWithAfero(t *testing.T) {
	appFS := afero.NewMemMapFs()
	appFS.MkdirAll("src/a", 0755)
	afero.WriteFile(appFS, "src/a/b", []byte("file b"), 0644)
	afero.WriteFile(appFS, "src/c", []byte("file c"), 0644)

	// on Windows, isHidden makes syscalls to get the file attributes
	// and returns errors for MemMapFs
	entries, _, err := GetEntries(appFS, "src", true, false)
	if err != nil {
		t.Fatal(err)
	}

	if len(entries) != 2 {
		t.Errorf("expecting %d entries, got %d", 2, len(entries))
	}
	a := entries[0]
	b := entries[1]
	if a.Name() != "a" {
		t.Errorf("expecting %s, got %s", "a", a.Name())
	}
	if b.Name() != "c" {
		t.Errorf("expecting %s, got %s", "c", b.Name())
	}
	if a.SizeStr != "1 entry" {
		t.Errorf("expecting %s, got %s", "1 entry", a.SizeStr)
	}
	if b.SizeStr != "6 B" {
		t.Errorf("expecting %s, got %s", "6 B", b.SizeStr)
	}
}
