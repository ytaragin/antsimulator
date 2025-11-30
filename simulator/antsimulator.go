package simulator

import (
	"fmt"
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type AntSimulator struct {
	clan           *AntClan
	blocks         []*Block
	pheromoneLayer *PheromoneLayer
	frameCounter   int
	width          int
	height         int
}

func (g *AntSimulator) Update() error {
	g.frameCounter++
	g.clan.Update()
	g.pheromoneLayer.Update()
	return nil
}

func (g *AntSimulator) Draw(screen *ebiten.Image) {
	g.pheromoneLayer.Draw(screen)
	for _, block := range g.blocks {
		block.Draw(screen)
	}
	g.clan.Draw(screen)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Frame: %d", g.frameCounter), g.width-100, 10)
}

func (g *AntSimulator) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.width, g.height
}

func NewSimulator(width, height int) *AntSimulator {
	sim := &AntSimulator{
		width:  width,
		height: height,
	}
	sim.init()
	return sim
}

func (g *AntSimulator) init() {
	g.blocks = g.createBlocks()
	g.pheromoneLayer = NewPheromoneLayer(g.width, g.height, 5.0)
	g.clan = NewAntClan(float64(g.width)/2, float64(g.height)/2, g)
}

func (g *AntSimulator) createBlocks() []*Block {
	// Create border blocks
	borderColor := color.RGBA{128, 128, 128, 255} // Gray
	borderThickness := 5.0
	screenWidth := float64(g.width)
	screenHeight := float64(g.height)

	blocks := []*Block{
		// Top border
		NewBlock(0, 0, screenWidth, borderThickness, borderColor),
		// Bottom border
		NewBlock(0, screenHeight-borderThickness, screenWidth, borderThickness, borderColor),
		// Left border
		NewBlock(0, 0, borderThickness, screenHeight, borderColor),
		// Right border
		NewBlock(screenWidth-borderThickness, 0, borderThickness, screenHeight, borderColor),
	}

	// Add a few random blocks
	for i := 0; i < 5; i++ {
		x := rand.Float64()*(screenWidth-40) + 20
		y := rand.Float64()*(screenHeight-40) + 20
		width := rand.Float64()*20 + 10
		height := rand.Float64()*20 + 10
		c := color.RGBA{
			uint8(rand.Intn(256)),
			uint8(rand.Intn(256)),
			uint8(rand.Intn(256)),
			255,
		}
		blocks = append(blocks, NewBlock(x, y, width, height, c))
	}

	return blocks
}

func (g *AntSimulator) CheckCollision(rect Rect) bool {
	for _, block := range g.blocks {
		if rect.Intersects(block.Rect) {
			return true
		}
	}
	return false
}
