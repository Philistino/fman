package nav

import (
	"io/fs"
	"time"

	"github.com/dgraph-io/ristretto"
	"github.com/nore-dev/fman/entry"
)

type NavState struct {
	path     string
	selected []string // TODO: should this be a map[string]struct{} instead?
}

// Path returns the path to the directory of the NavState
func (n NavState) Path() string {
	return n.path
}

// Selected returns the selected items in the directory
func (n NavState) Selected() []string {
	return n.selected
}

type Nav struct {
	cache *ristretto.Cache
	fsys  fs.FS
}

func NewNav(fsys fs.FS) *Nav {
	cache, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: 10_000, // number of keys to track frequency of. Authors recommend 10x number of items
		MaxCost:     1_000,  // going to use cost of 1 for each directory
		BufferItems: 64,     // number of keys per Get buffer. Authors recommend 64
	})
	if err != nil {
		panic(err)
	}
	return &Nav{
		cache: cache,
		fsys:  fsys,
	}
}

func (n *Nav) Set(path int, entry entry.Entry) {
	n.cache.SetWithTTL(path, entry, 1, time.Second*10)
	n.cache.Wait()
}
