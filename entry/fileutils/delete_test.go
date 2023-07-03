package fileutils

import (
	"context"
	"testing"

	"github.com/spf13/afero"
)

func TestRemoveMany(t *testing.T) {
	appFS := afero.NewMemMapFs()
	appFS.MkdirAll("src/a", 0755)
	afero.WriteFile(appFS, "src/a/b", []byte("file b"), 0644)
	afero.WriteFile(appFS, "src/c", []byte("file c"), 0644)

	errs := RemoveMany(context.Background(), appFS, []string{"src/a", "src/c"})
	for _, err := range errs {
		if err != nil {
			t.Errorf("expecting nil, got %q", err)
		}
	}
	_, err := appFS.Stat("src/a/b")
	if err == nil {
		t.Errorf("expecting error, got nil")
	}
	_, err = appFS.Stat("src/a")
	if err == nil {
		t.Errorf("expecting error, got nil")
	}
	_, err = appFS.Stat("src/c")
	if err == nil {
		t.Errorf("expecting error, got nil")
	}
}

func TestRemoveManyCancel(t *testing.T) {
	appFS := afero.NewMemMapFs()
	appFS.MkdirAll("src/a", 0755)
	afero.WriteFile(appFS, "src/a/b", []byte("file b"), 0644)
	afero.WriteFile(appFS, "src/c", []byte("file c"), 0644)

	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	errs := RemoveMany(ctx, appFS, []string{"src/a/b", "src/c"})
	for _, err := range errs {
		if err == nil {
			t.Errorf("expecting error, got nil")
		}
	}
	_, err := appFS.Stat("src/a/b")
	if err != nil {
		t.Errorf("expecting nil, got %q", err)
	}
	_, err = appFS.Stat("src/c")
	if err != nil {
		t.Errorf("expecting nil, got %q", err)
	}
}

func TestRemoveManyError(t *testing.T) {
	appFS := afero.NewMemMapFs()
	appFS.MkdirAll("src/a", 0755)
	afero.WriteFile(appFS, "src/a/b", []byte("file b"), 0644)
	afero.WriteFile(appFS, "src/c", []byte("file c"), 0644)

	errs := RemoveMany(context.Background(), appFS, []string{"src/a/b", "src/c", "src/d"})
	for i, err := range errs {
		if i < 2 && err != nil {
			t.Errorf("expecting nil, got %q", err)
		}
		if i == 2 && err == nil {
			t.Errorf("expecting error, got nil")
		}
	}
	_, err := appFS.Stat("src/a/b")
	if err == nil {
		t.Errorf("expecting error, got nil")
	}
	_, err = appFS.Stat("src/c")
	if err == nil {
		t.Errorf("expecting error, got nil")
	}
}
