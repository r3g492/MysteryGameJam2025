package raylib

import (
	"MysteryGameJam2025/game"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func DrawProjectile() {
	for i := 0; i < len(game.ProjectileList); i++ {
		projectile := game.ProjectileList[i]
		if projectile.Type == game.COMM {
			rl.DrawSphere(
				rl.Vector3{X: projectile.PosX, Y: projectile.PosY, Z: projectile.PosZ},
				0.5,
				rl.Green,
			)
		} else if projectile.Type == game.MISSILE {
			rl.DrawSphere(
				rl.Vector3{X: projectile.PosX, Y: projectile.PosY, Z: projectile.PosZ},
				0.5,
				rl.Red,
			)
		}
	}

}
