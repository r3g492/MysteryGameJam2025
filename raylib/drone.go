package raylib

import (
	"MysteryGameJam2025/game"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func DrawDrones() {
	for i := 0; i < len(game.DroneList); i++ {
		drone := game.DroneList[i]
		color := rl.Red
		if game.DroneTurn {
			color = rl.Green
		}
		rl.DrawSphere(
			rl.Vector3{X: drone.X, Y: drone.Y, Z: drone.Z},
			drone.ShowRadius,
			color,
		)
	}
}
