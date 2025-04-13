package game

var (
	gamePlayer Player
)

func CreatePlayer(
	player Player,
) {
	gamePlayer = player
}

func GetPlayer() Player {
	return gamePlayer
}
