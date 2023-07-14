package keys

import (
	"strings"

	"github.com/Philistino/fman/ui/theme/colors"
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
		key.WithKeys("up"),
		key.WithHelp("↑", "Move cursor up"),
	),
	MoveCursorDown: key.NewBinding(
		key.WithKeys("down"),
		key.WithHelp("↓", "Move cursor down"),
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
		key.WithKeys("left"),
		key.WithHelp("←", "Go to parent folder"),
	),
	GoToSelectedDirectory: key.NewBinding(
		key.WithKeys("right"),
		key.WithHelp("→", "Go to selected folder"),
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
		key.WithKeys("shift+up"),
		key.WithHelp("shift+↑", "Multi-select up"),
	),
	MultiSelectDown: key.NewBinding(
		key.WithKeys("shift+down"),
		key.WithHelp("shift+↓", "Multi-select down"),
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
	return nil
}

func (k KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Quit, k.ToggleHelp, k.ShowHiddenEntries, k.OpenFile},
		{k.MoveCursorUp, k.MoveCursorDown, k.MoveCursorToTop, k.MoveCursorToBottom},
		{k.GoToParentDirectory, k.GoToSelectedDirectory, k.GoToHomeDirectory, k.GoBack, k.GoForward},
		{k.MultiSelectAll, k.MultiSelectUp, k.MultiSelectDown, k.MultiSelectToTop, k.MultiSelectToBottom},
		{k.ScrollPreviewUp, k.ScrollPreviewDown},
	}
}

func (k *KeyMap) SetSize(width int, height int) {
	k.width = width
	k.height = height
}

func (k KeyMap) Width() int {
	return k.width
}

// renderGroup renders a group of key bindings as a help box string.
// It takes a slice of key bindings as input and returns a string.
func (k KeyMap) renderGroup(group []key.Binding) string {
	var (
		keys         []string
		descriptions []string
	)

	for _, kb := range group {
		if !kb.Enabled() {
			continue
		}
		keys = append(keys, kb.Help().Key)
		descriptions = append(descriptions, kb.Help().Desc)
	}

	// Calculate the width of the keys column
	var (
		maxKeyWidth int
		keyColWidth int
	)
	for _, key := range keys {
		if len(key) > maxKeyWidth {
			maxKeyWidth = len(key)
		}
	}
	if maxKeyWidth > 0 {
		keyColWidth = maxKeyWidth + 2
	}

	// Calculate the width of the descriptions column
	var (
		maxDescWidth int
		descColWidth int
	)
	for _, desc := range descriptions {
		if len(desc) > maxDescWidth {
			maxDescWidth = len(desc)
		}
	}
	if maxDescWidth > 0 {
		descColWidth = maxDescWidth + 2
	}

	// Calculate the width of the help box
	var (
		helpBoxWidth int
	)
	if keyColWidth > 0 && descColWidth > 0 {
		helpBoxWidth = keyColWidth + descColWidth
	}

	// Render the keys and descriptions
	var (
		keysStr         string
		descriptionsStr string
	)
	for i, key := range keys {
		desc := descriptions[i]
		if keyColWidth > 0 {
			key = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF00FF")).Render(key)
			key = lipgloss.NewStyle().Width(keyColWidth).Render(key)
			keysStr += key
		}
		if descColWidth > 0 {
			desc = lipgloss.NewStyle().Width(descColWidth).Render(desc)
			descriptionsStr += desc
		}
	}

	// Render the help box
	var (
		helpBoxStr string
	)
	if helpBoxWidth > 0 {
		helpBoxStr = lipgloss.NewStyle().Width(helpBoxWidth).Render(keysStr + descriptionsStr)
	} else {
		helpBoxStr = ""
	}

	return helpBoxStr
}

func (k KeyMap) ViewHelp(colors colors.Theme) string {

	groups := [][]key.Binding{
		{k.Quit, k.ToggleHelp, k.ShowHiddenEntries, k.OpenFile},
		{k.MoveCursorUp, k.MoveCursorDown, k.MoveCursorToTop, k.MoveCursorToBottom},
		{k.GoToParentDirectory, k.GoToSelectedDirectory, k.GoToHomeDirectory, k.GoBack, k.GoForward},
		{k.MultiSelectAll, k.MultiSelectUp, k.MultiSelectDown, k.MultiSelectToTop, k.MultiSelectToBottom},
		{k.ScrollPreviewUp, k.ScrollPreviewDown},
	}

	// Create a slice of text boxes, one for each chunk
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
			// lipgloss.NewStyle().Render(strings.Join(keys, "\n")),
			// lipgloss.NewStyle().Render("  "),
			// lipgloss.NewStyle().Render(strings.Join(descriptions, "\n")),
			strings.Join(keys, "\n"),
			"  ",
			strings.Join(descriptions, "\n"),
		)
		boxes = append(boxes, col)
	}

	// Loop continuously combining columns into rows
	// until the width is less than the given width or we have a single column
	//
	// This approach could definitely be optimized but it works and it is rarely called.
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
		rendered := lipgloss.NewStyle().Foreground(colors.TextColor).Render(lipgloss.JoinHorizontal(lipgloss.Top, innerBoxes...))
		if lipgloss.Width(rendered) <= k.width || len(innerBoxes) == 1 {
			return rendered
		}
		iteration++
	}
}
