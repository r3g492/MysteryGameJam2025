package raylib

import rl "github.com/gen2brain/raylib-go/raylib"

var (
	demonRadius float32 = 1000
	demonX      float32 = 0
	demonY      float32 = -1500
	demonZ      float32 = 0
)

func DrawDemon() {
	rl.DrawSphere(
		rl.Vector3{X: demonX, Y: demonY, Z: demonZ},
		demonRadius,
		rl.Black,
	)
}
