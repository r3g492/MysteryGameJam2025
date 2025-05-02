package game

var (
	gamePlayer *Player
)

type Player struct {
	X         float32
	Y         float32
	Z         float32
	MoveSpeed float32
}

func InitPlayer() *Player {
	gamePlayer = &Player{
		X:         0,
		Y:         70,
		Z:         50,
		MoveSpeed: 0.5,
	}
	return gamePlayer
}

func GetPlayer() *Player {
	return gamePlayer
}
func UnloadPlayer() {
	gamePlayer = nil
}
