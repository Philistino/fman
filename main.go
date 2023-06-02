package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/76creates/stickers"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	zone "github.com/lrstanley/bubblezone"
	"github.com/muesli/termenv"
	"github.com/nore-dev/fman/cfg"
	"github.com/nore-dev/fman/nav"

	"github.com/nore-dev/fman/keymap"
	"github.com/nore-dev/fman/message"

	"github.com/nore-dev/fman/model/buttonbar"
	"github.com/nore-dev/fman/model/dialog"
	"github.com/nore-dev/fman/model/entryinfo"
	"github.com/nore-dev/fman/model/infobar"
	"github.com/nore-dev/fman/model/list"
	"github.com/nore-dev/fman/model/toolbar"

	"github.com/nore-dev/fman/theme"
)

// The issue is the state isn't coming back from nav for new navigations

// This is the main model for the app. It does two jobs, acts like a message bus for the different
// components of the app, and composes the different UI components together.
type App struct {
	buttonBar buttonbar.ButtonBar
	list      list.List
	entryInfo entryinfo.EntryInfo
	toolbar   *toolbar.Toolbar
	infobar   infobar.Infobar
	dialog    dialog.Model

	width  int
	height int

	flexBox *stickers.FlexBox

	help     help.Model
	showHelp bool
	config   cfg.Cfg

	navigator *nav.Nav
}

func (app *App) Init() tea.Cmd {
	return tea.Batch(app.infobar.Init(), app.UpdatePath(), app.list.Init())
}

func (app *App) UpdatePath() tea.Cmd {
	return func() tea.Msg {
		absolutePath, _ := filepath.Abs(app.config.Path)
		return message.PathMsg{Path: absolutePath}
	}
}

func (app *App) manageSizes(height, width int) {
	app.width = width
	app.height = height
	app.flexBox.SetHeight(height - lipgloss.Height(app.toolbar.View()) - lipgloss.Height(app.toolbar.View()) - lipgloss.Height(app.buttonBar.View()))
	app.flexBox.SetWidth(width)
	app.flexBox.ForceRecalculate()
	app.list.SetWidth(app.flexBox.Row(0).Cell(0).GetWidth())
	app.list.SetHeight(app.flexBox.GetHeight())
	app.entryInfo.SetWidth(app.flexBox.Row(0).Cell(1).GetWidth())
	app.entryInfo.SetHeight(app.flexBox.GetHeight())
	app.help.Width = width
	app.toolbar.SetWidth(width)
}

func (app *App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		app.manageSizes(msg.Height, msg.Width)
	case message.UpdateDialogMsg:
		app.dialog.SetDialog(&msg.Dialog)
		return app, nil
	case message.NavBackMsg:
		cmd = message.HandleBackCmd(app.navigator, []string{app.list.SelectedEntry().Name()})
		cmds = append(cmds, cmd)
	case message.NavFwdMsg:
		cmd = message.HandleFwdCmd(app.navigator, []string{app.list.SelectedEntry().Name()})
		cmds = append(cmds, cmd)
	case message.NavUpMsg:
		cmd = message.HandleNavCmd(app.navigator, []string{app.list.SelectedEntry().Name()}, filepath.Dir(app.navigator.CurrentPath()))
		cmds = append(cmds, cmd)
	case message.NavHomeMsg:
		cmd = message.HandleNavCmd(app.navigator, []string{app.list.SelectedEntry().Name()}, "~")
		cmds = append(cmds, cmd)
	case message.NavDownMsg:
		name := app.list.SelectedEntry().Name()
		cmd = message.HandleNavCmd(app.navigator, []string{name}, filepath.Join(app.navigator.CurrentPath(), name))
		cmds = append(cmds, cmd)
	case tea.KeyMsg:
		if key.Matches(msg, keymap.Default.ToggleHelp) {
			app.showHelp = !app.showHelp
		}
		switch msg.String() {
		case "ctrl+c", "q":
			return app, tea.Quit
		}
	}

	var listCmd, toolbarCmd, entryCmd, infobarCmd, buttonBarCmd tea.Cmd

	app.list, listCmd = app.list.Update(msg)
	app.toolbar, toolbarCmd = app.toolbar.Update(msg)
	app.entryInfo, entryCmd = app.entryInfo.Update(msg)
	app.infobar, infobarCmd = app.infobar.Update(msg)
	app.buttonBar, buttonBarCmd = app.buttonBar.Update(msg)

	cmds = append(cmds, listCmd, toolbarCmd, entryCmd, infobarCmd, buttonBarCmd)

	return app, tea.Batch(cmds...)
}

