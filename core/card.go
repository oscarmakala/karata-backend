package core

type Suit string

//Enum values for the Suit type

const (
	Hearts   Suit = "Hearts"
	Diamonds Suit = "Diamonds"
	Clubs    Suit = "Clubs"
	Spades   Suit = "Spades"
	Joker    Suit = "Joker"
)

const (
	Eight = 8
	Plus2 = 2
	Seven = 7
)

// Rank represents the rank of a playing card
type Rank int32

// Card represents a playing card
type Card struct {
	Suit Suit
	Rank Rank
}

func (c *Card) equals(other *Card) bool {
	return c.Suit == other.Suit && c.Rank == other.Rank
}

// FindCardIndex finds the index of a card with a specified suit and value in the slice of cards.
func FindCardIndex(cards []Card, card Card) int {
	cardIndex := -1
	for i, c := range cards {
		if c == card {
			cardIndex = i
			break
		}
	}
	return cardIndex
}
