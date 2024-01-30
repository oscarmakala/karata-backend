package events

type EventType int

const (
	EventTypeStart      EventType = iota + 1 // New game round starting.
	EventTypeNextTurn                        // A game round has just completed.
	EventTypeCardPlayed                      //A move the player wishes to make and sends to the server.
	EventTypeRejected                        // Move was rejected.
	EventTypeOpponentLeft
	EventTypeUpdate = 7 // Opponent has left the game.
)
