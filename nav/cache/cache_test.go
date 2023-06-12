package cache

import (
	"context"
	"testing"
	"time"
)

func TestCachePrune(t *testing.T) {
	cache, _ := NewCache[string, int](
		context.Background(),
		2,
		time.Hour*5,
		nil,
		func(i int, j int) bool {
			return i < j
		},
	)
	cache.Set("a", 1)
	cache.pruneCache()
	cache.Set("b", 2)
	cache.Set("c", 3)
	cache.pruneCache()

	if cache.Size() != 2 {
		t.Errorf("Expected cache size of 2, got %d", cache.Size())
	}
	_, ok := cache.Get("a")
	if ok {
		t.Errorf("Expected cache to not contain key a")
	}
	_, ok = cache.Get("b")
	if !ok {
		t.Errorf("Expected cache to contain key b")
	}
	_, ok = cache.Get("c")
	if !ok {
		t.Errorf("Expected cache to contain key c")
	}
}

func TestCacheKeyLessFn(t *testing.T) {
	cache, _ := NewCache[int, string](
		context.Background(),
		2,
		time.Hour*5,
		func(i int, j int) bool {
			return i < j
		},
		nil,
	)
	cache.Set(3, "foo")
	cache.Set(2, "foo")
	cache.Set(1, "foo")
	cache.pruneCache()

	if cache.Size() != 2 {
		t.Errorf("Expected cache size of 2, got %d", cache.Size())
	}
	_, ok := cache.Get(1)
	if ok {
		t.Errorf("Expected cache to not contain key a")
	}
	_, ok = cache.Get(2)
	if !ok {
		t.Errorf("Expected cache to contain key b")
	}
	_, ok = cache.Get(3)
	if !ok {
		t.Errorf("Expected cache to contain key c")
	}
}

func TestCacheWaitForPrune(t *testing.T) {
	cache, _ := NewCache[string, int](
		context.Background(),
		2,
		time.Millisecond*5,
		nil,
		func(i int, j int) bool {
			return i < j
		},
	)
	cache.Set("a", 1)
	cache.Set("b", 2)
	cache.Set("c", 3)
	time.Sleep(time.Millisecond * 20)

	if cache.Size() != 2 {
		t.Errorf("Expected cache size of 2, got %d", cache.Size())
	}
	_, ok := cache.Get("a")
	if ok {
		t.Errorf("Expected cache to not contain key a")
	}
	_, ok = cache.Get("b")
	if !ok {
		t.Errorf("Expected cache to contain key b")
	}
	_, ok = cache.Get("c")
	if !ok {
		t.Errorf("Expected cache to contain key c")
	}
}

func TestCacheCancelContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cache, _ := NewCache[string, int](
		ctx,
		2,
		time.Millisecond*5,
		nil,
		func(i int, j int) bool {
			return i < j
		},
	)
	cancel()
	cache.Set("a", 1)
	cache.Set("b", 2)
	cache.Set("c", 3)
	time.Sleep(time.Millisecond * 20)

	if cache.Size() != 3 {
		t.Errorf("Expected cache size of 3, got %d", cache.Size())
	}
	_, ok := cache.Get("a")
	if !ok {
		t.Errorf("Expected cache to contain key a")
	}
	_, ok = cache.Get("b")
	if !ok {
		t.Errorf("Expected cache to contain key b")
	}
	_, ok = cache.Get("c")
	if !ok {
		t.Errorf("Expected cache to contain key c")
	}
}

func TestCacheDelete(t *testing.T) {
	cache, _ := NewCache[string, int](
		context.Background(),
		-1,
		time.Hour*5,
		nil,
		func(i int, j int) bool {
			return i < j
		},
	)
	cache.Set("a", 1)
	cache.Set("b", 2)
	cache.Set("c", 3)

	cache.Delete("c")

	if cache.Size() != 2 {
		t.Errorf("Expected cache size of 2, got %d", cache.Size())
	}
	_, ok := cache.Get("c")
	if ok {
		t.Errorf("Expected cache to not contain key c")
	}
}

func TestCacheNewTwoLess(t *testing.T) {
	less := func(i int, j int) bool {
		return i < j
	}

	_, err := NewCache[int, int](
		context.Background(),
		2,
		time.Millisecond*5,
		less,
		less,
	)

	if err == nil {
		t.Errorf("Expected err to be nil")
	}
}

func TestCacheNewNoLess(t *testing.T) {
	_, err := NewCache[int, int](
		context.Background(),
		2,
		time.Millisecond*5,
		nil,
		nil,
	)

	if err == nil {
		t.Errorf("Expected err to be nil")
	}
}
