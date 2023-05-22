package entry

import (
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
			got, _ := HighlightSyntax(tC.name, tC.preview)
			if got != tC.expected {
				t.Errorf("expecting %s, got %v", tC.expected, got)
			}
		})
	}
}

// reimplement with mock
// func TestGetEntries(t *testing.T) {

// 	testCases := []struct {
// 		desc         string
// 		path         string
// 		expectedSize int
// 	}{
// 		{
// 			desc:         "cur dir",
// 			path:         "./",
// 			expectedSize: 4,
// 		},
// 	}

// 	for _, tC := range testCases {
// 		t.Run(tC.desc, func(t *testing.T) {
// 			entries, _ := GetEntries(tC.path, true, false)
// 			if len(entries) != tC.expectedSize {
// 				t.Errorf("expecting %d entries, got %d", tC.expectedSize, len(entries))
// 			}
// 		})
// 	}
// }

// func TestSortEntries(t *testing.T) {
// 	tt := []struct {
// 		name      string
// 		inEntries []Entry
// 		want      []Entry
// 	}{
// 		{
// 			name: "first",
// 			inEntries: []Entry{
// 				{Name: "1 file", IsDir: false},
// 				{Name: "2 file", IsDir: false},
// 				{Name: "1 dir", IsDir: true},
// 				{Name: "2 dir", IsDir: true},
// 			},
// 			want: []Entry{
// 				{Name: "1 dir", IsDir: true},
// 				{Name: "2 dir", IsDir: true},
// 				{Name: "1 file", IsDir: false},
// 				{Name: "2 file", IsDir: false},
// 			},
// 		},
// 	}
// 	for _, tc := range tt {
// 		t.Run(tc.name, func(t *testing.T) {
// 			entries := sortEntries(tc.inEntries)
// 			for _, e := range entries {
// 				log.Println(e)
// 			}
// 			if reflect.DeepEqual(tc.want, entries) {
// 				t.Errorf("expecting %v entries, got %v", entries, len(entries))
// 			}
// 		})
// 		t.Error()
// 	}

// }
