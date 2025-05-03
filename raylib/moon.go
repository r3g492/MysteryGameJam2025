package raylib

import (
	embedWrapper "MysteryGameJam2025/embed"
	rl "github.com/gen2brain/raylib-go/raylib"
	"math"
)

var (
	moonMesh      rl.Mesh
	moonModel     rl.Model
	moonTexture   rl.Texture2D
	moonSpinDeg   float32 = 0
	moonSpinSpeed float32 = 10

	moonOrbitDeg   float32 = 0
	moonOrbitSpeed float32 = 5
)

var moonAxisY = rl.Vector3{Y: 1}

func InitMoon() {
	tex, _ := embedWrapper.LoadTextureFromEmbedded("earth.png", 500, 500) // texture path unchanged
	moonTexture = tex

	moonMesh = rl.GenMeshSphere(1, 64, 64)
	moonModel = rl.LoadModelFromMesh(moonMesh)

	for i := range moonModel.GetMaterials() {
		rl.SetMaterialTexture(&moonModel.GetMaterials()[i], rl.MapDiffuse, moonTexture)
	}
}

func UnloadMoon() {
	rl.UnloadModel(moonModel)
	rl.UnloadTexture(moonTexture)
}

func DrawMoon() {
	dt := rl.GetFrameTime()

	moonSpinDeg += moonSpinSpeed * dt
	if moonSpinDeg >= 360 {
		moonSpinDeg -= 360
	}

	moonOrbitDeg += moonOrbitSpeed * dt
	if moonOrbitDeg >= 360 {
		moonOrbitDeg -= 360
	}

	orbitRad := float32(20)
	rad := moonOrbitDeg * (math.Pi / 180)

	pos := rl.Vector3{
		X: float32(math.Cos(float64(rad))) * orbitRad,
		Y: 0,
		Z: float32(math.Sin(float64(rad))) * orbitRad,
	}

	rl.DrawSphere(
		pos,
		1,
		rl.DarkGray,
	)
}
