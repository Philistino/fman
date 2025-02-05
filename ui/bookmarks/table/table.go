package table

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	zone "github.com/lrstanley/bubblezone"
)

type tableSort struct {
	col int
	asc bool
}

// Table defines a state for the table widget.
type Table struct {
	KeyMap       KeyMap
	emptyMessage string

	cols     []Column
	rows     []Row
	cursor   int
	focus    bool
	styles   Styles
	selected map[int]struct{}

	start           int
	end             int
	height          int
	width           int
	zoneMgr         *zone.Manager
	lastClickedTime time.Time
	lastClickedIdx  int // list index of the last clicked item. Must be reset to -1 when the list is updated
	clickDelay      time.Duration

	sort tableSort
}

func (m *Table) Init() tea.Cmd {
	return nil
}

// Row represents one line in the table.
type Row []string

// Column defines the table structure.
type Column struct {
	Title string
	Width int
}

// KeyMap defines keybindings. It satisfies to the help.KeyMap interface, which
// is used to render the menu.
type KeyMap struct {
	LineUp              key.Binding
	LineDown            key.Binding
	PageUp              key.Binding
	PageDown            key.Binding
	HalfPageUp          key.Binding
	HalfPageDown        key.Binding
	GotoTop             key.Binding
	GotoBottom          key.Binding
	MultiSelectUp       key.Binding
	MultiSelectDown     key.Binding
	MultiSelectToTop    key.Binding
	MultiSelectToBottom key.Binding
}

// DefaultKeyMap returns a default set of keybindings.
func DefaultKeyMap() KeyMap {
	const spacebar = " "
	return KeyMap{
		LineUp: key.NewBinding(
			key.WithKeys("up", "k"),
			key.WithHelp("↑/k", "up"),
		),
		LineDown: key.NewBinding(
			key.WithKeys("down", "j"),
			key.WithHelp("↓/j", "down"),
		),
		PageUp: key.NewBinding(
			key.WithKeys("b", "pgup"),
			key.WithHelp("b/pgup", "page up"),
		),
		PageDown: key.NewBinding(
			key.WithKeys("f", "pgdown", spacebar),
			key.WithHelp("f/pgdn", "page down"),
		),
		HalfPageUp: key.NewBinding(
			key.WithKeys("u", "ctrl+u"),
			key.WithHelp("u", "½ page up"),
		),
		HalfPageDown: key.NewBinding(
			key.WithKeys("d", "ctrl+d"),
			key.WithHelp("d", "½ page down"),
		),
		GotoTop: key.NewBinding(
			key.WithKeys("home", "g"),
			key.WithHelp("g/home", "go to start"),
		),
		GotoBottom: key.NewBinding(
			key.WithKeys("end", "G"),
			key.WithHelp("G/end", "go to end"),
		),
		MultiSelectUp: key.NewBinding(
			key.WithKeys("shift+up", "shift+k"),
			key.WithHelp("shift+up/shift+k", "multi-select up"),
		),
		MultiSelectDown: key.NewBinding(
			key.WithKeys("shift+down", "shift+j"),
			key.WithHelp("shift+down/shift+j", "multi-select down"),
		),
		MultiSelectToTop: key.NewBinding(
			key.WithKeys("shift+home"),
			key.WithHelp("ctrl+shift+up", "multi-select to top"),
		),
		MultiSelectToBottom: key.NewBinding(
			key.WithKeys("shift+end"),
			key.WithHelp("ctrl+shift+down", "multi-select to bottom"),
		),
	}
}

// Styles contains style definitions for this list component. By default, these
// values are generated by DefaultStyles.
type Styles struct {
	Header   lipgloss.Style
	Cell     lipgloss.Style
	Selected lipgloss.Style
	Cursor   lipgloss.Style
	Wrapper  lipgloss.Style
	EvenCell lipgloss.Style // applied to cells in even rows
	OddCell  lipgloss.Style // applied to cells in odd rows
}

