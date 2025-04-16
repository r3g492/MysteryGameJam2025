package game

var (
	gamePlayer *Player
)

func CreatePlayer(player Player) *Player {
	gamePlayer = &player
	return gamePlayer
}

func DeletePlayer() {
	gamePlayer = nil
}
