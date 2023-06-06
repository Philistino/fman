package message

import (
	"bufio"
	"os"
	"os/exec"
	"runtime"
	"unicode/utf8"

	"github.com/Philistino/fman/entry"
	"github.com/Philistino/fman/model/dialog"
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

type UpdateDialogMsg struct {
	Dialog dialog.Dialog
}

// InternalCopyMsg is used to communicate to the main program
// that a "clipboard" copy operation is requested.
type InternalCopyMsg struct{}

// InternalCopyCmd is used to create a command that will
// communicate to the main program that a "clipboard" copy
// operation is requested.
func InternalCopyCmd() tea.Cmd {
	return func() tea.Msg {
		return InternalCopyMsg{}
	}
}

// CutMsg is used to communicate to the main program
// that a "clipboard" cut operation is requested.
type CutMsg struct{}

// CutCmd is used to create a command that will
// communicate to the main program that a "clipboard" cut
// operation is requested.
func CutCmd() tea.Cmd {
	return func() tea.Msg {
		return CutMsg{}
	}
}

// RenameMsg is used to communicate to the main program
// that a rename operation is requested.
type RenameMsg struct{}

// RenameCmd is used to create a command that will
// communicate to the main program that a rename
// operation is requested.
func RenameCmd() tea.Cmd {
	return func() tea.Msg {
		return RenameMsg{}
	}
}

// NewFileMsg is used to communicate to the main program
// that a new file operation is requested.
type NewFileMsg struct{}

// NewFileCmd is used to create a command that will
// communicate to the main program that a new file
// operation is requested.
func NewFileCmd() tea.Cmd {
	return func() tea.Msg {
		return NewFileMsg{}
	}
}

// DeleteMsg is used to communicate to the main program
// that a delete operation is requested.
type DeleteMsg struct{}

// DeleteCmd is used to create a command that will
// communicate to the main program that a delete
// operation is requested.
func DeleteCmd() tea.Cmd {
	return func() tea.Msg {
		return DeleteMsg{}
	}
}

type InternalPasteMsg struct{}

func InternalPasteCmd() tea.Cmd {
	return func() tea.Msg {
		return InternalPasteMsg{}
	}
}

// MkDirMsg is used to communicate to the main program
// that a new directory operation is requested.
type MkDirMsg struct{}

// MkDirCmd is used to create a command that will
// communicate to the main program that a new directory
// operation is requested.
func MkDirCmd() tea.Cmd {
	return func() tea.Msg {
		return MkDirMsg{}
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

// NavBackMsg is used to communicate to the main program
// that a navigation back operation is requested.
type NavBackMsg struct{}

// NavBackCmd is used to create a command that will
// communicate to the main program that a navigation back
// operation is requested.
func NavBackCmd() tea.Cmd {
	return func() tea.Msg {
		return NavBackMsg{}
	}
}

// NavFwdMsg is used to communicate to the main program
// that a navigation forward operation is requested.
type NavFwdMsg struct{}

// NavFwdCmd is used to create a command that will
// communicate to the main program that a navigation forward
// operation is requested.
func NavFwdCmd() tea.Cmd {
	return func() tea.Msg {
		return NavFwdMsg{}
	}
}

// NavUpMsg is used to communicate to the main program
// that a navigation to parent directory operation is requested.
type NavUpMsg struct{}

// NavUpCmd is used to create a command that will
// communicate to the main program that a navigation to the parent directory
// operation is requested.
func NavUpCmd() tea.Cmd {
	return func() tea.Msg {
		return NavUpMsg{}
	}
}

// NavHomeMsg is used to communicate to the main program
// that a navigation to the home directory operation is requested.
type NavHomeMsg struct{}

// NavHomeCmd is used to create a command that will
// communicate to the main program that a navigation to the home directory
// operation is requested.
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

func UpdateDialog(dialog *dialog.Dialog) tea.Cmd {
	return func() tea.Msg {
		return UpdateDialogMsg{Dialog: *dialog}
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
