package keys

import (
	"strings"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/lipgloss"
)

type KeyMap struct {
	Quit              key.Binding
	ToggleHelp        key.Binding
	ShowHiddenEntries key.Binding
	OpenFile          key.Binding

	MoveCursorUp       key.Binding
	MoveCursorDown     key.Binding
	MoveCursorToTop    key.Binding
	MoveCursorToBottom key.Binding

	MultiSelectUp       key.Binding
	MultiSelectDown     key.Binding
	MultiSelectToTop    key.Binding
	MultiSelectToBottom key.Binding
	MultiSelectAll      key.Binding

	GoToParentDirectory   key.Binding
	GoToSelectedDirectory key.Binding
	GoToHomeDirectory     key.Binding
	GoBack                key.Binding
	GoForward             key.Binding

	ScrollPreviewDown key.Binding
	ScrollPreviewUp   key.Binding

	CopyToClipboard key.Binding

	width  int
	height int
}

var Map = KeyMap{
	Quit: key.NewBinding(
		key.WithKeys("ctrl+q"),
		key.WithHelp("ctrl+q", "Quit"),
	),
	MoveCursorUp: key.NewBinding(
		key.WithKeys("w", "up"),
		key.WithHelp("↑/w", "Move cursor up"),
	),
	MoveCursorDown: key.NewBinding(
		key.WithKeys("s", "down"),
		key.WithHelp("↓/s", "Move cursor down"),
	),
	MoveCursorToTop: key.NewBinding(
		key.WithKeys("home"),
		key.WithHelp("home", "Move cursor to top"),
	),
	MoveCursorToBottom: key.NewBinding(
		key.WithKeys("end"),
		key.WithHelp("end", "Move cursor to bottom"),
	),
	GoToHomeDirectory: key.NewBinding(
		key.WithKeys("~"),
		key.WithHelp("~", "Go to home folder"),
	),
	GoToParentDirectory: key.NewBinding(
		key.WithKeys("a", "left"),
		key.WithHelp("←/a", "Go to parent folder"),
	),
	GoToSelectedDirectory: key.NewBinding(
		key.WithKeys("d", "right"),
		key.WithHelp("→/d", "Go to selected folder"),
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
		key.WithHelp(".", "Toggle show hidden"),
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
		key.WithHelp("?", "Toggle help"),
	),
	MultiSelectUp: key.NewBinding(
		key.WithKeys("shift+up", "shift+k"),
		key.WithHelp("shift+up/shift+k", "Multi-select up"),
	),
	MultiSelectDown: key.NewBinding(
		key.WithKeys("shift+down", "shift+j"),
		key.WithHelp("shift+down/shift+j", "Multi-select down"),
	),
	MultiSelectToTop: key.NewBinding(
		key.WithKeys("shift+home"),
		key.WithHelp("shift+home", "Multi-select to top"),
	),
	MultiSelectToBottom: key.NewBinding(
		key.WithKeys("shift+end"),
		key.WithHelp("shift+end", "Multi-select to bottom"),
	),
	MultiSelectAll: key.NewBinding(
		key.WithKeys("ctrl+a"),
		key.WithHelp("ctrl+a", "Select all"),
	),
}

func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{}
}

func (k KeyMap) FullHelp() [][]key.Binding {
	chunk1 := []key.Binding{k.Quit, k.ToggleHelp, k.ShowHiddenEntries, k.OpenFile}
	chunk2 := []key.Binding{k.MoveCursorUp, k.MoveCursorDown, k.MoveCursorToTop, k.MoveCursorToBottom}
	// chunk3 := []key.Binding{k.MultiSelectUp, k.MultiSelectDown, k.MultiSelectToTop, k.MultiSelectToBottom, k.MultiSelectAll}
	chunk4 := []key.Binding{k.GoToParentDirectory, k.GoToSelectedDirectory, k.GoToHomeDirectory, k.GoBack, k.GoForward}
	chunk5 := []key.Binding{k.ScrollPreviewUp, k.ScrollPreviewDown}

	return [][]key.Binding{
		chunk1,
		chunk2,
		chunk4,
		chunk5,
	}
}

func (k *KeyMap) SetSize(width int, height int) {
	k.width = width
	k.height = height
}

func (k KeyMap) Width() int {
	return k.width
}

func (k KeyMap) ViewHelp() string {
	chunk1 := []key.Binding{k.Quit, k.ToggleHelp, k.ShowHiddenEntries, k.OpenFile}
	chunk2 := []key.Binding{k.MoveCursorUp, k.MoveCursorDown, k.MoveCursorToTop, k.MoveCursorToBottom}
	chunk3 := []key.Binding{k.MultiSelectAll, k.MultiSelectUp, k.MultiSelectDown, k.MultiSelectToTop, k.MultiSelectToBottom}
	chunk4 := []key.Binding{k.GoToParentDirectory, k.GoToSelectedDirectory, k.GoToHomeDirectory, k.GoBack, k.GoForward}
	chunk5 := []key.Binding{k.ScrollPreviewUp, k.ScrollPreviewDown}

	groups := [][]key.Binding{chunk1, chunk2, chunk3, chunk4, chunk5}

	boxes := make([]string, 0, len(groups))
	for _, group := range groups {

		var (
			keys         []string
			descriptions []string
		)

		// Separate keys and descriptions into different slices
		for _, kb := range group {
			if !kb.Enabled() {
				continue
			}
			keys = append(keys, kb.Help().Key)
			descriptions = append(descriptions, kb.Help().Desc)
		}

		// Join the keys and descriptions into a single column
		col := lipgloss.JoinHorizontal(lipgloss.Top,
			lipgloss.NewStyle().Render(strings.Join(keys, "\n")),
			lipgloss.NewStyle().Render("  "),
			lipgloss.NewStyle().Render(strings.Join(descriptions, "\n")),
		)
		boxes = append(boxes, col)
	}

	// This approach could definitely be optimized but it works and it is rarely called
	// loop continuously combining columns until the width is less than the given width
	iteration := 0
	for {
		// define a slice of strings to hold the columns
		innerBoxes := []string{}
		var currentCol string
		j := 0 // j counts up to the number of iterations to know how many boxes to combine into each column
		for i, box := range boxes {

			// if we are on the first "row", set the current column to the current box
			// otherwise, append current box to the current column with a line between
			if j == 0 {
				currentCol = box
			} else {
				currentCol = currentCol + "\n\n" + box
			}

			// if we have reached the number of iterations, add the current column, which has been built up,
			// to the slice of columns with some horizontal separation (unless it is the last column)
			if j == iteration {
				horizontalSep := "   "
				if i == len(boxes)-1 {
					horizontalSep = ""
				}
				innerBoxes = append(innerBoxes, currentCol, horizontalSep)
				currentCol = ""
				j = 0
				continue
			}

			// if we reach the end with a box remaining, add it as a final column
			if i == len(boxes)-1 {
				innerBoxes = append(innerBoxes, currentCol)
			}
			j++
		}

		// if the width of the columns is fits in the given width or we only have one column, return the columns
		rendered := lipgloss.JoinHorizontal(lipgloss.Top, innerBoxes...)
		if lipgloss.Width(rendered) <= k.width || len(innerBoxes) == 1 {
			return rendered
		}
		iteration++
	}
}
