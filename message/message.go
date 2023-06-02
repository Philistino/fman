package message

import (
	"bufio"
	"os"
	"os/exec"
	"runtime"
	"unicode/utf8"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/nore-dev/fman/entry"
	"github.com/nore-dev/fman/model/dialog"
	"github.com/nore-dev/fman/nav"
)

type UpdateEntriesMsg struct {
	Parent bool
}

type ClearKeyMsg struct {
}

type PathMsg struct {
	Path string
}

type EntryMsg struct {
	Entry entry.Entry
}

type NewMessageMsg struct {
	Message string
}

type UpdateDialogMsg struct {
	Dialog dialog.Dialog
}

type ToggleShowHiddenMsg struct{}

func ToggleShowHiddenCmd() tea.Cmd {
	return func() tea.Msg {
		return ToggleShowHiddenMsg{}
	}
}

type NavBackMsg struct{}

func NavBackCmd() tea.Cmd {
	return func() tea.Msg {
		return NavBackMsg{}
	}
}

type NavFwdMsg struct{}

func NavFwdCmd() tea.Cmd {
	return func() tea.Msg {
		return NavFwdMsg{}
	}
}

type NavUpMsg struct{}

func NavUpCmd() tea.Cmd {
	return func() tea.Msg {
		return NavUpMsg{}
	}
}

type NavHomeMsg struct{}

func NavHomeCmd() tea.Cmd {
	return func() tea.Msg {
		return NavHomeMsg{}
	}
}

type NavDownMsg struct {
	Name string
}

func NavDownCmd(name string) tea.Cmd {
	return func() tea.Msg {
		return NavDownMsg{
			Name: name,
		}
	}
}

type NavOtherMsg struct {
	Path string
}

func NavOtherCmd(path string) tea.Cmd {
	return func() tea.Msg {
		return NavOtherMsg{
			Path: path,
		}
	}
}

func ChangePath(path string) tea.Cmd {
	return func() tea.Msg {
		return PathMsg{Path: path}
	}
}

func UpdateEntry(newEntry entry.Entry) tea.Cmd {
	return func() tea.Msg {
		return EntryMsg{Entry: newEntry}
	}
}

func SendMessage(message string) tea.Cmd {
	return func() tea.Msg {
		return NewMessageMsg{message}
	}
}

func UpdateDialog(dialog *dialog.Dialog) tea.Cmd {
	return func() tea.Msg {
		return UpdateDialogMsg{Dialog: *dialog}
	}
}

type DirChangedMsg struct {
	nav.DirState
}

func handleNav(state nav.DirState) tea.Msg {
	return DirChangedMsg{state}
}

// HandleFwdCmd fetches the forward state of the nav and returns a message to
// broadcast the new state
func HandleFwdCmd(navi *nav.Nav, currentSelected []string) tea.Cmd {
	return func() tea.Msg {
		return handleNav(
			navi.Forward(
				currentSelected,
			),
		)
	}
}

// HandleBackCmd fetches the last state of the nav and returns a message to
// broadcast the new state
func HandleBackCmd(navi *nav.Nav, currentSelected []string) tea.Cmd {
	return func() tea.Msg {
		return handleNav(
			navi.Back(
				currentSelected,
			),
		)
	}
}

// HandleNavCmd fetches a new state of the nav and returns a message to
// broadcast the new state
func HandleNavCmd(navi *nav.Nav, currentSelected []string, path string) tea.Cmd {
	return func() tea.Msg {
		return handleNav(
			navi.Go(
				path,
				currentSelected,
			),
		)
	}
}

// HandleReloadCmd reloads the current directory and returns a message to
// broadcast the returned state state
func HandleReloadCmd(navi *nav.Nav, currentSelected []string) tea.Cmd {
	return func() tea.Msg {
		return handleNav(
			navi.Reload(
				currentSelected,
			),
		)
	}
}

func openEditor(path string) tea.Cmd {
	const fallBackEditor = "nano"

	editor := os.Getenv("EDITOR")

	if editor == "" {
		editor = fallBackEditor
	}
	cmd := exec.Command(editor, path)
	return tea.ExecProcess(cmd, func(err error) tea.Msg {
		if err == nil {
			return tea.EnableMouseCellMotion
		}

		// Failed to open editor, open with default app instead
		cmd := exec.Command(detectOpenCommand(), path)
		cmd.Start()

		return tea.Batch(
			SendMessage(err.Error()),
			tea.EnableMouseCellMotion,
		)
	})
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

func isFileReadable(path string) bool {
	file, err := os.Open(path)

	if err != nil {
		return false
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	return utf8.ValidString(scanner.Text())
}
