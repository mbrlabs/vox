package gocraft

var (
	ColorWhite = NewColor(1, 1, 1, 1)
	ColorBlack = NewColor(0, 0, 0, 1)
	ColorRed   = NewColor(1, 0, 0, 1)
	ColorGreen = NewColor(0, 1, 0, 1)
	ColorBlue  = NewColor(0, 0, 1, 1)
	ColorTeal  = NewColor(0.5, 1, 1, 1)
)

type Color struct {
	R, G, B, A float32
}

func NewColor(r, g, b, a float32) *Color {
	return &Color{R: r, G: g, B: b, A: a}
}

func (c *Color) Copy() *Color {
	return NewColor(c.R, c.G, c.B, c.A)
}
