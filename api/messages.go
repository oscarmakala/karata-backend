package api

import "quana.co.tz/karata/core"

type Message interface {
}

type StartMessage struct {
	Board    core.Board
	Deadline int64
}

type Action int

type UpdateMessage struct {
	Deck        core.Deck
	GameOver    bool
	CurrentCard core.Card
	CurrentTurn string
	CardPlayed  core.Card
	Penalties   []core.Penalty
}

type DoneMessage struct {
	Winner string
}

type RpcFindMatchResponse struct {
	MatchIds []string
}
