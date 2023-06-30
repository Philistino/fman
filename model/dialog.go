package model

import (
	"strconv"

	"github.com/Philistino/fman/model/message"
	"github.com/Philistino/fman/theme"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	zone "github.com/lrstanley/bubblezone"
)

type dialogBox struct {
	height   int
	width    int
	focused  bool
	message  string
	choices  []string
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

	switch msg := msg.(type) {
	case message.AskDialogMsg:
		d.message = string(msg)
		d.choices = []string{"✖ cancel", "✓ confirm"}
		d.selected = 0
		d.focused = true
		return d, nil
	case message.AnswerDialogMsg:
		d.focused = false
		return d, message.AnswerDialogCmd(bool(msg))
	case tea.KeyMsg:
		if !d.focused {
			return d, nil
		}
		switch {
		case msg.String() == "left":
			if d.selected > 0 {
				d.selected--
			}
		case msg.String() == "right":
			if d.selected < len(d.choices)-1 {
				d.selected++
			}
		case msg.String() == "enter":
			d.focused = false
			if d.selected == 0 {
				return d, message.AnswerDialogCmd(false)
			}
			return d, message.AnswerDialogCmd(true)
		}
	case tea.MouseMsg:
		if msg.Type != tea.MouseLeft {
			return d, nil
		}
		for i := range d.choices {
			if zone.Get(d.zPrefix + strconv.Itoa(i)).InBounds(msg) {
				d.focused = false
				if i == 0 {
					return d, message.AnswerDialogCmd(false)
				}
				return d, message.AnswerDialogCmd(true)
			}
		}
	}
	return d, nil
}

func (d *dialogBox) View() string {
	activeBtnStyle := theme.ButtonStyle.Copy().Border(lipgloss.NormalBorder())
	inactiveBtnStyle := theme.ButtonStyle.Copy().Border(lipgloss.HiddenBorder())

	buttons := make([]string, len(d.choices))
	for i, choice := range d.choices {
		if i == d.selected {
			buttons[i] = zone.Mark(d.zPrefix+strconv.Itoa(i), activeBtnStyle.Render(choice))
		} else {
			buttons[i] = zone.Mark(d.zPrefix+strconv.Itoa(i), inactiveBtnStyle.Render(choice))
		}
	}
	renderedButtons := lipgloss.JoinHorizontal(lipgloss.Top, buttons...)
	content := lipgloss.JoinVertical(lipgloss.Center, d.message, renderedButtons)

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

// SetMessage sets the message of the dialog box
func (d *dialogBox) SetMessage(message string) {
	d.message = message
}

// SetChoices sets the choices of the dialog box
func (d *dialogBox) SetChoices(choices []string) {
	d.choices = choices
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
