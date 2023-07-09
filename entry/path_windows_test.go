//go:build windows
// +build windows

package entry

import "testing"

func TestInvalidFilename(t *testing.T) {
	cases := []struct {
		name string
		err  error
	}{
		{`asdf.txt`, nil},
		{`nul`, errInvalidFilenameWindowsReservedName},
		{`nul.txt`, errInvalidFilenameWindowsReservedName},
		{`nul.jpg.txt`, errInvalidFilenameWindowsReservedName},
		{`some.nul.jpg`, nil},
		{`foo>bar.txt`, errInvalidFilenameWindowsReservedChar},
		{`foo \bar.txt`, errInvalidFilenameWindowsReservedChar},
		{`foo.\bar.txt`, errInvalidFilenameWindowsReservedChar},
		{"", errInvalidFilenameEmptyWindows},
	}

	for _, tc := range cases {
		err := InvalidFilename(tc.name)
		if err != tc.err {
			t.Errorf("For %q, got %v, expected %v", tc.name, err, tc.err)
		}
	}
}
