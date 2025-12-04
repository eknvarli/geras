package engine

type Body struct {
	X, Y   float64
	W, H   float64
	VX, VY float64
	Mass   float64
	Static bool
}

func Integrate(b *Body, dt float64) {
	if b == nil || b.Static {
		return
	}
	b.X += b.VX * dt
	b.Y += b.VY * dt
}

func AABBIntersect(a, b *Body) bool {
	return !(a.X+a.W < b.X || b.X+b.W < a.X || a.Y+a.H < b.Y || b.Y+b.H < a.Y)
}
