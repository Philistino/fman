package model

import (
	"context"
	"path/filepath"

	"github.com/Philistino/fman/cfg"
	"github.com/Philistino/fman/model/infobar"
	"github.com/Philistino/fman/model/keys"
	"github.com/Philistino/fman/model/list"
	"github.com/Philistino/fman/model/message"
	"github.com/Philistino/fman/nav"
	"github.com/Philistino/fman/theme"
	"github.com/Philistino/fman/theme/colors"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	zone "github.com/lrstanley/bubblezone"
	"github.com/spf13/afero"
)

// This is the main model for the app. It does two jobs, acts like a message bus for the different
// components of the app, and composes the different UI components together.
type App struct {
	fileBtns   fileBtns
	list       list.List
	preview    *filePreview
	navBtns    *navBtns
	infobar    infobar.Infobar
	dialog     *dialogBox
	breadcrumb *breadCrumb

	width  int
	height int

	help     help.Model
	showHelp bool
	config   cfg.Cfg

	navi              *nav.Nav
	theme             colors.Theme
	internalClipboard []string // slice of paths to items in the "clipboard"
}

func (app *App) Init() tea.Cmd {
	// calling this here so the list does not show an empty dir before
	// the first filesystem read is complete
	// TODO: show loading directory message if this takes more than a tenth of a second

	// I am calling Reload here because it does not record the current state in the history
	// and there is no current state to leave behind. This is a little bit of a bandaid
	// over the history implementation not having a setter for initial state
	cmd := message.HandleReloadCmd(app.navi, []string{""}, "")
	msg := cmd() // get the initial DirChangedMsg
	load := func() tea.Msg {
		return msg
	}
	return tea.Batch(load, app.preview.Init())
}

func NewApp(cfg cfg.Cfg, selectedTheme colors.Theme, fsys afero.Fs) *App {
	absPath, err := filepath.Abs(filepath.Clean(cfg.Path))
	if err != nil {
		panic(err)
	}
	app := App{
		fileBtns:   newFileBtns(),
		list:       list.New(selectedTheme, *cfg.DoubleClickDelay),
		preview:    NewFilePreviewer(selectedTheme, *cfg.PreviewDelay),
		navBtns:    newNavBtns(),
		infobar:    infobar.New(),
		dialog:     NewDialogBox(),
		navi:       nav.NewNav(!*cfg.NoHidden, *cfg.DirsMixed, absPath, fsys, *cfg.PreviewDelay),
		breadcrumb: newBrdCrumb(),
		theme:      selectedTheme,
		config:     cfg,
	}
	app.help.FullSeparator = "   "
	app.help.ShowAll = true

	return &app
}

func (app *App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		app.manageSizes(msg.Height, msg.Width)
	case message.NavBackMsg:
		cmd = message.HandleBackCmd(app.navi, []string{app.list.SelectedEntryName()}, app.list.CursorName())
		cmds = append(cmds, cmd)
	case message.NavFwdMsg:
		cmd = message.HandleFwdCmd(app.navi, []string{app.list.SelectedEntryName()}, app.list.CursorName())
		cmds = append(cmds, cmd)
	case message.NavUpMsg:
		cmd = message.HandleNavCmd(app.navi, []string{app.list.SelectedEntryName()}, filepath.Dir(app.navi.CurrentPath()), app.list.CursorName())
		cmds = append(cmds, cmd)
	case message.NavHomeMsg:
		cmd = message.HandleNavCmd(app.navi, []string{app.list.SelectedEntryName()}, "~", app.list.CursorName())
		cmds = append(cmds, cmd)
	case message.NavDownMsg:
		name := app.list.SelectedEntryName()
		cmd = message.HandleNavCmd(app.navi, []string{name}, filepath.Join(app.navi.CurrentPath(), name), app.list.CursorName())
		cmds = append(cmds, cmd)
	case message.NavOtherMsg:
		cmd = message.HandleNavCmd(app.navi, []string{app.list.SelectedEntryName()}, msg.Path, app.list.CursorName())
		cmds = append(cmds, cmd)
	case message.InternalCopyMsg:
		cmd = app.setInternalClipboard()
		cmds = append(cmds, cmd)
	case message.InternalPasteMsg:
		cmd = app.paste()
		cmds = append(cmds, cmd)
	case message.ToggleShowHiddenMsg:
		app.navi.SetShowHidden(!app.navi.ShowHidden())
		cmd = message.HandleReloadCmd(app.navi, []string{app.list.SelectedEntryName()}, app.list.CursorName())
		cmds = append(cmds, cmd)
	case message.GetPreviewMsg:
		cmd = app.getPreviewCmd(msg.Ctx, msg.Path)
		cmds = append(cmds, cmd)
	case tea.KeyMsg:
		if key.Matches(msg, keys.Map.ToggleHelp) {
			// TODO Freeze components if showing help
			app.showHelp = !app.showHelp
		}
		switch msg.String() {
		case "ctrl+c", "q":
			return app, tea.Quit
		}
	}

	var listCmd, toolbarCmd, entryCmd, infobarCmd, buttonBarCmd, breadCrmbCmd, dialogCmd tea.Cmd

	app.list, listCmd = app.list.Update(msg)
	app.navBtns, toolbarCmd = app.navBtns.Update(msg)
	app.preview, entryCmd = app.preview.Update(msg)
	app.infobar, infobarCmd = app.infobar.Update(msg)
	app.fileBtns, buttonBarCmd = app.fileBtns.Update(msg)
	app.breadcrumb, breadCrmbCmd = app.breadcrumb.Update(msg)
	app.dialog, dialogCmd = app.dialog.Update(msg)

	cmds = append(cmds, listCmd, toolbarCmd, entryCmd, infobarCmd, buttonBarCmd, breadCrmbCmd, dialogCmd)

	return app, tea.Batch(cmds...)
}

