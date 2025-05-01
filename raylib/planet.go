package raylib

import (
	embedWrapper "MysteryGameJam2025/embed"
	"MysteryGameJam2025/game"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func DrawPlanets(
	planet game.Planet,
) {
	pos := rl.Vector3{X: planet.X, Y: planet.Y, Z: planet.Z}

	planetTexture, _ := embedWrapper.LoadTextureFromEmbedded(
		"earth.png",
		500,
		500,
	)

	planetMesh := rl.GenMeshSphere(
		planet.Radius,
		64,
		64,
	)
	planetModel := rl.LoadModelFromMesh(planetMesh)

	for i := range planetModel.GetMaterials() {
		rl.SetMaterialTexture(&planetModel.GetMaterials()[i], rl.MapDiffuse, planetTexture)
	}

	rl.DrawModel(planetModel, pos, 1.0, rl.White)
}
