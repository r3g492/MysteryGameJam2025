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
	"time"
)

var (
	screenWidth       float32 = 1600
	screenHeight      float32 = 900
	screenYaw         float32 = 0
	screenPitch       float32 = 0
	mouseSensitivity  float32 = 0.001
	bgm               rl.Sound
	gameState         int = START_MENU
	gameStartTime     time.Time
	eventIdx          = 0
	lastEventTime     time.Time
	mouseRay          rl.Ray
	hit               rl.Vector3
	shootDir          rl.Vector3
	modeClicked       bool
	whileCnt          int = 0
	EarthHit          int = 0
	EarthHitThreshold int = 10
)

const (
	START_MENU = iota
	IN_GAME
)

// projectile type
const (
	COMM = iota
	MISSILE
)

var currentProjectileType int = COMM

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
	bgm = embedWrapper.LoadSoundFromEmbedded("horror-tension-suspense-322304.mp3")
	explosionSound := embedWrapper.LoadSoundFromEmbedded("medium-explosion-40472.mp3")
	sonarSound := embedWrapper.LoadSoundFromEmbedded("sonar-107581.mp3")

	raylib.InitEarth()
	defer raylib.UnloadEarth()
	raylib.InitMoon()
	defer raylib.UnloadMoon()
	raylib.InitSun()
	defer raylib.UnloadSun()

	for !rl.WindowShouldClose() {
		modeClicked = false

		if gameState == START_MENU {
			if renderStart() {
				gameState = IN_GAME
				gameStartTime = time.Now()
			}
			continue
		}

		if eventIdx == 0 && time.Since(gameStartTime) > 1*time.Second {
			game.InputMessage("Hello?")
			eventIdx++
			lastEventTime = time.Now()
			rl.PlaySound(sonarSound)
		}
		if eventIdx == 1 && time.Since(lastEventTime) > 3*time.Second {
			game.InputMessage("Hello??")
			eventIdx++
			lastEventTime = time.Now()
			rl.PlaySound(sonarSound)
		}
		if eventIdx == 2 && time.Since(lastEventTime) > 3*time.Second {
			game.InputMessage("I am a friend.")
			eventIdx++
			lastEventTime = time.Now()
			rl.PlaySound(sonarSound)
		}

		if eventIdx == 3 && time.Since(lastEventTime) > 3*time.Second {
			game.InputMessage("They are coming.")
			eventIdx++
			lastEventTime = time.Now()
			rl.PlaySound(sonarSound)
		}

		if eventIdx == 4 && time.Since(lastEventTime) > 3*time.Second {
			game.InputMessage("Prepare your weapon.")
			eventIdx++
			lastEventTime = time.Now()
			rl.PlaySound(sonarSound)
		}

		if eventIdx == 5 && time.Since(lastEventTime) > 3*time.Second {
			game.InputMessage("Click the button below Earth.")
			eventIdx++
			lastEventTime = time.Now()
			rl.PlaySound(sonarSound)
		}

		if eventIdx == 6 && time.Since(lastEventTime) > 4*time.Second {
			eventIdx = 5
			lastEventTime = time.Now()
		}

		if eventIdx == 6 && currentProjectileType == MISSILE {
			game.InputMessage("Good.")
			eventIdx++
			lastEventTime = time.Now()
			rl.PlaySound(sonarSound)
		}

		if eventIdx == 7 && (time.Since(lastEventTime) > 4*time.Second) && currentProjectileType == MISSILE {
			game.InputMessage("Left click to shoot.")
			eventIdx++
			lastEventTime = time.Now()
			game.GenerateDrone(1, 0.4)
			rl.PlaySound(sonarSound)
			rl.PlaySound(bgm)
		}

		if eventIdx > 7 && !rl.IsSoundPlaying(bgm) {
			rl.PlaySound(bgm)
		}

		if eventIdx == 7 && time.Since(lastEventTime) > 30*time.Second && !game.EnemyAllDead() {
			eventIdx--
			lastEventTime = time.Now()
		}

		if eventIdx == 8 && time.Since(lastEventTime) > 4*time.Second && game.EnemyAllDead() {
			game.InputMessage("Well done.")
			eventIdx++
			lastEventTime = time.Now()
			rl.PlaySound(sonarSound)
		}

		if eventIdx == 9 && eventIdx <= 10 && time.Since(lastEventTime) > 2*time.Second {
			game.GenerateDrone(1, 0.4)
			eventIdx = 10
			lastEventTime = time.Now()
		}

		for eventIdx == 10 && time.Since(lastEventTime) > 3*time.Second {
			game.InputMessage("They are invaders ...")
			eventIdx = 11
			lastEventTime = time.Now()
			rl.PlaySound(sonarSound)
		}

		for eventIdx == 11 && time.Since(lastEventTime) > 3*time.Second {
			game.InputMessage("For me to help you ...")
			eventIdx = 12
			lastEventTime = time.Now()
			rl.PlaySound(sonarSound)
		}

		for eventIdx == 12 && time.Since(lastEventTime) > 3*time.Second {
			game.InputMessage("Send COMM to the moon.")
			eventIdx = 13
			lastEventTime = time.Now()
			rl.PlaySound(sonarSound)
		}

		if eventIdx == 13 && time.Since(lastEventTime) > 6*time.Second && currentProjectileType != COMM {
			eventIdx = 12
			lastEventTime = time.Now()
		}

		if eventIdx >= 13 {
			raylib.MoonCheck()
		}

		if eventIdx == 13 && game.MoonComm {
			game.InputMessage("Good.")
			eventIdx = 14
			lastEventTime = time.Now()
			whileCnt = 10
			rl.PlaySound(sonarSound)
		}

		if eventIdx == 14 && time.Since(lastEventTime) > 3*time.Second {
			game.GenerateDrone(2, 0.4)
			if whileCnt > 0 {
				whileCnt--
			} else {
				eventIdx = 15
			}
			lastEventTime = time.Now()
		}

		if eventIdx == 15 && time.Since(lastEventTime) > 3*time.Second {
			game.InputMessage("It's not enough ...")
			eventIdx = 16
			lastEventTime = time.Now()
			rl.PlaySound(sonarSound)
		}

		if eventIdx >= 15 {
			raylib.SunCheck()
		}

		if eventIdx == 16 && time.Since(lastEventTime) > 3*time.Second {
			game.InputMessage("Send COMM to The Sun!")
			eventIdx = 17
			lastEventTime = time.Now()
			rl.PlaySound(sonarSound)
		}

		if eventIdx == 17 && time.Since(lastEventTime) > 4*time.Second && game.SunComm {
			game.InputMessage("Now ...")
			eventIdx = 18
			lastEventTime = time.Now()
			rl.PlaySound(sonarSound)
		}

		if eventIdx == 18 && time.Since(lastEventTime) > 4*time.Second {
			game.InputMessage("I have ...")
			eventIdx = 19
			lastEventTime = time.Now()
			rl.PlaySound(sonarSound)
		}

		if eventIdx == 19 && time.Since(lastEventTime) > 4*time.Second {
			game.InputMessage("found you.")
			eventIdx = 20
			lastEventTime = time.Now()
			rl.PlaySound(sonarSound)
		}

		if eventIdx == 20 && time.Since(lastEventTime) > 4*time.Second {
			game.InputMessage("This is the last message.")
			eventIdx = 21
			lastEventTime = time.Now()
			game.StartCountdown(30)
		}

		if eventIdx >= 21 && !game.DroneTurn {
			if game.CommCheck() {
				game.TurnDrone()
				game.InputMessage("No!!!")
			}
		}

		if eventIdx == 21 && time.Since(lastEventTime) > 4*time.Second {
			game.GenerateDrone(24, 1)
			eventIdx = 21
			lastEventTime = time.Now()
		}

		if isGameEnded() || game.IsCountdownFinished() {
			if renderEnd() {
				return
			}
			continue
		}

		camera := updateCamera()

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)
		drawGradientSky(rl.Black, rl.LightGray)
		rl.BeginMode3D(camera)
		raylib.DrawEarth()
		raylib.DrawMoon()
		raylib.DrawSun()
		raylib.DrawProjectile()
		raylib.DrawDrones()
		raylib.DrawExplosions()
		dt := 1.0 / 60
		raylib.UpdateExplosions(float32(dt))

		if eventIdx >= 5 {
			getRayAndHit(camera)
			if currentProjectileType == COMM {
				showRay(rl.Green)
			} else if currentProjectileType == MISSILE {
				showRay(rl.Yellow)
			}
		}

		if eventIdx > 19 {
			raylib.DrawDemon()
		}

		rl.EndMode3D()

		renderModeChangeButton(camera)

		if game.CountDownBegin {
			secs := game.GetCountdown()
			rl.DrawText("Survive "+strconv.Itoa(secs)+"s", int32(screenWidth/2-50), int32(screenHeight/2-50), 10, rl.RayWhite)
		}

		renderMessage()
		exX, exY, exZ, explosionHappens := game.ProjectileCheck()
		if explosionHappens {
			raylib.AddExplosion(rl.Vector3{X: exX, Y: exY, Z: exZ})
			rl.PlaySound(explosionSound)
		}
		for explosionHappens {
			exX, exY, exZ, explosionHappens = game.ProjectileCheck()
			if explosionHappens {
				raylib.AddExplosion(rl.Vector3{X: exX, Y: exY, Z: exZ})
				rl.PlaySound(explosionSound)
			}
		}

		exX, exY, exZ, explosionHappens = game.EarthCheck()
		if explosionHappens {
			raylib.AddExplosion(rl.Vector3{X: exX, Y: exY, Z: exZ})
			rl.PlaySound(explosionSound)
			EarthHit++
			raylib.AddEarthHit(EarthHitThreshold)
			if EarthHit >= EarthHitThreshold {
				game.EarthHealth = 0
			}
		}
		for explosionHappens {
			exX, exY, exZ, explosionHappens = game.EarthCheck()
			if explosionHappens {
				raylib.AddExplosion(rl.Vector3{X: exX, Y: exY, Z: exZ})
				rl.PlaySound(explosionSound)
				EarthHit++
				raylib.AddEarthHit(EarthHitThreshold)
				if EarthHit >= EarthHitThreshold {
					game.EarthHealth = 0
				}
			}
		}

		game.MoveProjectile()
		if rl.IsMouseButtonPressed(rl.MouseButtonLeft) && !modeClicked && eventIdx >= 7 {
			game.AddProjectile(
				&game.Projectile{
					PosX:       0,
					PosY:       0,
					PosZ:       0,
					Type:       currentProjectileType,
					DirectionX: shootDir.X,
					DirectionY: shootDir.Y,
					DirectionZ: shootDir.Z,
					Speed:      1,
				})
			if currentProjectileType == MISSILE {
				rl.PlaySound(explosionSound)
			} else {
				rl.PlaySound(sonarSound)
			}
		}
		game.MoveDrone()
		rl.EndDrawing()
	}
}

