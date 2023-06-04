package infobar

import (
	"log"
	"testing"
	"time"
)

func TestDurations(t *testing.T) {
	ti := time.Now()
	minDuration := time.Second
	time.Sleep(time.Millisecond * 500)
	left := minDuration - time.Since(ti)
	log.Println(left)
	t.Error()
}
