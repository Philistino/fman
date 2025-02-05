package main

import (
	"log"
	"os"

	"github.com/Philistino/fman/cfg"
	"github.com/Philistino/fman/ui/app"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	zone "github.com/lrstanley/bubblezone"
	"github.com/muesli/termenv"
	"github.com/spf13/afero"

	"github.com/Philistino/fman/ui/theme"
)

func main() {

	zone.NewGlobal()
	defer zone.Close()

	cfg, err := cfg.LoadConfig()
	if err != nil {
		log.Println(err)
	}

	// TODO: move theme/icons to config and return them on the config struct
	selectedTheme := theme.GetActiveTheme(cfg.Theme)
	theme.SetIcons(cfg.Icons)
	theme.SetTheme(selectedTheme)

	// Set background color then reset it on quit
	bg := termenv.BackgroundColor()
	output := termenv.NewOutput(os.Stdout)
	output.SetBackgroundColor(termenv.RGBColor(lipgloss.Color(selectedTheme.BackgroundColor)))
	defer output.SetBackgroundColor(bg)

	a := app.NewApp(cfg, selectedTheme, afero.NewOsFs())
	p := tea.NewProgram(a, tea.WithAltScreen(), tea.WithMouseCellMotion(), tea.WithoutCatchPanics())
	_, err = p.Run()
	if err != nil {
		println("An error occured: ", err.Error())
	}
	if cfg.PrintPwdResult != nil && *cfg.PrintPwdResult {
		println(a.Navi.CurrentPath())
	}
}
