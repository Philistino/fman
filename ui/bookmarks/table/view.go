package table

import (
	"strconv"

	"github.com/charmbracelet/lipgloss"
	"github.com/mattn/go-runewidth"
)

// View renders the component.
func (m Table) View() string {

	if len(m.rows) == 0 {
		view := m.styles.Wrapper.Render(lipgloss.JoinVertical(
			lipgloss.Center,
			m.headersView(),
			lipgloss.Place(m.width, 1, lipgloss.Top, lipgloss.Center, m.emptyMessage),
		))
		if m.zoneMgr != nil {
			view = m.zoneMgr.Scan(view)
		}
		return view
	}

	renderedRows := make([]string, 0, m.height)
	for i := m.start; i <= m.end; i++ {
		renderedRows = append(renderedRows, m.renderRow(i))
	}

	view := m.styles.Wrapper.Render(lipgloss.JoinVertical(
		lipgloss.Center,
		m.headersView(),
		lipgloss.JoinVertical(lipgloss.Left, renderedRows...),
	))
	if m.zoneMgr != nil {
		view = m.zoneMgr.Scan(view)
	}
	return view
}

// headersView renders the column headers
func (m Table) headersView() string {
	var s = make([]string, 0, len(m.cols))
	for i, col := range m.cols {
		style := lipgloss.NewStyle().Width(col.Width).MaxWidth(col.Width).Inline(true)
		sortArrow := ""
		if i == m.sort.col {
			if m.sort.asc {
				sortArrow = "▲"
			} else {
				sortArrow = "▼"
			}
		}
		renderedCell := style.Render(runewidth.Truncate(col.Title, col.Width-len(sortArrow), "…") + sortArrow)
		renderedCell = m.styles.Header.Render(renderedCell)
		if m.zoneMgr != nil {
			renderedCell = m.zoneMgr.Mark("col"+strconv.Itoa(i), renderedCell)
		}
		s = append(s, renderedCell)
	}
	return lipgloss.JoinHorizontal(lipgloss.Left, s...)
}

func (m *Table) renderRow(rowID int) string {
	var s = make([]string, 0, len(m.cols))
	for i, value := range m.rows[rowID] {
		style := lipgloss.NewStyle().Width(m.cols[i].Width).MaxWidth(m.cols[i].Width).Inline(true)
		renderedCell := m.styles.Cell.Render(style.Render(runewidth.Truncate(value, m.cols[i].Width, "…")))
		s = append(s, renderedCell)
	}

	row := lipgloss.JoinHorizontal(lipgloss.Left, s...)

	_, selected := m.selected[rowID]
	switch {
	case rowID == m.cursor:
		row = m.styles.Cursor.Render(row)
	case selected:
		row = m.styles.Selected.Render(row)
	default:
	}

	if m.zoneMgr != nil {
		row = m.zoneMgr.Mark("row"+strconv.Itoa(rowID), row)
	}

	return row
}
