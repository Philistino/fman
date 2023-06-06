package entry

import (
	"os"
	"testing"
)

func TestIsZipFile(t *testing.T) {
	path := `fixtures/ziptest (2).zip`
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

func TestGetMimeType(t *testing.T) {
	path := `fixtures/ziptest (2).zip`
	f, err := os.Open(path)
	if err != nil {
		t.Fatalf("os.Open(%s) = %v; want nil", path, err)
	}
	defer f.Close()
	got, err := GetMimeType(f)
	if err != nil {
		t.Errorf("GetMimeType(%s) = %v; want nil", path, err)
	}
	if got != "application/zip" {
		t.Errorf("GetMimeType(%s) = %v; want application/zip", path, got)
	}

	path = `fixtures/text.txt`
	f, err = os.Open(path)
	if err != nil {
		t.Fatalf("os.Open(%s) = %v; want nil", path, err)
	}
	got, err = GetMimeType(f)
	if err != nil {
		t.Errorf("GetMimeType(%s) = %v; want nil", path, err)
	}
	if got != "text/plain; charset=utf-8" {
		t.Errorf("GetMimeType(%s) = %v; want text/plain; charset=utf-8", path, got)
	}
}
