package game

import (
	"math"
	"math/rand"
	"time"
)

type Drone struct {
	X          float32
	Y          float32
	Z          float32
	DirectionX float32
	DirectionY float32
	DirectionZ float32
	MoveSpeed  float32
	HitRadius  float32
	ShowRadius float32
}

var (
	DroneList []*Drone
	DroneTurn = false
)

func init() { rand.Seed(time.Now().UnixNano()) }

func GenerateDrone(
	howMany int,
	speed float32,
) {
	const (
		minRadius = 300.0
		maxRadius = 400.0
	)

	for i := 0; i < howMany; i++ {
		theta := rand.Float64() * 2 * math.Pi
		phi := math.Acos(1 - 2*rand.Float64())
		r := minRadius + rand.Float64()*(maxRadius-minRadius)

		x := float32(r * math.Sin(phi) * math.Cos(theta))
		y := float32(0)
		z := float32(r * math.Cos(phi))

		lenInv := 1 / float32(r)
		dx := -x * lenInv
		dy := -y * lenInv
		dz := -z * lenInv

		drone := &Drone{
			X:          x,
			Y:          0,
			Z:          z,
			DirectionX: dx,
			DirectionY: dy,
			DirectionZ: dz,
			MoveSpeed:  speed,
			HitRadius:  1.3,
			ShowRadius: 0.8,
		}
		DroneList = append(DroneList, drone)
	}
}
func EnemyAllDead() bool {
	return len(DroneList) == 0
}

func ProjectileCheck() (float32, float32, float32, bool) {
	for i := len(ProjectileList) - 1; i >= 0; i-- {
		p := ProjectileList[i]

		for j := len(DroneList) - 1; j >= 0; j-- {
			d := DroneList[j]

			dx := p.PosX - d.X
			dy := p.PosY - d.Y
			dz := p.PosZ - d.Z
			if dx*dx+dy*dy+dz*dz <= d.HitRadius*d.HitRadius && p.Type == MISSILE {
				ProjectileList = append(ProjectileList[:i], ProjectileList[i+1:]...)
				DroneList = append(DroneList[:j], DroneList[j+1:]...)
				return d.X, d.Y, d.Z, true
			}
		}
	}
	return 0, 0, 0, false
}

func CommCheck() bool {
	for i := len(ProjectileList) - 1; i >= 0; i-- {
		p := ProjectileList[i]

		for j := len(DroneList) - 1; j >= 0; j-- {
			d := DroneList[j]

			dx := p.PosX - d.X
			dy := p.PosY - d.Y
			dz := p.PosZ - d.Z
			if dx*dx+dy*dy+dz*dz <= d.HitRadius*d.HitRadius && p.Type == COMM {
				ProjectileList = append(ProjectileList[:i], ProjectileList[i+1:]...)
				DroneList = append(DroneList[:j], DroneList[j+1:]...)
				return true
			}
		}
	}
	return false
}

func EarthCheck() (float32, float32, float32, bool) {
	for j := len(DroneList) - 1; j >= 0; j-- {
		d := DroneList[j]

		dx := EarthPositionX - d.X
		dy := EarthPositionY - d.Y
		dz := EarthPositionZ - d.Z
		if dx*dx+dy*dy+dz*dz <= d.HitRadius*d.HitRadius*EarthRadius {
			DroneList = append(DroneList[:j], DroneList[j+1:]...)
			return d.X, d.Y, d.Z, true
		}
	}
	return 0, 0, 0, false
}

func MoveDrone() {
	for i := len(DroneList) - 1; i >= 0; i-- {
		d := DroneList[i]

		d.X += d.DirectionX * d.MoveSpeed
		d.Y += d.DirectionY * d.MoveSpeed
		d.Z += d.DirectionZ * d.MoveSpeed

		if d.X > 500 || d.X < -500 ||
			d.Y > 500 || d.Y < -500 ||
			d.Z > 500 || d.Z < -500 {
			DroneList = append(DroneList[:i], DroneList[i+1:]...)
		}
	}
}

func TurnDrone() {
	DroneTurn = true
	for i := len(DroneList) - 1; i >= 0; i-- {
		d := DroneList[i]

		d.DirectionX = 0
		d.DirectionY = 0
		d.DirectionZ = 0
	}
}
