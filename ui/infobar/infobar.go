package infobar

import (
	"log"
	"strings"
	"time"

	"github.com/Philistino/fman/entry/storage"
	"github.com/Philistino/fman/ui/infobar/queue"
	"github.com/Philistino/fman/ui/message"

	"github.com/Philistino/fman/ui/theme"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/dustin/go-humanize"
)

type TickMsg time.Time

const DEFAULT_MESSAGE = "--"

type Infobar struct {
	width           int                 // the width of the infobar, which should be the window width
	progressWidth   int                 // the width of the progress bar
	message         string              // the current message
	defaultDuration time.Duration       // the default duration to display a message
	minDuration     time.Duration       // the minimum duration to display a message
	startTime       time.Time           // the time the current message was displayed
	messageId       uint16              // Used to track notifications for clearing. its ok if this wraps back around. It just needs to be big enough to not collide with the most recent messages
	storageInfo     storage.StorageInfo // the disk storage info
	stack           *queue.Fifo[string] // a stack of messages to display
}

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

func (infobar *Infobar) clearMessageCmd(id uint16, t time.Duration) tea.Cmd {
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
		cmd = infobar.handleClearNotificationMsg(msg)
	case tea.WindowSizeMsg:
		infobar.width = msg.Width
	}
	return *infobar, cmd
}

func (i *Infobar) handleNewMessage(msg message.NewNotificationMsg) tea.Cmd {
	if time.Since(i.startTime) < i.minDuration {
		i.stack.Push(DEFAULT_MESSAGE)
		i.stack.Push(msg.Message)
		return i.clearMessageCmd(i.messageId, i.minDuration-time.Since(i.startTime))
	}

	if i.message == msg.Message {
		// set the current message to the default message for the minduration so it is seen between messages
		// and push the new message to the stack to be displayed after the minDuration for the default message has passed
		i.stack.Push(msg.Message)
		i.startTime = time.Now()
		i.message = DEFAULT_MESSAGE
		i.messageId++
		return i.clearMessageCmd(i.messageId, i.minDuration)
	}

	i.message = msg.Message
	i.startTime = time.Now()
	i.messageId++
	return i.clearMessageCmd(i.messageId, i.defaultDuration)
}

func (infobar *Infobar) handleClearNotificationMsg(msg clearNotificationMsg) tea.Cmd {
	if infobar.messageId != msg.id {
		return nil
	}
	if infobar.stack.IsEmpty() {
		infobar.message = DEFAULT_MESSAGE
		infobar.startTime = time.Now()
		infobar.messageId++
		return nil
	}

	noti, _ := infobar.stack.Pop()
	infobar.startTime = time.Now()
	infobar.message = *noti
	infobar.messageId++
	clearDuration := infobar.defaultDuration
	if !infobar.stack.IsEmpty() {
		clearDuration = infobar.minDuration
	}
	return infobar.clearMessageCmd(infobar.messageId, clearDuration)
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
		style.Width(infobar.width-(lipgloss.Width(progress)+lipgloss.Width(usedSpace)+lipgloss.Width(logo)+1)).Render(" "+infobar.message),
		progress,
		style.Render(usedSpace),
	)
}
