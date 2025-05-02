package raylib

import (
	embedWrapper "MysteryGameJam2025/embed"
	"MysteryGameJam2025/game"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	earthMesh    rl.Mesh
	earthModel   rl.Model
	earthTexture rl.Texture2D

	earthRot      float32 = 0                // current angle in degrees
	earthRotSpeed float32 = 10               // °/sec – tweak to taste
	axisY                 = rl.Vector3{Y: 1} // (0,1,0) – spin around Y
)

func InitEarth() {
	earthTextureResponse, _ := embedWrapper.LoadTextureFromEmbedded(
		"earth.png",
		500,
		500,
	)
	earthTexture = earthTextureResponse

	earthMesh = rl.GenMeshSphere(
		game.EarthRadius,
		64,
		64,
	)
	earthModel = rl.LoadModelFromMesh(earthMesh)

	for i := range earthModel.GetMaterials() {
		rl.SetMaterialTexture(&earthModel.GetMaterials()[i], rl.MapDiffuse, earthTexture)
	}

	game.EarthHealth = 100
}

func UnloadEarth() {
	rl.UnloadModel(earthModel)
	rl.UnloadTexture(earthTexture)
}

func DrawEarth() {
	earthRot += earthRotSpeed * rl.GetFrameTime()
	if earthRot >= 360 { // keep the value small
		earthRot -= 360
	}
	pos := rl.Vector3{X: game.EarthPositionX, Y: game.EarthPositionY, Z: game.EarthPositionZ}
	scale := rl.Vector3{X: 1, Y: 1, Z: 1}

	rl.DrawModelEx(earthModel, pos, axisY, earthRot, scale, rl.White)
}
