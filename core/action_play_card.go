package core

func (gs *ServiceImpl) ActionPlayCard(board *Board, cardToPlay *Card, playerID string) error {
	if cardToPlay == nil {
		return nil
	}

	playerIndex := FindPlayerIndex(board.Players, playerID)

	if playerIndex == -1 {
		return errNoSuchPlayer
	}

	//Find the index of the card in the player's hand
	cardIndex := FindCardIndex(board.Players[playerIndex].Hand, *cardToPlay)
	// Check if the card was found in the hand
	if cardIndex == -1 {
		return errNoSuchCardInHand
	}

	if canPlayCard(cardToPlay, board) == false {
		return errInvalidPlay
	}

	//Move the card from the hand to the dropzone
	board.DropZone = append(board.DropZone, board.Players[playerIndex].Hand[cardIndex])

	//Remove the card from the players hand
	board.Players[playerIndex].Hand = append(board.Players[playerIndex].Hand[:cardIndex], board.Players[playerIndex].Hand[cardIndex+1:]...)

	//add as the current card
	board.CurrentCard = *cardToPlay

	return nil
}

func canPlayCard(card *Card, board *Board) bool {
	currentCard := board.CurrentCard
	//if the same suit or same rank
	if card.Suit == currentCard.Suit || card.Rank == currentCard.Rank {
		//if its the first card played
		if len(board.DropZone) == 0 {
			return true
		}

		if card.Rank == currentCard.Rank {
			return true
		}
	}
	return false
}
