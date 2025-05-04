package game

const (
	COMM = iota
	MISSILE
)

type Projectile struct {
	Type       int
	PosX       float32
	PosY       float32
	PosZ       float32
	DirectionX float32
	DirectionY float32
	DirectionZ float32
	Speed      float32
}

var (
	ProjectileList []*Projectile
)

func MoveProjectile() {
	for i := len(ProjectileList) - 1; i >= 0; i-- {
		p := ProjectileList[i]

		p.PosX += p.DirectionX * p.Speed
		p.PosY += p.DirectionY * p.Speed
		p.PosZ += p.DirectionZ * p.Speed

		if p.PosX > 500 || p.PosX < -500 ||
			p.PosY > 500 || p.PosY < -500 ||
			p.PosZ > 500 || p.PosZ < -500 {
			ProjectileList = append(ProjectileList[:i], ProjectileList[i+1:]...)
		}
	}
}

func AddProjectile(p *Projectile) {
	ProjectileList = append(ProjectileList, p)
}
