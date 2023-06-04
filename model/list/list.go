package list

import (
	"time"

	"github.com/76creates/stickers"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/nore-dev/fman/entry"
	"github.com/nore-dev/fman/theme"
)

type List struct {
	entries []entry.Entry

	width  int
	height int

	cursorIdx int
	selected  map[int]struct{}

	flexBox *stickers.FlexBox

	maxEntryToShow int
	truncateLimit  int

	lastClickedTime time.Time
	clickDelay      float64

	theme theme.Theme

	lastKeyCharacter byte
	focused          bool
}

func New(theme theme.Theme) List {

	list := List{
		truncateLimit: 100,
		flexBox:       stickers.NewFlexBox(0, 0),
		clickDelay:    0.5,
		theme:         theme,
		focused:       true,
	}

	rows := []*stickers.FlexBoxRow{
		list.flexBox.NewRow().AddCells(
			[]*stickers.FlexBoxCell{
				stickers.NewFlexBoxCell(5, 1),
				stickers.NewFlexBoxCell(2, 1),
				stickers.NewFlexBoxCell(3, 1),
			},
		),
	}

	list.flexBox.AddRows(rows)

	return list
}

func (list *List) Init() tea.Cmd {
	return nil
}

func (list *List) SelectedEntry() entry.Entry {
	if len(list.entries) == 0 {
		return entry.Entry{}
	}
	return list.entries[list.cursorIdx]
}

// TODO: Change this when reimplementing the list
func (list *List) SelectedEntries() map[string]struct{} {
	if len(list.entries) == 0 {
		return nil
	}
	return map[string]struct{}{
		list.SelectedEntry().Name(): {},
	}
}

func (list *List) CursorName() string {
	return list.SelectedEntry().Name()
}

func (list *List) SetWidth(width int) {
	list.width = width
}

func (list *List) SetHeight(height int) {
	list.height = height
}

func (list *List) IsEmpty() bool {
	return len(list.entries) == 0
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func truncateText(str string, max int) string {
	// "hello world" -> "hello wo..."

	if len(str) > max {
		return str[:max-3] + "..."
	}

	return str
}

// Focused returns the focus state of the table.
func (m *List) Focused() bool {
	return m.focused
}

// Focus focuses the table, allowing the user to move around the rows and
// interact.
func (m *List) Focus() {
	m.focused = true
}

// Blur blurs the table, preventing selection or movement.
func (m *List) Blur() {
	m.focused = false
}
