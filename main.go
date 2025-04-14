package main

import (
	gameCamera "MysteryGameJam2025/camera"
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
	rl.InitWindow(int32(screenWidth), int32(screenHeight), "MysteryGameJam2025")
	rl.ToggleFullscreen()
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	// Create the player and retrieve its pointer/instance.
	game.CreatePlayer(game.Player{
		X:     0,
		Y:     0,
		Z:     0,
		Alive: true,
	})
	player := game.GetPlayer()

	for !rl.WindowShouldClose() {
		if player.Alive {
			mouseDelta := rl.GetMouseDelta()
			screenYaw += mouseDelta.X * 0.005
			screenPitch -= mouseDelta.Y * 0.005

			if screenPitch > 1.5 {
				screenPitch = 1.5
			} else if screenPitch < -1.5 {
				screenPitch = -1.5
			}

			if rl.IsKeyDown(rl.KeyW) {
				// Movement logic can go here.
			}

			gameCamera.PositionX = player.X
			gameCamera.PositionY = player.Y
			gameCamera.PositionZ = player.Z

			dirX := gameMath.Cos(screenPitch) * gameMath.Sin(screenYaw)
			dirY := gameMath.Sin(screenPitch)
			dirZ := gameMath.Cos(screenPitch) * gameMath.Cos(screenYaw)

			gameCamera.TargetX = player.X + dirX
			gameCamera.TargetY = player.Y + dirY
			gameCamera.TargetZ = player.Z + dirZ

			rl.DisableCursor()
		}

		camera := updateCamera()

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		rl.DrawText("Congrats! You created your first window!", 190, 200, 20, rl.LightGray)
		rl.DrawFPS(100, 100)

		rl.BeginMode3D(camera)
		rl.DrawGrid(100, 100)

		rl.DrawCube(rl.Vector3{X: 0, Y: 0.5, Z: 3}, 1, 1, 1, rl.Blue)
		rl.EndMode3D()

		rl.EndDrawing()
	}
}

func updateCamera() rl.Camera3D {
	return rl.Camera3D{
		Position: rl.Vector3{
			X: gameCamera.PositionX,
			Y: gameCamera.PositionY,
			Z: gameCamera.PositionZ,
		},
		Target: rl.Vector3{
			X: gameCamera.TargetX,
			Y: gameCamera.TargetY,
			Z: gameCamera.TargetZ,
		},
		Up: rl.Vector3{
			X: gameCamera.UpX,
			Y: gameCamera.UpY,
			Z: gameCamera.UpZ,
		},
		Fovy:       gameCamera.Fovy,
		Projection: rl.CameraPerspective,
	}
}
