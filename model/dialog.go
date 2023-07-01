package model

import (
	"strconv"

	"github.com/Philistino/fman/theme"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	zone "github.com/lrstanley/bubblezone"
)

// This assumes that no other AskDialogMsg will arrive while the current one is open

type AskDialogMsg interface {
	Message() string
	Options() []string
}

type AnswerDialogMsg struct {
	AskDialogMsg
	answerIdx int
}

func (a AnswerDialogMsg) AnswerIdx() int {
	return a.answerIdx
}

func (a AnswerDialogMsg) Answer() string {
	return a.Options()[a.answerIdx]
}

func AnswerDialogCmd(ask AskDialogMsg, answerIdx int) tea.Cmd {
	return func() tea.Msg {
		return AnswerDialogMsg{ask, answerIdx}
	}
}

type dialogBox struct {
	height   int
	width    int
	focused  bool
	message  AskDialogMsg
	selected int
	zPrefix  string
}

// NewDialogBox creates a new dialog box
func NewDialogBox() *dialogBox {
	return &dialogBox{
		zPrefix: "dialog",
		focused: false,
	}
}

func (d *dialogBox) Init() tea.Cmd {
	return nil
}

func (d *dialogBox) Update(msg tea.Msg) (*dialogBox, tea.Cmd) {

	_, ok := msg.(AskDialogMsg)
	if ok {
		d.message = msg.(AskDialogMsg)
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
			d.focused = false
			return d, AnswerDialogCmd(d.message, d.selected)
		}
	case tea.MouseMsg:
		if msg.Type != tea.MouseLeft {
			return d, nil
		}
		for i := range d.message.Options() {
			if zone.Get(d.zPrefix + strconv.Itoa(i)).InBounds(msg) {
				d.focused = false
				return d, AnswerDialogCmd(d.message, i)
			}
		}
	}
	return d, nil
}

func (d *dialogBox) View() string {
	activeBtnStyle := theme.ButtonStyle.Copy().Border(lipgloss.NormalBorder())
	inactiveBtnStyle := theme.ButtonStyle.Copy().Border(lipgloss.HiddenBorder())

	buttons := make([]string, len(d.message.Options()))
	for i, choice := range d.message.Options() {
		if i == d.selected {
			buttons[i] = zone.Mark(d.zPrefix+strconv.Itoa(i), activeBtnStyle.Render(choice))
		} else {
			buttons[i] = zone.Mark(d.zPrefix+strconv.Itoa(i), inactiveBtnStyle.Render(choice))
		}
	}
	renderedButtons := lipgloss.JoinHorizontal(lipgloss.Top, buttons...)
	content := lipgloss.JoinVertical(lipgloss.Center, d.message.Message(), renderedButtons)

	wrapperStyle := theme.EntryInfoStyle.Copy().Height(d.height-2).Width(d.width-margin).Align(lipgloss.Center, lipgloss.Center)
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
func (d *dialogBox) SetHeight(height int) {
	d.height = height
}

// SetWidth sets the width of the dialog box
func (d *dialogBox) SetWidth(width int) {
	d.width = width
}

// Focused returns the focus state of the model
func (m *dialogBox) Focused() bool {
	return m.focused
}

// Focus focuses the model, allowing interaction
func (m *dialogBox) Focus() {
	m.focused = true
}

// Blur freezes the model, preventing selection or movement
func (m *dialogBox) Blur() {
	m.focused = false
}
