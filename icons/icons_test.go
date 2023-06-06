// MIT License

// Copyright (c) 2020 Yash Handa

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

package icons

import (
	"log"
	"os"
	"testing"

	"github.com/Philistino/fman/icons/nerdicons"
	"github.com/mattn/go-runewidth"
)

func TestCheckWidth(t *testing.T) {
	for k, v := range iconSet {
		w := runewidth.StringWidth(v.Glyph())
		if w != 1 {
			t.Error(v.Glyph(), w, k)
		}
	}
}

func TestCheckNotDepricated(t *testing.T) {
	glyphToName := make(map[string]string)
	for k, v := range nerdicons.Icons {
		glyphToName[v] = k
	}
	missing := 0
	for k, v := range iconSet {
		if len(v.Glyph()) == 1 {
			continue
		}
		_, ok := glyphToName[v.Glyph()]
		if !ok {
			missing++
			t.Error(v.Glyph(), k)
		}
	}
	if missing > 0 {
		t.Error(missing)
	}
}

// func TestIconsOut2(t *testing.T) {
// 	path := `..`
// 	entries, err := os.ReadDir(path)
// 	if err != nil {
// 		t.Fail()
// 	}

// 	for _, entry := range entries {
// 		info, err := entry.Info()
// 		if err != nil {
// 			t.Fail()
// 		}
// 		icon := GetIconTerm(info, false)
// 		// icon := getGlyph(info, false)
// 		log.Println(icon + info.Name())
// 		// log.Println(lipgloss.Width(icon))
// 	}
// 	t.Error()
// }

func TestIconsOut3(t *testing.T) {
	path := `.`
	entries, err := os.ReadDir(path)
	if err != nil {
		t.Fail()
	}

	for _, entry := range entries {
		info, err := entry.Info()
		if err != nil {
			t.Fail()
		}
		v := GetIconForReal(info, false)
		log.Printf("%s%s\033[0m", v.ColorTerm(), v.Glyph())

		log.Println(GetIconTerm(info, false) + info.Name())
	}
	t.Error()
}

func TestRGBToHex(t *testing.T) {
	testcases := []struct {
		rgb  []int
		hex  string
		name string
	}{
		{[]int{0, 0, 0}, "#000000", "black"},
		{[]int{255, 255, 255}, "#ffffff", "white"},
		{[]int{255, 0, 0}, "#ff0000", "red"},
		{[]int{0, 255, 0}, "#00ff00", "green"},
		{[]int{0, 0, 255}, "#0000ff", "blue"},
		{[]int{255, 255, 0}, "#ffff00", "yellow"},
		{[]int{0, 255, 255}, "#00ffff", "cyan"},
		{[]int{255, 0, 255}, "#ff00ff", "magenta"},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			hex := rgbToHex(tc.rgb[0], tc.rgb[1], tc.rgb[2])
			if hex != tc.hex {
				t.Errorf("expected %s, got %s", tc.hex, hex)
			}
		})
	}
}
