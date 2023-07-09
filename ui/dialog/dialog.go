package dialog

import (
	"strconv"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	zone "github.com/lrstanley/bubblezone"
)

const margin = 2

// This assumes that no other AskDialogMsg will arrive while the current one is open.
// There is no queue of dialogs

// AskMsg represents a message that asks a question and has a list of options.
type AskMsg interface {
	ID() string
	Message() string
	Options() []string
}

// AnswerMsg represents a message that contains the index of the selected answer
// and implements the AskMsg interface.
type AnswerMsg struct {
	AskMsg
	answerIdx int
}

// AnswerIdx returns the index of the selected answer.
func (a AnswerMsg) AnswerIdx() int {
	return a.answerIdx
}

// Answer returns the selected answer.
func (a AnswerMsg) Answer() string {
	return a.Options()[a.answerIdx]
}

// AnswerCmd returns a command that sends an AnswerMsg with the given ask and answer.
func AnswerCmd(ask AskMsg, answerIdx int) tea.Cmd {
	return func() tea.Msg {
		return AnswerMsg{ask, answerIdx}
	}
}

// Dialog is a model that represents a dialog box that will display a question and
// and provide buttons for the user to select an answer.
// It accepts tea.msg objects that implements the AskMsg interface.
type Dialog struct {
	btnStyle     lipgloss.Style
	wrapperStyle lipgloss.Style
	height       int
	width        int
	focused      bool
	message      AskMsg
	selected     int
	zPrefix      string
}

// NewDialog creates a new dialog box
func NewDialog(btnStyle lipgloss.Style, wrapperStyle lipgloss.Style) *Dialog {
	return &Dialog{
		zPrefix:      "dialog",
		focused:      false,
		btnStyle:     btnStyle,
		wrapperStyle: wrapperStyle,
	}
}

// Init initializes the dialog box. It is called by the bubbletea framework.
func (d *Dialog) Init() tea.Cmd {
	return nil
}

// Update updates the dialog box. It is called by the bubbletea framework.
func (d *Dialog) Update(msg tea.Msg) (*Dialog, tea.Cmd) {
	_, ok := msg.(AnswerMsg)
	if ok {
		return d, nil
	}

	_, ok = msg.(AskMsg)
	if ok {
		d.message = msg.(AskMsg)
		d.selected = 0
		d.focused = true
		return d, nil
	}

	if !d.focused {
		return d, nil
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case msg.String() == "left":
			if d.selected > 0 {
				d.selected--
			}
		case msg.String() == "right":
			if d.selected < len(d.message.Options())-1 {
				d.selected++
			}
		case msg.String() == "enter":
			d.Blur()
			return d, AnswerCmd(d.message, d.selected)
		}
	case tea.MouseMsg:
		if msg.Type != tea.MouseLeft {
			return d, nil
		}
		for i := range d.message.Options() {
			if zone.Get(d.zPrefix + strconv.Itoa(i)).InBounds(msg) {
				d.Blur()
				return d, AnswerCmd(d.message, i)
			}
		}
	}
	return d, nil
}

// View renders the dialog box. It is called by the bubbletea framework.
func (d *Dialog) View() string {
	activeBtnStyle := d.btnStyle.Copy().Border(lipgloss.NormalBorder()).Padding(0, 2)
	inactiveBtnStyle := d.btnStyle.Copy().Border(lipgloss.HiddenBorder()).Padding(0, 2)
	buttons := make([]string, len(d.message.Options()))
	for i, choice := range d.message.Options() {
		style := inactiveBtnStyle
		if i == d.selected {
			style = activeBtnStyle
		}
		buttons[i] = zone.Mark(d.zPrefix+strconv.Itoa(i), style.Render(choice))
	}
	renderedButtons := lipgloss.JoinHorizontal(lipgloss.Center, buttons...)
	content := lipgloss.JoinVertical(lipgloss.Center, d.message.Message(), renderedButtons)
	wrapperStyle := d.wrapperStyle.Copy().Height(d.height).Width(d.width-margin).Align(lipgloss.Center, lipgloss.Center)
	return wrapperStyle.Render(content)

	// return theme.EntryInfoStyle.Render(
	// 	previewStyle.
	// 		MaxHeight(d.height).
	// 		Height(d.height - 2).
	// 		MaxWidth(d.width - margin).
	// 		Render(content),
	// )

	// return lipgloss.Place(
	// 	d.width-2,
	// 	d.height,
	// 	lipgloss.Center,
	// 	lipgloss.Center,
	// 	lipgloss.JoinVertical(lipgloss.Center, d.message, renderedButtons),
	// 	// lipgloss.WithWhitespaceChars("."),
	// 	lipgloss.WithWhitespaceForeground(theme.EvenItemStyle.GetBackground()),
	// )
}

// SetHeight sets the height of the dialog box
func (d *Dialog) SetHeight(height int) {
	d.height = height
}

// SetWidth sets the width of the dialog box
func (d *Dialog) SetWidth(width int) {
	d.width = width
}

// Focused returns the focus state of the model
func (m *Dialog) Focused() bool {
	return m.focused
}

// Focus focuses the model, allowing interaction
func (m *Dialog) Focus() {
	m.focused = true
}

// Blur freezes the model, preventing selection or movement
func (m *Dialog) Blur() {
	m.focused = false
}
