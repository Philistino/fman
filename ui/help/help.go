package help

import (
	"strings"

	"github.com/Philistino/fman/ui/focus"
	"github.com/Philistino/fman/ui/keys"
	"github.com/Philistino/fman/ui/theme/colors"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Help represents the help view.
type Help struct {
	focus.FocusField
	theme    colors.Theme
	keys     keys.KeyMap
	viewport viewport.Model
}

// New creates a new Help view.
func New(theme colors.Theme, keys keys.KeyMap, viewportStyle lipgloss.Style) Help {
	viewport := viewport.New(0, 0)
	viewport.Style = viewportStyle
	return Help{
		theme:    theme,
		keys:     keys,
		viewport: viewport,
	}
}

// Init initializes the Help model.
func Init() tea.Cmd {
	return nil
}

// Update updates the Help view.
func (h Help) Update(msg tea.Msg) (Help, tea.Cmd) {
	if !h.Focused() {
		return h, nil
	}
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case msg.String() == "up":
			h.viewport.LineUp(1)
		case msg.String() == "down":
			h.viewport.LineDown(1)
		}
	}
	h.viewport, cmd = h.viewport.Update(msg)
	return h, cmd
}

// View returns the Help view.
func (h Help) View() string {
	return h.viewport.View()
}

func (h *Help) setViewPortContent() {
	content := h.createView()
	h.viewport.SetContent(content)
	if h.viewport.TotalLineCount() <= h.viewport.VisibleLineCount() {
		h.viewport.Height = h.viewport.Style.GetVerticalPadding() + lipgloss.Height(content)
	}
}

// ViewHelp returns the Help view with the key bindings and descriptions.
func (h Help) createView() string {

	groups := h.keys.FullHelp()

	// Create a slice of text boxes, one for each group of key bindings
	boxes := make([]string, 0, len(groups))
	for _, group := range groups {
		boxes = append(boxes, renderGroup(group))
	}
	content := arrangeGroups(boxes, h.viewport.Width-h.viewport.Style.GetHorizontalPadding())
	return lipgloss.NewStyle().Foreground(h.theme.TextColor).Render(content)
}

// renderGroup takes a slice of key bindings and returns a string with the keys and descriptions
// of the enabled bindings. The keys and descriptions are joined horizontally with a separator.
func renderGroup(group []key.Binding) string {
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

	return lipgloss.JoinHorizontal(lipgloss.Top,
		strings.Join(keys, "\n"),
		"  ",
		strings.Join(descriptions, "\n"),
	)
}

// arrangeGroups takes a slice of text boxes and arranges them into columns and rows
// until the width is less than the given width or we have a single column.
// This function returns a string with the arranged text boxes.
func arrangeGroups(boxes []string, width int) string {

	// Loop continuously combining columns into rows
	// until the width is less than the given width or we have a single column
	//
	// This approach could definitely be optimized but it works and it is rarely called.
	iteration := 0
	for {
		// define a slice of strings to hold the columns
		innerBoxes := []string{}
		var currentCol strings.Builder
		j := 0 // j counts up to the number of iterations to know how many boxes to combine into each column
		for i, box := range boxes {

			// if we are on the first "row", set the current column to the current box
			// otherwise, append current box to the current column with a line between
			if j == 0 {
				currentCol.WriteString(box)
			} else {
				currentCol.WriteString("\n\n")
				currentCol.WriteString(box)
			}

			// if we have reached the number of iterations, add the current column, which has been built up,
			// to the slice of columns with some horizontal separation (unless it is the last column)
			if j == iteration {
				horizontalSep := "   "
				if i == len(boxes)-1 {
					horizontalSep = ""
				}
				innerBoxes = append(innerBoxes, currentCol.String(), horizontalSep)
				currentCol.Reset()
				j = 0
				continue
			}

			// if we reach the end with box(es) remaining, add as a final column
			if i == len(boxes)-1 {
				innerBoxes = append(innerBoxes, currentCol.String())
			}
			j++
		}

		// if the width of the columns is fits in the given width or we only have one column, return the columns
		rendered := lipgloss.JoinHorizontal(lipgloss.Top, innerBoxes...)
		if lipgloss.Width(rendered) <= width || len(innerBoxes) == 1 {
			return rendered
		}
		iteration++
	}
}

// SetSize sets the height and width of the Help view and updates the viewport.
func (h *Help) SetSize(height, width int) {
	h.viewport.Height = height
	h.viewport.Width = width
	h.setViewPortContent()
}
