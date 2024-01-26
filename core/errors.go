package core

import "github.com/heroiclabs/nakama-common/runtime"

var (
	errNoSuchPlayer     = runtime.NewError("No such Player", 100)
	errNoSuchCardInHand = runtime.NewError("No such card in player hand", 101)
	errInvalidPlay      = runtime.NewError("Invalid card play", 102)
)
