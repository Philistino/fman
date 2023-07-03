//go:build windows

package storage

import "testing"

// a bit of a null test here. Just making sure there is no error.
// It's hard to programmatically test disk space
func TestStorage(t *testing.T) {
	t.Parallel()
	_, err := GetStorageInfo()
	if err != nil {
		t.Error(err)
	}
}
