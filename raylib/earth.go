package raylib

import (
	embedWrapper "MysteryGameJam2025/embed"
	"MysteryGameJam2025/game"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var earthMesh rl.Mesh
var earthModel rl.Model
var earthTexture rl.Texture2D

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
	pos := rl.Vector3{X: game.EarthPositionX, Y: game.EarthPositionY, Z: game.EarthPositionZ}
	rl.DrawModel(earthModel, pos, 1.0, rl.White)
}
