package entry

import "testing"

func TestMounts(t *testing.T) {
	t.SkipNow()
	mounts, err := GetMounts()
	if err != nil {
		t.Error(err)
	}
	for _, mount := range mounts {
		t.Log(mount)
	}
	t.Error()
}