func (app *App) View() string {
	// Render the dialog if it is open
	if app.dialog.Dialog().IsOpen() {
		return zone.Scan(lipgloss.Place(
			app.width,
			app.height,
			lipgloss.Center,
			lipgloss.Center,
			app.dialog.View(),
		))
	}

	var view string
	switch {
	case app.list.IsEmpty():
		view = app.renderFull(theme.EmptyFolderStyle.Render("This folder is empty"))
	case app.showHelp:
		view = app.renderFull(theme.EmptyFolderStyle.Render(app.help.View(keymap.Default)))
	default:
		app.flexBox.ForceRecalculate()
		row := app.flexBox.Row(0)
		// Set content of list view
		row.Cell(0).SetContent(app.list.View())
		// Set content of entry view
		row.Cell(1).SetContent(app.entryInfo.View())
		view = zone.Mark("list", app.flexBox.Render())
	}

	return zone.Scan(lipgloss.JoinVertical(
		lipgloss.Top,
		app.buttonBar.View(),
		app.toolbar.View(),
		view,
		app.infobar.View(),
	))
}

func (app App) renderFull(str string) string {
	return lipgloss.Place(
		app.flexBox.GetWidth(),
		app.flexBox.GetHeight(),
		lipgloss.Center,
		lipgloss.Center,
		str,
		lipgloss.WithWhitespaceChars("."),
		lipgloss.WithWhitespaceForeground(app.list.Theme().EvenItemBgColor),
	)
}

func main() {
	// Initialize Bubblezone
	zone.NewGlobal()
	defer zone.Close()

	cfg, err := cfg.LoadConfig()
	if err != nil {
		log.Println(err)
	}
	absolutePath, _ := filepath.Abs(cfg.Path)
	navi, err := nav.NewNav(absolutePath, !*cfg.NoHidden, *cfg.DirsMixed)
	if err != nil {
		fmt.Printf("Error reading path: %s, %s", cfg.Path, err)
	}

	selectedTheme := theme.GetActiveTheme(cfg.Theme)
	theme.SetIcons(cfg.Icons)
	theme.SetTheme(selectedTheme)

	listX := list.New(&selectedTheme, navi.Entries())
	entryX := entryinfo.New(&selectedTheme, listX.SelectedEntry(), *cfg.PreviewDelay) // TODO be able to start the app without having a selected entry

	app := App{
		buttonBar: buttonbar.New(&selectedTheme),
		list:      listX,
		entryInfo: entryX,
		toolbar:   toolbar.New(),
		infobar:   infobar.New(),
		dialog:    dialog.New(),
		flexBox:   stickers.NewFlexBox(0, 0),
		config:    cfg,
		navigator: navi,
	}

	app.help.FullSeparator = "   "
	app.help.ShowAll = true

	rows := []*stickers.FlexBoxRow{
		app.flexBox.NewRow().AddCells(
			[]*stickers.FlexBoxCell{
				stickers.NewFlexBoxCell(7, 1).SetStyle(theme.ListStyle), // List
				stickers.NewFlexBoxCell(3, 1),                           // Entry Information
			},
		),
	}

	bg := termenv.BackgroundColor()

	// Set Background Color
	termenv.SetBackgroundColor(termenv.RGBColor(lipgloss.Color(selectedTheme.BackgroundColor)))

	// Reset background color to default
	defer func() {
		termenv.SetBackgroundColor(bg)
	}()

	app.flexBox.AddRows(rows)

	p := tea.NewProgram(&app, tea.WithAltScreen(), tea.WithMouseAllMotion())

	if err := p.Start(); err != nil {
		termenv.SetBackgroundColor(bg)
		println("An error occured: ", err.Error())

		os.Exit(1)
	}
}
