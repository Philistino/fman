package breadcrumb

import (
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/Philistino/fman/ui/message"
	"github.com/Philistino/fman/ui/theme"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	zone "github.com/lrstanley/bubblezone"
)

const pathSeparator = string(filepath.Separator)

var winRootRgx = regexp.MustCompile(`^[A-Za-z]:/?$`) // matches windows root paths like C:/ or D:

// PathError is an interface that can be used to return a path and an error
type PathError interface {
	Path() string
	Error() error
}

// BreadCrumb is the model for the path breadcrumb
type BreadCrumb struct {
	path      string
	width     int
	viewParts []string
	focused   bool
}

// newBreadCrumb creates a new breadcrumb
func NewBreadCrumb() *BreadCrumb {
	return &BreadCrumb{focused: true}
}

// Init initializes the model
func (breadcrumb *BreadCrumb) Init() tea.Cmd {
	return nil
}

// Update updates the model
func (breadcrumb *BreadCrumb) Update(msg tea.Msg) (*BreadCrumb, tea.Cmd) {
	if !breadcrumb.focused {
		return breadcrumb, nil
	}

	_, ok := msg.(PathError)
	if ok {
		msg := msg.(PathError)
		if msg.Error() != nil {
			return breadcrumb, nil
		}
		breadcrumb.path = msg.Path()
		breadcrumb.updateView(msg.Path())
		return breadcrumb, nil
	}

	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.MouseMsg:
		cmd = breadcrumb.handleMouseMsg(msg)
	}
	return breadcrumb, cmd
}

// View renders the model
func (breadcrumb *BreadCrumb) View() string {
	parts := make([]string, 0, len(breadcrumb.viewParts))
	for i, part := range breadcrumb.viewParts {
		parts = append(parts, zone.Mark(strconv.Itoa(i), part))
	}
	return lipgloss.NewStyle().MarginLeft(2).Render(strings.Join(parts, ""))
}

// handleMouseMsg handles mouse clicks on the breadcrumb
func (breadcrumb *BreadCrumb) handleMouseMsg(msg tea.MouseMsg) tea.Cmd {
	if msg.Type != tea.MouseLeft {
		return nil
	}
	pathParts := strings.SplitAfter(breadcrumb.path, pathSeparator) // TODO: does split after work here on unix?
	clicked := false
	var viewPartClicked int
	for i := 0; i < len(breadcrumb.viewParts); i++ {
		if !zone.Get(strconv.Itoa(i)).InBounds(msg) {
			continue
		}
		viewPartClicked = i
		clicked = true
		break
	}
	if !clicked {
		return nil
	}
	if len(breadcrumb.viewParts) < len(pathParts) {
		viewPartClicked = len(pathParts) - len(breadcrumb.viewParts) + viewPartClicked
	}
	path := filepath.Join(pathParts[:viewPartClicked+1]...)
	return message.NavOtherCmd(path)
}

// updateView creates the renderable breadcrumb for the given path
// and updates the view attribute. This could probably be optimized a bit
// but it's only called once per directory change instead of on every
// call to View()
func (breadcrumb *BreadCrumb) updateView(path string) {

	// if the path is a root path, just return the root rendered
	if winRootRgx.MatchString(path) {
		breadcrumb.viewParts = []string{theme.PathStyle.Render(strings.Replace(path, pathSeparator, "", 1))}
		return
	}
	if path == pathSeparator {
		breadcrumb.viewParts = []string{theme.PathStyle.Render(path)}
		return
	}

	pathParts := strings.Split(path, pathSeparator)

	if strings.HasPrefix(path, pathSeparator) {
		pathParts = append([]string{pathSeparator}, pathParts...)
	}

	separator := theme.ArrowStyle.Render(string(theme.GetActiveIconTheme().BreadcrumbArrowIcon))

	// reverse the parts so we prioritize directories closer to the current
	// directory over ones closer to root
	reverse(pathParts)

	totalLength := 0
	parts := make([]string, 0, len(pathParts))
	for i, part := range pathParts {

		if part == "" {
			continue
		}

		partRendered := theme.PathStyle.Render(part)
		if i != 0 {
			partRendered = partRendered + separator
		}

		partWidth := lipgloss.Width(partRendered)
		totalLength += partWidth
		if totalLength+12 > breadcrumb.width { // +12 seems to be a magic number
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
func (b *BreadCrumb) SetWidth(width int) {
	b.width = width
	b.updateView(b.path)
}

// reverse reverses a slice
func reverse[S ~[]E, E any](s S) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// Blur unfocuses the breadcrumb
func (breadcrumb *BreadCrumb) Blur() {
	breadcrumb.focused = false
}

// Focus focuses the breadcrumb
func (breadcrumb *BreadCrumb) Focus() {
	breadcrumb.focused = true
}