func renderMessage() {
	msg, show, remaining := game.MessageResponse()
	if !show {
		return
	}
	alpha := float32(remaining) / float32(game.MessageDuration*1000)
	if alpha < 0 {
		alpha = 0
	} else if alpha > 1 {
		alpha = 1
	}

	fontSize := int32(100)
	textWidth := rl.MeasureText(msg, fontSize)
	midX := int32(screenWidth/2) - textWidth/2
	midY := int32(screenHeight/2) - fontSize/2 - 200

	faded := rl.Fade(rl.White, alpha)
	rl.DrawText(msg, midX, midY, fontSize, faded)
}

func renderModeChangeButton(camera rl.Camera3D) {
	earthScreen := rl.GetWorldToScreen(
		rl.Vector3{X: 0, Y: 0, Z: 0},
		camera,
	)
	if earthScreen.X > 0 && earthScreen.X < screenWidth &&
		earthScreen.Y > 0 && earthScreen.Y < screenHeight {

		bw, bh := int32(200), int32(16)
		bx := int32(earthScreen.X) - bw/2
		by := int32(earthScreen.Y) + 50

		var label string
		var btnColor rl.Color
		if currentProjectileType == COMM {
			label = "COMM MODE: Click to switch"
			btnColor = rl.DarkGreen
		} else {
			label = "MISSILE MODE: Click to switch"
			btnColor = rl.Maroon
		}

		rl.DrawRectangle(bx, by, bw, bh, btnColor)
		tw := rl.MeasureText(label, 8)
		rl.DrawText(label, bx+(bw-tw)/2, by+bh/2-9, 8, rl.White)

		if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
			mp := rl.GetMousePosition()
			if mp.X >= float32(bx) && mp.X <= float32(bx+bw) &&
				mp.Y >= float32(by) && mp.Y <= float32(by+bh) {
				modeClicked = true
				if currentProjectileType == COMM {
					currentProjectileType = MISSILE
				} else {
					currentProjectileType = COMM
				}
			}
		}
	}
}

