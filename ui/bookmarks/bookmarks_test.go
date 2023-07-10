package bookmark_ui

import (
	"context"
	"testing"

	"github.com/Philistino/fman/bookmarks"
)

func TestBookmarksStartWithNoBookmarks(t *testing.T) {
	querier, _ := bookmarks.NewQueries(context.Background(), ":memory:")
	defer querier.Close()

	marks := NewBookmarks(querier, 'p', true, 0)
	marks.Init()
	marks.Focus()

	// should not have any bookmarks
	if len(marks.paths) != 0 {
		t.Fatalf("expected 0 paths, got %d", len(marks.paths))
	}
	// add some bookmarks
	marks, _ = marks.Update(BookmarkCmd([]string{"Bingo", "Bango"})())

	// should have 2 bookmarks
	if len(marks.paths) != 2 {
		t.Fatalf("expected 2 paths, got %d", len(marks.paths))
	}
}

func TestBookmarksStartWithBookmarks(t *testing.T) {
	querier, _ := bookmarks.NewQueries(context.Background(), ":memory:")
	defer querier.Close()
	querier.CreateBookmarks(context.Background(), []string{"Bingo", "Bango"})

	marks := NewBookmarks(querier, 'p', true, 0)
	marks.Init()
	marks.Focus()
	// should have 2 bookmarks
	if len(marks.paths) != 2 {
		t.Fatalf("expected 2 paths, got %d", len(marks.paths))
	}

	// add some bookmarks
	marks, _ = marks.Update(BookmarkCmd([]string{"Foo", "Bar"})())

	// should have 4 bookmarks
	if len(marks.paths) != 4 {
		t.Fatalf("expected 4 paths, got %d", len(marks.paths))
	}

	// delete some bookmarks
	marks, _ = marks.Update(UnbookmarkCmd([]string{"Foo", "Bar"})())

	// should have 2 bookmarks
	if len(marks.paths) != 2 {
		t.Fatalf("expected 2 paths, got %d", len(marks.paths))
	}
}

func TestShowHide(t *testing.T) {
	querier, _ := bookmarks.NewQueries(context.Background(), ":memory:")
	defer querier.Close()

	marks := NewBookmarks(querier, 'p', true, 0)
	marks.Init()

	// should be hidden initially
	if !marks.hidden {
		t.Fatalf("expected hidden to be true, got %t", marks.hidden)
	}

	// show
	marks.Show()
	if marks.hidden {
		t.Fatalf("expected hidden to be false, got %t", marks.hidden)
	}

	// hide
	marks.Hide()
	if !marks.hidden {
		t.Fatalf("expected hidden to be true, got %t", marks.hidden)
	}
}

func TestSize(t *testing.T) {
	querier, _ := bookmarks.NewQueries(context.Background(), ":memory:")
	defer querier.Close()

	marks := NewBookmarks(querier, 'p', true, 0)
	marks.Init()

	marks.SetHeight(100)
	marks.SetWidth(10)

	if marks.height != 100 {
		t.Errorf("expected height to be 100, got %d", marks.height)
	}

	if marks.width != 10 {
		t.Errorf("expected width to be 10, got %d", marks.width)
	}
}
