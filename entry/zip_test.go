package entry

import (
	"testing"
)

func TestIsZipFile(t *testing.T) {
	path := `fixtures/ziptest.zip`
	got, err := IsZipFile(path)
	if err != nil {
		t.Errorf("IsZipFile(%s) = %v; want nil", path, err)
	}
	if !got {
		t.Errorf("IsZipFile(%s) = %v; want true", path, got)
	}

	path = `fixtures/text.txt`
	got, err = IsZipFile(path)
	if err != nil {
		t.Errorf("IsZipFile(%s) = %v; want nil", path, err)
	}
	if got {
		t.Errorf("IsZipFile(%s) = %v; want true", path, got)
	}
}
