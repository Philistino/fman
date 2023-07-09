package navbtns

import (
	"github.com/Philistino/fman/ui/message"
	"github.com/Philistino/fman/ui/theme"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	zone "github.com/lrstanley/bubblezone"
)

type ActiveNavBtns interface {
	BackActive() bool
	ForwardActive() bool
	UpActive() bool
}

// NavBtns handle the back/forward/up navigation buttons in the toolbar
type NavBtns struct {
	backActive bool
	fwdActive  bool
	upActive   bool
	focused    bool
}

func NewNavBtns() *NavBtns {
	return &NavBtns{focused: true}
}

func (toolbar *NavBtns) Init() tea.Cmd {
	return nil
}

func (toolbar *NavBtns) Update(msg tea.Msg) (*NavBtns, tea.Cmd) {
	if !toolbar.focused {
		return toolbar, nil
	}

	_, ok := msg.(ActiveNavBtns)
	if ok {
		msg := msg.(ActiveNavBtns)
		toolbar.backActive = msg.BackActive()
		toolbar.fwdActive = msg.ForwardActive()
		toolbar.upActive = msg.UpActive()
		return toolbar, nil
	}

	switch msg := msg.(type) {
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

func (toolbar *NavBtns) View() string {

	var backBtn, fwdBtn, upBtn string

	icons := theme.GetActiveIconTheme()

	if toolbar.backActive && toolbar.focused {
		backBtn = zone.Mark("back", theme.ButtonStyle.Render(string(icons.LeftArrowIcon)))
	} else {
		backBtn = zone.Mark("back", theme.InactiveButtonStyle.Render(string(icons.LeftArrowIcon)))
	}

	if toolbar.fwdActive && toolbar.focused {
		fwdBtn = zone.Mark("forward", theme.ButtonStyle.Render(string(icons.RightArrowIcon)))
	} else {
		fwdBtn = zone.Mark("forward", theme.InactiveButtonStyle.Render(string(icons.RightArrowIcon)))
	}

	if toolbar.upActive && toolbar.focused {
		upBtn = zone.Mark("up", theme.ButtonStyle.Render(string(icons.UpArrowIcon)))
	} else {
		upBtn = zone.Mark("up", theme.InactiveButtonStyle.Render(string(icons.UpArrowIcon)))
	}

	return lipgloss.JoinHorizontal(
		lipgloss.Left,
		backBtn,
		fwdBtn,
		upBtn,
	)
}

// Blur unfocuses the toolbar
func (toolbar *NavBtns) Blur() {
	toolbar.focused = false
}

// Focus focuses the toolbar
func (toolbar *NavBtns) Focus() {
	toolbar.focused = true
}
