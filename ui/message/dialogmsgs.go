package message

import tea "github.com/charmbracelet/bubbletea"

type AskDialogGeneric struct {
	id      string
	message string
	options []string
}

func (d AskDialogGeneric) ID() string {
	return d.id
}

func (d AskDialogGeneric) Message() string {
	return d.message
}

func (d AskDialogGeneric) Options() []string {
	return d.options
}

func AskDialogCmd(id string, message string, options []string) tea.Cmd {
	return func() tea.Msg {
		return AskDialogGeneric{
			id:      id,
			message: message,
			options: options,
		}
	}
}
