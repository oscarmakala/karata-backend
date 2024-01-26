package tests

import (
	"encoding/json"
	"fmt"
	"quana.co.tz/karata/core"
	"testing"
)

func TestCreateGame(t *testing.T) {
	service := &core.ServiceImpl{}
	var players = initTestPlayer()
	board := service.ActionStartGame(players, 5)

	fmt.Printf("Deck has cards: %v\n", len(board.Deck.Cards))
	fmt.Printf("Current Card is: %v\n", board.CurrentCard)

	printBoard(board)

	for _, player := range board.Players {
		if player.UserID == board.CurrentTurn {
			topCard := player.Hand[len(player.Hand)-1]
			err := service.ActionPlayCard(&board, &topCard, board.CurrentTurn)
			if err != nil {
				fmt.Printf("error occured %v\n", err.Error())
				return
			}
			printBoard(board)

			deckCard := board.Deck.TopCard()
			err = service.ActionDrawCard(&board, board.CurrentTurn, &deckCard)
			if err != nil {
				fmt.Printf("error occured %v\n", err.Error())
				return
			}
			printBoard(board)
			break
		}
	}
}

func TestNextTurn(t *testing.T) {
	service := &core.ServiceImpl{}
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

func printBoard(board core.Board) {
	boardJson, err := json.Marshal(board)
	if err != nil {
		fmt.Printf("error encoding label: %v\n", err)
	}
	fmt.Println(string(boardJson))
}
