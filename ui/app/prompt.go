package app

import (
	"context"
	"fmt"

	"github.com/Philistino/fman/entry"
	"github.com/Philistino/fman/ui/infobar"
	"github.com/Philistino/fman/ui/message"
	tea "github.com/charmbracelet/bubbletea"
)

const (
	promptNewFile = "New file"
	promptNewDir  = "New directory"
	promptRename  = "Rename"
)

// fileNameValidator returns a function that validates a filename.
// It checks that the filename is valid and that it is not already taken
// by another entry in the same directory.
func (app *App) fileNameValidator() func(string) error {
	takenNames := app.list.EntryNames()
	return func(name string) error {
		if err := entry.InvalidFilename(name); err != nil {
			return err
		}
		for _, taken := range takenNames {
			if taken == name {
				return fmt.Errorf("name already taken")
			}
		}
		return nil
	}

}

// promptInput prompts the user to enter a new name for a file or directory.
// It returns a tea.Cmd that will display a prompt in the infobar.
// The type of prompt depends on the type of message passed in.
func (app *App) promptInput(msg tea.Msg) tea.Cmd {
	app.list.Blur()
	app.fileBtns.Blur()
	app.navBtns.Blur()
	app.breadcrumb.Blur()

	switch msg.(type) {
	case message.NewFileMsg:
		return infobar.PromptAskCmd(promptNewFile, "New file", app.fileNameValidator())
	case message.MkDirMsg:
		return infobar.PromptAskCmd(promptNewDir, "New folder", app.fileNameValidator())
	case message.RenameMsg:
		return infobar.PromptAskCmd(promptRename, "New name", app.fileNameValidator())
	}
	return nil
}

// handleInput handles the user's response to a prompt.
func (app *App) handleInput(msg infobar.PromptAnswerMsg) tea.Cmd {

	if msg.Cancelled {
		app.list.Focus()
		app.fileBtns.Focus()
		app.navBtns.Focus()
		app.breadcrumb.Focus()
		return nil
	}

	// Should all of these be cmds so they can be run in the background?
	// show spinner while running?
	var errs []error
	switch msg.ID {
	case promptNewFile:
		errs = app.Navi.MkFile(context.Background(), msg.Message)
	case promptNewDir:
		errs = app.Navi.MkDir(context.Background(), msg.Message)
	case promptRename:
		errs = app.Navi.Rename(context.Background(), app.list.SelectedEntryName(), msg.Message)
	}
	return app.handleErrorsAndReload(errs)
}
