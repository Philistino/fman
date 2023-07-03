package breadcrumb

import (
	"errors"
	"path/filepath"
	"strings"
	"testing"

	zone "github.com/lrstanley/bubblezone"
)

// TODO: test truncation

type pathError struct {
	path string
	err  error
}

func (p pathError) Path() string {
	return p.path
}

func (p pathError) Error() error {
	return p.err
}

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
			b := NewBreadCrumb()
			b.Init()
			b.SetWidth(1000)
			b.Update(pathError{path: tc.path, err: nil})
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

// TestError tests that a msg with an error is ignored
func TestError(t *testing.T) {
	zone.NewGlobal()
	b := NewBreadCrumb()
	b.SetWidth(1000)
	want := "start"
	ignore := "ignore"
	b.Update(pathError{path: want, err: nil})
	if !strings.Contains(b.View(), want) {
		t.Errorf("TestError failed: '%s' not found in View()", want)
	}
	b.Update(pathError{path: ignore, err: errors.New("Fake error")})
	if strings.Contains(b.View(), ignore) {
		t.Errorf("TestError failed: '%s' found in View()", ignore)
	}
	if !strings.Contains(b.View(), want) {
		t.Errorf("TestError failed: '%s' not found in View()", want)
	}
}

func TestUpdateViewTruncate(t *testing.T) {
	testcases := []struct {
		desc         string
		path         string
		width        int
		wantContains []string
		wantLen      int
	}{
		{
			desc:         "path is /one/two/three",
			path:         filepath.Join(pathSeparator, "one", "two", "three"),
			width:        30,
			wantContains: []string{"two", "three"},
			wantLen:      2,
		},
		// {
		// 	desc:         "path is C:/one/two/three",
		// 	path:         filepath.Join("C:", pathSeparator, "one", "two", "three"),
		// 	wantContains: []string{"C:", "one", "two", "three"},
		// 	wantLen:      4,
		// },
	}
	zone.NewGlobal()
	defer zone.Close()
	for _, tc := range testcases {
		t.Run(tc.desc, func(t *testing.T) {
			b := NewBreadCrumb()
			b.SetWidth(tc.width)
			b.updateView(tc.path)
			if len(b.viewParts) != tc.wantLen {
				t.Errorf("TestUpdateViewTruncate failed for testcase '%s': want len %d, got %d, %q", tc.desc, tc.wantLen, len(b.viewParts), b.viewParts)
			}
			for i, want := range tc.wantContains {
				if !strings.Contains(b.viewParts[i], want) {
					t.Errorf("TestUpdateViewTruncate failed for testcase '%s': '%s' not found", tc.desc, want)
				}
				if !strings.Contains(b.View(), want) {
					t.Errorf("TestUpdateViewTruncate failed for testcase '%s': '%s' not found in View()", tc.desc, want)
				}
			}
		})
	}
}
