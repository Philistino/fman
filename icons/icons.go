package icons

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"
)

// iconInfo (icon information)
type iconInfo struct {
	i string
	c [3]uint8 // represents the color in rgb (default 0,0,0 is black)
	e bool     // whether or not the file is executable [true = is executable]
}

func (i *iconInfo) GetGlyph() string {
	return i.i
}

func (i *iconInfo) GetColor() string {
	if i.e {
		return "\033[38;2;76;175;080m"
	}
	return fmt.Sprintf("\033[38;2;%03d;%03d;%03dm", i.c[0], i.c[1], i.c[2])
}

func (i *iconInfo) MakeExe() {
	i.e = true
}

// TODO: this isn't correct
func GetIcon(info fs.FileInfo, hidden bool) string {
	i := getGlyph(info, hidden)
	if i.e {
		return fmt.Sprintf("%s\033[38;2;76;175;080m\033[0m", i.GetGlyph())
	}
	return fmt.Sprintf("%s\033[38;2;%03d;%03d;%03dm\033[0m", i.GetGlyph(), i.c[0], i.c[1], i.c[2])
	// return fmt.Sprintf("%s%s\033[0m", i.GetColor(), i.GetGlyph())
}

func GetIcon2(info fs.FileInfo, hidden bool) string {
	i := getGlyph(info, hidden)
	return fmt.Sprintf("%s%s\033[0m", i.GetColor(), i.GetGlyph())
}

func getGlyph(info fs.FileInfo, hidden bool) iconInfo {
	var i iconInfo
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
		if i.GetGlyph() == "\U000f0224" {
			i = iconDefault["exe"]
		}
		i.MakeExe()
	}

	return i
}
