package infobar

import (
	"log"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/dustin/go-humanize"
	"github.com/nore-dev/fman/model/message"
	"github.com/nore-dev/fman/storage"
	"github.com/nore-dev/fman/theme"
)

type Infobar struct {
	width           int                 // the width of the infobar, which should be the window width
	progressWidth   int                 // the width of the progress bar
	message         string              // the current message
	defaultDuration time.Duration       // the default duration to display a message
	minDuration     time.Duration       // the minimum gap to display the default message between messages
	startTime       time.Time           // the time the current message was displayed
	messageIdx      int                 // Used to track messages for clearing when multiple messages are to be displayed in a short time
	storageInfo     storage.StorageInfo // the disk storage info
	stack           Stack[string]       // a stack of messages to display
}

type TickMsg time.Time

const DEFAULT_MESSAGE = "--"

func New() Infobar {
	info, err := storage.GetStorageInfo()
	if err != nil {
		log.Println("An error occurred while getting storage info", err.Error())
	}
	return Infobar{
		progressWidth:   20,
		defaultDuration: time.Second * 2,
		minDuration:     time.Millisecond * 500,
		message:         DEFAULT_MESSAGE,
		storageInfo:     info,
		startTime:       time.Now(),
	}
}

func (infobar *Infobar) Init() tea.Cmd {
	return nil
}

func (infobar *Infobar) Message() string {
	return infobar.message
}

type clearMessageMsg struct {
	idx int
}

func (infobar *Infobar) clearMessage(idx int, t time.Duration) tea.Cmd {
	return func() tea.Msg {
		// It's ok to sleep here without a context. In bubbletea this will
		// not block quitting the program.
		time.Sleep(t)
		return clearMessageMsg{idx: idx}
	}
}

func (infobar *Infobar) Update(msg tea.Msg) (Infobar, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case message.NewNotificationMsg:
		cmd = infobar.handleNewMessage(msg)
	case clearMessageMsg:
		if infobar.messageIdx != msg.idx {
			return *infobar, nil
		}
		if !infobar.stack.IsEmpty() {
			noti, err := infobar.stack.Pop()
			if err != nil {
				log.Println("An error occurred while popping a notification from the stack", err.Error())
			}
			infobar.startTime = time.Now()
			infobar.message = *noti
			infobar.messageIdx++
			cmd = infobar.clearMessage(infobar.messageIdx, infobar.defaultDuration)
		}
		infobar.message = DEFAULT_MESSAGE
		infobar.startTime = time.Now()
		infobar.messageIdx++
	case tea.WindowSizeMsg:
		infobar.width = msg.Width
	}
	return *infobar, cmd
}

func (i *Infobar) handleNewMessage(msg message.NewNotificationMsg) tea.Cmd {
	if time.Since(i.startTime) < i.minDuration {
		// the current message was not displayed long enough so push the new message to the stack
		// and clear the current message after the minGapDuration has passed
		i.stack.Push(msg.Message)
		return i.clearMessage(i.messageIdx, i.minDuration-time.Since(i.startTime))
	}

	if i.message == msg.Message {
		// set the current message to the default message for the minduration so it is seen between messages
		// and push the new message to the stack to be displayed after the minDuration for the default message has passed
		i.stack.Push(msg.Message)
		i.startTime = time.Now()
		i.message = DEFAULT_MESSAGE
		i.messageIdx++
		return i.clearMessage(i.messageIdx, i.minDuration)
	}

	i.message = msg.Message
	i.startTime = time.Now()
	i.messageIdx++
	return i.clearMessage(i.messageIdx, i.defaultDuration)
}

func renderProgress(width int, usedSpace uint64, totalSpace uint64) string {
	usedWidth := (int(usedSpace) * width / int(totalSpace))
	usedStr := strings.Repeat("█", int(width-usedWidth))
	return theme.ProgressStyle.Width(width).Render(usedStr)
}

func (infobar *Infobar) View() string {
	info := infobar.storageInfo
	logo := theme.LogoStyle.Render(string(theme.GetActiveIconTheme().GopherIcon) + "FMAN")
	progress := renderProgress(infobar.progressWidth, info.AvailableSpace, info.TotalSpace)
	style := theme.InfobarStyle
	usedSpace := lipgloss.JoinHorizontal(lipgloss.Center, " ", humanize.Bytes(info.AvailableSpace), "/", humanize.Bytes(info.TotalSpace), " free")
	return lipgloss.JoinHorizontal(
		lipgloss.Center,
		logo,
		style.Width(infobar.width-(lipgloss.Width(progress)+lipgloss.Width(usedSpace)+lipgloss.Width(logo)+1)).Render(" "+infobar.Message()),
		progress,
		style.Render(usedSpace),
	)
}
