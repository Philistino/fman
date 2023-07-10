package table

// Option is used to set options in New
type Option func(*Table)

// WithCursor sets the initial cursor position.
func WithCursor(n int) Option {
	return func(m *Table) {
		m.cursor = n
		m.selected = map[int]struct{}{n: {}}
	}
}

// WithSort sets the column to sort by and the sort order.
func WithSort(col int, asc bool) Option {
	return func(m *Table) {
		m.sort.col = col
		m.sort.asc = asc
	}
}

// WithEmptyMessage sets the message to display when the table has no rows.
// For example:
//
//	table := New(WithEmptyMessage("No results found"))
func WithEmptyMessage(text string) Option {
	return func(m *Table) {
		m.emptyMessage = text
	}
}

// WithColumns sets the table columns (headers).
func WithColumns(cols []Column) Option {
	return func(m *Table) {
		m.cols = cols
	}
}

// WithRows sets the table rows (data).
func WithRows(rows []Row) Option {
	return func(m *Table) {
		m.rows = rows
	}
}

// WithHeight sets the height of the table.
func WithHeight(h int) Option {

	return func(m *Table) {
		m.SetHeight(h)
	}
}

// WithWidth sets the width of the table.
func WithWidth(w int) Option {
	return func(m *Table) {
		m.width = w
	}
}

// WithFocused sets the focus state of the table.
func WithFocused(f bool) Option {
	return func(m *Table) {
		m.focus = f
	}
}

// WithStyles sets the table styles.
func WithStyles(s Styles) Option {
	return func(m *Table) {
		m.styles = s
	}
}

// WithKeyMap sets the key map.
func WithKeyMap(km KeyMap) Option {
	return func(m *Table) {
		m.KeyMap = km
	}
}
