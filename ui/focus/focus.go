package focus

// FocusField is a struct that holds a boolean field.
// It is used as an embedded struct to determine if a component is focused.
type FocusField struct {
	focused bool
}

// NewFocus returns a new Focus struct.
func NewFocus(focus bool) FocusField {
	return FocusField{
		focused: focus,
	}
}

// Focused returns the focused field.
func (f *FocusField) Focused() bool {
	return f.focused
}

// Focus sets the focused field to true.
func (f *FocusField) Focus() {
	f.focused = true
}

// Blur sets the focused field to false.
func (f *FocusField) Blur() {
	f.focused = false
}

// ToggleFocus toggles the focused field.
func (f *FocusField) ToggleFocus() {
	f.focused = !f.focused
}
