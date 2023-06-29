package list

import (
	"time"

	"github.com/76creates/stickers"
	"github.com/Philistino/fman/entry"
	"github.com/Philistino/fman/theme/colors"
	tea "github.com/charmbracelet/bubbletea"
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
	lastClickedIdx  int // list index of the last clicked item. Must be reset to -1 when the list is updated
	clickDelay      time.Duration

	theme colors.Theme

	lastKeyCharacter byte
	focused          bool
}

func New(theme colors.Theme, doubleClickDelay int) List {

	list := List{
		entries:          []entry.Entry{},
		width:            0,
		height:           0,
		cursorIdx:        0,
		selected:         map[int]struct{}{},
		flexBox:          stickers.NewFlexBox(0, 0),
		maxEntryToShow:   0,
		truncateLimit:    100,
		lastClickedTime:  time.Time{},
		lastClickedIdx:   -1,
		clickDelay:       time.Duration(time.Millisecond * time.Duration(doubleClickDelay)),
		theme:            theme,
		lastKeyCharacter: 0,
		focused:          true,
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

// THIS CAN CAUSE A PANIC IF THE LIST IS EMPTY
// The zero value of an entry is has a nil pointer for fs.FileInfo
func (list *List) SelectedEntry() entry.Entry {
	if len(list.entries) == 0 {
		return entry.Entry{}
	}
	return list.entries[list.cursorIdx]
}

func (list *List) SelectedEntryName() string {
	if len(list.entries) == 0 {
		return ""
	}
	return list.entries[list.cursorIdx].Name()
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
	return list.SelectedEntryName()
}

func (list *List) SetWidth(width int) {
	list.width = width
	list.flexBox.SetWidth(width)
}

func (list *List) SetHeight(height int) {
	list.height = height
	list.flexBox.SetHeight(height)
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
