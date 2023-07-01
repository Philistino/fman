package model

import (
	"path/filepath"
	"strings"
	"testing"

	zone "github.com/lrstanley/bubblezone"
)

// TODO: test truncation

func TestUpdateView(t *testing.T) {
	testcases := []struct {
		desc         string
		path         string
		wantContains []string
		wantLen      int
	}{
		{
			desc:         "path is empty",
			path:         "",
			wantContains: []string{},
			wantLen:      0,
		},
		{
			desc:         "path is /",
			path:         "/",
			wantContains: []string{"/"},
			wantLen:      1,
		},
		{
			desc:         "path is separator",
			path:         pathSeparator,
			wantContains: []string{pathSeparator},
			wantLen:      1,
		},
		{
			desc:         "path is /one/two/three",
			path:         filepath.Join(pathSeparator, "one", "two", "three"),
			wantContains: []string{pathSeparator, "one", "two", "three"},
			wantLen:      4,
		},
		{
			desc:         "path is C:/one/two/three",
			path:         filepath.Join("C:", pathSeparator, "one", "two", "three"),
			wantContains: []string{"C:", "one", "two", "three"},
			wantLen:      4,
		},
		{
			desc:         "path is C:/",
			path:         "C:/",
			wantContains: []string{"C:"},
			wantLen:      1,
		},
		{
			desc:         "path is C:",
			path:         "C:",
			wantContains: []string{"C:"},
			wantLen:      1,
		},
	}
	zone.NewGlobal()
	defer zone.Close()
	for _, tc := range testcases {
		t.Run(tc.desc, func(t *testing.T) {
			b := newBrdCrumb()
			b.SetWidth(1000)
			b.updateView(tc.path)
			if len(b.viewParts) != tc.wantLen {
				t.Errorf("TestUpdateView failed for testcase '%s': want len %d, got %d, %q", tc.desc, tc.wantLen, len(b.viewParts), b.viewParts)
			}
			for i, want := range tc.wantContains {
				if !strings.Contains(b.viewParts[i], want) {
					t.Errorf("TestUpdateView failed for testcase '%s': '%s' not found", tc.desc, want)
				}
				if !strings.Contains(b.View(), want) {
					t.Errorf("TestUpdateView failed for testcase '%s': '%s' not found in View()", tc.desc, want)
				}
			}
		})
	}
}
