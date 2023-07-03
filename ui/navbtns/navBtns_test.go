package navbtns

import (
	"strings"
	"testing"

	"github.com/Philistino/fman/ui/theme"
	zone "github.com/lrstanley/bubblezone"
)

// Test Update function with ActiveNavBtns message
type activeMsg struct {
	backActive    bool
	forwardActive bool
	upActive      bool
}

func (msg activeMsg) BackActive() bool {
	return msg.backActive
}
func (msg activeMsg) ForwardActive() bool {
	return msg.forwardActive
}
func (msg activeMsg) UpActive() bool {
	return msg.upActive
}

func TestView(t *testing.T) {
	zone.NewGlobal()
	navBtns := NewNavBtns()

	// Test Init function
	cmd := navBtns.Init()
	if cmd != nil {
		t.Errorf("Init() should return nil, got %v", cmd)
	}

	// // Test View function
	view := navBtns.View()
	icons := theme.GetActiveIconTheme()
	backBtn := string(icons.LeftArrowIcon)
	fwdBtn := string(icons.RightArrowIcon)
	upBtn := string(icons.UpArrowIcon)
	for _, btn := range []string{backBtn, fwdBtn, upBtn} {
		if !strings.Contains(view, btn) {
			t.Errorf("View() should return a string containing %s, got %s", btn, view)
		}
	}

	// Test Update function
	msg := activeMsg{
		backActive:    true,
		forwardActive: true,
		upActive:      true,
	}
	updatedNavBtns, cmd := navBtns.Update(msg)
	if updatedNavBtns.backActive != true || updatedNavBtns.fwdActive != true || updatedNavBtns.upActive != true {
		t.Errorf("Update() should update backActive, fwdActive, and upActive to true, got backActive=%v, fwdActive=%v, upActive=%v", updatedNavBtns.backActive, updatedNavBtns.fwdActive, updatedNavBtns.upActive)
	}
	if cmd != nil {
		t.Errorf("Update() should return nil, got %v", cmd)
	}

	// Test View function again
	view = navBtns.View()
	for _, btn := range []string{backBtn, fwdBtn, upBtn} {
		if !strings.Contains(view, btn) {
			t.Errorf("View() should return a string containing %s, got %s", btn, view)
		}
	}
}

func TestNavBtns(t *testing.T) {
	zone.NewGlobal()
	navBtns := NewNavBtns()

	// Test Init function
	cmd := navBtns.Init()
	if cmd != nil {
		t.Errorf("Init() should return nil, got %v", cmd)
	}

	msg := activeMsg{
		backActive:    true,
		forwardActive: true,
		upActive:      true,
	}
	updatedNavBtns, cmd := navBtns.Update(msg)
	if updatedNavBtns.backActive != true || updatedNavBtns.fwdActive != true || updatedNavBtns.upActive != true {
		t.Errorf("Update() should update backActive, fwdActive, and upActive to true, got backActive=%v, fwdActive=%v, upActive=%v", updatedNavBtns.backActive, updatedNavBtns.fwdActive, updatedNavBtns.upActive)
	}
	if cmd != nil {
		t.Errorf("Update() should return nil, got %v", cmd)
	}

}
