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
	// rl.ToggleFullscreen()
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	player := game.CreatePlayer(
		game.Player{
			X:         0,
			Y:         0,
			Z:         0,
			MoveSpeed: 0.01,
			Alive:     true,
		})
	defer game.DeletePlayer()

	for !rl.WindowShouldClose() {
		if player.Alive {
			mouseDelta := rl.GetMouseDelta()
			screenYaw -= mouseDelta.X * 0.005
			screenPitch -= mouseDelta.Y * 0.005

			if screenPitch > 1.5 {
				screenPitch = 1.5
			} else if screenPitch < -1.5 {
				screenPitch = -1.5
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

		if rl.IsKeyDown(rl.KeySpace) {
			player.Y += 1
		}
		if rl.IsKeyDown(rl.KeyLeftControl) {
			player.Y -= 1
		}

		if rl.IsKeyDown(rl.KeyW) {
			player.X += gameMath.Sin(screenYaw) * player.MoveSpeed
			player.Z += gameMath.Cos(screenYaw) * player.MoveSpeed
		}

		if rl.IsKeyDown(rl.KeyS) {
			player.X -= gameMath.Sin(screenYaw) * player.MoveSpeed
			player.Z -= gameMath.Cos(screenYaw) * player.MoveSpeed
		}

		if rl.IsKeyDown(rl.KeyA) {
			player.X += gameMath.Sin(screenYaw+1.57) * player.MoveSpeed
			player.Z += gameMath.Cos(screenYaw+1.57) * player.MoveSpeed
		}

		if rl.IsKeyDown(rl.KeyD) {
			player.X += gameMath.Sin(screenYaw-1.57) * player.MoveSpeed
			player.Z += gameMath.Cos(screenYaw-1.57) * player.MoveSpeed
		}

		rl.BeginDrawing()
		// draw 2d stuffs
		rl.ClearBackground(rl.RayWhite)
		drawGradientSky(rl.SkyBlue, rl.White)

		// draw uis
		rl.DrawFPS(100, 100)

		rl.BeginMode3D(camera)
		// draw 3d stuffs
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

func drawGradientSky(topColor, bottomColor rl.Color) {
	width := int32(screenWidth)
	height := int32(screenHeight)

	for i := int32(0); i < height; i++ {
		t := float32(i) / float32(height)
		color := rl.Color{
			R: uint8(float32(topColor.R)*(1-t) + float32(bottomColor.R)*t),
			G: uint8(float32(topColor.G)*(1-t) + float32(bottomColor.G)*t),
			B: uint8(float32(topColor.B)*(1-t) + float32(bottomColor.B)*t),
			A: 255,
		}
		rl.DrawLine(0, i, width, i, color)
	}
}
