package api

import "quana.co.tz/karata/core"

type OpCode int

const (
	Start      OpCode = iota + 1 // New game round starting.
	NextTurn                     // A game round has just completed.
	CardPlayed                   //A move the player wishes to make and sends to the server.
	Rejected                     // Move was rejected.
	OpponentLeft
	Update = 7 // Opponent has left the game.
)

type Message interface {
}

type StartMessage struct {
	Board    core.Board
	Deadline int64
}

type UpdateMessage struct {
	Deck        core.Deck
	GameOver    bool
	CurrentCard core.Card
	CurrentTurn string
	Penalties   []core.Penalty
}

type DoneMessage struct {
	Winner string
}

type CardPlayedMessage struct {
	Card   core.Card
	UserId string
}

type RpcFindMatchResponse struct {
	MatchIds []string
}
