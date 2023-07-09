package infobar

import (
	"time"

	"github.com/Philistino/fman/ui/infobar/queue"
	"github.com/Philistino/fman/ui/message"
	tea "github.com/charmbracelet/bubbletea"
)

const DEFAULT_MESSAGE = "--"

type notifications struct {
	width           int                 // the width of the infobar, which should be the window width
	progressWidth   int                 // the width of the progress bar
	message         string              // the current message
	defaultDuration time.Duration       // the default duration to display a message
	minDuration     time.Duration       // the minimum duration to display a message
	startTime       time.Time           // the time the current message was displayed
	messageId       uint16              // Used to track notifications for clearing. its ok if this wraps back around. It just needs to be big enough to not collide with the most recent messages
	stack           *queue.Fifo[string] // a stack of messages to display
}

func newNotifications() notifications {
	return notifications{
		width:           0,
		progressWidth:   20,
		message:         DEFAULT_MESSAGE,
		defaultDuration: time.Second * 1,
		minDuration:     time.Millisecond * 250,
		startTime:       time.Now(),
		messageId:       0,
		stack:           queue.NewFifo[string](),
	}
}

func (m *notifications) Init() tea.Cmd {
	return nil
}

func (m *notifications) Message() string {
	return m.message
}

type clearNotificationMsg struct {
	id uint16
}

func (m *notifications) clearMessageCmd(id uint16, t time.Duration) tea.Cmd {
	return func() tea.Msg {
		time.Sleep(t)
		return clearNotificationMsg{id: id}
	}
}

func (m *notifications) Update(msg tea.Msg) (notifications, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case message.DirChangedMsg:
		if msg.Error() != nil {
			cmd = message.NewNotificationCmd("Error: " + msg.Error().Error())
			return *m, cmd
		}
		m.stack.Clear()
	case message.NewNotificationMsg:
		cmd = m.handleNewMessage(msg)
	case clearNotificationMsg:
		cmd = m.handleClearNotificationMsg(msg)
	}
	return *m, cmd
}

func (m *notifications) handleNewMessage(msg message.NewNotificationMsg) tea.Cmd {
	if time.Since(m.startTime) < m.minDuration {
		m.stack.Push(DEFAULT_MESSAGE)
		m.stack.Push(msg.Message)
		return m.clearMessageCmd(m.messageId, m.minDuration-time.Since(m.startTime))
	}

	if m.message == msg.Message {
		// set the current message to the default message for the minduration so it is seen between messages
		// and push the new message to the stack to be displayed after the minDuration for the default message has passed
		m.stack.Push(msg.Message)
		m.startTime = time.Now()
		m.message = DEFAULT_MESSAGE
		m.messageId++
		return m.clearMessageCmd(m.messageId, m.minDuration)
	}

	m.message = msg.Message
	m.startTime = time.Now()
	m.messageId++
	return m.clearMessageCmd(m.messageId, m.defaultDuration)
}

func (m *notifications) handleClearNotificationMsg(msg clearNotificationMsg) tea.Cmd {
	if m.messageId != msg.id {
		return nil
	}
	if m.stack.IsEmpty() {
		m.message = DEFAULT_MESSAGE
		m.startTime = time.Now()
		m.messageId++
		return nil
	}

	noti, _ := m.stack.Pop()
	m.startTime = time.Now()
	m.message = *noti
	m.messageId++
	clearDuration := m.defaultDuration
	if !m.stack.IsEmpty() {
		clearDuration = m.minDuration
	}
	return m.clearMessageCmd(m.messageId, clearDuration)
}

func (m *notifications) View() string {
	return m.message
}
