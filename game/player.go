package game

var (
	gamePlayer *Player
)

type Player struct {
	X         float32
	Y         float32
	Z         float32
	MoveSpeed float32
	Alive     bool
}

func InitPlayer() *Player {
	gamePlayer = &Player{
		X:         0,
		Y:         10,
		Z:         0,
		MoveSpeed: 0.5,
		Alive:     true,
	}
	return gamePlayer
}

func GetPlayer() *Player {
	return gamePlayer
}
func UnloadPlayer() {
	gamePlayer = nil
}
