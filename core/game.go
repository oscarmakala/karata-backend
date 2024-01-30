package core

type GameHandler interface {
	ActionStartGame(players []PlayerState, nbCards int) Board
	ActionPlayCard(board *Board, cardToPlay *Card, playerID string) error
	ActionNextTurn(board *Board, playerID string, isPlaying bool) error
	ActionDrawCard(board *Board, playerID string, cardToDraw []Card) error
}
