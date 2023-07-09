package app

import (
	"context"
	"fmt"
	"path/filepath"
	"sort"
	"strings"

	"github.com/Philistino/fman/cfg"
	"github.com/Philistino/fman/entry"
	"github.com/Philistino/fman/nav"
	"github.com/Philistino/fman/ui/breadcrumb"
	"github.com/Philistino/fman/ui/dialog"
	"github.com/Philistino/fman/ui/filebtns"
	"github.com/Philistino/fman/ui/infobar"
	"github.com/Philistino/fman/ui/keys"
	"github.com/Philistino/fman/ui/list"
	"github.com/Philistino/fman/ui/message"
	"github.com/Philistino/fman/ui/navbtns"
	"github.com/Philistino/fman/ui/preview"
	"github.com/Philistino/fman/ui/theme"
	"github.com/Philistino/fman/ui/theme/colors"
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
	fileBtns   filebtns.FileBtns
	list       list.List
	preview    *preview.FilePreview
	navBtns    *navbtns.NavBtns
	infobar    infobar.Infobar2
	dialog     *dialog.Dialog
	breadcrumb *breadcrumb.BreadCrumb

	width  int
	height int

	help     help.Model
	showHelp bool
	config   cfg.Cfg

	Navi              *nav.Nav
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
	cmd := message.HandleReloadCmd(app.Navi, []string{""}, "")
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
		fileBtns:   filebtns.NewFileBtns(),
		list:       list.New(selectedTheme, *cfg.DoubleClickDelay),
		preview:    preview.NewFilePreviewer(selectedTheme, *cfg.PreviewDelay),
		navBtns:    navbtns.NewNavBtns(),
		infobar:    infobar.New2(),
		dialog:     dialog.NewDialog(theme.ButtonStyle, theme.EntryInfoStyle),
		Navi:       nav.NewNav(!*cfg.NoHidden, *cfg.DirsMixed, absPath, fsys, *cfg.PreviewDelay, *cfg.DryRun),
		breadcrumb: breadcrumb.NewBreadCrumb(),
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
		cmd = message.HandleBackCmd(app.Navi, []string{app.list.SelectedEntryName()}, app.list.CursorName())
		cmds = append(cmds, cmd)
	case message.NavFwdMsg:
		cmd = message.HandleFwdCmd(app.Navi, []string{app.list.SelectedEntryName()}, app.list.CursorName())
		cmds = append(cmds, cmd)
	case message.NavUpMsg:
		cmd = message.HandleNavCmd(app.Navi, []string{app.list.SelectedEntryName()}, filepath.Dir(app.Navi.CurrentPath()), app.list.CursorName())
		cmds = append(cmds, cmd)
	case message.NavHomeMsg:
		cmd = message.HandleNavCmd(app.Navi, []string{app.list.SelectedEntryName()}, "~", app.list.CursorName())
		cmds = append(cmds, cmd)
	case message.NavDownMsg:
		name := app.list.SelectedEntryName()
		cmd = message.HandleNavCmd(app.Navi, []string{name}, filepath.Join(app.Navi.CurrentPath(), name), app.list.CursorName())
		cmds = append(cmds, cmd)
	case message.NavOtherMsg:
		cmd = message.HandleNavCmd(app.Navi, []string{app.list.SelectedEntryName()}, msg.Path, app.list.CursorName())
		cmds = append(cmds, cmd)
	case message.InternalCopyMsg:
		cmd = app.clipboardCopy()
		cmds = append(cmds, cmd)
	case message.InternalPasteMsg:
		cmd = app.clipboardPaste()
		cmds = append(cmds, cmd)
	case message.ToggleShowHiddenMsg:
		app.Navi.SetShowHidden(!app.Navi.ShowHidden())
		cmd = message.HandleReloadCmd(app.Navi, []string{app.list.SelectedEntryName()}, app.list.CursorName())
		cmds = append(cmds, cmd)
	case message.GetPreviewMsg:
		cmd = app.getPreviewCmd(msg.Ctx, msg.Path)
		cmds = append(cmds, cmd)
	case message.DeleteMsg:
		cmd = app.handleDeleteCmd()
		cmds = append(cmds, cmd)
	case dialog.AnswerMsg:
		cmd = app.handleDialogAnswer(msg)
		cmds = append(cmds, cmd)
	case message.NewFileMsg, message.MkDirMsg, message.RenameMsg:
		cmd = app.promptName(msg)
		cmds = append(cmds, cmd)
	case infobar.PromptAnswerMsg:
		cmd = app.handleInfobarPromptAnswer(msg)
		cmds = append(cmds, cmd)
	case tea.KeyMsg:
		if key.Matches(msg, keys.Map.ToggleHelp) {
			// TODO Freeze components if showing help
			app.showHelp = !app.showHelp
		}
		switch msg.String() {
		case "ctrl+c", "ctrl+q":
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
	app.dialog.SetHeight(height - lipgloss.Height(app.navBtns.View()) - lipgloss.Height(app.navBtns.View()) - lipgloss.Height(app.fileBtns.View()) - 2)
	app.dialog.SetWidth(width - listWidth)
	app.help.Width = width
	app.breadcrumb.SetWidth(width - lipgloss.Width(app.navBtns.View()))
}

// clipboardCopy sets the internal clipboard to the selected entries
func (app *App) clipboardCopy() tea.Cmd {
	selected := app.list.SelectedEntries()
	clipboard := make([]string, 0, len(selected))
	dir := app.Navi.CurrentPath()
	for name := range selected {
		clipboard = append(clipboard, filepath.Join(dir, name))
	}
	app.internalClipboard = clipboard
	return message.NewNotificationCmd("Copied!")
}

// TODO: make this real
func (app *App) clipboardPaste() tea.Cmd {
	return message.NewNotificationCmd("Paste!")
}

func (app *App) getPreviewCmd(ctx context.Context, path string) tea.Cmd {
	return func() tea.Msg {
		prv := app.Navi.GetPreview(ctx, path)
		return preview.PreviewReadyMsg{
			Path:    path,
			Preview: prv.Content,
			Err:     prv.Err,
		}
	}
}

func (app *App) handleDeleteCmd() tea.Cmd {
	app.list.Blur()
	entries := app.list.SelectedEntries()
	if len(entries) == 0 {
		return message.NewNotificationCmd("No entries selected")
	}
	entryNames := make([]string, 0, len(entries))
	for k := range entries {
		entryNames = append(entryNames, k)
	}
	sort.Strings(entryNames)

	return message.AskDialogCmd(
		"Delete",
		fmt.Sprintf("Permanently delete?\n%s", strings.Join(entryNames, ", ")),
		[]string{"Cancel", "Confirm"},
	)
}

func (app *App) handleDialogAnswer(msg dialog.AnswerMsg) tea.Cmd {
	if msg.ID() == "Delete" && msg.Answer() == "Confirm" {
		return app.deleteEntries()
	}
	return nil
}

func (app *App) deleteEntries() tea.Cmd {
	entries := app.list.SelectedEntries()
	entryNames := make([]string, 0, len(entries))
	for k := range entries {
		entryNames = append(entryNames, k)
	}
	sort.Strings(entryNames)
	errs := app.Navi.Delete(context.Background(), entryNames)
	cmds := make([]tea.Cmd, 0, len(errs))
	for _, err := range errs {
		cmds = append(cmds, message.NewNotificationCmd(err.Error()))
	}
	cmd := message.HandleReloadCmd(app.Navi, []string{app.list.SelectedEntryName()}, app.list.CursorName())
	cmds = append(cmds, cmd)
	app.list.Focus()
	return tea.Batch(cmds...)
}

const (
	promptNewFile = "New file"
	promptNewDir  = "New directory"
	promptRename  = "Rename"
)

// promptName prompts the user to enter a new name for a file or directory.
// It returns a tea.Cmd that will display a prompt in the infobar.
// The type of prompt depends on the type of message passed in.
func (app *App) promptName(msg tea.Msg) tea.Cmd {
	app.list.Blur()
	app.fileBtns.Blur()
	app.navBtns.Blur()
	app.breadcrumb.Blur()

	takenNames := app.list.EntryNames()

	validator := func(s string) error {
		if err := entry.InvalidFilename(s); err != nil {
			return err
		}
		for _, name := range takenNames {
			if name == s {
				return fmt.Errorf("name already taken")
			}
		}
		return nil
	}

	switch msg.(type) {
	case message.NewFileMsg:
		return infobar.PromptAskCmd(promptNewFile, "New file", validator)
	case message.MkDirMsg:
		return infobar.PromptAskCmd(promptNewDir, "New folder", validator)
	case message.RenameMsg:
		return infobar.PromptAskCmd(promptNewDir, "New name", validator)
	}
	return nil
}

// handleInfobarPromptAnswer
func (app *App) handleInfobarPromptAnswer(msg infobar.PromptAnswerMsg) tea.Cmd {
	app.list.Focus()
	app.fileBtns.Focus()
	app.navBtns.Focus()
	app.breadcrumb.Focus()

	if msg.Cancelled {
		return nil
	}

	switch msg.ID {
	case promptNewFile:
		return message.NewNotificationCmd("New file: " + msg.Message)
	case promptNewDir:
		return message.NewNotificationCmd("New directory: " + msg.Message)
	case promptRename:
		return message.NewNotificationCmd("Rename: " + msg.Message)
	}
	return nil
}
