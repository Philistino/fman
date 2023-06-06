package nav

import (
	"testing"
)

func TestHandleCursor(t *testing.T) {
	testcases := []struct {
		name    string
		srcPath string
		dstPath string
		want    string
	}{
		{
			name:    "Parent",
			srcPath: "/a/b",
			dstPath: "/a",
			want:    "b",
		},
		{
			name:    "Windows",
			srcPath: "C:/Users/Jimbo/Documents/GitHub",
			dstPath: "C:/Users/Jimbo",
			want:    "Documents",
		},
		{
			name:    "Unix",
			srcPath: "/a/b/c/d/e",
			dstPath: "/a/b",
			want:    "c",
		},
		{
			name:    "Different root",
			srcPath: "/a/b/c/d/e",
			dstPath: "/d/e/f/g/h",
			want:    "",
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			n := NewNav(true, true, "Bingo")
			n.currentPath = tc.srcPath
			got := n.handleCursor(tc.dstPath)
			if got != tc.want {
				t.Errorf("handleCursor(%s, %s) = %s; want %s", tc.srcPath, tc.dstPath, got, tc.want)
			}
		})
	}
}

func TestHandleCursorWithHistory(t *testing.T) {
	testcases := []struct {
		name    string
		srcPath string
		dstPath string
		want    string
	}{
		{
			name:    "Windows",
			srcPath: "C:/Users/Jimbo",
			dstPath: "C:/Users/Jimbo/Documents/GitHub",
			want:    "Github",
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			n := NewNav(true, true, "Bingo")
			n.cursorHist[tc.dstPath] = tc.want
			n.currentPath = tc.srcPath
			got := n.handleCursor(tc.dstPath)
			if got != tc.want {
				t.Errorf("handleCursor(%s, %s) = %s; want %s", tc.srcPath, tc.dstPath, got, tc.want)
			}
		})
	}
}
