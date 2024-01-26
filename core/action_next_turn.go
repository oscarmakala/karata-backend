package core

func (gs *ServiceImpl) ActionNextTurn(board *Board, playerID string, isPlaying bool) error {

	if isPlaying == false {
		return nil
	}
	currentPlayerIndex := FindPlayerIndex(board.Players, playerID)
	if currentPlayerIndex == -1 {
		return errNoSuchPlayer
	}

	CheckIfCurrentPlayerIsForcedToTakeCards(board, playerID)

	nextPlayerIndex := IncreasePlayerIndex(board, currentPlayerIndex)
	CheckIfNextPlayerIsForcedToTakeCards(board, nextPlayerIndex)
	//empty dropzone
	board.DropZone = board.DropZone[:0]
	return nil
}

func CheckIfCurrentPlayerIsForcedToTakeCards(board *Board, playerID string) {
	numberOfCards := 0
	penaltyAction := None
	if len(board.DropZone) == 0 {
		penaltyAction = Draw
	}
	board.Penalties = append(board.Penalties, Penalty{
		UserId:        playerID,
		NumberOfCards: numberOfCards,
		Penalty:       penaltyAction,
	})

}

func CheckIfNextPlayerIsForcedToTakeCards(board *Board, indexOfPlayer int) {
	player := board.Players[indexOfPlayer]
	if player.hasCardsToDefend(board.CurrentCard) {
		board.CanDefend = true
	} else {
		TakePenaltyCards(player, board)
	}
}

func TakePenaltyCards(player PlayerState, board *Board) {
	factor := len(board.DropZone)
	switch board.CurrentCard.Rank {
	case Plus2:
		//take two
		board.Penalties = append(board.Penalties, Penalty{
			UserId:        player.UserID,
			NumberOfCards: 2 * factor,
			Penalty:       Draw,
		})
	case Seven:
		board.Penalties = append(board.Penalties, Penalty{
			UserId:        player.UserID,
			NumberOfCards: 0,
			Penalty:       Skip,
		})
	}
}

func IncreasePlayerIndex(board *Board, playerIndex int) int {
	return (playerIndex + 1) % len(board.PlayerOrder)

}
