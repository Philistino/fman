package model

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	zone "github.com/lrstanley/bubblezone"
	"github.com/nore-dev/fman/theme"
)

type ToolbarModel struct {
	path string
}

func NewToolbarModel() ToolbarModel {
	return ToolbarModel{}
}

func (toolbar ToolbarModel) Init() tea.Cmd {

	return nil
}

func (toolbar ToolbarModel) Update(msg tea.Msg) (ToolbarModel, tea.Cmd) {

	switch msg := msg.(type) {
	case PathMsg:
		toolbar.path = msg.path

	case tea.MouseMsg:
		if msg.Type != tea.MouseLeft {
			return toolbar, nil
		}

		if zone.Get("forward").InBounds(msg) {
			return toolbar, func() tea.Msg {
				return UpdateEntriesMsg{}
			}
		}

		if zone.Get("back").InBounds(msg) {
			return toolbar, func() tea.Msg {
				return UpdateEntriesMsg{parent: true}
			}
		}

		return toolbar, nil
	}

	return toolbar, nil
}

func (toolbar ToolbarModel) View() string {

	view := lipgloss.JoinHorizontal(lipgloss.Left,
		zone.Mark("back", theme.ButtonStyle.Render("←")),
		zone.Mark("forward", theme.ButtonStyle.Render("→")),
	)

	return lipgloss.JoinHorizontal(lipgloss.Center, view, toolbar.path)
}
