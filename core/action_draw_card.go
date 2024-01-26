package core

// ActionDrawCard DrawCard uses Client-Side Prediction for drawing cards and reconciliation is done on the ser
func (gs *ServiceImpl) ActionDrawCard(board *Board, playerID string, cardToDraw *Card) error {

	//Find the player's hand
	playerIndex := FindPlayerIndex(board.Players, playerID)
	if playerIndex == -1 {
		return errNoSuchPlayer
	}

	topCard := board.Deck.TopCard()
	//compare the client's prediction and with the actual result and correct
	if *cardToDraw != topCard {
		*cardToDraw = topCard
	}

	//add card to hand
	board.Players[playerIndex].drawCard(*cardToDraw)
	//remove card from deck
	board.Deck.removeCard(*cardToDraw)
	return nil
}
