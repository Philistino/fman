package icons

import (
	"errors"
)

const (
	Dark BackGroundColor = iota
	Light
)

type color struct {
	Name    string `json:"name" yaml:"name" toml:"name" xml:"name" ini:"name" csv:"name"`
	DarkBg  string `json:"darkBg" yaml:"darkBg" toml:"darkBg" xml:"darkBg" ini:"darkBg" csv:"darkBg"`
	LightBg string `json:"lightBg" yaml:"lightBg" toml:"lightBg" xml:"lightBg" ini:"lightBg" csv:"lightBg"`
}

var errInvalidFormat = errors.New("invalid format")

// https://stackoverflow.com/a/54200713
func hexToRGB(s string) ([3]uint8, error) {
	var c [3]uint8
	var err error

	if s[0] != '#' {
		return c, errInvalidFormat
	}

	hexToByte := func(b byte) byte {
		switch {
		case b >= '0' && b <= '9':
			return b - '0'
		case b >= 'a' && b <= 'f':
			return b - 'a' + 10
		case b >= 'A' && b <= 'F':
			return b - 'A' + 10
		}
		err = errInvalidFormat
		return 0
	}

	switch len(s) {
	case 7:
		c[0] = hexToByte(s[1])<<4 + hexToByte(s[2])
		c[1] = hexToByte(s[3])<<4 + hexToByte(s[4])
		c[2] = hexToByte(s[5])<<4 + hexToByte(s[6])
	case 4:
		c[0] = hexToByte(s[1]) * 17
		c[1] = hexToByte(s[2]) * 17
		c[2] = hexToByte(s[3]) * 17
	default:
		err = errInvalidFormat
	}
	return c, err
}

type BackGroundColor uint8

func (c color) RGB(bgColor BackGroundColor) [3]uint8 {
	if bgColor == Dark {
		rgb, _ := hexToRGB(c.DarkBg)
		return rgb
	}
	rgb, _ := hexToRGB(c.LightBg)
	return rgb
}

func (c color) Hex(bgColor BackGroundColor) string {
	if bgColor == Dark {
		return c.DarkBg
	}
	return c.LightBg
}

var (
	nerdIconsMaroon = color{
		Name:    "maroon",
		LightBg: "#8F5536",
		DarkBg:  "#8F5536",
	}
	nerdIconsLmaroon = color{
		Name:    "light maroon",
		LightBg: "#CE7A4E",
		DarkBg:  "#CE7A4E",
	}
	nerdIconsDmaroon = color{
		Name:    "dark maroon",
		LightBg: "#72584B",
		DarkBg:  "#72584B",
	}
	nerdIconsRed = color{
		Name:    "red",
		LightBg: "#AC4142",
		DarkBg:  "#AC4142",
	}
	nerdIconsLred = color{
		Name:    "light red",
		LightBg: "#EB595A",
		DarkBg:  "#EB595A",
	}
	nerdIconsDred = color{
		Name:    "dark red",
		LightBg: "#843031",
		DarkBg:  "#843031",
	}
	nerdIconsRedAlt = color{
		Name:    "red alt",
		LightBg: "#ce5643",
		DarkBg:  "#ce5643",
	}
	nerdIconsGreen = color{
		Name:    "green",
		LightBg: "#90A959",
		DarkBg:  "#90A959",
	}
	nerdIconsLgreen = color{
		Name:    "light green",
		LightBg: "#3D6837",
		DarkBg:  "#C6E87A",
	}
	nerdIconsDgreen = color{
		Name:    "dark green",
		LightBg: "#6D8143",
		DarkBg:  "#6D8143",
	}
	nerdIconsYellow = color{
		Name:    "yellow",
		LightBg: "#FFCC0E",
		DarkBg:  "#FFD446",
	}
	nerdIconsLyellow = color{
		Name:    "light yellow",
		LightBg: "#FF9300",
		DarkBg:  "#FFC16D",
	}
	nerdIconsDyellow = color{
		Name:    "dark yellow",
		LightBg: "#B48D56",
		DarkBg:  "#B48D56",
	}
	nerdIconsBlue = color{
		Name:    "blue",
		LightBg: "#6A9FB5",
		DarkBg:  "#6A9FB5",
	}
	nerdIconsBlueAlt = color{
		Name:    "blue alt",
		LightBg: "#2188b6",
		DarkBg:  "#2188b6",
	}
	nerdIconsLblue = color{
		Name:    "light blue",
		LightBg: "#677174",
		DarkBg:  "#8FD7F4",
	}
	nerdIconsDblue = color{
		Name:    "dark blue",
		LightBg: "#446674",
		DarkBg:  "#446674",
	}
	nerdIconsPink = color{
		Name:    "pink",
		LightBg: "#FC505B",
		DarkBg:  "#F2B4B8",
	}
	nerdIconsLpink = color{
		Name:    "light pink",
		LightBg: "#FF505B",
		DarkBg:  "#FFBDC1",
	}
	nerdIconsDpink = color{
		Name:    "dark pink",
		LightBg: "#B6567E",
		DarkBg:  "#B6567E",
	}
	nerdIconsSilver = color{
		Name:    "silver",
		LightBg: "#716E68",
		DarkBg:  "#716E68",
	}
	nerdIconsLsilver = color{
		Name:    "light silver",
		LightBg: "#7F7869",
		DarkBg:  "#B9B6AA",
	}
	nerdIconsDsilver = color{
		Name:    "dark silver",
		LightBg: "#838484",
		DarkBg:  "#838484",
	}
	nerdIconsCyan = color{
		Name:    "cyan",
		LightBg: "#75B5AA",
		DarkBg:  "#75B5AA",
	}
	nerdIconsLcyan = color{
		Name:    "light cyan",
		DarkBg:  "#A5FDEC",
		LightBg: "#2C7D6E",
	}
	nerdIconsDcyan = color{
		Name:    "dark cyan",
		LightBg: "#48746D",
		DarkBg:  "#48746D",
	}
	nerdIconsCyanAlt = color{
		Name:    "cyan alt",
		DarkBg:  "#61dafb",
		LightBg: "#0595bd",
	}
	nerdIconsPurple = color{
		Name:    "purple",
		DarkBg:  "#AA759F",
		LightBg: "#68295B",
	}
	nerdIconsPurpleAlt = color{
		Name:    "purple alt",
		DarkBg:  "#5D54E1",
		LightBg: "#5D54E1",
	}
	nerdIconsLpurple = color{
		Name:    "light purple",
		DarkBg:  "#E69DD6",
		LightBg: "#E69DD6",
	}
	nerdIconsDpurple = color{
		Name:    "dark purple",
		DarkBg:  "#694863",
		LightBg: "#694863",
	}
	nerdIconsOrange = color{
		Name:    "orange",
		DarkBg:  "#D4843E",
		LightBg: "#D4843E",
	}
	nerdIconsLorange = color{
		Name:    "light orange",
		DarkBg:  "#FFA500",
		LightBg: "#FFA500",
	}
	nerdIconsDorange = color{
		Name:    "dark orange",
		DarkBg:  "#915B2D",
		LightBg: "#915B2D",
	}
)
