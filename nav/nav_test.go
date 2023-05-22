package nav

import (
	"testing"

	"github.com/nore-dev/fman/entry"
)

func BenchmarkWait(b *testing.B) {
	n := NewNav(nil)
	entry := entry.Entry{}
	for i := 0; i < b.N; i++ {
		n.Set(i, entry)
	}
}
