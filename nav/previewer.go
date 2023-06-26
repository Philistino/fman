package nav

import (
	"context"
	"time"

	"github.com/Philistino/fman/entry"
	"github.com/Philistino/fman/nav/cache"
	"github.com/spf13/afero"
)

type PreviewHandler struct {
	readDelay int
	maxBytes  int
	cache     *cache.Cache[string, entry.Preview]
}

// NewPreviewHandler creates a new PreviewHandler
// ctx is used to cancel periodic pruning of the cache
// previewDelay is the delay in milliseconds to wait before reading a file
// maxBytes is the maximum number of bytes to read from a file to form the preview
// cacheSize is the maximum number of previews to store in the cache
// pruneInterval is the interval at which to prune the cache
func NewPreviewHandler(ctx context.Context, previewDelay int, maxBytes int, cacheSize int, pruneInterval time.Duration) *PreviewHandler {
	prevCache, _ := cache.NewCache[string, entry.Preview](
		ctx,
		cacheSize,
		pruneInterval,
		nil,
		func(i entry.Preview, j entry.Preview) bool {
			return i.ReadTime.Before(j.ReadTime)
		},
	)
	return &PreviewHandler{
		readDelay: previewDelay,
		maxBytes:  maxBytes,
		cache:     prevCache,
	}
}

func (ph *PreviewHandler) GetPreview(ctx context.Context, fsys afero.Fs, path string) entry.Preview {
	preview, ok := ph.cache.Get(path)
	if ok {
		preview = entry.CreatePreview(ctx, fsys, preview, ph.maxBytes)
		ph.cache.Set(path, preview)
		return preview
	}

	if ph.readDelay >= 0 {
		time.Sleep(time.Millisecond * time.Duration(ph.readDelay))
	}

	// return early if context is cancelled
	if ctx.Err() != nil {
		return preview
	}

	preview.Path = path
	preview = entry.CreatePreview(ctx, fsys, preview, ph.maxBytes)
	ph.cache.Set(path, preview)
	return preview
}
