package simulator

import (
	"image/color"
	"math"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Ant struct {
	Speed       float64
	TargetX     float64
	TargetY     float64
	TargetRange float64 // Maximum distance to pick a new target
	HasTarget   bool
	Rect        Rect
	clan        *AntClan
}

func NewAnt(x, y float64, clan *AntClan) *Ant {
	ant := &Ant{
		Speed:       0.5,
		TargetRange: 50.0,
		HasTarget:   false,
		Rect:        NewRect(x, y, 2, 2),
		clan:        clan,
	}
	ant.pickNewTarget()
	return ant
}

func (a *Ant) Update() {
	if !a.HasTarget {
		a.pickNewTarget()
	}

	// Calculate distance to target
	dx := a.TargetX - a.Rect.X
	dy := a.TargetY - a.Rect.Y
	distance := math.Sqrt(dx*dx + dy*dy)

	// If we've reached the target (within 2 pixels), pick a new target
	if distance < 2.0 {
		a.pickNewTarget()
		return
	}

	// Calculate new position
	newX := a.Rect.X + (dx/distance)*a.Speed
	newY := a.Rect.Y + (dy/distance)*a.Speed

	// Create a rect for the new position
	newRect := NewRect(newX, newY, a.Rect.Width, a.Rect.Height)

	// Check if the new position would collide with any blocks
	if a.clan.simulator.CheckCollision(newRect) {
		// Collision detected, pick a new target instead of moving
		a.pickNewTarget()
		return
	}

	// No collision, move to new position
	a.Rect.X = newX
	a.Rect.Y = newY
}

func (a *Ant) pickNewTarget() {
	// Pick a random target within range
	angle := rand.Float64() * 2 * math.Pi
	distance := rand.Float64() * a.TargetRange

	a.TargetX = a.Rect.X + math.Cos(angle)*distance
	a.TargetY = a.Rect.Y + math.Sin(angle)*distance

	a.HasTarget = true
}

func (a *Ant) Draw(screen *ebiten.Image) {
	// Draw a 1x1 red square using the Rect
	vector.FillRect(screen, float32(a.Rect.X), float32(a.Rect.Y), float32(a.Rect.Width), float32(a.Rect.Height), color.RGBA{255, 0, 0, 255}, false)
}
