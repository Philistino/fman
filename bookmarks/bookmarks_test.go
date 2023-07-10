package bookmarks

import (
	"context"
	"testing"
)

func TestBookMarks(t *testing.T) {
	ctx := context.Background()
	q, err := NewQueries(ctx, ":memory:")
	if err != nil {
		t.Fatal(err)
	}
	defer q.Close()

	err = q.CreateBookmarks(ctx, []string{"Bingo", "Bango"})
	if err != nil {
		t.Fatal(err)
	}

	// test duplicate
	err = q.CreateBookmark(ctx, "Bingo")
	if err != nil {
		t.Fatal(err)
	}

	paths, err := q.GetBookmarks(ctx)
	if err != nil {
		t.Fatal(err)
	}

	if len(paths) != 2 {
		t.Fatalf("expected 2 paths, got %d", len(paths))
	}

	if paths[0] != "Bingo" {
		t.Fatalf("expected Bingo, got %s", paths[0])
	}

	if paths[1] != "Bango" {
		t.Fatalf("expected Bango, got %s", paths[1])
	}

	err = q.DeleteBookMarks(ctx, []string{"Bingo", "DNE"})
	if err != nil {
		t.Fatal(err)
	}

	paths, err = q.GetBookmarks(ctx)
	if err != nil {
		t.Fatal(err)
	}

	if len(paths) != 1 {
		t.Fatalf("expected 1 path, got %d", len(paths))
	}

	if paths[0] != "Bango" {
		t.Fatalf("expected Bango, got %s", paths[0])
	}
}

func TestConnectionError(t *testing.T) {
	ctx := context.Background()
	_, err := NewQueries(ctx, "/tmp/does/not/exist38383838383")
	if err == nil {
		t.Fatal("expected error")
	}
}
