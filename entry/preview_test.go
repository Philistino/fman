package entry

import (
	"context"
	"io"
	"os"
	"strings"
	"testing"
)

func TestHighlightSyntax(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		desc     string
		name     string
		preview  string
		expected string
	}{
		{
			desc:     "empty",
			name:     "",
			preview:  "",
			expected: "",
		},
		{
			desc:     "go",
			name:     "go",
			preview:  "package main\n\nfunc main()\n{\n}\n",
			expected: "\x1b[1m\x1b[37mpackage main\x1b[0m\x1b[1m\x1b[37m\n\x1b[0m\x1b[1m\x1b[37m\n\x1b[0m\x1b[1m\x1b[37mfunc main()\x1b[0m\x1b[1m\x1b[37m\n\x1b[0m\x1b[1m\x1b[37m{\x1b[0m\x1b[1m\x1b[37m\n\x1b[0m\x1b[1m\x1b[37m}\x1b[0m\x1b[1m\x1b[37m\n\x1b[0m",
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got, _ := highlightSyntax(tC.name, tC.preview)
			if got != tC.expected {
				t.Errorf("expecting %s, got %v", tC.expected, got)
			}
		})
	}
}

func TestReadBytes(t *testing.T) {
	testcases := []struct {
		name     string
		reader   io.Reader
		maxBytes int
		wantStr  string
	}{
		{
			name:     "read all",
			reader:   strings.NewReader("1234567890"),
			maxBytes: 1_000_000,
			wantStr:  "1234567890",
		},
		{
			name:     "read 10 bytes",
			reader:   strings.NewReader("1234567890"),
			maxBytes: 10,
			wantStr:  "1234567890",
		},
		{
			name:     "read 5 bytes",
			reader:   strings.NewReader("1234567890"),
			maxBytes: 5,
			wantStr:  "12345",
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			gotStr, err := readNBytes(context.Background(), tc.reader, tc.maxBytes)
			if err != nil {
				t.Errorf("readNBytes() error = %v", err)
			}
			if gotStr != tc.wantStr {
				t.Errorf("readNBytes() got string = %v, want %v", gotStr, tc.wantStr)
			}
		})
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
			got, err := GetMimeTypeByRead(f)
			if err != nil {
				t.Errorf("GetMimeType(%s) = %v; want nil", tc.path, err)
			}
			if got != tc.want {
				t.Errorf("GetMimeType(%s) = %v; want %v", tc.path, got, tc.want)
			}
		})
	}
}
