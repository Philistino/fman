package nav

import (
	"context"
	"sort"
	"sync"
	"time"

	"github.com/Philistino/fman/entry"
)

type PreviewHandler struct {
	mu        sync.Mutex
	readDelay int
	maxBytes  int
	cache     map[string]entry.Preview
	cacheSize int
}

func NewPreviewHandler(previewDelay int, maxBytes int, cacheSize int) *PreviewHandler {
	previewer := &PreviewHandler{
		readDelay: previewDelay,
		cache:     make(map[string]entry.Preview),
		maxBytes:  maxBytes,
		cacheSize: cacheSize,
	}
	go previewer.pruneCache(context.Background(), time.Second*5)
	return previewer
}

func (ph *PreviewHandler) GetPreview(ctx context.Context, path string) entry.Preview {
	preview, ok := ph.cache[path]
	if ok {
		preview = entry.CreatePreview(ctx, preview, ph.maxBytes)
		ph.mu.Lock()
		defer ph.mu.Unlock()
		ph.cache[path] = preview
		return preview
	}

	if ph.readDelay >= 0 {
		time.Sleep(time.Millisecond * time.Duration(ph.readDelay))
	}
	preview.Path = path

	// return early if context is cancelled
	if ctx.Err() != nil {
		return preview
	}
	preview = entry.CreatePreview(ctx, preview, ph.maxBytes)
	// return early if context is cancelled
	if ctx.Err() != nil {
		return preview
	}
	ph.mu.Lock()
	defer ph.mu.Unlock()
	ph.cache[path] = preview
	return preview
}

func (ph *PreviewHandler) ClearCache() {
	ph.mu.Lock()
	defer ph.mu.Unlock()
	ph.cache = make(map[string]entry.Preview)
}

func (ph *PreviewHandler) PruneCache() {
	if len(ph.cache) < ph.cacheSize {
		return
	}
	ph.mu.Lock()
	defer ph.mu.Unlock()
	previews := make([]entry.Preview, 0, len(ph.cache))
	for _, preview := range ph.cache {
		previews = append(previews, preview)
	}

	sort.Slice(previews, func(i, j int) bool {
		return previews[i].ReadTime.Before(previews[j].ReadTime)
	})
	for i := 0; i < len(previews)-ph.cacheSize; i++ {
		delete(ph.cache, previews[i].Path)
	}
}

func (ph *PreviewHandler) pruneCache(ctx context.Context, interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			ticker.Stop()
			return
		case <-ticker.C:
			ph.PruneCache()
		}
	}
}
