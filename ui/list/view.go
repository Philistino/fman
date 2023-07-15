package list

import (
	"fmt"
	"strings"

	"github.com/Philistino/fman/icons"
	"github.com/Philistino/fman/ui/theme"
	"github.com/charmbracelet/lipgloss"
	"github.com/mattn/go-runewidth"
	"github.com/muesli/termenv"
)

func (list *List) View() string {
	list.flexBox.ForceRecalculate()

	cellsLength := list.flexBox.Row(0).CellsLen()
	contents := make([]strings.Builder, cellsLength)

	// Write List headers
	contents[0].WriteRune(theme.GetActiveIconTheme().NameIcon)
	contents[0].WriteString(termenv.String(" Name").Italic().String())
	contents[0].WriteByte('\n')

	contents[1].WriteRune(theme.GetActiveIconTheme().SizeIcon)
	contents[1].WriteString(termenv.String(" Size").Italic().String())
	contents[1].WriteByte('\n')

	contents[2].WriteRune(theme.GetActiveIconTheme().TimeIcon)
	contents[2].WriteString(termenv.String(" Modify Time").Italic().String())
	contents[2].WriteByte('\n')

	if len(list.entries) == 0 {
		for i := 0; i < cellsLength; i++ {
			list.flexBox.Row(0).Cell(i).SetContent(contents[i].String())
		}
		return lipgloss.JoinVertical(lipgloss.Center, list.flexBox.Render(), lipgloss.Place(list.flexBox.GetWidth(), 1, lipgloss.Center, lipgloss.Center, "This directory is empty"))
	}

	for index := list.table.start; index <= list.table.end; index++ {
		entry := list.entries[index]
		content := make([]strings.Builder, cellsLength)

		if entry.SymlinkName != "" {
			content[0].WriteRune(theme.GetActiveIconTheme().SymlinkIcon)
		} else if entry.IsDir() {
			i := icons.GetIconForReal(entry, entry.IsHidden)
			content[0].WriteString(fmt.Sprintf("%s%s\033[39m", i.ColorTerm(), i.Glyph()))

		} else {
			content[0].WriteString(icons.GetIconTerm(entry, entry.IsHidden))
		}

		content[0].WriteRune(' ')
		name := runewidth.Truncate(entry.Name(), list.truncateLimit-2, "...")
		content[0].WriteString(name)
		content[1].WriteString(entry.SizeStr)
		content[2].WriteString(entry.ModifyTime)

		var style lipgloss.Style
		for i := 0; i < cellsLength; i++ {
			if index == list.table.Cursor() {
				style = theme.SelectedItemStyle
			} else if index%2 == 0 {
				style = theme.EvenItemStyle
			}

			// IDK
			// offset := 0
			// if i == 2 {
			// 	offset = 2
			// }

			// style = style.Width(list.flexBox.Row(0).Cell(i).GetWidth() - offset)
			style = style.Width(list.flexBox.Row(0).Cell(i).GetWidth())

			if i == 0 && entry.SymlinkName != "" {
				style = style.Bold(true).Underline(true)
			} else {
				style = style.UnsetBold().UnsetUnderline()
			}

			// Colors
			if index == list.table.Cursor() {
				style = style.Foreground(list.theme.SelectedItemFgColor)
			} else if entry.IsHidden {
				style = style.Foreground(list.theme.HiddenFileColor)
				if entry.IsDir() {
					style = style.Foreground(list.theme.HiddenFolderColor)
				}
			} else if entry.IsDir() {
				style = style.Foreground(list.theme.FolderColor)
			} else {
				style = style.Foreground(list.theme.TextColor)
			}

			if i != 0 && index != list.table.Cursor() {
				style = style.Foreground(list.theme.TextColor)
			}

			contents[i].WriteString(style.Render(content[i].String()))
			contents[i].WriteByte('\n')
		}
	}

	for i := 0; i < cellsLength; i++ {
		list.flexBox.Row(0).Cell(i).SetContent(contents[i].String())
	}

	return lipgloss.NewStyle().Margin(1, 0).Render(list.flexBox.Render())
}