// DefaultStyles returns a set of default style definitions for this table.
func DefaultStyles() Styles {
	return Styles{
		Selected: lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("212")),
		Header:   lipgloss.NewStyle().Bold(true).Padding(0, 1),
		Cell:     lipgloss.NewStyle().Padding(0, 1),
		Cursor:   lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("212")),
		Wrapper: lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("240")),
		EvenCell: lipgloss.NewStyle().Padding(0, 1),
		OddCell:  lipgloss.NewStyle().Padding(0, 1),
	}
}

// SetStyles sets the table styles.
func (m *Table) SetStyles(s Styles) {
	m.styles = s
}

// NewTable creates a new model for the table widget.
func NewTable(doubleClickDelay int, opts ...Option) Table {

	m := Table{
		cursor:          0,
		selected:        map[int]struct{}{0: {}},
		KeyMap:          DefaultKeyMap(),
		styles:          DefaultStyles(),
		zoneMgr:         zone.New(),
		lastClickedTime: time.Time{},
		lastClickedIdx:  -1,
		clickDelay:      time.Duration(time.Millisecond * time.Duration(doubleClickDelay)),
	}
	for _, opt := range opts {
		opt(&m)
	}

	m.updateViewport()

	return m
}

// Update is the Bubble Tea update loop.
func (m Table) Update(msg tea.Msg) (Table, tea.Cmd) {
	if !m.focus {
		return m, nil
	}
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.KeyMap.LineUp):
			m.moveUp(1, false)
		case key.Matches(msg, m.KeyMap.LineDown):
			m.moveDown(1, false)
		case key.Matches(msg, m.KeyMap.PageUp):
			m.moveUp(m.height, false)
		case key.Matches(msg, m.KeyMap.PageDown):
			m.moveDown(m.height, false)
		case key.Matches(msg, m.KeyMap.HalfPageUp):
			m.moveUp(m.height/2, false)
		case key.Matches(msg, m.KeyMap.HalfPageDown):
			m.moveDown(m.height/2, false)
		case key.Matches(msg, m.KeyMap.LineDown):
			m.moveDown(1, false)
		case key.Matches(msg, m.KeyMap.GotoTop):
			m.gotoTop()
		case key.Matches(msg, m.KeyMap.GotoBottom):
			m.gotoBottom()
		case key.Matches(msg, m.KeyMap.MultiSelectUp):
			m.moveUp(1, true)
		case key.Matches(msg, m.KeyMap.MultiSelectDown):
			m.moveDown(1, true)
		case key.Matches(msg, m.KeyMap.MultiSelectToTop):
			m.multiSelectToTop()
		case key.Matches(msg, m.KeyMap.MultiSelectToBottom):
			m.multiSelectToBottom()
		}
	case tea.MouseMsg:
		if msg.Type != tea.MouseLeft {
			return m, nil
		}
		for i, col := range m.cols {
			if m.zoneMgr.Get("col" + strconv.Itoa(i)).InBounds(msg) {
				log.Println("Clicked on column", col.Title)
				// m.sortBy = col.Title
				// m.sortAsc = !m.sortAsc
				// m.Sort()
			}
		}

		for i := m.start; i <= m.end; i++ {
			if m.zoneMgr.Get("row" + strconv.Itoa(i)).InBounds(msg) {
				log.Println("in bounds", i, m.rows[i])
			}
		}

	}

	return m, nil
}

// SelectedRows returns the indexes of the selected rows.
func (m Table) SelectedRows() []int {
	rows := make([]int, 0, len(m.selected))
	for k := range m.selected {
		rows = append(rows, k)
	}
	sort.Ints(rows)
	return rows
}

// SelectedRowsValues returns the values of the selected rows.
// The values are sorted randomly.
func (m Table) SelectedRowsValues() [][]string {
	rows := make([][]string, 0, len(m.selected))
	for k := range m.selected {
		rows = append(rows, m.rows[k])
	}
	return rows
}

