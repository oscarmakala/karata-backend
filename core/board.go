package core

type PenaltyAction string

const (
	None    PenaltyAction = "None"
	Draw    PenaltyAction = "Draw"
	Skip    PenaltyAction = "Skip"
	Reverse PenaltyAction = "Reverse"
)

type Board struct {
	Deck                Deck
	CurrentCard         Card
	NoOfCardsToDeal     int
	DropZone            []Card
	Players             []PlayerState
	ActivePlayerId      string
	Penalties           []Penalty
	CanDefend           bool
	DrawnCards          []Card
	PlayerTurnDirection int
}

type Penalty struct {
	UserId        string
	NumberOfCards int
	Penalty       PenaltyAction
}
