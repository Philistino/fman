package entry

import (
	"context"
	"io"
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
			desc:    "go",
			name:    "go",
			preview: "package main\n\nfunc main()\n{\n}\n",
			expected: `[1m[37mpackage main[0m[1m[37m
[0m[1m[37m
[0m[1m[37mfunc main()[0m[1m[37m
[0m[1m[37m{[0m[1m[37m
[0m[1m[37m}[0m[1m[37m
[0m`,
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

func TestReadLines2(t *testing.T) {
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
				t.Errorf("readLines2() error = %v", err)
			}
			if gotStr != tc.wantStr {
				t.Errorf("readLines2() got string = %v, want %v", gotStr, tc.wantStr)
			}
		})
	}
}
