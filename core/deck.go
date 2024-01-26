package core

import (
	"math/rand"
	"time"
)

type Deck struct {
	Cards []Card
}

func (deck *Deck) createStack(lowest int32) {
	deck.addSuit(Hearts, lowest)
	deck.addSuit(Diamonds, lowest)
	deck.addSuit(Clubs, lowest)
	deck.addSuit(Spades, lowest)
}

func (deck *Deck) addSuit(suit Suit, lowest int32) {
	for i := lowest; i <= 13; i++ {
		deck.Cards = append(deck.Cards, Card{Suit: suit, Rank: Rank(i)})
	}
}

/**
 * shuffle Shuffles the stack.
 */
func (deck *Deck) shuffle() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	shuffleCards := make([]Card, 0)

	for len(deck.Cards) > 0 {
		index := rand.Intn(len(deck.Cards))
		shuffleCards = append(shuffleCards, deck.Cards[index])
		deck.Cards = append(deck.Cards[:index], deck.Cards[index+1:]...)
	}
	deck.Cards = shuffleCards
}

// TopCard  returns the top card of the deck.
func (deck *Deck) TopCard() Card {
	return deck.Cards[len(deck.Cards)-1]
}

func (deck *Deck) deal(nbCards int, cards *[]Card) {
	for i := 0; i < nbCards; i++ {
		card := deck.TopCard()
		*cards = append(*cards, card)
		deck.removeCard(card)
	}
}

func (deck *Deck) gameCard() Card {
	gameCard := make([]Card, 0)
	deck.deal(1, &gameCard)
	return gameCard[0]
}

/**
 * removeCard removes the first matching card from the stack.
 * Parameters:
 * 	card (Card): The card to remove
 */
func (deck *Deck) removeCard(card Card) {
	for i := 0; i < len(deck.Cards); i++ {
		c := deck.Cards[i]
		if c.equals(&card) {
			deck.Cards = append(deck.Cards[:i], deck.Cards[i+1:]...)
		}
	}
}
