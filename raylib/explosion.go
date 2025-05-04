package raylib

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"math"
)

type Explosion struct {
	Pos   rl.Vector3
	Timer float32 // seconds left
}

var explosions []*Explosion

const (
	explosionLife   = 0.6
	explosionRadius = 15
	boomLines       = 12
)

func AddExplosion(pos rl.Vector3) {
	explosions = append(explosions, &Explosion{
		Pos:   pos,
		Timer: explosionLife,
	})
}

func UpdateExplosions(dt float32) {
	for i := len(explosions) - 1; i >= 0; i-- {
		e := explosions[i]
		e.Timer -= dt
		if e.Timer <= 0 {
			explosions = append(explosions[:i], explosions[i+1:]...)
		}
	}
}

func DrawExplosions() {
	for _, e := range explosions {
		alpha := e.Timer / explosionLife
		col := rl.Fade(rl.Red, alpha) // fade out
		r := explosionRadius * (1 - alpha)

		for i := 0; i < boomLines; i++ {
			angle := float32(i) * 2 * math32Pi / boomLines
			dir := rl.Vector3{X: r * math32Cos(angle), Z: r * math32Sin(angle)}
			rl.DrawLine3D(e.Pos, rl.Vector3{X: e.Pos.X + dir.X, Y: e.Pos.Y, Z: e.Pos.Z + dir.Z}, col)
		}
	}
}

const math32Pi = 3.14159274

func math32Sin(x float32) float32 { return float32(math.Sin(float64(x))) }
func math32Cos(x float32) float32 { return float32(math.Cos(float64(x))) }
