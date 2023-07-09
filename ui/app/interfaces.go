package app

import (
	"github.com/Philistino/fman/ui/dialog"
	"github.com/Philistino/fman/ui/message"
	"github.com/Philistino/fman/ui/navbtns"
)

// checking that concrete types implement the relevant interfaces

var _ dialog.AskMsg = new(message.AskDialogGeneric)

var _ navbtns.ActiveNavBtns = new(message.DirChangedMsg)
