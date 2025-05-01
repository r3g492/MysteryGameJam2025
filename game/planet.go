package game

type Planet struct {
	X           float32
	Y           float32
	Z           float32
	Radius      float32
	isDestroyed bool
	isIce       bool
	isOcean     bool
}

var (
	planets           []Planet
	alienHomeWorldIdx int = 0
)

func GenerateCandidates(
	numberOfPlanets int,
) []Planet {
	var answer []Planet
	for i := 0; i < numberOfPlanets; i++ {
		answer = append(answer,
			Planet{},
		)
	}

	return answer
}

func GetPlanets() []Planet {
	return planets
}

func HasWon() bool {
	if len(planets) > 0 && planets[alienHomeWorldIdx].isDestroyed {
		return true
	}
	return false
}
