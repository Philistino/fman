package model

import (
	"strings"
	"testing"

	"github.com/Philistino/fman/model/message"
	tea "github.com/charmbracelet/bubbletea"
	zone "github.com/lrstanley/bubblezone"
)

// TestDialogFocus tests the focus behaviour
func TestDialogFocus(t *testing.T) {
	dialog := NewDialogBox()
	if dialog.focused {
		t.Errorf("dialog should not be focused")
	}

	dialog, _ = dialog.Update(tea.KeyMsg(tea.Key{Type: tea.KeyRight}))
	if dialog.selected != 0 {
		t.Errorf("dialog should not have changed")
	}

	dialog, _ = dialog.Update(message.AskDialogMsg("test"))
	if !dialog.focused {
		t.Errorf("dialog should be focused")
	}
	dialog, _ = dialog.Update(message.AnswerDialogMsg(true))
	if dialog.focused {
		t.Errorf("dialog should not be focused")
	}
}

// TestDialogKeys tests the key behaviour
func TestDialogKeys(t *testing.T) {
	dialog := NewDialogBox()
	dialog, _ = dialog.Update(message.AskDialogMsg("test"))
	if dialog.selected != 0 {
		t.Errorf("dialog should have selected the first choice")
	}

	dialog, _ = dialog.Update(tea.KeyMsg(tea.Key{Type: tea.KeyDown}))
	if dialog.selected != 0 {
		t.Errorf("dialog should not have changed")
	}
	dialog, _ = dialog.Update(tea.KeyMsg(tea.Key{Type: tea.KeyUp}))
	if dialog.selected != 0 {
		t.Errorf("dialog should not have changed")
	}

	dialog, _ = dialog.Update(tea.KeyMsg(tea.Key{Type: tea.KeyRight}))
	if dialog.selected != 1 {
		t.Errorf("dialog should have selected the second choice")
	}
	dialog, _ = dialog.Update(tea.KeyMsg(tea.Key{Type: tea.KeyRight}))
	if dialog.selected != 1 {
		t.Errorf("dialog should not have changed")
	}

	dialog, _ = dialog.Update(tea.KeyMsg(tea.Key{Type: tea.KeyLeft}))
	if dialog.selected != 0 {
		t.Errorf("dialog should have selected the first choice")
	}
	dialog, _ = dialog.Update(tea.KeyMsg(tea.Key{Type: tea.KeyLeft}))
	if dialog.selected != 0 {
		t.Errorf("dialog should not have changed")
	}
}

func TestDialogSelectionFalse(t *testing.T) {
	dialog := NewDialogBox()
	dialog, _ = dialog.Update(message.AskDialogMsg("test"))

	dialog, cmd := dialog.Update(tea.KeyMsg(tea.Key{Type: tea.KeyEnter}))
	if dialog.focused {
		t.Errorf("dialog should not be focused")
	}
	if cmd == nil {
		t.Errorf("dialog should have returned a command")
	}
	answer, ok := cmd().(message.AnswerDialogMsg)
	if !ok {
		t.Errorf("dialog should have returned a message.AnswerDialogCmd")
	}
	if answer {
		t.Errorf("dialog should have returned false")
	}
}

func TestDialogSelectionTrue(t *testing.T) {
	dialog := NewDialogBox()
	dialog, _ = dialog.Update(message.AskDialogMsg("test"))
	dialog, _ = dialog.Update(tea.KeyMsg(tea.Key{Type: tea.KeyRight}))
	dialog, cmd := dialog.Update(tea.KeyMsg(tea.Key{Type: tea.KeyEnter}))
	if dialog.focused {
		t.Errorf("dialog should not be focused")
	}
	if cmd == nil {
		t.Errorf("dialog should have returned a command")
	}
	answer, ok := cmd().(message.AnswerDialogMsg)
	if !ok {
		t.Errorf("dialog should have returned a message.AnswerDialogCmd")
	}
	if !answer {
		t.Errorf("dialog should have returned true")
	}
}

func TestDialogView(t *testing.T) {
	zone.NewGlobal()
	dialog := NewDialogBox()
	options := []string{"Bingo", "Bango"}
	dialog, _ = dialog.Update(message.AskDialogMsg(""))
	dialog.SetChoices(options)
	msgTxt := "This is a test"
	dialog.SetMessage(msgTxt)
	dialog, _ = dialog.Update(tea.KeyMsg(tea.Key{Type: tea.KeyRight}))
	dialog, _ = dialog.Update(tea.KeyMsg(tea.Key{Type: tea.KeyEnter}))
	dialog.SetHeight(200)
	dialog.SetWidth(300)
	view := dialog.View()
	if !strings.Contains(view, msgTxt) {
		t.Errorf("dialog view should contain the message")
	}
	if !strings.Contains(view, options[0]) {
		t.Errorf("dialog view should contain the confirm button")
	}
	if !strings.Contains(view, options[1]) {
		t.Errorf("dialog view should contain the cancel button")
	}
}

func TestDialogDefaultsView(t *testing.T) {
	zone.NewGlobal()
	dialog := NewDialogBox()
	msgTxt := "This is a test"
	dialog, _ = dialog.Update(message.AskDialogMsg(msgTxt))
	dialog, _ = dialog.Update(tea.KeyMsg(tea.Key{Type: tea.KeyRight}))
	dialog, _ = dialog.Update(tea.KeyMsg(tea.Key{Type: tea.KeyEnter}))
	dialog.SetHeight(200)
	dialog.SetWidth(300)
	view := dialog.View()
	if !strings.Contains(view, msgTxt) {
		t.Errorf("dialog view should contain the message")
	}
	if !strings.Contains(view, "confirm") {
		t.Errorf("dialog view should contain the confirm button")
	}
	if !strings.Contains(view, "cancel") {
		t.Errorf("dialog view should contain the cancel button")
	}
}
