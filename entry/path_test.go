package entry

import "testing"

func TestWindowsInvalidFilename(t *testing.T) {
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
		{`foo \bar.txt`, errInvalidFilenameWindowsSpacePeriod},
		{`foo.\bar.txt`, errInvalidFilenameWindowsSpacePeriod},
		{`foo.d\bar.txt`, nil},
		{`foo.d\bar .txt`, nil},
		{`foo.d\bar. txt`, nil},
		{"", errInvalidFilenameEmpty},
	}

	for _, tc := range cases {
		err := WindowsInvalidFilename(tc.name)
		if err != tc.err {
			t.Errorf("For %q, got %v, expected %v", tc.name, err, tc.err)
		}
	}
}
