package keys

import "github.com/charmbracelet/bubbles/key"

// Moved this in here temporarily. Should be moved to /model when all components are in the same package

type KeyMap struct {
	Quit                  key.Binding
	MoveCursorUp          key.Binding
	MoveCursorDown        key.Binding
	GoToTop               key.Binding
	GoToBottom            key.Binding
	GoToHomeDirectory     key.Binding
	GoToParentDirectory   key.Binding
	GoToSelectedDirectory key.Binding
	GoBack                key.Binding
	GoForward             key.Binding
	ScrollPreviewDown     key.Binding
	ScrollPreviewUp       key.Binding
	CopyToClipboard       key.Binding
	OpenFile              key.Binding
	ShowHiddenEntries     key.Binding
	ToggleHelp            key.Binding
	MultiSelectUp         key.Binding
	MultiSelectDown       key.Binding
	MultiSelectAll        key.Binding
	MultiSelectToTop      key.Binding
	MultiSelectToBottom   key.Binding
}

var Map = KeyMap{
	Quit: key.NewBinding(
		key.WithKeys("ctrl+q", "esc"),
		key.WithHelp("ctrl+q/esc", "Quit"),
	),
	MoveCursorUp: key.NewBinding(
		key.WithKeys("w", "up"),
		key.WithHelp("w/↑", "Move cursor up"),
	),
	MoveCursorDown: key.NewBinding(
		key.WithKeys("s", "down"),
		key.WithHelp("s/↓", "Move cursor down"),
	),
	GoToTop: key.NewBinding(
		key.WithKeys("ctrl+up"),
		key.WithHelp("ctrl+↑", "Go to Top"),
	),
	GoToBottom: key.NewBinding(
		key.WithKeys("ctrl+down"),
		key.WithHelp("ctrl+↓", "Go to Bottom"),
	),
	GoToHomeDirectory: key.NewBinding(
		key.WithKeys("~"),
		key.WithHelp("~", "Go to Home Directory"),
	),
	GoToParentDirectory: key.NewBinding(
		key.WithKeys("a", "left"),
		key.WithHelp("a/←", "Go to parent directory"),
	),
	GoToSelectedDirectory: key.NewBinding(
		key.WithKeys("d", "right"),
		key.WithHelp("d/→", "Go to selected directory"),
	),
	GoBack: key.NewBinding(
		key.WithKeys("alt+left"),
		key.WithHelp("alt+←", "Go back"),
	),
	GoForward: key.NewBinding(
		key.WithKeys("alt+right"),
		key.WithHelp("alt+→", "Go forward"),
	),
	CopyToClipboard: key.NewBinding(
		key.WithKeys("c"),
		key.WithHelp("c", "Copy to clipboard"),
	),
	OpenFile: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "Open file"),
	),
	ShowHiddenEntries: key.NewBinding(
		key.WithKeys("."),
		key.WithHelp(".", "Show hidden"),
	),
	ScrollPreviewDown: key.NewBinding(
		key.WithKeys("shift+down"),
		key.WithHelp("shift+↓", "Scroll preview down"),
	),
	ScrollPreviewUp: key.NewBinding(
		key.WithKeys("shift+up"),
		key.WithHelp("shift+↑", "Scroll preview up"),
	),
	ToggleHelp: key.NewBinding(
		key.WithKeys("?"),
	),
	MultiSelectUp: key.NewBinding(
		key.WithKeys("shift+up", "shift+k"),
		key.WithHelp("shift+up/shift+k", "multi-select up"),
	),
	MultiSelectDown: key.NewBinding(
		key.WithKeys("shift+down", "shift+j"),
		key.WithHelp("shift+down/shift+j", "multi-select down"),
	),
	MultiSelectToTop: key.NewBinding(
		key.WithKeys("shift+home"),
		key.WithHelp("ctrl+shift+up", "multi-select to top"),
	),
	MultiSelectToBottom: key.NewBinding(
		key.WithKeys("shift+end"),
		key.WithHelp("ctrl+shift+down", "multi-select to bottom"),
	),
	MultiSelectAll: key.NewBinding(
		key.WithKeys("ctrl+a"),
		key.WithHelp("ctrl+a", "select all"),
	),
}

func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{}
}

func (k KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.MoveCursorUp, k.MoveCursorDown, k.ScrollPreviewUp, k.ScrollPreviewDown},
		{k.GoToTop, k.GoToBottom, k.GoBack, k.GoForward},
		{k.GoToHomeDirectory, k.GoToParentDirectory, k.GoToSelectedDirectory},
		{k.OpenFile, k.ShowHiddenEntries, k.CopyToClipboard},
	}
}
