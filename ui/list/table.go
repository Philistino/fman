package list

import (
	"fmt"
	"sort"
)

// Table defines a state for the table widget.
type Table struct {
	nRows    int
	cursor   int
	selected map[int]struct{}

	start     int
	end       int
	height    int
	maxHeight int
}

// NewTable creates a new model for the table widget.
func NewTable() Table {

	m := Table{
		cursor:   0,
		selected: map[int]struct{}{0: {}},
	}
	m.updateViewport()
	return m
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

// NRows returns the current number of rows.
func (m Table) NRows() int {
	return m.nRows
}

// SetNRows sets the number of rows that the table has
func (m *Table) SetNRows(n int) {
	m.nRows = n
	toShow := clamp(m.nRows, 0, m.maxHeight)
	grow := toShow - m.height
	m.height = toShow
	m.setHeight(grow)
}

// SetHeight sets the number of rows that will be displayed.
func (m *Table) SetHeight(h int) {
	m.maxHeight = h
	grow := h - m.height
	if m.nRows < h {
		m.height = m.nRows
	} else {
		m.height = h
	}
	m.setHeight(grow)
}

func (m *Table) setHeight(grow int) {
	switch {
	case grow == 0:
		//no op
		return
	case grow > 0:
		// grow as much down as possible and then add above if needed
		m.end = clamp(m.start+m.height-1, m.start, m.nRows-1)
		m.start = clamp(m.end-m.height+1, 0, m.end)
	case grow < 0:
		// remove from end until reaching the cursor, then remove from the top
		if m.end+grow <= m.cursor {
			m.end = m.cursor
		} else {
			m.end = m.end + grow
		}
		m.start = clamp(m.end-m.height+1, 0, m.nRows-1)
	}
	m.updateViewport()
}

// SetSelected sets the rows at the given indexes as selected.
func (m *Table) SetSelected(idxs []int) error {
	m.ClearSelected()
	for _, idx := range idxs {
		if idx < 0 || idx >= m.nRows-1 {
			return fmt.Errorf("index out of bounds: %d", idx)
		}
		m.selected[idx] = struct{}{}
	}
	return nil
}

// Height returns the viewport height of the table.
func (m Table) Height() int {
	return m.height
}

// Cursor returns the index of the selected row.
func (m Table) Cursor() int {
	return m.cursor
}

// SetCursor sets the cursor position in the table.
// If you only want one item selected, call ClearSelected before this method.
func (m *Table) SetCursor(n int) {
	if m.nRows == 0 {
		return
	}
	m.cursor = clamp(n, 0, m.nRows-1)
	m.updateViewport()
}

// ClearSelected unselects all rows
func (m *Table) ClearSelected() {
	for k := range m.selected {
		delete(m.selected, k)
	}
}

// manageSelected manages the selected rows in the table based on the given parameters.
// It moves the cursor up or down by the given number of rows (n), and updates the selected rows
// based on whether multi-select is enabled and whether the cursor moved up or down.
func (m *Table) manageSelected(n int, multi bool, down bool) {

	// record existing cursor position
	preCursor := m.cursor

	// to make this method useable for both up and down movements
	if down {
		n = n * -1
	}
	m.cursor = clamp(m.cursor-n, 0, m.nRows-1)

	if !multi {
		m.ClearSelected()
		m.selected[m.cursor] = struct{}{}
		return
	}

	// if the cursor didn't move (e.g., at top of list),
	// and we are multi-selecting, return early, its a noop
	if m.cursor == preCursor {
		return
	}

	// if the new cursor was already in the map, the user is actually de-selecting
	// the prior extrema of the multi-select so we delete it from the map and return
	_, ok := m.selected[m.cursor]
	if ok {
		delete(m.selected, preCursor)
		return
	}

	// we are extending the multiselect
	m.selected[m.cursor] = struct{}{}
}

// updateViewport updates the list content based on the previously defined
// columns and rows.
func (m *Table) updateViewport() {
	switch {
	case m.height == 0:
		m.start = m.cursor
		m.end = m.cursor
	case m.start <= m.cursor && m.cursor <= m.end:
		// do nothing
	case m.cursor > m.end:
		m.end = m.cursor
		m.start = m.end - m.height + 1
	case m.cursor < m.start:
		m.start = m.cursor
		m.end = m.start + m.height - 1
	}
}

// MoveUp moves the selection up by any number of rows.
// It can not go above the first row.
func (m *Table) MoveUp(n int, multi bool) {
	m.manageSelected(n, multi, false)
	m.updateViewport()
}

// MoveDown moves the selection down by any number of rows.
// It can not go below the last row.
func (m *Table) MoveDown(n int, multi bool) {
	m.manageSelected(n, multi, true)
	m.updateViewport()
}

// MultiSelectToTop selects all rows from the cursor to the top.
// And sets the cursor to the first row.
func (m *Table) MultiSelectToTop() {
	for i := 0; i <= m.cursor; i++ {
		m.selected[i] = struct{}{}
	}
	m.SetCursor(0)
	m.updateViewport()
}

// MultiSelectToBottom selects all rows from the cursor to the bottom.
// And sets the cursor to the final row.
func (m *Table) MultiSelectToBottom() {
	for i := m.cursor; i < m.nRows; i++ {
		m.selected[i] = struct{}{}
	}
	m.SetCursor(m.nRows - 1)
	m.updateViewport()
}

// GoToTop moves the cursor to the first row.
func (m *Table) GoToTop() {
	m.MoveUp(m.cursor, false)
}

// GoToBottom moves the cursor to the last row.
func (m *Table) GoToBottom() {
	m.MoveDown(m.nRows, false)
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
