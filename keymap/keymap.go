package keymap

import "github.com/charmbracelet/bubbles/key"

type KeyMap struct {
	MoveCursorUp   key.Binding
	MoveCursorDown key.Binding

	GoToTop    key.Binding
	GoToBottom key.Binding

	GoToHomeDirectory key.Binding

	GoToParentDirectory   key.Binding
	GoToSelectedDirectory key.Binding

	GoBack    key.Binding
	GoForward key.Binding

	ScrollPreviewDown key.Binding
	ScrollPreviewUp   key.Binding

	CopyToClipboard key.Binding

	OpenFile key.Binding

	ShowHiddenEntries key.Binding

	ToggleHelp key.Binding
}

var Default = KeyMap{
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
		key.WithKeys("m"),
		key.WithHelp("m", "Show hidden entries"),
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
