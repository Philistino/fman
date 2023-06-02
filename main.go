package main

import (
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	zone "github.com/lrstanley/bubblezone"
	"github.com/muesli/termenv"
	"github.com/nore-dev/fman/cfg"
	"github.com/nore-dev/fman/model"

	"github.com/nore-dev/fman/theme"
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

	// Reset background color to default and reset on quit
	bg := termenv.BackgroundColor()
	termenv.SetBackgroundColor(termenv.RGBColor(lipgloss.Color(selectedTheme.BackgroundColor)))
	defer func() {
		termenv.SetBackgroundColor(bg)
	}()

	app := model.NewApp(cfg, selectedTheme)
	p := tea.NewProgram(app, tea.WithAltScreen(), tea.WithMouseAllMotion())

	if err := p.Start(); err != nil {
		termenv.SetBackgroundColor(bg)
		println("An error occured: ", err.Error())
		os.Exit(1)
	}
}
