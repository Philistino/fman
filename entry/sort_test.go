package entry

import (
	"os"
	"reflect"
	"testing"
	"time"
)

type fakeFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
	isDir   bool
}

func (f fakeFileInfo) Name() string       { return f.name }
func (f fakeFileInfo) Size() int64        { return f.size }
func (f fakeFileInfo) Mode() os.FileMode  { return f.mode }
func (f fakeFileInfo) ModTime() time.Time { return f.modTime }
func (f fakeFileInfo) IsDir() bool        { return f.isDir }
func (f fakeFileInfo) Sys() interface{}   { return nil }

func TestSort(t *testing.T) {
	testcases := []struct {
		name      string
		sortT     SortOrder
		entries   []Entry
		wantNames []string
	}{
		{
			name: "naturalSort",
			sortT: SortOrder{
				method:     NaturalSort,
				dirsFirst:  false,
				dirsOnly:   false,
				showHidden: true,
				reverse:    false,
				ignoreDiac: true,
			},
			entries: []Entry{
				{FileInfo: fakeFileInfo{name: "10b"}},
				{FileInfo: fakeFileInfo{name: "2b"}},
				{FileInfo: fakeFileInfo{name: "a"}},
				{FileInfo: fakeFileInfo{name: "c"}},
			},
			wantNames: []string{"2b", "10b", "a", "c"},
		},
		{
			name: "name sort",
			sortT: SortOrder{
				method:     NameSort,
				dirsFirst:  false,
				dirsOnly:   false,
				showHidden: true,
				reverse:    false,
				ignoreDiac: true,
			},
			entries: []Entry{
				{FileInfo: fakeFileInfo{name: "B"}},
				{FileInfo: fakeFileInfo{name: "a"}},
				{FileInfo: fakeFileInfo{name: "c"}},
			},
			wantNames: []string{"B", "a", "c"},
		},
		{
			name: "size sort",
			sortT: SortOrder{
				method:     SizeSort,
				dirsFirst:  false,
				dirsOnly:   false,
				showHidden: true,
				reverse:    false,
				ignoreDiac: true,
			},
			entries: []Entry{
				{FileInfo: fakeFileInfo{name: "a"}, SizeInt: 3},
				{FileInfo: fakeFileInfo{name: "b"}, SizeInt: 2},
				{FileInfo: fakeFileInfo{name: "c"}, SizeInt: 1},
			},
			wantNames: []string{"c", "b", "a"},
		},
		{
			name: "mod time sort",
			sortT: SortOrder{
				method:     MtimeSort,
				dirsFirst:  false,
				dirsOnly:   false,
				showHidden: true,
				reverse:    false,
				ignoreDiac: true,
			},
			entries: []Entry{
				{FileInfo: fakeFileInfo{name: "c", modTime: time.Now().Add(time.Second * 2)}},
				{FileInfo: fakeFileInfo{name: "a", modTime: time.Now().Add(time.Second * 3)}},
				{FileInfo: fakeFileInfo{name: "b", modTime: time.Now().Add(time.Second * 1)}},
			},
			wantNames: []string{"b", "c", "a"},
		},
		{
			name: "ext sort",
			sortT: SortOrder{
				method:     ExtSort,
				dirsFirst:  false,
				dirsOnly:   false,
				showHidden: true,
				reverse:    false,
				ignoreDiac: true,
			},
			entries: []Entry{
				{FileInfo: fakeFileInfo{name: "z", modTime: time.Now().Add(time.Second * 1)}},
				{FileInfo: fakeFileInfo{name: "a", modTime: time.Now().Add(time.Second * 1)}},
				{FileInfo: fakeFileInfo{name: "a.c", modTime: time.Now().Add(time.Second * 1)}},
				{FileInfo: fakeFileInfo{name: "a.b", modTime: time.Now().Add(time.Second * 2)}},
				{FileInfo: fakeFileInfo{name: "a.a", modTime: time.Now().Add(time.Second * 3)}},
			},
			wantNames: []string{"a", "z", "a.a", "a.b", "a.c"},
		},
		{
			name: "natural sort, dirs first",
			sortT: SortOrder{
				method:    NameSort,
				dirsFirst: true,
				dirsOnly:  false,
				reverse:   false,
			},
			entries: []Entry{
				{FileInfo: fakeFileInfo{name: "a", isDir: false}},
				{FileInfo: fakeFileInfo{name: "z", isDir: true}},
				{FileInfo: fakeFileInfo{name: "b", isDir: false}},
			},
			wantNames: []string{"z", "a", "b"},
		},
		{
			name: "natural sort, dirs first, reverse",
			sortT: SortOrder{
				method:    NameSort,
				dirsFirst: true,
				dirsOnly:  false,
				reverse:   true,
			},
			entries: []Entry{
				{FileInfo: fakeFileInfo{name: "a", isDir: false}},
				{FileInfo: fakeFileInfo{name: "z", isDir: true}},
				{FileInfo: fakeFileInfo{name: "b", isDir: false}},
			},
			wantNames: []string{"b", "a", "z"},
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			got := sortEntries("", tc.entries, tc.sortT)
			gotNames := make([]string, len(got))
			for i, entry := range got {
				gotNames[i] = entry.Name()
			}
			if !reflect.DeepEqual(gotNames, tc.wantNames) {
				t.Errorf("got %v, want %v", gotNames, tc.wantNames)
			}
		})
	}
}

func TestNaturalLess(t *testing.T) {
	t.Parallel()
	tests := []struct {
		s1  string
		s2  string
		exp bool
	}{
		{"foo", "bar", false},
		{"bar", "baz", true},
		{"foo", "123", false},
		{"foo1", "foobar", true},
		{"foo1", "foo10", true},
		{"foo2", "foo10", true},
		{"foo1", "foo10bar", true},
		{"foo2", "foo10bar", true},
		{"foo1bar", "foo10bar", true},
		{"foo2bar", "foo10bar", true},
		{"foo1bar", "foo10", true},
		{"foo2bar", "foo10", true},
	}

	for _, test := range tests {
		if got := naturalLess(test.s1, test.s2); got != test.exp {
			t.Errorf("at input '%s' and '%s' expected '%t' but got '%t'", test.s1, test.s2, test.exp, got)
		}
	}
}
