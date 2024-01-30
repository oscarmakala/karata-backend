package core

type PlayerState struct {
	UserID string
	Hand   []Card
	Name   string
}

func (m *PlayerState) drawCard(card Card) {
	m.Hand = append(m.Hand, card)
}

func (m *PlayerState) removeFromHand(cardIndex int) {
	m.Hand = append(m.Hand[:cardIndex], m.Hand[cardIndex+1:]...)
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

func (m *PlayerState) FindCardIndex(card Card) int {
	cardIndex := -1
	for i, c := range m.Hand {
		if c == card {
			cardIndex = i
			break
		}
	}
	return cardIndex
}
