//go:build !windows
// +build !windows

package entry

import "testing"

func TestInvalidFilename(t *testing.T) {
	cases := []struct {
		name string
		err  error
	}{
		{`asdf.txt`, nil},
		{`nul.txt`, nil},
		{`nul.jpg.txt`, nil},
		{`some.nul.jpg`, nil},
		{`foo /bar.txt`, errInvalidFilenameReservedCharUnix},
		{`foo.\bar.txt`, nil},
		{"", errInvalidFilenameEmptyUnix},
	}

	for _, tc := range cases {
		err := InvalidFilename(tc.name)
		if err != tc.err {
			t.Errorf("For %q, got %v, expected %v", tc.name, err, tc.err)
		}
	}
}
