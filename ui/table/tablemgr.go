package table

import (
	"reflect"
	"time"

	"github.com/Philistino/fman/entry"
	"github.com/Philistino/fman/ui/message"
	"github.com/Philistino/fman/ui/theme"
	"github.com/Philistino/fman/ui/theme/colors"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type TableMgr struct {
	Table
	entries      []entry.Entry
	loadingDelay time.Duration
	spinner      spinner.Model
	loading      bool
	sort         tableSort
}

func NewTableMgr(colorScheme colors.Theme, loadingDelay int, doubleClickDelay int) *TableMgr {
	spin := spinner.New()
	spin.Spinner = spinner.Dot
	spin.Style = lipgloss.NewStyle().Foreground(colorScheme.SelectedItemBgColor)

	cols := []Column{
		{"Name", 50},
		{"Size", 10},
		{"Modified", 10},
	}

	styles := Styles{
		Wrapper:  lipgloss.NewStyle().Padding(0, 0, 1, 0),
		Selected: theme.SelectedItemStyle,
		EvenCell: theme.EvenItemStyle,
		Cursor:   theme.SelectedItemStyle,
	}

	mgr := TableMgr{
		Table:        NewTable(doubleClickDelay, WithColumns(cols), WithStyles(styles)),
		loadingDelay: time.Duration(time.Millisecond * time.Duration(loadingDelay)),
		spinner:      spin,
	}
	return &mgr
}

// Update is the Bubble Tea update loop.
func (m *TableMgr) Update(msg tea.Msg) (*TableMgr, tea.Cmd) {

	switch msg := msg.(type) {
	case spinner.TickMsg:
		spin, cmd := m.spinner.Update(msg)
		m.spinner = spin
		return m, cmd
	case message.DirChangedMsg:
		// log.Println("TableMgr: DirChangedMsg")
		selected := m.selected
		m.handlePathChange(msg)
		if !reflect.DeepEqual(selected, m.selected) { // lazy
			selected := make(map[string]struct{}, len(m.selected))
			for i := range m.selected {
				selected[m.rows[i][0]] = struct{}{}
			}
			return m, message.SelectedCmd(selected)
		}
	}
	table, cmd := m.Table.Update(msg)
	m.Table = table
	return m, cmd
}

func (m *TableMgr) handlePathChange(newDir message.DirChangedMsg) tea.Cmd {
	if newDir.Error() != nil {
		return nil
	}
	m.SetRows(m.entriesToRows(newDir.Entries()))

	m.selected = make(map[int]struct{})
	m.entries = newDir.Entries()
	selected := newDir.Selected()
	matched := false
	for i, entry := range m.entries {
		// set the cursor
		if entry.Name() == newDir.Cursor() {
			// list.cursorIdx = i
			m.selected[i] = struct{}{}
			matched = true
			continue
		}
		// set the selected entries
		_, ok := selected[entry.Name()]
		if !ok {
			continue
		}
		m.selected[i] = struct{}{}
	}
	if !matched {
		m.SetCursor(0)
	}

	// if len(list.entries) == 0 {
	// 	return message.EmptyDirCmd()
	// }
	return nil
	// return message.NewEntryCmd(list.SelectedEntry())
}

// entriesToRows creates a slice of rows from a slice of entries.
func (m TableMgr) entriesToRows(entries []entry.Entry) []Row {
	rows := []Row{}
	for _, e := range entries {
		// i := icons.GetIconForReal(e, e.IsHidden)
		rows = append(
			rows,
			Row{
				// lipgloss.NewStyle().Foreground(lipgloss.Color(i.ColorHex())).Render(i.Glyph()) + lipgloss.NewStyle().Render(" "+e.Name()),
				e.Name(),
				// fmt.Sprintf("%s%s\033[39m", i.ColorTerm(), i.Glyph()),
				e.SizeStr,
				e.ModifyTime,
			},
		)
	}
	return rows
}

// View renders the table.
func (m TableMgr) View() string {
	// log.Println("TableMgr: View")
	if m.loading {
		return m.spinner.View()
	}
	return m.Table.View()
}
