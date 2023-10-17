package infobar

import (
	"fmt"
	"log"

	"github.com/Philistino/fman/ui/message"
	"github.com/Philistino/fman/ui/theme"
	tea "github.com/charmbracelet/bubbletea"
)

type itemTracker struct {
	nItems    int
	nSelected int
}

func (m itemTracker) Init() tea.Cmd {
	return nil
}

func (m itemTracker) Update(msg tea.Msg) (itemTracker, tea.Cmd) {
	switch msg := msg.(type) {
	case message.DirChangedMsg:
		m.nItems = len(msg.Entries())
		m.nSelected = len(msg.Selected())
	case message.SelectedMsg:
		log.Println("selected msg")
		m.nSelected = len(msg.Selected())
	}
	return m, nil
}

func (m itemTracker) View() string {
	style := theme.InfobarStyle.Copy().UnsetWidth().Padding(0, 1)
	items := fmt.Sprintf("%d items", m.nItems)
	var selected string
	if m.nSelected > 0 {
		selected = fmt.Sprintf("%d selected │", m.nSelected)
	}
	return style.Render(selected, items)
}
