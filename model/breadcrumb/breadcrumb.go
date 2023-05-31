package breadcrumb

import (
	"path/filepath"
	"strconv"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	zone "github.com/lrstanley/bubblezone"
	"github.com/nore-dev/fman/message"
	"github.com/nore-dev/fman/theme"
)

type Breadcrumb struct {
	path      string
	width     int
	viewParts []string
}

func New() *Breadcrumb {
	return &Breadcrumb{}
}

func (breadcrumb *Breadcrumb) Init() tea.Cmd {
	return nil
}

func (breadcrumb *Breadcrumb) Update(msg tea.Msg) (*Breadcrumb, tea.Cmd) {
	switch msg := msg.(type) {
	case message.PathMsg:
		breadcrumb.path = msg.Path
		breadcrumb.updateView()
	case tea.MouseMsg:
		if msg.Type != tea.MouseLeft {
			return breadcrumb, nil
		}

		pathParts := strings.SplitAfter(breadcrumb.path, string(filepath.Separator))

		// Quick Path Jump
		// Mouse Support
		for i := 0; i < len(pathParts); i++ {
			if zone.Get(strconv.Itoa(i)).InBounds(msg) {
				newPath := filepath.Join(pathParts[:i+1]...)
				breadcrumb.path = newPath
				return breadcrumb, message.ChangePath(breadcrumb.path)
			}
		}
	}

	return breadcrumb, nil
}

func (breadcrumb *Breadcrumb) View() string {
	parts := make([]string, 0, len(breadcrumb.viewParts))
	for i, part := range breadcrumb.viewParts {
		parts = append(parts, zone.Mark(strconv.Itoa(i), part))
	}
	return lipgloss.NewStyle().MarginLeft(2).Render(strings.Join(parts, ""))
}

// updateView creates the renderable breadcrumb for the given path
// and updates the view attribute. This could probably be optimized a bit
// but it's only called once per directory change instead of on every
// call to View()
func (breadcrumb *Breadcrumb) updateView() {
	pathParts := strings.Split(breadcrumb.path, string(filepath.Separator))
	separator := theme.ArrowStyle.Render(string(theme.GetActiveIconTheme().BreadcrumbArrowIcon))

	// reverse the parts so we prioritize directories closer to the current
	// directory over ones closer to root
	reverse(pathParts)

	totalLength := 0
	parts := make([]string, 0, len(pathParts))
	for i, part := range pathParts {

		// this was in the old implementation. Is this only for root?
		if pathParts[i] == "" {
			continue
		}

		partRendered := theme.PathStyle.Render(part)
		if i != 0 {
			partRendered = partRendered + separator
		}

		partWidth := lipgloss.Width(partRendered)
		totalLength += partWidth
		if totalLength > breadcrumb.width+2 { // +2 for left margin
			break
		}
		parts = append(parts, partRendered)
	}
	reverse(parts)

	breadcrumb.viewParts = parts
}

// SetWidth sets the max allowable width for the view.
// This should be called on every change of the terminal window width.
// The width is not managed at the Breadcrumb model level because
// the breadcrumb shares the same row with other renderables and the relative
// widths should be managed above the level of this model
func (b *Breadcrumb) SetWidth(width int) {
	b.width = width
	b.updateView()
}

func reverse[S ~[]E, E any](s S) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
