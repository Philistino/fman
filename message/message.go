package message

import (
	"os"
	"os/exec"
	"runtime"

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
	Path     string
	Entries  []entry.Entry
	Selected map[string]struct{}
	err      error
}

func (d DirChangedMsg) Error() error {
	return d.err
}

func handleNav(entries []entry.Entry, state nav.NavState, err error) tea.Msg {
	if err != nil {
		return SendMessage(err.Error())
	}
	return DirChangedMsg{
		Path:     state.Path(),
		Entries:  entries,
		Selected: state.Selected(),
	}
}

func HandleFwdCmd(navi *nav.Nav, currentSelected []string) tea.Cmd {
	return func() tea.Msg {
		return handleNav(
			navi.Forward(
				currentSelected,
			),
		)
	}
}

func HandleBackCmd(navi *nav.Nav, currentSelected []string) tea.Cmd {
	return func() tea.Msg {
		return handleNav(
			navi.Back(
				currentSelected,
			),
		)
	}
}

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
