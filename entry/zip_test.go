package entry

import (
	"os"
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

func TestGetMimeType(t *testing.T) {

	testCases := []struct {
		path string
		want string
	}{
		{
			path: `fixtures/ziptest.zip`,
			want: "application/zip",
		},
		{
			path: `fixtures/text.txt`,
			want: "text/plain; charset=utf-8",
		},
		// {
		// 	path: `fixtures/ziptest.tar.gz`,
		// 	want: "application/gzip",
		// },
	}
	for _, tc := range testCases {
		t.Run(tc.path, func(t *testing.T) {
			f, err := os.Open(tc.path)
			if err != nil {
				t.Fatalf("os.Open(%s) = %v; want nil", tc.path, err)
			}
			defer f.Close()
			got, err := GetMimeType(f)
			if err != nil {
				t.Errorf("GetMimeType(%s) = %v; want nil", tc.path, err)
			}
			if got != tc.want {
				t.Errorf("GetMimeType(%s) = %v; want %v", tc.path, got, tc.want)
			}
		})
	}
}
