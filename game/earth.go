package game

var (
	EarthPositionX float32 = 0
	EarthPositionY float32 = 0
	EarthPositionZ float32 = 0
	EarthRadius    float32 = 5
	EarthHealth    int32   = 100
)

func HasLost() bool {
	if EarthHealth <= 0 {
		return true
	}
	return false
}
