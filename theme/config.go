package theme

import "github.com/Philistino/fman/theme/colors"

type ColorScheme int

const (
	Brogrammer ColorScheme = iota
	CatppuccinFrappe
	CatppuccinLatte
	CatppuccinMacchiato
	CatppuccinMocha
	Dracula
	Everblush
	Gruvbox
	Nord
)

func (c ColorScheme) String() string {
	return [...]string{
		"brogrammer",
		"catppuccin-frappe",
		"catppuccin-latte",
		"catppuccin-macchiato",
		"catppuccin-mocha",
		"dracula",
		"everblush",
		"gruvbox",
		"nord",
	}[c]
}

type ThemeMap map[string]colors.Theme

var themes = ThemeMap{
	"brogrammer":           colors.BrogrammerTheme,
	"catppuccin-frappe":    colors.CatppuccinThemeFrappe,
	"catppuccin-latte":     colors.CatppuccinThemeLatte,
	"catppuccin-macchiato": colors.CatppuccinThemeMacchiato,
	"catppuccin-mocha":     colors.CatppuccinThemeMocha,
	"dracula":              colors.DraculaTheme,
	"everblush":            colors.EverblushTheme,
	"gruvbox":              colors.GruvboxTheme,
	"nord":                 colors.NordTheme,
}

// GetActiveTheme tries to match provided flag value for --theme against
// an existing ThemeMap and returns default theme if theme
// name does not match any records
// in the ThemeMap (due to a typo for example)
func GetActiveTheme(themeNameCandidate string) colors.Theme {
	if _, ok := themes[themeNameCandidate]; ok {
		return themes[themeNameCandidate]
	}
	return colors.DraculaTheme
}
