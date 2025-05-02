package raylib

import (
	embedWrapper "MysteryGameJam2025/embed"
	"MysteryGameJam2025/game"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	moonMesh    rl.Mesh
	moonModel   rl.Model
	moonTexture rl.Texture2D

	moonRot      float32 = 0                // current angle in degrees
	moonRotSpeed float32 = 10               // °/sec – tweak to taste
	moonAxisY            = rl.Vector3{Y: 1} // (0,1,0) – spin around Y
)

func InitMoon() {
	moonTextureResponse, _ := embedWrapper.LoadTextureFromEmbedded(
		"earth.png",
		500,
		500,
	)
	moonTexture = moonTextureResponse

	moonMesh = rl.GenMeshSphere(
		game.MoonRadius,
		64,
		64,
	)
	moonModel = rl.LoadModelFromMesh(moonMesh)

	for i := range moonModel.GetMaterials() {
		rl.SetMaterialTexture(&earthModel.GetMaterials()[i], rl.MapDiffuse, moonTexture)
	}
}

func UnloadMoon() {
	rl.UnloadModel(moonModel)
	rl.UnloadTexture(moonTexture)
}

func DrawMoon() {
	moonRot += moonRotSpeed * rl.GetFrameTime()
	if moonRot >= 360 {
		moonRot -= 360
	}
	pos := rl.Vector3{X: game.MoonPositionX, Y: game.MoonPositionY, Z: game.MoonPositionZ}
	scale := rl.Vector3{X: 1, Y: 1, Z: 1}

	rl.DrawModelEx(moonModel, pos, moonAxisY, moonRot, scale, rl.DarkGray)
}
