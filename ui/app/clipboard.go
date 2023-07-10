package app

import (
	"github.com/Philistino/fman/ui/message"
	tea "github.com/charmbracelet/bubbletea"
)

// clipboardCopy sets the internal clipboard to the selected entries
func (app *App) handleCopy(msg tea.Msg) tea.Cmd {
	var cmd tea.Cmd
	var cut bool
	switch msg.(type) {
	case message.InternalCopyMsg:
		cut = false
		cmd = message.NewNotificationCmd("Copied!")
	case message.CutMsg:
		cut = true
		cmd = message.NewNotificationCmd("Cut!")
	}
	app.Navi.ClipboardCopy(app.list.SelectedEntries(), cut)
	return cmd
}

// // TODO: make this real
// func (app *App) handlePaste() tea.Cmd {
// 	if app.clipboard.Empty() {
// 		return message.NewNotificationCmd("Nothing to paste")
// 	}

// }
