package icons

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"
)

// https://raw.githubusercontent.com/ryanoasis/nerd-fonts/master/css/nerd-fonts-generated.css

type Icon struct {
	glyph      string   // codepoint
	rgb        [3]uint8 // represents the color in rgb (default 0,0,0 is black)
	executable bool     // whether or not the file is executable [true = is executable]
	name       string   // name of the icon
}

// Glyph returns the codepoint of the icon
func (i *Icon) Glyph() string {
	return i.glyph
}

// ColorTerm returns the color escape sequence for the icon for terminal output.
// Note that this does not include the reset escape sequence.
func (i *Icon) ColorTerm() string {
	if i.executable {
		return "\033[38;2;76;175;080m"
	}
	return fmt.Sprintf("\033[38;2;%03d;%03d;%03dm", i.rgb[0], i.rgb[1], i.rgb[2])
}

// ColorHex returns the color of the icon in hex format.
func (i *Icon) ColorHex() string {
	return rgbToHex(int(i.rgb[0]), int(i.rgb[1]), int(i.rgb[2]))
}

// ColorRGB returns the color of the icon in rgb format.
func (i *Icon) ColorRGB() [3]uint8 {
	return i.rgb
}

// rgbToHex converts rgb values to hex values.
func rgbToHex(r, g, b int) string {
	return fmt.Sprintf("#%02x%02x%02x", r, g, b)
}

// makeExe sets the executable flag to true.
func (i *Icon) makeExe() {
	i.executable = true
}

// GetIconTerm returns the icon with escape sequences for colored
// output in the terminal.
func GetIconTerm(info fs.FileInfo, hidden bool) string {
	// The escape sequence '\033[39m' sets the foreground color back to default to default.
	i := GetIconForReal(info, hidden)
	return fmt.Sprintf("%s%s\033[39m", i.ColorTerm(), i.Glyph())
}

// GetIconForReal returns the icon for the given file.
func GetIconForReal(info fs.FileInfo, hidden bool) Icon {
	var i Icon
	var ok bool
	name := info.Name()
	ext := filepath.Ext(name)

	switch info.IsDir() {
	case true:
		i, ok = iconDir[strings.ToLower(name)]
		if ok {
			break
		}
		if hidden {
			i = iconDefault["hiddendir"]
			break
		}
		i = iconDefault["dir"]
		return i
	default:
		i, ok = iconFileName[strings.ToLower(name)]
		if ok {
			break
		}

		// a special admiration for goLang
		if ext == ".go" && strings.HasSuffix(name, "_test") {
			i = iconSet["go-test"]
			break
		}

		t := strings.Split(name, ".")
		if len(t) > 1 && t[0] != "" {
			i, ok = iconSubExt[strings.ToLower(t[len(t)-1]+ext)]
			if ok {
				break
			}
		}

		i, ok = iconExt[strings.ToLower(strings.TrimPrefix(ext, "."))]
		if ok {
			break
		}

		if hidden {
			i = iconDefault["hiddenfile"]
			break
		}
		i = iconDefault["file"]
	}

	// this will change icon color to green if the file is executable
	if info.Mode()&1000000 > 0 {
		if i == iconDefault["file"] {
			i = iconDefault["exe"]
		}
		i.makeExe()
	}

	return i
}
