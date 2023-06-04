package model

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	zone "github.com/lrstanley/bubblezone"
	"github.com/nore-dev/fman/message"
	"github.com/nore-dev/fman/theme"
)

// navBtns handle the back/forward/up navigation buttons in the toolbar

type navBtns struct {
	backActive bool
	fwdActive  bool
	upActive   bool
}

func newNavBtns() *navBtns {
	return &navBtns{}
}

func (toolbar *navBtns) Init() tea.Cmd {
	return nil
}

func (toolbar *navBtns) Update(msg tea.Msg) (*navBtns, tea.Cmd) {
	switch msg := msg.(type) {
	case message.DirChangedMsg:
		toolbar.backActive = msg.BackActive()
		toolbar.fwdActive = msg.ForwardActive()
		toolbar.upActive = msg.UpActive()
	case tea.MouseMsg:
		if msg.Type != tea.MouseLeft {
			return toolbar, nil
		}
		if zone.Get("forward").InBounds(msg) && toolbar.fwdActive {
			return toolbar, message.NavFwdCmd()
		}
		if zone.Get("back").InBounds(msg) && toolbar.backActive {
			return toolbar, message.NavBackCmd()
		}
		if zone.Get("up").InBounds(msg) && toolbar.upActive {
			return toolbar, message.NavUpCmd()
		}
	}
	return toolbar, nil
}

func (toolbar *navBtns) View() string {

	var backBtn, fwdBtn, upBtn string

	icons := theme.GetActiveIconTheme()

	if toolbar.backActive {
		backBtn = zone.Mark("back", theme.ButtonStyle.Render(string(icons.LeftArrowIcon)))
	} else {
		backBtn = zone.Mark("back", theme.InactiveButtonStyle.Render(string(icons.LeftArrowIcon)))
	}

	if toolbar.fwdActive {
		fwdBtn = zone.Mark("forward", theme.ButtonStyle.Render(string(icons.RightArrowIcon)))
	} else {
		fwdBtn = zone.Mark("forward", theme.InactiveButtonStyle.Render(string(icons.RightArrowIcon)))
	}

	if toolbar.upActive {
		upBtn = zone.Mark("up", theme.ButtonStyle.Render(string(icons.UpArrowIcon)))
	} else {
		upBtn = zone.Mark("up", theme.InactiveButtonStyle.Render(string(icons.UpArrowIcon)))
	}

	return lipgloss.JoinHorizontal(lipgloss.Left,
		backBtn,
		fwdBtn,
		upBtn,
	)
}

// func (toolbar *navBtns) SetWidth(width int) {
// 	toolbar.breadcrumb.SetWidth(width - lipgloss.Width(theme.ButtonStyle.Render(string(theme.GetActiveIconTheme().LeftArrowIcon))))
// }
