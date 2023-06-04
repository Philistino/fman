package model

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	zone "github.com/lrstanley/bubblezone"
	"github.com/nore-dev/fman/message"
	"github.com/nore-dev/fman/theme"
)

// fileBtns handle the file manipulation buttons in the toolbar

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

type fileBtns struct {
	id            string
	width         int
	fileSelected  bool
	clipBoardFull bool
	focused       bool
	// selectedIsArchive bool
}

func newFileBtns() fileBtns {
	return fileBtns{
		id:            "buttonbar",
		fileSelected:  false,
		clipBoardFull: false,
		focused:       true,
	}
}

func (m fileBtns) Init() tea.Cmd {
	return nil
}

func (m fileBtns) Update(msg tea.Msg) (fileBtns, tea.Cmd) {
	if !m.focused {
		return m, nil
	}

	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
	case message.EntryMsg:
		m.fileSelected = true
	case message.DirChangedMsg:
		m.fileSelected = false
	case message.InternalCopyMsg:
		m.clipBoardFull = true
	case tea.MouseMsg:
		if msg.Type != tea.MouseLeft {
			return m, nil
		}
		switch {
		case zone.Get(m.id + "new file").InBounds(msg):
			cmd = message.NewFileCmd()
		case zone.Get(m.id + "new folder").InBounds(msg):
			cmd = message.MkDirCmd()
		case zone.Get(m.id+"cut").InBounds(msg) && m.fileSelected:
			cmd = message.CutCmd()
		case zone.Get(m.id+"copy").InBounds(msg) && m.fileSelected:
			cmd = message.InternalCopyCmd()
		case zone.Get(m.id+"paste").InBounds(msg) && m.clipBoardFull:
			cmd = message.InternalPasteCmd()
		case zone.Get(m.id+"rename").InBounds(msg) && m.fileSelected:
			cmd = message.RenameCmd()
		case zone.Get(m.id+"delete").InBounds(msg) && m.fileSelected:
			cmd = message.DeleteCmd()
			// case zone.Get(m.id + "compress").InBounds(msg):
			// case zone.Get(m.id + "extract").InBounds(msg):
		}
	}
	return m, cmd
}

func (m fileBtns) View() string {
	sectionWrapper := lipgloss.NewStyle().
		Padding(0, 1, 1, 1)

	var newFile, newFolder, cut, copy, paste, rename, delete string

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

	if m.clipBoardFull {
		paste = theme.ButtonStyle.Render("Paste")
	} else {
		paste = theme.InactiveButtonStyle.Render("Paste")
	}

	// var compress, extract string
	// if !m.fileSelected {
	// 	compress = theme.InactiveButtonStyle.Render("Compress")
	// 	extract = theme.InactiveButtonStyle.Render("Extract")
	// } else if m.selectedIsArchive {
	// 	compress = theme.InactiveButtonStyle.Render("Compress")
	// 	extract = theme.ButtonStyle.Render("Extract")
	// } else {
	// 	compress = theme.ButtonStyle.Render("Compress")
	// 	extract = theme.InactiveButtonStyle.Render("Extract")
	// }

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
		// sectionWrapper.Copy().BorderLeft(false).BorderRight(false).PaddingRight(0).Render(
		// 	lipgloss.JoinHorizontal(
		// 		lipgloss.Top,
		// 		zone.Mark(m.id+"compress", compress),
		// 		zone.Mark(m.id+"extract", extract),
		// 	),
		// ),
	)
	return lipgloss.JoinHorizontal(lipgloss.Left, buttons)
}
