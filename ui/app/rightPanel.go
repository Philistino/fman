package app

import (
	"github.com/Philistino/fman/ui/dialog"
	"github.com/Philistino/fman/ui/preview"
	tea "github.com/charmbracelet/bubbletea"
)

type rightPanelState uint8

const (
	rightPanelStatePreviewing rightPanelState = iota
	rightPanelStateHidden
	rightPanelStateDialog
)

type rightPanel struct {
	state       rightPanelState
	filePreview *preview.FilePreview
	dialog      *dialog.Dialog
}

func (model *rightPanel) Update(msg tea.Msg) (*rightPanel, tea.Cmd) {
	return model, nil
}

func (model *rightPanel) View() string {
	switch model.state {
	case rightPanelStatePreviewing:
		return model.filePreview.View()
	case rightPanelStateHidden:
		return ""
	case rightPanelStateDialog:
		return model.dialog.View()
	}
	return ""
}
