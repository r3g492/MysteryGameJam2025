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

	earthRot      float32 = 0
	earthRotSpeed float32 = 10
	axisY                 = rl.Vector3{Y: 1}
	earthDamage   float32 = 0
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
	if earthRot >= 360 {
		earthRot -= 360
	}

	g := uint8(255 * (1 - earthDamage))
	b := g
	tint := rl.Color{R: 255, G: g, B: b, A: 255}

	pos := rl.Vector3{X: game.EarthPositionX, Y: game.EarthPositionY, Z: game.EarthPositionZ}
	scale := rl.Vector3{X: 1, Y: 1, Z: 1}
	rl.DrawModelEx(earthModel, pos, axisY, earthRot, scale, tint)
}

func AddEarthHit(threshold int) {
	step := 1.0 / float32(threshold)
	earthDamage += step
	if earthDamage > 1 {
		earthDamage = 1
	}
}
