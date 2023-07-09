package infobar

import (
	"time"

	"github.com/Philistino/fman/ui/theme"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type PromptAskMsg struct {
	// The message to display
	ID          string
	Placeholder string
	Validator   func(string) error
}

func PromptAskCmd(id string, placeholder string, validator func(string) error) tea.Cmd {
	return func() tea.Msg {
		return PromptAskMsg{ID: id, Placeholder: placeholder, Validator: validator}
	}
}

// PromptAnswerMsg is sent when the user answers the prompt
type PromptAnswerMsg struct {
	ID        string
	Message   string
	Cancelled bool
}

// String returns the message as a string
func (m PromptAnswerMsg) String() string {
	return m.Message
}

// PromptAnswerCmd is used to send the answer from the prompt
func PromptAnswerCmd(id string, value string, cancelled bool) tea.Cmd {
	return func() tea.Msg {
		return PromptAnswerMsg{ID: id, Message: value, Cancelled: cancelled}
	}
}

type prompt struct {
	width     int
	textInput textinput.Model
	askMsg    PromptAskMsg
}

func (m prompt) Init() tea.Cmd {
	return textinput.Blink
}

func newPrompt() prompt {
	ti := textinput.New()
	ti.CharLimit = 156
	ti.Width = 20
	ti.PromptStyle = theme.InfobarStyle.Copy()
	ti.TextStyle = theme.InfobarStyle.Copy()
	ti.PlaceholderStyle = theme.InfobarStyle.Copy()
	ti.Cursor.Style = theme.SelectedItemStyle.Copy().Foreground(theme.GetActiveTheme("dracula").SelectedItemBgColor).Background(theme.GetActiveTheme("dracula").SelectedItemFgColor)
	return prompt{
		textInput: ti,
	}
}

type resetPlaceholderMsg struct{}

func resetPlaceholderCmd() tea.Cmd {
	return func() tea.Msg {
		time.Sleep(1000 * time.Millisecond)
		return resetPlaceholderMsg{}
	}
}

func (m prompt) Update(msg tea.Msg) (prompt, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case PromptAskMsg:
		m.textInput.Reset()
		m.askMsg = msg
		m.textInput.Placeholder = msg.Placeholder + " (press ESC to cancel)"
		m.textInput.Focus()
	case PromptAnswerMsg:
		m.textInput.Reset()
	case resetPlaceholderMsg:
		m.textInput.Placeholder = m.askMsg.Placeholder + " (press ESC to cancel)"
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			if m.askMsg.Validator != nil {
				if err := m.askMsg.Validator(m.textInput.Value()); err != nil {
					m.textInput.Reset()
					cmd = resetPlaceholderCmd()
					m.textInput.Placeholder = err.Error()
					return m, cmd
				}
			}
			m.textInput.Blur()
			cmd = PromptAnswerCmd(m.askMsg.ID, m.textInput.Value(), false)
			m.textInput.Reset()
			return m, cmd
		case tea.KeyEsc:
			m.textInput.Blur()
			cmd = PromptAnswerCmd(m.askMsg.ID, m.textInput.Value(), true)
			m.textInput.Reset()
			return m, cmd
		}
	}
	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m prompt) View() string {
	return m.textInput.View()
}
