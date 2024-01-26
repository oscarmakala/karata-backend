package core

type PenaltyAction string

const (
	None    PenaltyAction = "None"
	Draw    PenaltyAction = "Draw"
	Skip    PenaltyAction = "Skip"
	Reverse PenaltyAction = "Reverse"
)

type Board struct {
	Deck            Deck
	CurrentCard     Card
	NoOfCardsToDeal int
	DropZone        []Card
	Players         []PlayerState
	CurrentTurn     string
	PlayerOrder     []string
	Penalties       []Penalty
	CanDefend       bool
}

type Penalty struct {
	UserId        string
	NumberOfCards int
	Penalty       PenaltyAction
}

type PlayerState struct {
	UserID string
	Hand   []Card
	Name   string
}

func (m *PlayerState) drawCard(card Card) {
	m.Hand = append(m.Hand, card)
}

func (m *PlayerState) hasCardsToDefend(topCard Card) bool {
	cards := make([]Card, 0)
	for _, card := range m.Hand {
		if m.checkIfCardCardCanRescueFromPenalty(card, topCard) {
			cards = append(cards, card)
		}
	}
	return len(cards) > 0
}

func (m *PlayerState) checkIfCardCardCanRescueFromPenalty(c Card, topCard Card) bool {
	switch topCard.Rank {
	case Plus2, Eight, Seven:
		if c.Rank == topCard.Rank {
			return true
		}
	}
	return false
}
