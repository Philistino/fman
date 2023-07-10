package focus

import "testing"

type other struct {
	FocusField
}

func TestFocus(t *testing.T) {
	f := NewFocus(false)
	if f.Focused() {
		t.Error("expected false, got true")
	}

	f.Focus()
	if !f.Focused() {
		t.Error("expected true, got false")
	}

	f.Blur()
	if f.Focused() {
		t.Error("expected false, got true")
	}
}

func TestFocusEmbedded(t *testing.T) {
	o := other{
		FocusField: NewFocus(false),
	}
	if o.Focused() {
		t.Error("expected false, got true")
	}

	o.Focus()
	if !o.Focused() {
		t.Error("expected true, got false")
	}

	o.Blur()
	if o.Focused() {
		t.Error("expected false, got true")
	}
}
