package dialog

import (
	"strings"
	"testing"

	"github.com/Philistino/fman/ui/message"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	zone "github.com/lrstanley/bubblezone"
)

func TestDialogFocus(t *testing.T) {
	zone.NewGlobal()
	dialog := NewDialog(lipgloss.NewStyle(), lipgloss.NewStyle())
	dialog.Init()
	if dialog.Focused() {
		t.Errorf("dialog should not be focused")
	}

	dialog.Focus()
	if !dialog.Focused() {
		t.Errorf("dialog should be focused")
	}

	dialog.Blur()
	if dialog.Focused() {
		t.Errorf("dialog should not be focused")
	}

	dialog, _ = dialog.Update(tea.KeyMsg(tea.Key{Type: tea.KeyRight}))
	if dialog.selected != 0 {
		t.Errorf("dialog should not have changed")
	}
	msg := message.AskDialogCmd("", "Is go the best?", []string{"Yes", "No"})()
	dialog, _ = dialog.Update(msg)
	if !dialog.focused {
		t.Errorf("dialog should be focused")
	}
	dialog, _ = dialog.Update(tea.KeyMsg(tea.Key{Type: tea.KeyEnter}))
	if dialog.Focused() {
		t.Errorf("dialog should not be focused")
	}
}

func TestDialogKeys(t *testing.T) {
	zone.NewGlobal()

	dialog := NewDialog(lipgloss.NewStyle(), lipgloss.NewStyle())
	dialog, _ = dialog.Update(message.AskDialogCmd("", "Is go the best?", []string{"Yes", "No"})())
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
	zone.NewGlobal()
	dialog := NewDialog(lipgloss.NewStyle(), lipgloss.NewStyle())
	dialog, _ = dialog.Update(message.AskDialogCmd("", "Is go the best?", []string{"Yes", "No"})())

	dialog, cmd := dialog.Update(tea.KeyMsg(tea.Key{Type: tea.KeyEnter}))
	if dialog.focused {
		t.Errorf("dialog should not be focused")
	}
	if cmd == nil {
		t.Errorf("dialog should have returned a command")
	}
	answer, ok := cmd().(AnswerMsg)
	if !ok {
		t.Errorf("dialog should have returned a message.AnswerDialogCmd")
	}
	if answer.Answer() != "Yes" {
		t.Errorf("dialog should have returned Yes")
	}
	if answer.AnswerIdx() != 0 {
		t.Errorf("dialog should have returned index of 0")
	}
}

func TestDialogSelectionTrue(t *testing.T) {
	zone.NewGlobal()
	dialog := NewDialog(lipgloss.NewStyle(), lipgloss.NewStyle())
	dialog, _ = dialog.Update(message.AskDialogCmd("", "Is go the best?", []string{"Yes", "No"})())
	dialog, _ = dialog.Update(tea.KeyMsg(tea.Key{Type: tea.KeyRight}))
	dialog, cmd := dialog.Update(tea.KeyMsg(tea.Key{Type: tea.KeyEnter}))
	if dialog.Focused() {
		t.Errorf("dialog should not be focused")
	}
	if cmd == nil {
		t.Errorf("dialog should have returned a command")
	}
	answer, ok := cmd().(AnswerMsg)
	if !ok {
		t.Errorf("dialog should have returned a message.AnswerDialogCmd")
	}
	if answer.Answer() != "No" {
		t.Errorf("dialog should have returned No")
	}
	if answer.AnswerIdx() != 1 {
		t.Errorf("dialog should have returned index of 1")
	}
}

func TestDialogView(t *testing.T) {
	zone.NewGlobal()
	dialog := NewDialog(lipgloss.NewStyle(), lipgloss.NewStyle())
	options := []string{"Bingo", "Bango"}
	msgTxt := "This is a test"
	dialog, _ = dialog.Update(message.AskDialogCmd("", msgTxt, options)())
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
