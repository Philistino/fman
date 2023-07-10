package bookmark_ui

import (
	"context"
	"path/filepath"
	"sort"
	"strings"

	"github.com/Philistino/fman/bookmarks"
	"github.com/Philistino/fman/ui/focus"
	"github.com/Philistino/fman/ui/message"
	"github.com/Philistino/fman/ui/table"
	"github.com/Philistino/fman/ui/theme"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// TODO: Only bookmark directories, not files

type Bookmarks struct {
	focus.FocusField
	height   int
	width    int
	hidden   bool
	quierier *bookmarks.Querier
	table    table.Table
	paths    []string
	pinIcon  rune
	zPrefix  string
}

func NewBookmarks(quierier *bookmarks.Querier, pinIcon rune, hidden bool, doubleClickDelay int) *Bookmarks {
	styles := table.Styles{
		Header:   lipgloss.NewStyle(),
		Cell:     lipgloss.NewStyle(),
		Selected: lipgloss.NewStyle().Foreground(lipgloss.Color("205")),
		Cursor:   lipgloss.NewStyle().Foreground(lipgloss.Color("205")),
		EvenCell: lipgloss.NewStyle(),
		OddCell:  lipgloss.NewStyle(),
		Wrapper:  theme.EntryInfoStyle.Copy(),
	}

	table := table.NewTable(
		doubleClickDelay,
		table.WithStyles(styles),
		table.WithColumns([]table.Column{{Title: "", Width: 25}}),
		table.WithFocused(true),
	)

	return &Bookmarks{
		quierier: quierier,
		hidden:   hidden,
		pinIcon:  pinIcon,
		zPrefix:  "bookmarks",
		table:    table,
	}
}

func (m *Bookmarks) Init() tea.Cmd {
	return m.getBookMarks()
}

// getBookMarks retrieves the bookmarks from the querier and sorts them alphabetically.
// It sets the paths field to the sorted bookmarks.
func (m *Bookmarks) getBookMarks() tea.Cmd {
	marks, err := m.quierier.GetBookmarks(context.Background())
	if err != nil {
		return message.NewNotificationCmd("Error loading bookmarks " + err.Error())
	}
	sort.Slice(marks, func(i, j int) bool {
		return strings.ToLower(marks[i]) < strings.ToLower(marks[j])
	})
	m.paths = marks
	m.SetRows()
	m.table.ClearSelected()
	m.table.SetCursor(0)
	return nil
}

// addBookmarks adds the given paths to the bookmarks.
func (m *Bookmarks) addBookmarks(paths []string) tea.Cmd {
	err := m.quierier.CreateBookmarks(context.Background(), paths)
	if err != nil {
		return message.NewNotificationCmd("Error adding bookmarks " + err.Error())
	}
	return m.getBookMarks()
}

// deleteBookmarks deletes the given paths from the bookmarks.
func (m *Bookmarks) deleteBookmarks(paths []string) tea.Cmd {
	err := m.quierier.DeleteBookMarks(context.Background(), paths)
	if err != nil {
		return message.NewNotificationCmd("Error deleting bookmarks " + err.Error())
	}
	return m.getBookMarks()
}

func (m *Bookmarks) Update(msg tea.Msg) (*Bookmarks, tea.Cmd) {
	if !m.Focused() {
		return m, nil
	}

	var cmd tea.Cmd
	switch msg := msg.(type) {
	case BookmarkMsg:
		cmd = m.addBookmarks(msg.paths)
		return m, cmd
	case UnbookmarkMsg:
		cmd = m.deleteBookmarks(msg.paths)
		return m, cmd
	}
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m *Bookmarks) View() string {
	if m.hidden {
		return ""
	}
	return m.table.View()
}

func (m *Bookmarks) SetRows() {
	rows := make([]table.Row, len(m.paths))
	for i, path := range m.paths {
		rows[i] = table.Row{filepath.Base(path)}
	}
	m.table.SetRows(rows)
}

func (m *Bookmarks) SetWidth(w int) {
	m.width = w
	m.table.SetWidth(w)
}

func (m *Bookmarks) SetHeight(h int) {
	m.height = h
	m.table.SetHeight(h)
}

func (m *Bookmarks) Hide() {
	m.hidden = true
}

func (m *Bookmarks) Show() {
	m.hidden = false
}
