// MIT License

// Copyright (c) 2016 Gökçehan Kara

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package entry

import (
	"os"
	"reflect"
	"testing"
)

func TestIsRoot(t *testing.T) {
	t.Parallel()
	sep := string(os.PathSeparator)
	if !isRoot(sep) {
		t.Errorf(`"%s" is root`, sep)
	}

	paths := []string{
		"",
		"~",
		"foo",
		"foo/bar",
		"foo/bar",
		"/home",
		"/home/user",
	}

	for _, p := range paths {
		if isRoot(p) {
			t.Errorf("'%s' is not root", p)
		}
	}
}

func TestRuneSliceWidth(t *testing.T) {
	t.Parallel()
	tests := []struct {
		rs  []rune
		exp int
	}{
		{[]rune{'a', 'b'}, 2},
		{[]rune{'ı', 'ş'}, 2},
		{[]rune{'世', '界'}, 4},
		{[]rune{'世', 'a', '界', 'ı'}, 6},
	}

	for _, test := range tests {
		if got := runeSliceWidth(test.rs); got != test.exp {
			t.Errorf("at input '%v' expected '%d' but got '%d'", test.rs, test.exp, got)
		}
	}
}

func TestEscape(t *testing.T) {
	t.Parallel()
	tests := []struct {
		s   string
		exp string
	}{
		{"", ""},
		{"foo", "foo"},
		{"foo bar", `foo\ bar`},
		{"foo  bar", `foo\ \ bar`},
		{`foo\bar`, `foo\\bar`},
		{`foo\ bar`, `foo\\\ bar`},
		{`foo;bar`, `foo\;bar`},
		{`foo#bar`, `foo\#bar`},
		{`foo\tbar`, `foo\\tbar`},
		{"foo\tbar", "foo\\\tbar"},
		{`foo\`, `foo\\`},
	}

	for _, test := range tests {
		if got := escape(test.s); !reflect.DeepEqual(got, test.exp) {
			t.Errorf("at input '%v' expected '%v' but got '%v'", test.s, test.exp, got)
		}
	}
}

func TestUnescape(t *testing.T) {
	t.Parallel()
	tests := []struct {
		s   string
		exp string
	}{
		{"", ""},
		{"foo", "foo"},
		{`foo\ bar`, "foo bar"},
		{`foo\ \ bar`, "foo  bar"},
		{`foo\\bar`, `foo\bar`},
		{`foo\\\ bar`, `foo\ bar`},
		{`foo\;bar`, `foo;bar`},
		{`foo\#bar`, `foo#bar`},
		{`foo\\tbar`, `foo\tbar`},
		{"foo\\\tbar", "foo\tbar"},
		{`foo\`, `foo\`},
	}

	for _, test := range tests {
		if got := unescape(test.s); !reflect.DeepEqual(got, test.exp) {
			t.Errorf("at input '%v' expected '%v' but got '%v'", test.s, test.exp, got)
		}
	}
}

func TestHumanizeSize(t *testing.T) {
	t.Parallel()
	tests := []struct {
		i   int64
		exp string
	}{
		{0, "0B"},
		{9, "9B"},
		{99, "99B"},
		{999, "999B"},
		{1000, "1.0K"},
		{1023, "1.0K"},
		{1025, "1.0K"},
		{1049, "1.0K"},
		{1050, "1.0K"},
		{1099, "1.0K"},
		{9999, "9.9K"},
		{10000, "10K"},
		{10100, "10K"},
		{10500, "10K"},
		{1000000, "1.0M"},
	}

	for _, test := range tests {
		if got := humanizeSize(test.i); got != test.exp {
			t.Errorf("at input '%d' expected '%s' but got '%s'", test.i, test.exp, got)
		}
	}
}
