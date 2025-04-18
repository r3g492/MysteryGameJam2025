package main

import (
	gameCamera "MysteryGameJam2025/camera"
	embedWrapper "MysteryGameJam2025/embed"
	"MysteryGameJam2025/game"
	gameMath "MysteryGameJam2025/math"
	rl "github.com/gen2brain/raylib-go/raylib"
	"math"
)

var (
	screenWidth      float32 = 1600
	screenHeight     float32 = 900
	screenYaw        float32 = 0
	screenPitch      float32 = 0
	mouseSensitivity float32 = 0.001
	pitchThreshold   float32 = 1.5
	zoomSensitivity  float32 = 2.0
	fovMin           float32 = 70
	fovMax           float32 = 100

	// UI toggle state
	showControls bool = false

	// UI dimensions
	controlsTabWidth  int32 = 30
	controlsTabHeight int32 = 100
	panelPadding      int32 = 10
)

const HalfPi float32 = 1.57
const PlayerHighestPoint float32 = 10
const PlayerLowestPoint float32 = 0.05

func main() {
	rl.InitWindow(int32(screenWidth), int32(screenHeight), "MysteryGameJam2025")
	// rl.ToggleFullscreen()
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	player := game.CreatePlayer(
		game.Player{
			X:         0,
			Y:         10,
			Z:         0,
			MoveSpeed: 0.5,
			Alive:     true,
		})
	defer game.DeletePlayer()
	setInitialTarget(player)
	rl.InitAudioDevice()
	bgm := embedWrapper.LoadSoundFromEmbedded("calm-space-music-312291.mp3")

	for !rl.WindowShouldClose() {
		if !rl.IsSoundPlaying(bgm) {
			rl.PlaySound(bgm)
		}

		handleControlsToggle()

		camera := updateCamera()
		cameraZoomControl()
		cameraNormalControl(player)
		if player.Alive && rl.IsMouseButtonDown(rl.MouseButtonRight) {
			cameraClickedControl(player)
		} else {
			rl.ShowCursor()
		}
		cameraEdgeCheck(player)

		rl.BeginDrawing()
		// draw 2d stuffs
		rl.ClearBackground(rl.RayWhite)
		drawGradientSky(rl.Black, rl.LightGray)

		// draw uis
		rl.DrawFPS(100, 100)

		rl.BeginMode3D(camera)
		// draw 3d stuffs
		rl.DrawGrid(1000, 1)

		rl.DrawCube(rl.Vector3{X: 0, Y: 0.5, Z: 3}, 1, 1, 1, rl.Blue)
		rl.EndMode3D()

		drawControlsUI()

		rl.EndDrawing()
	}
}

func cameraEdgeCheck(player *game.Player) {
	if player.Y > PlayerHighestPoint {
		player.Y = PlayerHighestPoint
	}

	if player.Y < PlayerLowestPoint {
		player.Y = PlayerLowestPoint
	}
}

func cameraZoomControl() {
	wheelMove := rl.GetMouseWheelMove()
	gameCamera.Fovy -= wheelMove * zoomSensitivity
	if gameCamera.Fovy < fovMin {
		gameCamera.Fovy = fovMin
	} else if gameCamera.Fovy > fovMax {
		gameCamera.Fovy = fovMax
	}
}

func cameraNormalControl(player *game.Player) {
	gameCamera.PositionX = player.X
	gameCamera.PositionY = player.Y
	gameCamera.PositionZ = player.Z

	dirX := gameMath.Cos(screenPitch) * gameMath.Sin(screenYaw)
	dirY := gameMath.Sin(screenPitch)
	dirZ := gameMath.Cos(screenPitch) * gameMath.Cos(screenYaw)

	gameCamera.TargetX = player.X + dirX
	gameCamera.TargetY = player.Y + dirY
	gameCamera.TargetZ = player.Z + dirZ

	if rl.IsKeyDown(rl.KeyW) {
		player.X += gameMath.Sin(screenYaw) * player.MoveSpeed
		player.Z += gameMath.Cos(screenYaw) * player.MoveSpeed
	}

	if rl.IsKeyDown(rl.KeyS) {
		player.X -= gameMath.Sin(screenYaw) * player.MoveSpeed
		player.Z -= gameMath.Cos(screenYaw) * player.MoveSpeed
	}

	if rl.IsKeyDown(rl.KeyA) {
		player.X += gameMath.Sin(screenYaw+HalfPi) * player.MoveSpeed
		player.Z += gameMath.Cos(screenYaw+HalfPi) * player.MoveSpeed
	}

	if rl.IsKeyDown(rl.KeyD) {
		player.X += gameMath.Sin(screenYaw-HalfPi) * player.MoveSpeed
		player.Z += gameMath.Cos(screenYaw-HalfPi) * player.MoveSpeed
	}
}

func cameraClickedControl(player *game.Player) {
	mouseDelta := rl.GetMouseDelta()
	screenYaw -= mouseDelta.X * mouseSensitivity
	screenPitch -= mouseDelta.Y * mouseSensitivity

	if screenPitch > pitchThreshold {
		screenPitch = pitchThreshold
	} else if screenPitch < -pitchThreshold {
		screenPitch = -pitchThreshold
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

func setInitialTarget(player *game.Player) {
	initialDir := rl.Vector3{
		X: 0 - player.X,
		Y: 0 - player.Y,
		Z: 0 - player.Z,
	}
	screenYaw = float32(math.Atan2(float64(initialDir.X), float64(initialDir.Z)))
	horizontalDist := float32(math.Sqrt(float64(initialDir.X*initialDir.X + initialDir.Z*initialDir.Z)))
	screenPitch = float32(math.Atan2(float64(initialDir.Y), float64(horizontalDist)))

	mouseDelta := rl.GetMouseDelta()
	screenYaw -= mouseDelta.X * mouseSensitivity
	screenPitch -= mouseDelta.Y * mouseSensitivity

	if screenPitch > pitchThreshold {
		screenPitch = pitchThreshold
	} else if screenPitch < -pitchThreshold {
		screenPitch = -pitchThreshold
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

func handleControlsToggle() {
	if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		m := rl.GetMousePosition()
		if m.X <= float32(controlsTabWidth) && m.Y <= float32(controlsTabHeight) && !showControls {
			showControls = true
		} else if m.X <= float32(controlsTabWidth+400) && m.Y <= float32(controlsTabHeight+50) && showControls {
			showControls = false
		}
	}
}

func drawControlsUI() {
	if showControls {
		rl.DrawRectangle(0, 0, controlsTabWidth+400, controlsTabHeight+50, rl.Gray)
		rl.DrawText("<<", 5, 5, 20, rl.White)
		y := panelPadding
		y += 30
		rl.DrawText("W/A/S/D = Move forward/left/back/right", panelPadding, y, 18, rl.White)
		y += 25
		rl.DrawText("Scroll wheel = Zoom", panelPadding, y, 18, rl.White)
		y += 25
		rl.DrawText("Right click + drag = Free look", panelPadding, y, 18, rl.White)
	} else {
		rl.DrawRectangle(0, 0, controlsTabWidth, controlsTabHeight, rl.Gray)
		rl.DrawText("controls", 5, 5, 20, rl.White)
	}
}
