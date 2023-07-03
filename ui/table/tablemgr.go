package model

import (
	"fmt"
	"reflect"
	"time"

	"github.com/Philistino/fman/entry"
	"github.com/Philistino/fman/icons"
	"github.com/Philistino/fman/ui/message"
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

func NewTableMgr(theme colors.Theme, loadingDelay int, doubleClickDelay int) TableMgr {
	spin := spinner.New()
	spin.Spinner = spinner.Dot
	spin.Style = lipgloss.NewStyle().Foreground(theme.SelectedItemBgColor)
	return TableMgr{
		Table:        NewTable(doubleClickDelay),
		loadingDelay: time.Duration(time.Millisecond * time.Duration(loadingDelay)),
		spinner:      spin,
	}
}

// Update is the Bubble Tea update loop.
func (m TableMgr) Update(msg tea.Msg) (TableMgr, tea.Cmd) {

	switch msg := msg.(type) {

	case spinner.TickMsg:
		spin, cmd := m.spinner.Update(msg)
		m.spinner = spin
		return m, cmd
	case message.DirChangedMsg:
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
	return m, nil
}

func (list *TableMgr) handlePathChange(newDir message.DirChangedMsg) tea.Cmd {
	if newDir.Error() != nil {
		return nil
	}
	list.selected = make(map[int]struct{})
	list.entries = newDir.Entries()
	selected := newDir.Selected()
	matched := false
	for i, entry := range list.entries {
		// set the cursor
		if entry.Name() == newDir.Cursor() {
			// list.cursorIdx = i
			list.selected[i] = struct{}{}
			matched = true
			continue
		}
		// set the selected entries
		_, ok := selected[entry.Name()]
		if !ok {
			continue
		}
		list.selected[i] = struct{}{}
	}
	if !matched {
		// list.cursorIdx = 0
	}

	if len(list.entries) == 0 {
		return message.EmptyDirCmd()
	}
	return nil
	// return message.NewEntryCmd(list.SelectedEntry())
}

// entriesToRows creates a slice of rows from a slice of entries.
func (m TableMgr) entriesToRows(entries []entry.Entry) []Row {
	rows := []Row{}
	for _, e := range entries {
		i := icons.GetIconForReal(e, e.IsHidden)
		rows = append(
			rows,
			Row{
				fmt.Sprintf("%s%s\033[39m", i.ColorTerm(), i.Glyph()),
				e.Name(),
				e.SizeStr,
				e.ModifyTime,
			},
		)
	}
	return rows
}
