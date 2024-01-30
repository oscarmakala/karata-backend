package tests

import (
	"encoding/json"
	"fmt"
	"quana.co.tz/karata/core"
	"testing"
)

func TestCreateGame(t *testing.T) {
	service := &core.GameHandlerImpl{}
	var players = initTestPlayer()
	board := service.ActionStartGame(players, 5)

	fmt.Printf("Deck has cards: %v\n", len(board.Deck.Cards))
	fmt.Printf("Current Card is: %v\n", board.CurrentCard)

	printBoard(board, true)

	for _, player := range board.Players {
		if player.UserID == board.ActivePlayerId {
			topCard := player.Hand[len(player.Hand)-1]
			err := service.ActionPlayCard(&board, &topCard, board.ActivePlayerId)
			if err != nil {
				fmt.Printf("error occured %v\n", err.Error())
				return
			}

			printBoard(board, false)

			//deckCard := board.Deck.TopCard()
			//err = service.ActionDrawCard(&board, board.ActivePlayerId, &deckCard)
			//if err != nil {
			//	fmt.Printf("error occured %v\n", err.Error())
			//	return
			//}
			//printBoard(board)
			break
		}
	}
}

func TestParsingBoardJson(t *testing.T) {
	board := &core.Board{}
	err := json.Unmarshal([]byte(TestDeck), board)
	if err != nil {
		fmt.Printf("error occured %v\n", err.Error())
		return
	}
	fmt.Printf("%v", &board.CurrentCard)
}

func TestNextTurnWhenNoCardsPlayed(t *testing.T) {
	gameHandler := &core.GameHandlerImpl{}

	board := &core.Board{}
	_ = json.Unmarshal([]byte(TestDeck), board)

	printBoard(*board, false)
	err := gameHandler.ActionNextTurn(board, "abc", true)
	if err != nil {
		fmt.Printf("error occured %v\n", err.Error())
		return
	}
	printBoard(*board, false)
}

func TestNextTurnWhenCardPlayedIsPlus2NoDefense(t *testing.T) {
	gameHandler := &core.GameHandlerImpl{}

	board := &core.Board{}
	_ = json.Unmarshal([]byte(KTestNextTurnWhenCardPlayedIsPlus2NoDefense), board)

	printBoard(*board, false)
	err := gameHandler.ActionNextTurn(board, board.ActivePlayerId, true)
	if err != nil {
		fmt.Printf("error occured %v\n", err.Error())
		return
	}
	printBoard(*board, false)
}

func TestDrawCardsWhenTakingPenalty(t *testing.T) {
	gameHandler := &core.GameHandlerImpl{}

	board := &core.Board{}
	_ = json.Unmarshal([]byte(KTestDrawAfterPenalty), board)
	printBoard(*board, false)

	drawCard := []core.Card{
		{
			Rank: 5,
			Suit: core.Hearts,
		},
		{
			Rank: 13,
			Suit: core.Spades,
		},
		{
			Rank: 1,
			Suit: core.Spades,
		},
		{
			Rank: 1,
			Suit: core.Diamonds,
		},
	}

	err := gameHandler.ActionDrawCard(board, board.Penalties[0].UserId, drawCard)
	if err != nil {
		fmt.Printf("error occured %v\n", err.Error())
		return
	}
	printBoard(*board, true)
}

func TestNextTurnWhenCardPlayedIsPlus4NoDefense(t *testing.T) {
	gameHandler := &core.GameHandlerImpl{}

	board := &core.Board{}
	_ = json.Unmarshal([]byte(KTestNextTurnWhenCardPlayedIsPlus4NoDefense), board)

	printBoard(*board, false)
	err := gameHandler.ActionNextTurn(board, board.ActivePlayerId, true)
	if err != nil {
		fmt.Printf("error occured %v\n", err.Error())
		return
	}
	printBoard(*board, true)
}

func TestNextTurn(t *testing.T) {
	service := &core.GameHandlerImpl{}
	var players = initTestPlayer()
	board := service.ActionStartGame(players, 5)

	fmt.Printf("Deck has cards: %v\n", len(board.Deck.Cards))
	fmt.Printf("Current Card is: %v\n", board.CurrentCard)
	currentPlayer := "abc"

	err := service.ActionNextTurn(&board, currentPlayer, true)
	if err != nil {
		fmt.Printf("error occured %v\n", err.Error())
		return
	}

}

func initTestPlayer() []core.PlayerState {
	return []core.PlayerState{
		{UserID: "abc", Name: "tom", Hand: make([]core.Card, 0)},
		{UserID: "efg", Name: "Mark", Hand: make([]core.Card, 0)},
	}
}

func printBoard(board core.Board, showJson bool) {
	buff, err := json.Marshal(board)
	if err != nil {
		fmt.Printf("error encoding label: %v\n", err)
	}
	fmt.Println()
	fmt.Println("***Start**")
	if showJson == true {
		fmt.Printf("JSON:  %s\n", string(buff))
	}
	fmt.Printf("Penalties         : %v\n", board.Penalties)
	fmt.Printf("Current Player: %v\n", board.ActivePlayerId)
	fmt.Printf("Can Defend    : %v\n", board.CanDefend)
	fmt.Printf("Cards in Deck : %v\n", len(board.Deck.Cards))
	fmt.Println("***End***")
}
