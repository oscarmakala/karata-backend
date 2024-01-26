package core

import "sort"

func (gs *ServiceImpl) ActionStartGame(players []PlayerState, nbCards int) Board {
	//initialize the deck of cards
	deck := &Deck{make([]Card, 0)}

	//add cards to the deck
	deck.createStack(1)

	//shuffle the deck
	deck.shuffle()

	for position := range players {
		deck.deal(nbCards, &players[position].Hand)
	}
	//initialize player order
	playerOrder := initializePlayOrderClockwise(players)
	currentTurn := playerOrder[0]

	return Board{
		Deck:            *deck,
		CurrentCard:     deck.gameCard(),
		NoOfCardsToDeal: nbCards,
		DropZone:        make([]Card, 0),
		Players:         players,
		CurrentTurn:     currentTurn,
		PlayerOrder:     playerOrder,
		Penalties:       make([]Penalty, 0),
	}
}

func initializePlayOrderClockwise(players []PlayerState) []string {
	var playerIDs []string
	for _, player := range players {
		playerIDs = append(playerIDs, player.UserID)
	}
	// Step 2: Sort player IDs clockwise
	sort.Slice(playerIDs, func(i, j int) bool {
		// Implement clockwise sorting logic based on your requirements
		// Here, we assume that player IDs are integers, and we sort them in ascending order
		// You may need to adapt this logic based on your actual player ID structure
		return playerIDs[i] < playerIDs[j]
	})
	return playerIDs
}
