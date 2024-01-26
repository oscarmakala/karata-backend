package core

func FindPlayerIndex(players []PlayerState, playerID string) int {
	playerIndex := -1
	//Find the player
	for position := range players {
		if players[position].UserID == playerID {
			playerIndex = position
			break
		}
	}
	return playerIndex
}