// Rows returns the current rows.
func (m Table) Rows() []Row {
	return m.rows
}

// SetRows sets a new rows state.
func (m *Table) SetRows(r []Row) {
	m.rows = r
	m.updateViewport()
}

// FromValues create the table rows from a simple string. It uses `\n` by
// default for getting all the rows and the given separator for the fields on
// each row.
func (m *Table) FromValues(value, separator string) {
	rows := []Row{}
	for _, line := range strings.Split(value, "\n") {
		r := Row{}
		for _, field := range strings.Split(line, separator) {
			r = append(r, field)
		}
		rows = append(rows, r)
	}

	m.SetRows(rows)
}

// SetColumns sets a new columns state.
func (m *Table) SetColumns(c []Column) {
	m.cols = c
}

// SetWidth sets the width of the viewport of the table.
func (m *Table) SetWidth(w int) {
	m.width = w
}

// SetHeight sets the number of rows that will be displayed.
// If height is set to zero, the table will become unresponsive (blurred)
// because the user would not be able to see what is happening with the table.
func (m *Table) SetHeight(h int) {
	m.styles.Wrapper = m.styles.Wrapper.Height(h - m.styles.Wrapper.GetVerticalFrameSize())
	rowHeight := h - lipgloss.Height(m.headersView()) - m.styles.Wrapper.GetVerticalFrameSize()
	grow := rowHeight - m.height
	// height is the number of rows showing so it is "1 indexed"
	m.height = clamp(rowHeight, 0, len(m.rows))
	switch {
	case grow == 0:
		//no op
		return
	case grow > 0:
		// grow as much down as possible and then add above if needed
		m.end = clamp(m.start+m.height-1, m.start, len(m.rows)-1)
		m.start = clamp(m.end-m.height+1, 0, m.end)
	case grow < 0:
		// remove from end until reaching the cursor, then remove from the top
		if m.end+grow <= m.cursor {
			m.end = m.cursor
		} else {
			m.end = m.end + grow
		}
		m.start = clamp(m.end-m.height+1, 0, len(m.rows)-1)
	}
	m.updateViewport()
}

// SetSelected sets the rows at the given indexes as selected.
func (m *Table) SetSelected(idxs []int) error {
	for _, idx := range idxs {
		if idx < 0 || idx >= len(m.rows)-1 {
			return fmt.Errorf("index out of bounds: %d", idx)
		}
		m.selected[idx] = struct{}{}
	}
	return nil
}

func (m *Table) SetSort(col int, asc bool) {
	m.sort.col = col
	m.sort.asc = asc
}

// Height returns the viewport height of the table.
func (m Table) Height() int {
	return m.height
}

// Width returns the viewport width of the table.
func (m Table) Width() int {
	return m.width
}

// Cursor returns the index of the selected row.
func (m Table) Cursor() int {
	return m.cursor
}

// CursorValue returns the row where the cursor is
func (m Table) CursorValue() []string {
	return m.rows[m.cursor]
}

// SetCursor sets the cursor position in the table.
// If you only want one item selected, call ClearSelected before this method.
func (m *Table) SetCursor(n int) {
	if len(m.rows) == 0 {
		return
	}
	m.cursor = clamp(n, 0, len(m.rows)-1)
	m.updateViewport()
}

// ClearSelected unselects all rows
func (m *Table) ClearSelected() {
	for k := range m.selected {
		delete(m.selected, k)
	}
}

// Focused returns the focus state of the table.
func (m Table) Focused() bool {
	return m.focus
}

// Focus focuses the table, allowing the user to move around the rows and
// interact.
func (m *Table) Focus() {
	m.focus = true
}

// Blur blurs the table, preventing selection or movement.
func (m *Table) Blur() {
	m.focus = false
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func clamp(v, low, high int) int {
	return min(max(v, low), high)
}
