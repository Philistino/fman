package infobar

import (
	"log"
	"strings"
	"time"

	"github.com/Philistino/fman/storage"
	"github.com/Philistino/fman/ui/infobar/queue"
	"github.com/Philistino/fman/ui/message"

	"github.com/Philistino/fman/ui/theme"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/dustin/go-humanize"
)

type Infobar struct {
	width           int                 // the width of the infobar, which should be the window width
	progressWidth   int                 // the width of the progress bar
	message         string              // the current message
	defaultDuration time.Duration       // the default duration to display a message
	minDuration     time.Duration       // the minimum gap to display the default message between messages
	startTime       time.Time           // the time the current message was displayed
	messageId       uint16              // Used to track messages for clearing when multiple messages are to be displayed in a short time
	storageInfo     storage.StorageInfo // the disk storage info
	stack           *queue.Fifo[string] // a stack of messages to display
}

type TickMsg time.Time

const DEFAULT_MESSAGE = "--"

func New() Infobar {
	info, err := storage.GetStorageInfo()
	if err != nil {
		log.Println("An error occurred while getting storage info", err.Error())
	}
	return Infobar{
		width:           0,
		progressWidth:   20,
		message:         DEFAULT_MESSAGE,
		defaultDuration: time.Second * 1,
		minDuration:     time.Millisecond * 250,
		startTime:       time.Now(),
		messageId:       0,
		storageInfo:     info,
		stack:           queue.NewFifo[string](),
	}
}

func (infobar *Infobar) Init() tea.Cmd {
	return nil
}

func (infobar *Infobar) Message() string {
	return infobar.message
}

type clearNotificationMsg struct {
	id uint16
}

func (infobar *Infobar) clearMessage(id uint16, t time.Duration) tea.Cmd {
	return func() tea.Msg {
		time.Sleep(t)
		return clearNotificationMsg{id: id}
	}
}

func (infobar *Infobar) Update(msg tea.Msg) (Infobar, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case message.DirChangedMsg:
		if msg.Error() != nil {
			cmd = message.NewNotificationCmd("Error: " + msg.Error().Error())
			return *infobar, cmd
		}
		infobar.stack.Clear()
	case message.NewNotificationMsg:
		cmd = infobar.handleNewMessage(msg)
	case clearNotificationMsg:
		if infobar.messageId != msg.id {
			return *infobar, nil
		}
		if infobar.stack.IsEmpty() {
			infobar.message = DEFAULT_MESSAGE
			infobar.startTime = time.Now()
			infobar.messageId++
			return *infobar, nil
		}

		noti, err := infobar.stack.Pop()
		if err != nil {
			log.Println("An error occurred while popping a notification from the stack", err.Error())
		}
		infobar.startTime = time.Now()
		infobar.message = *noti
		infobar.messageId++
		clearDuration := infobar.defaultDuration
		if !infobar.stack.IsEmpty() {
			clearDuration = infobar.minDuration
		}
		cmd = infobar.clearMessage(infobar.messageId, clearDuration)

	case tea.WindowSizeMsg:
		infobar.width = msg.Width
	}
	return *infobar, cmd
}

func (i *Infobar) handleNewMessage(msg message.NewNotificationMsg) tea.Cmd {
	if time.Since(i.startTime) < i.minDuration {
		// the current message was not displayed long enough so push the new message to the stack
		// and clear the current message after the minGapDuration has passed
		i.stack.Push(DEFAULT_MESSAGE)
		i.stack.Push(msg.Message)
		return i.clearMessage(i.messageId, i.minDuration-time.Since(i.startTime))
	}

	if i.message == msg.Message {
		// set the current message to the default message for the minduration so it is seen between messages
		// and push the new message to the stack to be displayed after the minDuration for the default message has passed
		i.stack.Push(msg.Message)
		i.startTime = time.Now()
		i.message = DEFAULT_MESSAGE
		i.messageId++
		return i.clearMessage(i.messageId, i.minDuration)
	}

	i.message = msg.Message
	i.startTime = time.Now()
	i.messageId++
	return i.clearMessage(i.messageId, i.defaultDuration)
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
