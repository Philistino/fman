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
	path  string
	width int
	view  string
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
	return breadcrumb.view
}

func (breadcrumb *Breadcrumb) updateView() {

	// strBuilder := strings.Builder{}

	pathParts := strings.Split(breadcrumb.path, string(filepath.Separator))

	separator := theme.ArrowStyle.Render(string(theme.GetActiveIconTheme().BreadcrumbArrowIcon))

	reverse(pathParts)

	totalLength := 0

	parts := make([]string, 0, len(pathParts))

	for i, part := range pathParts {

		if pathParts[i] == "" {
			continue
		}

		// strBuilder.WriteString()
		partRendered := theme.PathStyle.Render(zone.Mark(strconv.Itoa(i), part))
		if i != 0 {
			partRendered = partRendered + separator
		}
		partWidth := lipgloss.Width(partRendered)

		totalLength += partWidth
		if totalLength > breadcrumb.width+2 { // +2 for left margin
			break
		}
		parts = append(parts, partRendered)
		// if i != len(pathParts)-1 {
		// 	strBuilder.WriteString(separator)
		// }
	}
	reverse(parts)

	breadcrumb.view = lipgloss.NewStyle().MarginLeft(2).Render(strings.Join(parts, ""))
}

func (b *Breadcrumb) SetWidth(width int) {
	b.width = width
	b.updateView()
}

func reverse[S ~[]E, E any](s S) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
