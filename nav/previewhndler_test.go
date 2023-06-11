package nav

import (
	"testing"
	"time"

	"github.com/Philistino/fman/entry"
)

func TestPruneCache(t *testing.T) {
	previewer := NewPreviewHandler(0, 100, 2)
	previewer.cache["a"] = entry.Preview{Path: "a", ReadTime: time.Now()}
	previewer.cache["b"] = entry.Preview{Path: "b", ReadTime: time.Now()}
	previewer.PruneCache()
	previewer.cache["c"] = entry.Preview{Path: "c", ReadTime: time.Now()}
	previewer.PruneCache()

	if len(previewer.cache) != 2 {
		t.Errorf("Expected cache size of 2, got %d", len(previewer.cache))
	}
	_, ok := previewer.cache["a"]
	if ok {
		t.Errorf("Expected cache to not contain key a")
	}
	_, ok = previewer.cache["b"]
	if !ok {
		t.Errorf("Expected cache to contain key b")
	}
	_, ok = previewer.cache["c"]
	if !ok {
		t.Errorf("Expected cache to contain key c")
	}

}
