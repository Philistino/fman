package cache

import "sync"

type SyncMap[K comparable, V any] struct {
	mu    sync.Mutex
	cache map[K]V
}

func NewSyncMap[K comparable, V any](size int) *SyncMap[K, V] {
	if size < 0 {
		size = 0
	}
	return &SyncMap[K, V]{
		cache: make(map[K]V, size),
	}
}

// Set sets the value associated with the given key.
func (s *SyncMap[K, V]) Set(key K, value V) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.cache[key] = value
}

// Get returns the value associated with the given key and whether or not the key exists in the cache.
func (s *SyncMap[K, V]) Get(key K) (V, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	v, ok := s.cache[key]
	return v, ok
}

// Delete deletes the value associated with the given key.
func (s *SyncMap[K, V]) Delete(key K) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.cache, key)
}

// Delete deletes the value associated with the given key.
func (s *SyncMap[K, V]) DeleteMany(keys ...K) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, key := range keys {
		delete(s.cache, key)
	}
}

// Keys returns a slice of all the keys in the cache.
// Remember that the order of the values is not guaranteed. If you need to
// maintain the order, use KeysAndValues instead.
func (s *SyncMap[K, V]) Keys() []K {
	s.mu.Lock()
	defer s.mu.Unlock()
	keys := make([]K, 0, len(s.cache))
	for key := range s.cache {
		keys = append(keys, key)
	}
	return keys
}

// Values returns a slice of all the values in the map.
// Remember that the order of the values is not guaranteed. If you need to
// maintain the order, use KeysAndValues instead.
func (s *SyncMap[K, V]) Values() []V {
	s.mu.Lock()
	defer s.mu.Unlock()
	values := make([]V, 0, len(s.cache))
	for _, value := range s.cache {
		values = append(values, value)
	}
	return values
}

// KeysAndValues returns slices of all the keys and values in the map.
func (s *SyncMap[K, V]) KeysAndValues() ([]K, []V) {
	s.mu.Lock()
	defer s.mu.Unlock()
	keys := make([]K, 0, len(s.cache))
	values := make([]V, 0, len(s.cache))
	for key, value := range s.cache {
		keys = append(keys, key)
		values = append(values, value)
	}
	return keys, values
}

// Size returns the current number of elements in the map.
func (s *SyncMap[K, V]) Size() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return len(s.cache)
}
