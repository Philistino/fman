package model

import (
	"context"
	"path/filepath"
	"time"

	"github.com/76creates/stickers"
	"github.com/Philistino/fman/cfg"
	"github.com/Philistino/fman/model/dialog"
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
)

// This is the main model for the app. It does two jobs, acts like a message bus for the different
// components of the app, and composes the different UI components together.
type App struct {
	fileBtns   fileBtns
	list       list.List
	preview    *filePreview
	navBtns    *navBtns
	infobar    infobar.Infobar
	dialog     dialog.Model
	breadcrumb *breadCrumb

	width  int
	height int

	flexBox *stickers.FlexBox

	help     help.Model
	showHelp bool
	config   cfg.Cfg

	navi              *nav.Nav
	theme             colors.Theme
	internalClipboard []string // slice of paths to items in the "clipboard"
	PreHandler        *nav.PreviewHandler
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
	return func() tea.Msg {
		return msg
	}
}

func NewApp(cfg cfg.Cfg, selectedTheme colors.Theme) *App {
	absPath, _ := filepath.Abs(cfg.Path)
	absPath = filepath.ToSlash(absPath)
	app := App{
		fileBtns:   newFileBtns(),
		list:       list.New(selectedTheme, *cfg.DoubleClickDelay),
		preview:    NewFilePreviewer(selectedTheme, *cfg.PreviewDelay),
		navBtns:    newNavBtns(),
		infobar:    infobar.New(),
		dialog:     dialog.New(),
		flexBox:    stickers.NewFlexBox(0, 0),
		navi:       nav.NewNav(!*cfg.NoHidden, *cfg.DirsMixed, absPath),
		breadcrumb: newBrdCrumb(),
		theme:      selectedTheme,
		config:     cfg,
		PreHandler: nav.NewPreviewHandler(
			context.Background(),
			*cfg.PreviewDelay,
			50_000, // 50 kB
			100,
			time.Second*time.Duration(30),
		),
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
	app.flexBox.AddRows(rows)

	return &app
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
		cmd = message.HandleBackCmd(app.navi, []string{app.list.SelectedEntry().Name()}, app.list.CursorName())
		cmds = append(cmds, cmd)
	case message.NavFwdMsg:
		cmd = message.HandleFwdCmd(app.navi, []string{app.list.SelectedEntry().Name()}, app.list.CursorName())
		cmds = append(cmds, cmd)
	case message.NavUpMsg:
		cmd = message.HandleNavCmd(app.navi, []string{app.list.SelectedEntry().Name()}, filepath.Dir(app.navi.CurrentPath()), app.list.CursorName())
		cmds = append(cmds, cmd)
	case message.NavHomeMsg:
		cmd = message.HandleNavCmd(app.navi, []string{app.list.SelectedEntry().Name()}, "~", app.list.CursorName())
		cmds = append(cmds, cmd)
	case message.NavDownMsg:
		name := app.list.SelectedEntry().Name()
		cmd = message.HandleNavCmd(app.navi, []string{name}, filepath.Join(app.navi.CurrentPath(), name), app.list.CursorName())
		cmds = append(cmds, cmd)
	case message.NavOtherMsg:
		cmd = message.HandleNavCmd(app.navi, []string{app.list.SelectedEntry().Name()}, msg.Path, app.list.CursorName())
		cmds = append(cmds, cmd)
	case message.InternalCopyMsg:
		cmd = app.setInternalClipboard()
		cmds = append(cmds, cmd)
	case message.InternalPasteMsg:
		cmd = app.paste()
		cmds = append(cmds, cmd)
	case message.ToggleShowHiddenMsg:
		app.navi.SetShowHidden(!app.navi.ShowHidden())
		cmd = message.HandleReloadCmd(app.navi, []string{app.list.SelectedEntry().Name()}, app.list.CursorName())
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

	var listCmd, toolbarCmd, entryCmd, infobarCmd, buttonBarCmd, breadCrmbCmd tea.Cmd

	app.list, listCmd = app.list.Update(msg)
	app.navBtns, toolbarCmd = app.navBtns.Update(msg)
	app.preview, entryCmd = app.preview.Update(msg)
	app.infobar, infobarCmd = app.infobar.Update(msg)
	app.fileBtns, buttonBarCmd = app.fileBtns.Update(msg)
	app.breadcrumb, breadCrmbCmd = app.breadcrumb.Update(msg)

	cmds = append(cmds, listCmd, toolbarCmd, entryCmd, infobarCmd, buttonBarCmd, breadCrmbCmd)

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
		view = app.renderFull(theme.EmptyFolderStyle.Render(app.help.View(keys.Map)))
	default:
		app.flexBox.ForceRecalculate()
		row := app.flexBox.Row(0)
		// Set content of list view
		row.Cell(0).SetContent(app.list.View())
		// Set content of entry view
		row.Cell(1).SetContent(app.preview.View())
		view = zone.Mark("list", app.flexBox.Render())
	}

	secondRow := lipgloss.JoinHorizontal(lipgloss.Center, app.navBtns.View(), app.breadcrumb.View())

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
		app.flexBox.GetWidth(),
		app.flexBox.GetHeight(),
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
	app.flexBox.SetHeight(height - lipgloss.Height(app.navBtns.View()) - lipgloss.Height(app.navBtns.View()) - lipgloss.Height(app.fileBtns.View()))
	app.flexBox.SetWidth(width)
	app.flexBox.ForceRecalculate()
	app.list.SetWidth(app.flexBox.Row(0).Cell(0).GetWidth())
	app.list.SetHeight(app.flexBox.GetHeight())
	app.preview.SetWidth(app.flexBox.Row(0).Cell(1).GetWidth())
	app.preview.SetHeight(app.flexBox.GetHeight())
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
		preview := app.PreHandler.GetPreview(ctx, path)
		return previewReadyMsg{
			Path:    path,
			Preview: preview.Content,
			Err:     preview.Err,
		}
	}
}
