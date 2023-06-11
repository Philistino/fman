package entryinfo

import (
	"context"
	"io"
	"strings"
	"testing"
)

func TestReadLines(t *testing.T) {
	str := "1\n2\n3\n4\n5\n6\n7\n8\n9\n10\n11\n12\n13\n14\n15\n16\n17\n18\n19\n20"
	testcases := []struct {
		name     string
		reader   io.Reader
		nLines   int
		maxBytes int
		wantStr  string
		wantEOF  bool
	}{
		{
			name:     "read all",
			reader:   strings.NewReader(str),
			nLines:   100,
			maxBytes: 1_000_000,
			wantStr:  str,
			wantEOF:  true,
		},
		{
			name:     "read 10 lines",
			reader:   strings.NewReader(str),
			nLines:   10,
			maxBytes: 1_000_000,
			wantStr:  "1\n2\n3\n4\n5\n6\n7\n8\n9\n10",
			wantEOF:  false,
		},
		{
			name:     "read 10 lines hit maxBytes",
			reader:   strings.NewReader(str),
			nLines:   10,
			maxBytes: 10,
			wantStr:  "1\n2\n3\n4\n5",
			wantEOF:  false,
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			gotStr, gotEOF, err := readLines(context.Background(), tc.reader, tc.nLines, tc.maxBytes)
			if err != nil {
				t.Errorf("readLines() error = %v", err)
				return
			}
			if gotStr != tc.wantStr {
				t.Errorf("readLines() got string = %v, want %v", gotStr, tc.wantStr)
			}
			if gotEOF != tc.wantEOF {
				t.Errorf("readLines() got EOF = %v, want %v", gotEOF, tc.wantEOF)
			}
		})
	}
}
