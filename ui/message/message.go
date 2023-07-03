package message

import (
	"bufio"
	"context"
	"os"
	"os/exec"
	"runtime"
	"unicode/utf8"

	"github.com/Philistino/fman/entry"
	"github.com/Philistino/fman/nav"
	tea "github.com/charmbracelet/bubbletea"
)

type ClearKeyMsg struct{}

type NewEntryMsg struct {
	Entry entry.Entry
}

func NewEntryCmd(newEntry entry.Entry) tea.Cmd {
	return func() tea.Msg {
		return NewEntryMsg{Entry: newEntry}
	}
}

type SelectedMsg struct {
	selected map[string]struct{}
}

func SelectedCmd(selected map[string]struct{}) tea.Cmd {
	return func() tea.Msg {
		return SelectedMsg{selected}
	}
}

// type EmptyDirMsg struct{}

// func EmptyDirCmd() tea.Cmd {
// 	return func() tea.Msg {
// 		return EmptyDirMsg{}
// 	}
// }

type GetPreviewMsg struct {
	Ctx  context.Context
	Path string
}

func GetPreviewCmd(ctx context.Context, path string) tea.Cmd {
	return func() tea.Msg {
		return GetPreviewMsg{ctx, path}
	}
}

type NewNotificationMsg struct {
	Message string
}

// NewNotificationCmd is used to create a command that will
// display a notification.
func NewNotificationCmd(message string) tea.Cmd {
	return func() tea.Msg {
		return NewNotificationMsg{message}
	}
}

// ToggleShowHiddenMsg is used to communicate to the main program
// that a toggle show hidden operation is requested.
type ToggleShowHiddenMsg struct{}

// ToggleShowHiddenCmd is used to create a command that will
// communicate to the main program that a toggle show hidden
// operation is requested.
func ToggleShowHiddenCmd() tea.Cmd {
	return func() tea.Msg {
		return ToggleShowHiddenMsg{}
	}
}

// DirChangedMsg is used to communicate that the CWD has changed
type DirChangedMsg struct {
	nav.DirState
}

func handleNav(state nav.DirState) tea.Msg {
	return DirChangedMsg{state}
}

// HandleFwdCmd fetches the forward state of the nav and returns a message to
// broadcast the new state
func HandleFwdCmd(navi *nav.Nav, currentSelected []string, cursor string) tea.Cmd {
	return func() tea.Msg {
		return handleNav(
			navi.Forward(
				currentSelected,
				cursor,
			),
		)
	}
}

// HandleBackCmd fetches the last state of the nav and returns a message to
// broadcast the new state
func HandleBackCmd(navi *nav.Nav, currentSelected []string, cursor string) tea.Cmd {
	return func() tea.Msg {
		return handleNav(
			navi.Back(
				currentSelected,
				cursor,
			),
		)
	}
}

// HandleNavCmd fetches a new state of the nav and returns a message to
// broadcast the new state
func HandleNavCmd(navi *nav.Nav, currentSelected []string, path string, cursor string) tea.Cmd {
	return func() tea.Msg {
		return handleNav(
			navi.Go(
				path,
				cursor,
				currentSelected,
			),
		)
	}
}

// HandleReloadCmd reloads the current directory and returns a message to
// broadcast the returned state state
func HandleReloadCmd(navi *nav.Nav, currentSelected []string, cursor string) tea.Cmd {
	return func() tea.Msg {
		return handleNav(
			navi.Reload(
				currentSelected,
				cursor,
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
			NewNotificationCmd(err.Error()),
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
