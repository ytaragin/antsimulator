package simulator

import "github.com/hajimehoshi/ebiten/v2"

type AntClan struct {
	ants      []*Ant
	nest      *Nest
	simulator *AntSimulator
}

func NewAntClan(x, y float64, simulator *AntSimulator) *AntClan {
	clan := &AntClan{
		ants:      []*Ant{},
		simulator: simulator,
	}
	clan.nest = NewNest(x, y, 10, clan, 10, 60) // 10 ants total, spawn every 60 frames (1 per second at 60 FPS)
	return clan
}

func (c *AntClan) Update() {
	c.nest.Update()
	for _, ant := range c.ants {
		ant.Update()
	}
}

func (c *AntClan) Draw(screen *ebiten.Image) {
	for _, ant := range c.ants {
		ant.Draw(screen)
	}
	c.nest.Draw(screen)
}

func (c *AntClan) AddAnt(ant *Ant) {
	c.ants = append(c.ants, ant)
}
