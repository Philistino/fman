package model

import (
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/76creates/stickers"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/nore-dev/fman/entry"
	"github.com/nore-dev/fman/theme"
)

type ListModel struct {
	entries []entry.Entry

	path string

	Width int

	selected_index int
	flexBox        *stickers.FlexBox

	maxEntryToShow int

	initialized bool
}

type UpdateEntriesMsg struct {
	parent bool
}

type PathMsg struct {
	path string
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func detectOpenCommand() string {
	switch runtime.GOOS {
	case "linux":
		return "xdg-open"
	case "darwin":
		return "open"
	}

	return "start"
}

func NewListModel() ListModel {

	path, _ := filepath.Abs(".")

	entries, err := entry.GetEntries(path)

	if err != nil {
		panic(err)
	}

	list := ListModel{
		path:           path,
		entries:        entries,
		flexBox:        stickers.NewFlexBox(0, 0),
		maxEntryToShow: 23,
		initialized:    false,
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

func (list ListModel) Init() tea.Cmd {
	return nil
}

func (list *ListModel) getEntriesAbove() {
	list.path = filepath.Dir(list.path)
	entries, err := entry.GetEntries(list.path)

	if err != nil {
		panic(err)
	}

	list.entries = entries
}

func (list *ListModel) getEntriesBelow() {
	if !list.SelectedEntry().IsDir {
		return
	}

	list.path = filepath.Join(list.path, list.SelectedEntry().Name)

	entries, err := entry.GetEntries(list.path)

	if err != nil {
		panic(err)
	}

	list.entries = entries
}

func (list *ListModel) restrictIndex() {
	if list.selected_index < 0 {
		list.selected_index = len(list.entries) - 1
	} else if list.selected_index >= len(list.entries) {
		list.selected_index = 0
	}
}

func (list ListModel) Update(msg tea.Msg) (ListModel, tea.Cmd) {

	fullPath := filepath.Join(list.path, list.SelectedEntry().Name)

	switch msg := msg.(type) {
	case UpdateEntriesMsg:
		if msg.parent {
			list.getEntriesAbove()
		} else {
			list.getEntriesBelow()
		}

		list.restrictIndex()

		return list, func() tea.Msg {
			return PathMsg{list.path}
		}

	case tea.KeyMsg:
		switch msg.String() {
		case "w", "up", "j": // Select entry above
			list.selected_index -= 1

		case "s", "down", "k": // Select entry below
			list.selected_index += 1

		case "a", "left", "h": // Get entries from parent directory
			return list, func() tea.Msg {
				return UpdateEntriesMsg{parent: true}
			}
		case "d", "right", "l": // If the selected entry is a directory. Get entries under that directory
			return list, func() tea.Msg {
				return UpdateEntriesMsg{}
			}
		case "enter": // Open file with default application
			cmd := exec.Command(detectOpenCommand(), fullPath)
			cmd.Run()
		}

	case tea.WindowSizeMsg:
		list.flexBox.SetWidth(list.Width)
	}

	list.restrictIndex()

	if !list.initialized {
		list.initialized = true

		return list, func() tea.Msg {
			return PathMsg{list.path}
		}
	}

	return list, nil

}

func (list ListModel) View() string {
	list.flexBox.ForceRecalculate()

	if len(list.entries) == 0 {
		return "Empty"
	}

	cellsLength := list.flexBox.Row(0).CellsLen()
	contents := make([]strings.Builder, cellsLength)

	// Write List headers
	contents[0].WriteString(theme.BoldStyle.Render("Name"))
	contents[0].WriteByte('\n')

	contents[1].WriteString(theme.BoldStyle.Render("Size"))
	contents[1].WriteByte('\n')

	contents[2].WriteString(theme.BoldStyle.Render("Modify Time"))
	contents[2].WriteByte('\n')

	startIndex := max(0, list.selected_index-list.maxEntryToShow)

	for index := startIndex; index < len(list.entries); index++ {
		entry := list.entries[index]

		content := make([]strings.Builder, cellsLength)

		content[0].WriteString(entry.Name)
		content[1].WriteString(entry.Size)
		content[2].WriteString(entry.ModifyTime)

		for i := 0; i < cellsLength; i++ {

			style := lipgloss.NewStyle()
			offset := 0

			if index == list.selected_index {
				style = theme.SelectedItemStyle
			} else if index%2 == 0 {
				style = theme.EvenItemStyle
			}

			// IDK
			if i == 2 {
				offset = 2
			}

			style.Width(list.flexBox.Row(0).Cell(i).GetWidth() - offset)

			contents[i].WriteString(style.Render(content[i].String()))
			contents[i].WriteByte('\n')
		}
	}

	for i := 0; i < cellsLength; i++ {
		list.flexBox.Row(0).Cell(i).SetContent(contents[i].String())
	}

	return list.flexBox.Render()
}

func (list ListModel) SelectedEntry() entry.Entry {

	if len(list.entries) == 0 {
		return entry.Entry{}
	}

	return list.entries[list.selected_index]
}
