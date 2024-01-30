package core

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

var _ GameHandler = &GameHandlerImpl{}

type GameHandlerImpl struct{}

func (g GameHandlerImpl) ActionStartGame(players []PlayerState, nbCards int) Board {
	//initialize the deck of cards
	deck := &Deck{make([]Card, 0)}

	//add cards to the deck
	deck.createStack(1)

	//shuffle the deck
	deck.shuffle()

	for position := range players {
		deck.deal(nbCards, &players[position].Hand)
	}

	board := Board{
		Deck:                *deck,
		CurrentCard:         deck.gameCard(),
		NoOfCardsToDeal:     nbCards,
		DropZone:            make([]Card, 0),
		Players:             players,
		Penalties:           make([]Penalty, 0),
		DrawnCards:          make([]Card, 0),
		PlayerTurnDirection: 1,
	}

	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	board.ActivePlayerId = players[random.Intn(len(players))].UserID
	return board
}

func (g GameHandlerImpl) ActionPlayCard(board *Board, cardToPlay *Card, playerID string) error {
	if cardToPlay == nil {
		return nil
	}

	playerIndex := FindPlayerIndex(board.Players, playerID)
	if playerIndex == -1 {
		return errNoSuchPlayer
	}

	//Find the index of the card in the player's hand
	cardIndex := board.Players[playerIndex].FindCardIndex(*cardToPlay)
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
	board.Players[playerIndex].removeFromHand(cardIndex)
	//add as the current card
	board.CurrentCard = *cardToPlay

	return nil
}

func (g GameHandlerImpl) ActionNextTurn(board *Board, playerID string, isPlaying bool) error {
	if isPlaying == false {
		return nil
	}

	currentPlayerIndex := FindPlayerIndex(board.Players, playerID)
	fmt.Printf("currentPlayerIndex is %v\n", currentPlayerIndex)
	if currentPlayerIndex == -1 {
		return errNoSuchPlayer
	}

	CheckIfCurrentPlayerIsForcedToTakeCards(*board, playerID)

	nextPlayerIndex := IncreasePlayerIndex(*board, currentPlayerIndex)

	CheckIfNextPlayerIsForcedToTakeCards(board, nextPlayerIndex)

	//empty dropzone
	board.DropZone = board.DropZone[:0]
	return nil
}

func CheckIfCurrentPlayerIsForcedToTakeCards(board Board, playerID string) {
	if len(board.DropZone) == 0 {
		board.Penalties = append(board.Penalties, Penalty{
			UserId:        playerID,
			NumberOfCards: 1,
			Penalty:       Draw,
		})
	}
}

func (g GameHandlerImpl) ActionDrawCard(board *Board, playerID string, cardToDraw []Card) error {
	//Find the player's hand
	playerIndex := FindPlayerIndex(board.Players, playerID)
	if playerIndex == -1 {
		return errNoSuchPlayer
	}

	if evaluateDrawnCards(board, cardToDraw) == false {
		return nil
	}

	for _, card := range cardToDraw {
		//if okay draw the first top cards
		board.Deck.removeCard(card)
		board.Players[playerIndex].drawCard(card)
		board.DrawnCards = append(board.DrawnCards, card)
	}

	if len(board.Penalties) > 0 {
		//clear the penalties from previous state
		board.Penalties = board.Penalties[:0]
	}
	return nil
}

func evaluateDrawnCards(board *Board, draw []Card) bool {
	if len(draw) == 0 {
		return false
	}

	//get the top cards
	topCards := board.Deck.Cards[len(board.Deck.Cards)-4:]
	foundCount := 0
	for _, topCard := range topCards {
		for _, drawCard := range draw {
			if topCard.equals(&drawCard) {
				foundCount++
				break
			}
		}
	}
	return foundCount == len(draw)
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
func CheckIfNextPlayerIsForcedToTakeCards(board *Board, currentPlayerIndex int) {
	player := board.Players[currentPlayerIndex]
	if player.hasCardsToDefend(board.CurrentCard) {
		board.CanDefend = true
	} else {
		TakePenaltyCards(player, board)
	}
}

func TakePenaltyCards(player PlayerState, board *Board) {
	totalCards := len(board.DropZone)
	switch board.CurrentCard.Rank {
	case Plus2:
		//take two
		board.Penalties = append(board.Penalties, Penalty{
			UserId:        player.UserID,
			NumberOfCards: 2 * totalCards,
			Penalty:       Draw,
		})
		break
	case Seven:
		board.Penalties = append(board.Penalties, Penalty{
			UserId:        player.UserID,
			NumberOfCards: 0,
			Penalty:       Skip,
		})
		break
	}
}

func IncreasePlayerIndex(board Board, playerIndex int) int {
	playerIndex = (playerIndex + board.PlayerTurnDirection) % len(board.Players)
	fmt.Printf("%v\n", playerIndex)
	if playerIndex < 0 {
		playerIndex += len(board.Players)
		fmt.Printf("%v\n", playerIndex)
	}
	fmt.Printf("result %v\n", playerIndex)
	return playerIndex
}
