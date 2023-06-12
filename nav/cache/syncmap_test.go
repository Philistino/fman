package cache

import "testing"

func TestSyncMap(t *testing.T) {
	s := NewSyncMap[int, string](-1)

	// Set
	s.Set(1, "foo")
	s.Set(2, "bar")
	s.Set(3, "baz")
	s.Set(4, "qux")

	// Get
	if v, ok := s.Get(1); !ok || v != "foo" {
		t.Errorf("expected %v, got %v", "foo", v)
	}
	if v, ok := s.Get(2); !ok || v != "bar" {
		t.Errorf("expected %v, got %v", "bar", v)
	}
	if v, ok := s.Get(3); !ok || v != "baz" {
		t.Errorf("expected %v, got %v", "baz", v)
	}
	if v, ok := s.Get(4); !ok || v != "qux" {
		t.Errorf("expected %v, got %v", "qux", v)
	}

	// Delete
	s.Delete(1)
	if _, ok := s.Get(1); ok {
		t.Errorf("expected %v, got %v", false, ok)
	}

	// DeleteMany
	s.DeleteMany(2, 3)
	if _, ok := s.Get(2); ok {
		t.Errorf("expected %v, got %v", false, ok)
	}
	if _, ok := s.Get(3); ok {
		t.Errorf("expected %v, got %v", false, ok)
	}

	// Keys
	keys := s.Keys()
	if len(keys) != 1 {
		t.Errorf("expected %v, got %v", 1, len(keys))
	}
	if keys[0] != 4 {
		t.Errorf("expected %v, got %v", 4, keys[0])
	}

	// Values
	values := s.Values()
	if len(values) != 1 {
		t.Errorf("expected %v, got %v", 1, len(values))
	}
	if values[0] != "qux" {
		t.Errorf("expected %v, got %v", "qux", values[0])
	}

	// KeysAndValues
	keys, values = s.KeysAndValues()
	if len(keys) != 1 {
		t.Errorf("expected %v, got %v", 1, len(keys))
	}
	if keys[0] != 4 {
		t.Errorf("expected %v, got %v", 4, keys[0])
	}
	if len(values) != 1 {
		t.Errorf("expected %v, got %v", 1, len(values))
	}
	if values[0] != "qux" {
		t.Errorf("expected %v, got %v", "qux", values[0])
	}

	// Size
	if s.Size() != 1 {
		t.Errorf("expected %v, got %v", 1, s.Size())
	}
}
