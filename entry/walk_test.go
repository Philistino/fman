package entry

import (
	"context"
	"fmt"
	"path/filepath"
	"testing"

	"github.com/spf13/afero"
)

func TestWalk(t *testing.T) {
	t.Parallel()
	fsys := afero.NewMemMapFs()
	dirs := []string{string(filepath.Separator), "dir1", "dir2", "dir3", "dir4", "dir5", "dir6", "dir7", "dir8", "dir9"}
	dir := filepath.Join(dirs...)
	fsys.MkdirAll(dir, 0755)
	_, err := fsys.Create(filepath.Join(dir, "file1.txt"))
	if err != nil {
		t.Fatal(err)
	}
	tt := []struct {
		name     string
		start    string
		maxDepth int
		wantLen  int
	}{
		{"Depth -1 from root", dirs[0], -1, 11},
		{"Depth -1 from dir5", filepath.Join(dirs[0:6]...), -1, 6},
		{"Depth 5 from root", dirs[0], 5, 6},
		{"Depth 5 from dir1", filepath.Join(dirs[0:2]...), 5, 6},
		{"Depth 0 from root", dirs[0], 0, 1},
		{"Depth 0 from dir1", filepath.Join(dirs[0:2]...), 0, 1},
		{"Depth 1 from dir1", filepath.Join(dirs[0:2]...), 1, 2},
		{"Depth 1 from root", dirs[0], 1, 2}, // root + dir1
	}
	for _, tc := range tt {
		files, _, err := WalkDown(context.Background(), fsys, tc.start, tc.maxDepth, 10, true)
		if err != nil {
			t.Errorf("Errored: WalkDir(): %s, %q", tc.name, err)
		}
		got := 0
		entries := make([]string, 0, 10)
		for file := range files {
			got++
			entries = append(entries, file.Name())
		}
		if got != tc.wantLen {
			t.Errorf("test fail WalkDir() '%s', wanted: %d, got: %d, entries: %v", tc.name, tc.wantLen, got, entries)
		}
	}
}

func TestWalkChan(t *testing.T) {
	t.Parallel()
	fsys := afero.NewMemMapFs()
	dirs := []string{string(filepath.Separator), "dir1", "dir2", "dir3", "dir4", "dir5", "dir6", "dir7", "dir8", "dir9"}
	for i := 0; i < 10; i++ {
		dirs[1] = fmt.Sprintf("dir%d", i)
		dir := filepath.Join(dirs...)
		fsys.MkdirAll(dir, 0755)
		_, err := fsys.Create(filepath.Join(dir, "file.txt"))
		if err != nil {
			t.Fatal(err)
		}
	}

	files, errCh, err := WalkDown(context.Background(), fsys, string(filepath.Separator), 99, 2, true)
	if err != nil {
		t.Error("TestWalkChan errored", err)
		return
	}

	// test files
	got := 0
	for range files {
		got++
	}
	want := 101 // 100 listed above plus root
	if want != got {
		t.Errorf("TestWalk failed. Wanted %d files, got %d files", want, got)
	}

	// test errors
	got = 0
	for range errCh {
		got++
	}
	want = 0
	if want != got {
		t.Errorf("Wanted %d files, got %d files", want, got)
	}
}

func TestWalkUp(t *testing.T) {
	t.Parallel()
	fsys := afero.NewMemMapFs()
	dirs := []string{string(filepath.Separator), "dir1", "dir2", "dir3", "dir4", "dir5", "dir6", "dir7", "dir8", "dir9"}
	dir := filepath.Join(dirs...)
	fsys.MkdirAll(dir, 0755)
	_, err := fsys.Create(filepath.Join(dir, "file1.txt"))
	if err != nil {
		t.Fatal(err)
	}
	tt := []struct {
		name     string
		start    string
		maxDepth int
		wantLen  int
	}{
		{"Depth 0 from root", dirs[0], 0, 1},
		{"Depth 1 from root", dirs[0], 1, 1},
		{"Depth -1 from root", dirs[0], -1, 1},
		{"Depth 2 from dir5", filepath.Join(dirs[0:6]...), 2, 3},
		{"Depth 1 from dir1", filepath.Join(dirs[0:2]...), 1, 2},
		{"Depth 5 from dir9", filepath.Join(dirs[0:10]...), 5, 6},
		{"Depth 99 from dir9", filepath.Join(dirs[0:10]...), 99, 10},
		{"Depth -1 from dir9", filepath.Join(dirs[0:10]...), -1, 10},
	}
	for _, tc := range tt {
		files, errCh, err := WalkUp(context.Background(), fsys, tc.start, tc.maxDepth, -1, true)
		if err != nil {
			t.Errorf("Errored: WalkDir(): %s, %q", tc.name, err)
		}
		got := 0
		entries := make([]string, 0, 10)
		for file := range files {
			got++
			entries = append(entries, file.Name())
		}
		for err := range errCh {
			t.Errorf("Errored: WalkDir(): %s, %q", tc.name, err)
		}
		if got != tc.wantLen {
			t.Errorf("test fail WalkDir() '%s', wanted: %d, got: %d, entries: %v", tc.name, tc.wantLen, got, entries)
		}
	}
}
