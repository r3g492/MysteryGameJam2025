package main

import (
	gameCamera "MysteryGameJam2025/camera"
	embedWrapper "MysteryGameJam2025/embed"
	"MysteryGameJam2025/game"
	gameMath "MysteryGameJam2025/math"
	"MysteryGameJam2025/raylib"
	rl "github.com/gen2brain/raylib-go/raylib"
	"math"
	"strconv"
)

var (
	screenWidth             float32 = 1600
	screenHeight            float32 = 900
	screenYaw               float32 = 0
	screenPitch             float32 = 0
	mouseSensitivity        float32 = 0.001
	pitchMaxThreshold       float32 = 0
	pitchMinThreshold       float32 = 1.5
	zoomSensitivity         float32 = 2.0
	fovMin                  float32 = 70
	fovMax                  float32 = 100
	cameraMovementThreshold float32 = 300

	// UI toggle state
	showControls bool = false

	// UI dimensions
	controlsTabWidth  int32 = 30
	controlsTabHeight int32 = 100
	panelPadding      int32 = 10
	alienRevealed     bool  = false

	calmBgm      rl.Sound
	battleBgm    rl.Sound
	curBgm       *rl.Sound
	survivalTime int = 60
)

const HalfPi float32 = 1.57
const PlayerHighestPoint float32 = 100
const PlayerLowestPoint float32 = 80

// START game state
const (
	START = iota
)

// projectile type
const (
	SCAN = iota
	MISSILE
)

var currentProjectileType int = SCAN

func main() {
	rl.InitWindow(int32(screenWidth), int32(screenHeight), "MysteryGameJam2025")
	monitor := rl.GetCurrentMonitor()
	fullScreenWidth := rl.GetMonitorWidth(monitor)
	fullScreenHeight := rl.GetMonitorHeight(monitor)
	screenWidth = float32(fullScreenWidth)
	screenHeight = float32(fullScreenHeight)
	rl.SetWindowSize(fullScreenWidth, fullScreenHeight)
	rl.ToggleFullscreen()
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	game.InitPlayer()
	defer game.UnloadPlayer()
	player := game.GetPlayer()

	setInitialTarget(player)
	rl.InitAudioDevice()
	calmBgm = embedWrapper.LoadSoundFromEmbedded("calm-space-music-312291.mp3")
	battleBgm = embedWrapper.LoadSoundFromEmbedded("horror-tension-suspense-322304.mp3")
	initBgm()

	raylib.InitEarth()
	defer raylib.UnloadEarth()
	raylib.InitMoon()
	defer raylib.UnloadMoon()
	raylib.InitSun()
	defer raylib.UnloadSun()

	for !rl.WindowShouldClose() {
		game.StartCountdown(60)
		if rl.IsKeyPressed(rl.KeyF5) {
			alienRevealed = !alienRevealed
		}

		if isGameEnded() || game.IsCountdownFinished() {
			if renderEnd() {
				return
			}
			continue
		} else {
			updateBgm(alienRevealed)
		}

		camera := updateCamera()

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)
		// draw 2d stuffs
		drawGradientSky(rl.Black, rl.LightGray)

		rl.BeginMode3D(camera)
		rl.DrawSphere(rl.Vector3{X: 5, Y: 5, Z: 0}, 0.5, rl.Red)
		raylib.DrawEarth()
		raylib.DrawMoon()
		raylib.DrawSun()

		if rl.IsMouseButtonDown(rl.MouseButtonLeft) {
			mouse := rl.GetMousePosition()
			ray := rl.GetScreenToWorldRay(mouse, camera)
			if ray.Direction.Y != 0 {
				t := -ray.Position.Y / ray.Direction.Y
				if t > 0 {
					hit := rl.Vector3{
						X: ray.Position.X + ray.Direction.X*t,
						Y: 0,
						Z: ray.Position.Z + ray.Direction.Z*t,
					}
					rl.DrawLine3D(
						rl.Vector3{X: 0, Y: 0, Z: 0},
						hit,
						rl.Green,
					)
					rl.DrawSphere(hit, 0.3, rl.Yellow)
				}
			}
		}
		if rl.IsMouseButtonReleased(rl.MouseButtonLeft) {

		}

		rl.EndMode3D()

		secs := game.GetCountdown()
		rl.DrawText("Survive "+strconv.Itoa(secs)+"s", int32(screenWidth/2-50), int32(screenHeight/2-50), 10, rl.RayWhite)

		// rl.DrawFPS(100, 100)

		rl.EndDrawing()
	}
}

func renderEnd() bool {
	rl.BeginDrawing()
	rl.ClearBackground(rl.RayWhite)
	label := "Victory"
	if game.HasLost() {
		label = "Defeat"
	}
	sw := int32(screenWidth)
	sh := int32(screenHeight)
	bw, bh := int32(300), int32(100)
	bx, by := (sw-bw)/2, (sh-bh)/2

	rl.DrawRectangle(bx, by, bw, bh, rl.DarkGray)
	tw := rl.MeasureText(label, 48)
	rl.DrawText(label, bx+(bw-tw)/2, by+(bh/2-24), 48, rl.White)

	if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		mp := rl.GetMousePosition()
		if mp.X >= float32(bx) && mp.X <= float32(bx+bw) &&
			mp.Y >= float32(by) && mp.Y <= float32(by+bh) {
			rl.EndDrawing()
			return true
		}
	}
	rl.EndDrawing()
	return false
}

func isGameEnded() bool {
	if game.HasLost() {
		return true
	}
	return false
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

func initBgm() {
	curBgm = &calmBgm
	rl.PlaySound(*curBgm) // start with the calm theme
}

func updateBgm(alienRevealed bool) {
	// decide which track we *should* be on
	want := &calmBgm
	if alienRevealed {
		want = &battleBgm
	}

	// did we switch state?
	if want != curBgm {
		rl.StopSound(*curBgm) // stop the old one
		curBgm = want
		rl.PlaySound(*curBgm) // and start the new one
	} else if !rl.IsSoundPlaying(*curBgm) {
		rl.PlaySound(*curBgm) // restart if it finished (manual “loop”)
	}
}