func getRayAndHit(cam rl.Camera3D) {
	mouse := rl.GetMousePosition()
	mouseRay = rl.GetScreenToWorldRay(mouse, cam)

	if mouseRay.Direction.Y == 0 {
		return
	}

	t := -mouseRay.Position.Y / mouseRay.Direction.Y
	if t <= 0 {
		return
	}

	hit = rl.Vector3{
		X: mouseRay.Position.X + mouseRay.Direction.X*t,
		Y: 0,
		Z: mouseRay.Position.Z + mouseRay.Direction.Z*t,
	}

	shootDir = rl.Vector3{X: hit.X, Y: hit.Y, Z: hit.Z}
	length := float32(math.Sqrt(float64(shootDir.X*shootDir.X +
		shootDir.Y*shootDir.Y +
		shootDir.Z*shootDir.Z)))
	if length != 0 {
		shootDir.X /= length
		shootDir.Y /= length
		shootDir.Z /= length
	}
}

func showRay(
	rayColor rl.Color,
) {
	rl.DrawLine3D(
		rl.Vector3{X: 0, Y: 0, Z: 0},
		hit,
		rayColor,
	)
	rl.DrawSphere(hit, 0.3, rl.Yellow)
}

func renderStart() bool {
	rl.BeginDrawing()
	rl.ClearBackground(rl.Black)
	label := "Start"
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
