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
	if a.SizeStr != "1 item" {
		t.Errorf("expecting %s, got %s", "1 entry", a.SizeStr)
	}
	if b.SizeStr != "6 B" {
		t.Errorf("expecting %s, got %s", "6 B", b.SizeStr)
	}
}

// func TestModTimeChange(t *testing.T) {
// 	file, err := os.CreateTemp("", "test")
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	defer os.Remove(file.Name())
// 	info, err := os.Stat(file.Name())
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	preMod := info.ModTime()
// 	log.Println(preMod)

// 	_, err = file.WriteString("test")
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	file.Close()
// 	time.Sleep(1000 * time.Millisecond)

// 	otherInfo, err := os.Stat(file.Name())
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	log.Println(otherInfo.ModTime(), preMod)
// 	if !otherInfo.ModTime().After(preMod) {
// 		t.Errorf("expecting %s, got %s", "modified", "not modified")
// 	}
// 	t.Error()
// }
