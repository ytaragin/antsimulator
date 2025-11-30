package simulator

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Nest struct {
	X, Y         float64
	Radius       float64
	clan         *AntClan
	totalAnts    int
	spawnRate    int // frames between spawns
	antsSpawned  int
	frameCounter int
}

func NewNest(x, y, radius float64, clan *AntClan, totalAnts, spawnRate int) *Nest {
	return &Nest{
		X:            x,
		Y:            y,
		Radius:       radius,
		clan:         clan,
		totalAnts:    totalAnts,
		spawnRate:    spawnRate,
		antsSpawned:  0,
		frameCounter: 0,
	}
}

func (n *Nest) Draw(screen *ebiten.Image) {
	// Draw a blue circle
	vector.FillCircle(screen, float32(n.X), float32(n.Y), float32(n.Radius), color.RGBA{0, 0, 255, 255}, false)
}

func (n *Nest) Update() {
	n.ManagePopulation()
}

func (n *Nest) ManagePopulation() {
	if n.antsSpawned >= n.totalAnts {
		return
	}

	n.frameCounter++
	if n.frameCounter >= n.spawnRate {
		n.SpawnAnt()
		n.frameCounter = 0
	}
}

func (n *Nest) SpawnAnt() {
	ant := NewAnt(n.X, n.Y, n.clan)
	n.clan.AddAnt(ant)
	n.antsSpawned++
}
