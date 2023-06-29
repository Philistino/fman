package list

import (
	"path/filepath"
	"time"

	"github.com/Philistino/fman/entry"
	"github.com/Philistino/fman/model/keys"
	"github.com/Philistino/fman/model/message"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	zone "github.com/lrstanley/bubblezone"
)

func (list *List) clearLastKey() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return message.ClearKeyMsg{}
	})
}

func (list *List) restrictIndex() {
	if list.cursorIdx < 0 {
		list.cursorIdx = len(list.entries) - 1
	} else if list.cursorIdx >= len(list.entries) {
		list.cursorIdx = 0
	}
}

func getFullPath(entry entry.Entry, path string) string {
	if entry.SymLinkPath != "" {
		return entry.SymLinkPath
	}

	return filepath.Join(path, entry.Name())
}

func (list *List) handlePathChange(newDir message.DirChangedMsg) tea.Cmd {
	if newDir.Error() != nil {
		return message.NewNotificationCmd(newDir.Error().Error())
	}
	list.selected = make(map[int]struct{})
	list.entries = newDir.Entries()
	selected := newDir.Selected()
	matched := false
	for i, entry := range list.entries {
		// set the cursor
		if entry.Name() == newDir.Cursor() {
			list.cursorIdx = i
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
		list.cursorIdx = 0
	}

	list.restrictIndex()
	list.flexBox.ForceRecalculate()
	if len(list.entries) == 0 {
		return nil
	}
	return message.NewEntryCmd(list.SelectedEntry())
}

func (list *List) handleMouseClick(msg tea.MouseMsg) tea.Cmd {
	if msg.Type != tea.MouseLeft || !zone.Get("list").InBounds(msg) {
		return nil
	}
	x, y := zone.Get("list").Pos(msg)
	offset := 2
	if (y < offset || y > len(list.entries)+offset-1) || x > list.width {
		return nil
	}
	list.cursorIdx = y + max(0, list.cursorIdx-list.maxEntryToShow) - offset

	// Double click
	time := time.Now()
	if time.Sub(list.lastClickedTime) < list.clickDelay && list.SelectedEntry().IsDir() && list.cursorIdx == list.lastClickedIdx {

		// If the user doesn't have permission to access the directory, return a notification
		if list.SelectedEntry().SizeStr == "Access Denied" {
			return message.NewNotificationCmd("Access Denied")
		}

		// Send message to update the preview pane
		list.lastClickedIdx = -1 // reset the last clicked index
		return message.NavDownCmd(list.SelectedEntry().Name())
	}
	list.lastClickedTime = time
	list.lastClickedIdx = list.cursorIdx

	// Send message to update the preview pane
	return message.NewEntryCmd(list.SelectedEntry())
}

func (list *List) resizeList() {
	list.flexBox.SetWidth(list.width)
	list.flexBox.SetHeight(list.height)
	list.flexBox.ForceRecalculate()
	list.truncateLimit = list.flexBox.Row(0).Cell(0).GetWidth() - 1
	list.maxEntryToShow = list.height * 3 / 4
}

func (list *List) Update(msg tea.Msg) (List, tea.Cmd) {
	if !list.focused {
		return *list, nil
	}

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		list.resizeList()
	case tea.MouseMsg:
		return *list, list.handleMouseClick(msg)
	case message.DirChangedMsg:
		return *list, list.handlePathChange(msg)
	case message.ClearKeyMsg:
		list.lastKeyCharacter = ' '
		return *list, list.clearLastKey()
	case tea.KeyMsg:
		switch {

		// Move this elsewhere TODO!!!
		// case key.Matches(msg, keys.Map.OpenFile): // Open file with default application
		// 	path := getFullPath(list.SelectedEntry(), list.path)
		// 	// If the file can be readable open the default editor for editing
		// 	if !list.SelectedEntry().IsDir() && isFileReadable(path) {
		// 		return *list, list.openEditor(path)
		// 	}
		// 	cmd := exec.Command(detectOpenCommand(), path)
		// 	cmd.Start()
		// 	return *list, nil

		// Move this elsewhere TODO!!!
		// case key.Matches(msg, keys.Map.CopyToClipboard): // Copy path to the clipboard
		// 	path := getFullPath(list.SelectedEntry(), list.path)
		// 	clipboard.WriteAll(path)
		// 	return *list, message.SendMessage("Copied!")

		case key.Matches(msg, keys.Map.GoToTop): // Move to the beginning of the list
			list.cursorIdx = 0
		case key.Matches(msg, keys.Map.GoToBottom): // Move to the end of the list
			list.cursorIdx = len(list.entries) - 1
		case key.Matches(msg, keys.Map.MoveCursorUp): // Select entry above
			if len(list.entries) == 0 {
				return *list, nil
			}
			if list.cursorIdx == 0 {
				return *list, nil
			}
			list.cursorIdx--
			list.restrictIndex()
			return *list, message.NewEntryCmd(list.SelectedEntry())
		case key.Matches(msg, keys.Map.MoveCursorDown): // Select entry below
			if len(list.entries) == 0 {
				return *list, nil
			}
			if list.cursorIdx == len(list.entries)-1 {
				return *list, nil
			}
			list.cursorIdx += 1
			list.restrictIndex()
			return *list, message.NewEntryCmd(list.SelectedEntry())
		case key.Matches(msg, keys.Map.GoToParentDirectory): // Get entries from parent directory
			return *list, message.NavUpCmd()
		case key.Matches(msg, keys.Map.GoToSelectedDirectory): // If the selected entry is a directory. Get entries under that directory
			if len(list.entries) == 0 {
				return *list, nil
			}
			if !list.SelectedEntry().IsDir() {
				return *list, nil
			}
			if list.SelectedEntry().SizeStr == "Access Denied" {
				return *list, message.NewNotificationCmd("Access Denied")
			}
			return *list, message.NavDownCmd(list.SelectedEntry().Name())

		// TODO: Move this elsewhere
		case key.Matches(msg, keys.Map.GoBack):
			return *list, message.NavBackCmd()
		case key.Matches(msg, keys.Map.GoForward):
			return *list, message.NavFwdCmd()
		case key.Matches(msg, keys.Map.GoToHomeDirectory): // Move to the home directory
			return *list, message.NavHomeCmd()
		case key.Matches(msg, keys.Map.ShowHiddenEntries): // Show hidden files
			return *list, message.ToggleShowHiddenCmd()
		}
	}
	list.restrictIndex()
	return *list, nil

}
