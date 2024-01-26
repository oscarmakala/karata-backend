package core

type Service interface {
	ActionStartGame(players []PlayerState, nbCards int)
	ActionPlayCard(board *Board, cardToPlay *Card, playerID string)
	ActionNextTurn(board *Board, playerID string, isPlaying bool)
	ActionDrawCard(board *Board, playerID string, cardToDraw *Card)
}

type ServiceImpl struct{}
