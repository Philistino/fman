package theme

type ThemeMap map[string]Theme

var themes = ThemeMap{
	"brogrammer":           BrogrammerTheme,
	"catppuccin-frappe":    CatppuccinThemeFrappe,
	"catppuccin-latte":     CatppuccinThemeLatte,
	"catppuccin-macchiato": CatppuccinThemeMacchiato,
	"catppuccin-mocha":     CatppuccinThemeMocha,
	"dracula":              DefaultTheme,
	"default":              DefaultTheme,
	"everblush":            EverblushTheme,
	"gruvbox":              GruvboxTheme,
	"nord":                 NordTheme,
}

// Tries to match provided flag value for --theme against
// an existing ThemeMap and returns default theme if theme
// name does not match any records
// in the ThemeMap (due to a typo for example)
func GetActiveTheme(themeNameCandidate string) (theme Theme) {
	if _, ok := themes[themeNameCandidate]; ok {
		return themes[themeNameCandidate]
	}
	return DefaultTheme
}
