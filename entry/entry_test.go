package entry

import (
	"log"
	"os"
	"path/filepath"
	"testing"
)

// cleanRoot returns a relative path from root to dir.
func cleanRoot(dir string) (string, error) {
	// if reading at root AKA ".", no cleaning required
	if dir == "." {
		return dir, nil
	}
	root := "/"
	// on unix, VolumeName will return ""
	if vol := filepath.VolumeName(dir); vol != "" {
		// on windows, set root to the volume name and slash
		root = vol + "/"
	}
	rel, err := filepath.Rel(root, dir)
	if err != nil {
		return "", err
	}
	rel = filepath.ToSlash(rel)
	return rel, nil
}

// Exists returns whether the given file or directory exists
func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// reimplement with mock
func TestGetEntries(t *testing.T) {

	path, err := filepath.Abs(".")
	if err != nil {
		t.Fatal(err)
	}

	// path = `C:\Users\tyman\OneDrive\Desktop\explain`

	path, err = cleanRoot(path)
	if err != nil {
		t.Fatal(err)
	}
	path = "c:/" + path
	log.Println(path)

	fsys := os.DirFS("/")
	entries, errMap, err := GetEntries(fsys, path, true, false)
	for key, value := range errMap {
		log.Println(key, value)
	}
	if err != nil {
		t.Fatal(err)
	}
	for _, e := range entries {
		log.Println(e.Name(), e.IsDir())
	}
	if len(entries) != 18 {
		t.Errorf("expecting %d entries, got %d", 18, len(entries))
	}
	t.Error()
}

// func TestSortEntries(t *testing.T) {
// 	tt := []struct {
// 		name      string
// 		inEntries []Entry
// 		want      []Entry
// 	}{
// 		{
// 			name: "first",
// 			inEntries: []Entry{
// 				{Name: "1 file", IsDir: false},
// 				{Name: "2 file", IsDir: false},
// 				{Name: "1 dir", IsDir: true},
// 				{Name: "2 dir", IsDir: true},
// 			},
// 			want: []Entry{
// 				{Name: "1 dir", IsDir: true},
// 				{Name: "2 dir", IsDir: true},
// 				{Name: "1 file", IsDir: false},
// 				{Name: "2 file", IsDir: false},
// 			},
// 		},
// 	}
// 	for _, tc := range tt {
// 		t.Run(tc.name, func(t *testing.T) {
// 			entries := sortEntries(tc.inEntries)
// 			for _, e := range entries {
// 				log.Println(e)
// 			}
// 			if reflect.DeepEqual(tc.want, entries) {
// 				t.Errorf("expecting %v entries, got %v", entries, len(entries))
// 			}
// 		})
// 		t.Error()
// 	}

// }
