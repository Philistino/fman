package buttonbar

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	zone "github.com/lrstanley/bubblezone"
	"github.com/nore-dev/fman/message"
	"github.com/nore-dev/fman/model/dialog"
	"github.com/nore-dev/fman/theme"
)

// have to activate and deactivate buttons based on state
//	back and forward

// Buttons:
// 	- cut activate when item select
//	- copy activate when item select
//	- paste activate when item in clipboard
//	- rename activate when item select
// 	- delete activate when item select
//	- new file always active
//	- new folder always active
// 	- compress activate when item select
//	- extract activate when compressed item select

type ButtonBar struct {
	id                     string
	width                  int
	fileSelected           bool
	clipBoardEmpty         bool
	onlyCompressedSelected bool
}

func New() ButtonBar {
	return ButtonBar{}
}

func (m ButtonBar) Init() tea.Cmd {
	return nil
}

func (m ButtonBar) Update(msg tea.Msg) (ButtonBar, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
	case tea.MouseMsg:
		if msg.Type != tea.MouseLeft {
			return m, nil
		}
		switch {
		case zone.Get(m.id + "new file").InBounds(msg):
			d := dialog.Default("new file")
			d.SetTitle("Create new file?")
			cmd = message.UpdateDialog(&d)
			log.Println("Clicked new file!")
		case zone.Get(m.id + "new folder").InBounds(msg):
			log.Println("Clicked new folder!")
		case zone.Get(m.id + "cut").InBounds(msg):
			log.Println("Clicked cut!")
		case zone.Get(m.id + "copy").InBounds(msg):
			log.Println("Clicked copy!")
		case zone.Get(m.id + "paste").InBounds(msg):
			log.Println("Clicked paste!")
		case zone.Get(m.id + "rename").InBounds(msg):
			log.Println("Clicked rename!")
		case zone.Get(m.id + "delete").InBounds(msg):
			log.Println("Clicked delete!")
		case zone.Get(m.id + "compress").InBounds(msg):
			log.Println("Clicked compress!")
		case zone.Get(m.id + "extract").InBounds(msg):
		}
		return m, cmd
	}
	return m, cmd
}

func (m ButtonBar) View() string {
	sectionWrapper := lipgloss.NewStyle().
		Padding(0, 1, 1, 1)

	var newFile, newFolder, cut, copy, paste, rename, delete, compress, extract string

	newFile = theme.ButtonStyle.Render("New File")
	newFolder = theme.ButtonStyle.Render("New Folder")

	if m.fileSelected {
		cut = theme.ButtonStyle.Render("Cut")
		copy = theme.ButtonStyle.Render("Copy")
		rename = theme.ButtonStyle.Render("Rename")
		delete = theme.ButtonStyle.Render("Delete")
	} else {
		cut = theme.InactiveButtonStyle.Render("Cut")
		copy = theme.InactiveButtonStyle.Render("Copy")
		rename = theme.InactiveButtonStyle.Render("Rename")
		delete = theme.InactiveButtonStyle.Render("Delete")
	}

	if m.onlyCompressedSelected {
		compress = theme.InactiveButtonStyle.Render("Compress")
		extract = theme.ButtonStyle.Render("Extract")
	} else {
		compress = theme.ButtonStyle.Render("Compress")
		extract = theme.InactiveButtonStyle.Render("Extract")
	}

	buttons := lipgloss.JoinHorizontal(
		lipgloss.Top,
		sectionWrapper.Copy().BorderLeft(false).PaddingLeft(0).Render(
			lipgloss.JoinHorizontal(
				lipgloss.Top,
				zone.Mark(m.id+"new file", newFile),
				zone.Mark(m.id+"new folder", newFolder),
			),
		),
		sectionWrapper.Copy().BorderLeft(false).Render(
			lipgloss.JoinHorizontal(
				lipgloss.Top,
				zone.Mark(m.id+"cut", cut),
				zone.Mark(m.id+"copy", copy),
				zone.Mark(m.id+"paste", paste),
				zone.Mark(m.id+"rename", rename),
				zone.Mark(m.id+"delete", delete),
			),
		),
		sectionWrapper.Copy().BorderLeft(false).BorderRight(false).PaddingRight(0).Render(
			lipgloss.JoinHorizontal(
				lipgloss.Top,
				zone.Mark(m.id+"compress", compress),
				zone.Mark(m.id+"extract", extract),
			),
		),
	)
	return lipgloss.JoinHorizontal(lipgloss.Left, buttons)
}
