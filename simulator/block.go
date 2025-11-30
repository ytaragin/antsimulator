package simulator

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Block struct {
	Rect  Rect
	Color color.Color
}

func NewBlock(x, y, width, height float64, color color.Color) *Block {
	return &Block{
		Rect: Rect{
			X:      x,
			Y:      y,
			Width:  width,
			Height: height,
		},
		Color: color,
	}
}

func (b *Block) Draw(screen *ebiten.Image) {
	vector.FillRect(screen,
		float32(b.Rect.X),
		float32(b.Rect.Y),
		float32(b.Rect.Width),
		float32(b.Rect.Height),
		b.Color,
		true)
}