func (app *App) View() string {

	var view string
	switch {
	case app.dialog.Focused():
		view = lipgloss.JoinHorizontal(
			lipgloss.Top,
			app.list.View(),
			app.dialog.View(),
		)
	case app.showHelp:
		view = app.renderFull(theme.EmptyFolderStyle.Render(app.help.View(keys.Map)))
	default:
		view = lipgloss.JoinHorizontal(
			lipgloss.Top,
			app.list.View(),
			app.preview.View(),
		)
		view = zone.Mark("list", view)
	}

	secondRow := lipgloss.JoinHorizontal(lipgloss.Top, app.navBtns.View(), app.breadcrumb.View())

	return zone.Scan(lipgloss.JoinVertical(
		lipgloss.Top,
		app.fileBtns.View(),
		secondRow,
		view,
		app.infobar.View(),
	))
}

func (app *App) renderFull(str string) string {
	return lipgloss.Place(
		app.width,
		app.height,
		lipgloss.Center,
		lipgloss.Center,
		str,
		lipgloss.WithWhitespaceChars("."),
		lipgloss.WithWhitespaceForeground(app.theme.EvenItemBgColor),
	)
}

func (app *App) manageSizes(height, width int) {
	app.width = width
	app.height = height
	app.list.SetHeight(height - lipgloss.Height(app.navBtns.View()) - lipgloss.Height(app.navBtns.View()) - lipgloss.Height(app.fileBtns.View()) - 2) // maybe remove the -1?
	listWidth := (width * 2) / 3
	app.list.SetWidth(listWidth)
	app.preview.SetHeight(height - lipgloss.Height(app.navBtns.View()) - lipgloss.Height(app.navBtns.View()) - lipgloss.Height(app.fileBtns.View()))
	app.preview.SetWidth(width - listWidth)
	app.dialog.SetHeight(height - lipgloss.Height(app.navBtns.View()) - lipgloss.Height(app.navBtns.View()) - lipgloss.Height(app.fileBtns.View()))
	app.dialog.SetWidth(width - listWidth)
	app.help.Width = width
	app.breadcrumb.SetWidth(width - lipgloss.Width(app.navBtns.View()))
}

// setInternalClipboard sets the internal clipboard to the selected entries
func (app *App) setInternalClipboard() tea.Cmd {
	selected := app.list.SelectedEntries()
	clipboard := make([]string, 0, len(selected))
	dir := app.navi.CurrentPath()
	for name := range selected {
		clipboard = append(clipboard, filepath.Join(dir, name))
	}
	app.internalClipboard = clipboard
	return message.NewNotificationCmd("Copied!")
}

// TODO: make this real
func (app *App) paste() tea.Cmd {
	return message.NewNotificationCmd("Paste!")
}

func (app *App) getPreviewCmd(ctx context.Context, path string) tea.Cmd {
	return func() tea.Msg {
		preview := app.navi.GetPreview(ctx, path)
		return previewReadyMsg{
			Path:    path,
			Preview: preview.Content,
			Err:     preview.Err,
		}
	}
}
