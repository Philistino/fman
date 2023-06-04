package breadcrumb

import (
	"strings"
	"testing"
)

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
			desc:         "path is /one/two/three",
			path:         "/one/two/three",
			wantContains: []string{"/", "one", "two", "three"},
			wantLen:      4,
		},
		{
			desc:         "path is C:/one/two/three",
			path:         "C:/one/two/three",
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
	for _, tc := range testcases {
		t.Run(tc.desc, func(t *testing.T) {
			b := New()
			b.SetWidth(1000)
			b.updateView(tc.path)
			if len(b.viewParts) != tc.wantLen {
				t.Errorf("TestUpdateView failed for testcase '%s': want len %d, got %d", tc.desc, tc.wantLen, len(b.viewParts))
			}
			for _, want := range tc.wantContains {
				found := false
				for _, v := range b.viewParts {
					if !strings.Contains(v, want) {
						continue
					}
					found = true
					break
				}
				if !found {
					t.Errorf("TestUpdateView failed for testcase '%s': %s not found", tc.desc, want)
				}
			}
		})
	}
}
