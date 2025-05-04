package raylib

import (
	embedWrapper "MysteryGameJam2025/embed"
	"MysteryGameJam2025/game"
	rl "github.com/gen2brain/raylib-go/raylib"
	"math"
)

var (
	sunMesh      rl.Mesh
	sunModel     rl.Model
	sunTexture   rl.Texture2D
	sunSpinDeg   float32 = 0
	sunSpinSpeed float32 = 1

	sunOrbitDeg   float32 = 170
	sunOrbitSpeed float32 = 2
	sunRadius     float32 = 20
	sunOrbitRad   float32 = 100
)

func InitSun() {
	tex, _ := embedWrapper.LoadTextureFromEmbedded("earth.png", 500, 500)
	sunTexture = tex

	sunMesh = rl.GenMeshSphere(sunRadius, 64, 64)
	sunModel = rl.LoadModelFromMesh(sunMesh)

	for i := range sunModel.GetMaterials() {
		rl.SetMaterialTexture(&sunModel.GetMaterials()[i], rl.MapDiffuse, sunTexture)
	}
}

func UnloadSun() {
	rl.UnloadModel(sunModel)
	rl.UnloadTexture(sunTexture)
}

func DrawSun() {
	dt := rl.GetFrameTime()

	sunSpinDeg += sunSpinSpeed * dt
	if sunSpinDeg >= 360 {
		sunSpinDeg -= 360
	}

	sunOrbitDeg += sunOrbitSpeed * dt
	if sunOrbitDeg >= 360 {
		sunOrbitDeg -= 360
	}

	orbitRad := sunOrbitRad
	rad := sunOrbitDeg * (math.Pi / 180)

	pos := rl.Vector3{
		X: float32(math.Cos(float64(rad))) * orbitRad,
		Y: 0,
		Z: float32(math.Sin(float64(rad))) * orbitRad,
	}

	rl.DrawSphere(
		pos,
		sunRadius,
		rl.Yellow,
	)
}

func GetSunCollision() (x float32, z float32, radius float32) {
	orbitRad := sunOrbitRad
	rad := sunOrbitDeg * (math.Pi / 180)

	x = float32(math.Cos(float64(rad))) * orbitRad
	z = float32(math.Sin(float64(rad))) * orbitRad
	radius = sunRadius

	return x, z, radius
}

func SunCheck() bool {
	x, z, radius := GetSunCollision()
	for i := len(game.ProjectileList) - 1; i >= 0; i-- {
		p := game.ProjectileList[i]

		dx := p.PosX - x
		dy := p.PosY - 0
		dz := p.PosZ - z
		if dx*dx+dy*dy+dz*dz <= radius*radius && p.Type == game.COMM {
			game.ProjectileList = append(game.ProjectileList[:i], game.ProjectileList[i+1:]...)
			game.SunComm = true
			return true
		}
	}
	return false
}
