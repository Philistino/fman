package nav

import (
	"log"
	"path/filepath"
	"testing"
)

func TestIsRoot(t *testing.T) {
	path := "C:/Users/daniel/Documents/GitHub/go-nav/test/test.txt"
	for i := 0; i < 10; i++ {
		path = filepath.Dir(path)
		log.Println(path)
	}
	t.Error(path)
}
