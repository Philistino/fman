package cache

import (
	"context"
	"fmt"
	"sort"
	"sync"
	"time"
)

// Cache is a bounded cache that prunes itself periodically.
// The cache is not bounded to the given size, it is pruned down to the cacheSize at a given interval.
// During a pruning operation, other methods block so this implementation
// is not optimized for really large cache sizes, slow less functions, or constant write/reads.
type Cache[K comparable, V any] struct {
	mu            sync.Mutex
	cache         *SyncMap[K, V]
	cacheSize     int
	pruneInterval time.Duration
	lessV         func(i, j V) bool
	lessK         func(i, j K) bool
}

// NewCache creates a new bounded cache with the given size and prune interval.
// The cache is not guaranteed to be bounded to the given size.
// The cache may grow to be larger than the given size and it will be pruned down to the given size at the next pruneInterval.
//
// The ctx is used to stop the prune goroutine when the context is canceled.
// The cacheSize is the number of entries to keep in the cache on a prune operation. If the cacheSize is <1, the cache will not be pruned.
// The pruneInterval is the time between prune operations.
// The lessKeys and lessValues functions are used to sort the keys and values respectively to determine which keys to remove on a prune operation.
// If lessKeys is nil, lessValues is used to sort the keys. lessKeys is prioritized over lessValues.
// If both lessKeys and lessValues are nil, an error is returned.
// If lessKeys and lessValues are both non-nil, an error is returned.
//
// Because values may be repeated within the map, we can get inconsistent pruning if we sort by values that are non-unique.
// For instance, if there are multiple values with the same sort order, we cannot know which keys will be removed on a prune operation.
func NewCache[K comparable, V any](
	ctx context.Context,
	cacheSize int,
	pruneInterval time.Duration,
	lessKeys func(i, j K) bool,
	lessValues func(i, j V) bool,
) (*Cache[K, V], error) {

	if lessValues == nil && lessKeys == nil {
		return nil, fmt.Errorf("lessKeys and lessValues functions cannot both be nil")
	}
	if lessValues != nil && lessKeys != nil {
		return nil, fmt.Errorf("lessKeys and lessValues functions cannot both be nil")
	}

	cache := &Cache[K, V]{
		cache:         NewSyncMap[K, V](0),
		cacheSize:     cacheSize,
		pruneInterval: pruneInterval,
		lessV:         lessValues,
		lessK:         lessKeys,
	}
	if cacheSize > 0 {
		go cache.pruneCachePeriodic(ctx, pruneInterval)
	}
	return cache, nil
}

// Set sets the value associated with the given key.
func (bc *Cache[K, V]) Set(key K, value V) {
	bc.mu.Lock()
	defer bc.mu.Unlock()
	bc.cache.Set(key, value)
}

// Get returns the value associated with the given key and whether or not the key exists in the cache.
func (bc *Cache[K, V]) Get(key K) (V, bool) {
	bc.mu.Lock()
	defer bc.mu.Unlock()
	v, ok := bc.cache.Get(key)
	return v, ok
}

// Delete deletes the value associated with the given key.
func (bc *Cache[K, V]) Delete(key K) {
	bc.mu.Lock()
	defer bc.mu.Unlock()
	bc.cache.Delete(key)
}

// Size returns the current number of elements in the cache.
func (pc *Cache[K, V]) Size() int {
	pc.mu.Lock()
	defer pc.mu.Unlock()
	return pc.cache.Size()
}

// pruneCache removes entries from the cache until
// the cache size is equal to the cacheSize.
// The entries are removed by sorting the keys or values
// and removing the first n entries where n is the
// number of entries over the cacheSize.
func (bc *Cache[K, V]) pruneCache() {
	bc.mu.Lock()
	defer bc.mu.Unlock()
	if bc.cache.Size() <= bc.cacheSize {
		return
	}

	keys, values := bc.cache.KeysAndValues()

	// Sort keys by lessK if it exists, otherwise sort by lessV
	sort.Slice(keys, func(i, j int) bool {
		if bc.lessK != nil {
			return bc.lessK(keys[i], keys[j])
		}
		return bc.lessV(values[i], values[j])
	})
	bc.cache.DeleteMany(keys[:len(keys)-bc.cacheSize]...)
}

// pruneCachePeriodic calls pruneCache at the given interval until the context is canceled.
func (pc *Cache[K, V]) pruneCachePeriodic(ctx context.Context, interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			pc.pruneCache()
		}
	}
}
