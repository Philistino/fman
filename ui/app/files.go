package app

import (
	"context"
	"fmt"
	"sort"
	"strings"

	"github.com/Philistino/fman/ui/dialog"
	"github.com/Philistino/fman/ui/message"
	"github.com/Philistino/fman/ui/preview"
	tea "github.com/charmbracelet/bubbletea"
)

func (app *App) handleDialogAnswer(msg dialog.AnswerMsg) tea.Cmd {
	if msg.ID() == "Delete" && msg.Answer() == "Confirm" {
		return app.deleteEntries()
	}
	return nil
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

func (app *App) deleteEntries() tea.Cmd {
	entries := app.list.SelectedEntries()
	entryNames := make([]string, 0, len(entries))
	for k := range entries {
		entryNames = append(entryNames, k)
	}
	sort.Strings(entryNames)
	errs := app.Navi.Delete(context.Background(), entryNames)
	return app.handleErrorsAndReload(errs)
}

func (app *App) handleErrorsAndReload(errs []error) tea.Cmd {
	cmds := make([]tea.Cmd, 0, len(errs)+1)
	for _, err := range errs {
		if err != nil {
			cmds = append(cmds, message.NewNotificationCmd(err.Error()))
		}
	}

	cmd := message.HandleReloadCmd(app.Navi, []string{app.list.SelectedEntryName()}, app.list.CursorName())
	cmds = append(cmds, cmd)

	app.list.Focus()
	app.fileBtns.Focus()
	app.navBtns.Focus()
	app.breadcrumb.Focus()

	return tea.Batch(cmds...)
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
