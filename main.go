package main

import (
	"MysteryGameJam2025/game"
	gameMath "MysteryGameJam2025/math"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	screenWidth  float32 = 1600
	screenHeight float32 = 900
	screenYaw    float32 = 0
	screenPitch  float32 = 0
)

func main() {
	rl.InitWindow(
		int32(screenWidth),
		int32(screenHeight),
		"MysteryGameJam2025",
	)
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	camera := rl.Camera3D{
		Position: rl.Vector3{
			X: 10,
			Y: 10,
			Z: 0,
		},
		Target: rl.Vector3{
			X: 0,
			Y: 0,
			Z: 0,
		},
		Up: rl.Vector3{
			X: 1,
			Y: 10,
			Z: 0,
		},
		Fovy:       90,
		Projection: 0,
	}

	game.CreatePlayer(
		game.Player{
			X: 0, Y: 0, Z: 0,
			Alive: true,
		},
	)
	var player = game.GetPlayer()

	for !rl.WindowShouldClose() {
		camera = rl.Camera3D{
			Position: rl.Vector3{
				X: 10,
				Y: 10,
				Z: 0,
			},
			Target: rl.Vector3{
				X: player.X,
				Y: player.Y,
				Z: player.Z,
			},
			Up: rl.Vector3{
				X: 1,
				Y: 10,
				Z: 0,
			},
			Fovy:       90,
			Projection: 0,
		}

		if player.Alive {
			mouseDelta := rl.GetMouseDelta()
			screenYaw += mouseDelta.X * 0.005
			screenPitch -= mouseDelta.Y * 0.005

			if screenPitch > 1.5 {
				screenPitch = 1.5
			} else if screenPitch < -1.5 {
				screenPitch = -1.5
			}

			radius := float32(10.0)
			camX := radius * gameMath.Cos(screenPitch) * gameMath.Sin(screenYaw)
			camY := radius * gameMath.Sin(screenPitch)
			camZ := radius * gameMath.Cos(screenPitch) * gameMath.Cos(screenYaw)

			camera.Position = rl.Vector3{
				X: player.X + camX,
				Y: player.Y + camY,
				Z: player.Z + camZ,
			}
			camera.Target = rl.Vector3{
				X: player.X,
				Y: player.Y,
				Z: player.Z,
			}
			camera.Up = rl.NewVector3(0, 1, 0)

			var cursorRay = rl.GetScreenToWorldRay(
				rl.Vector2{
					X: screenWidth / 2,
					Y: screenHeight / 2,
				},
				camera,
			)

			camera.Target = rl.Vector3{
				X: player.X + cursorRay.Direction.X*100,
				Y: player.Y + cursorRay.Direction.Y*100,
				Z: player.Z + cursorRay.Direction.Z*100,
			}

			// rl.DrawLine3D(player.Position(), target, rl.Red)
		}

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		rl.DrawText("Congrats! You created your first window!", 190, 200, 20, rl.LightGray)
		rl.DrawFPS(100, 100)
		rl.BeginMode3D(camera)
		rl.DrawGrid(10, 10)
		rl.EndMode3D()
		rl.EndDrawing()
	}
}
