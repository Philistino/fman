package list

import (
	"bufio"
	"os"
	"time"
	"unicode/utf8"

	"github.com/76creates/stickers"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/nore-dev/fman/entry"
	"github.com/nore-dev/fman/theme"
)

type List struct {
	entries []entry.Entry

	width  int
	height int

	selected_index int
	flexBox        *stickers.FlexBox

	maxEntryToShow int
	truncateLimit  int

	lastClickedTime time.Time
	clickDelay      float64

	theme *theme.Theme

	lastKeyCharacter byte

	lastDirectory string
}

func New(theme *theme.Theme, entries []entry.Entry) List {

	list := List{
		entries:       entries,
		truncateLimit: 100,
		flexBox:       stickers.NewFlexBox(0, 0),
		clickDelay:    0.5,
		theme:         theme,
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
	return list.clearLastKey()
}

func (list *List) SelectedEntry() entry.Entry {

	if len(list.entries) == 0 {
		return entry.Entry{}
	}

	return list.entries[list.selected_index]
}

func (list *List) Theme() *theme.Theme {
	return list.theme
}

func (list *List) Width() int {
	return list.width
}

func (list *List) SetWidth(width int) {
	list.width = width
}

func (list *List) Height() int {
	return list.height
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

func isFileReadable(path string) bool {
	file, err := os.Open(path)

	if err != nil {
		return false
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	return utf8.ValidString(scanner.Text())
}
