package model

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
	m.cursor = clamp(m.cursor-n, 0, len(m.rows)-1)

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

// moveUp moves the selection up by any number of rows.
// It can not go above the first row.
func (m *Table) moveUp(n int, multi bool) {
	m.manageSelected(n, multi, false)
	m.updateViewport()
}

// moveDown moves the selection down by any number of rows.
// It can not go below the last row.
func (m *Table) moveDown(n int, multi bool) {
	m.manageSelected(n, multi, true)
	m.updateViewport()
}

// multiSelectToTop selects all rows from the cursor to the top.
// And sets the cursor to the first row.
func (m *Table) multiSelectToTop() {
	for i := 0; i <= m.cursor; i++ {
		m.selected[i] = struct{}{}
	}
	m.SetCursor(0)
	m.updateViewport()
}

// multiSelectToBottom selects all rows from the cursor to the bottom.
// And sets the cursor to the final row.
func (m *Table) multiSelectToBottom() {
	for i := m.cursor; i < len(m.rows); i++ {
		m.selected[i] = struct{}{}
	}
	m.SetCursor(len(m.rows) - 1)
	m.updateViewport()
}

// gotoTop moves the cursor to the first row.
func (m *Table) gotoTop() {
	m.moveUp(m.cursor, false)
}

// gotoBottom moves the cursor to the last row.
func (m *Table) gotoBottom() {
	m.moveDown(len(m.rows), false)
}
