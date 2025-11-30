package simulator

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type PheromoneLayer struct {
	Width       int
	Height      int
	CellSize    float64
	Grid        map[string][][]float64
	GridWidth   int
	GridHeight  int
	DecayRate   float64
	MaxStrength float64
}

func NewPheromoneLayer(width, height int, cellSize float64) *PheromoneLayer {
	gridWidth := int(float64(width) / cellSize)
	gridHeight := int(float64(height) / cellSize)

	grid := make(map[string][][]float64)

	return &PheromoneLayer{
		Width:       width,
		Height:      height,
		CellSize:    cellSize,
		Grid:        grid,
		GridWidth:   gridWidth,
		GridHeight:  gridHeight,
		DecayRate:   0.01,
		MaxStrength: 100.0,
	}
}

func (p *PheromoneLayer) AddPheromone(x, y, strength float64, pheroType string) {
	gridX := int(x / p.CellSize)
	gridY := int(y / p.CellSize)

	if p.inGrid(gridX, gridY) {
		// Initialize the 2D array for this pheromone type if it doesn't exist
		if _, exists := p.Grid[pheroType]; !exists {
			p.initPheromoneType(pheroType)
		}
		current := p.Grid[pheroType][gridY][gridX]
		p.Grid[pheroType][gridY][gridX] = math.Min(current+strength, p.MaxStrength)
	}
}

func (p *PheromoneLayer) GetPheromone(x, y float64, pheroType string) float64 {
	gridX := int(x / p.CellSize)
	gridY := int(y / p.CellSize)

	if p.inGrid(gridX, gridY) {
		if grid, exists := p.Grid[pheroType]; exists {
			return grid[gridY][gridX]
		}
	}
	return 0.0
}

func (p *PheromoneLayer) Update() {
	for pheroType, grid := range p.Grid {
		hasNonZero := false
		for y := range grid {
			for x := range grid[y] {
				grid[y][x] -= p.DecayRate
				if grid[y][x] < 0 {
					grid[y][x] = 0
				} else if grid[y][x] > 0 {
					hasNonZero = true
				}
			}
		}
		// Remove the entire grid if all values are zero
		if !hasNonZero {
			delete(p.Grid, pheroType)
		}
	}
}

func (p *PheromoneLayer) Draw(screen *ebiten.Image) {
	for pheroType, grid := range p.Grid {
		for y := range grid {
			for x := range grid[y] {
				strength := grid[y][x]
				if strength > 0 {
					intensity := uint8((strength / p.MaxStrength) * 255)
					// Use different colors for different pheromone types
					c := p.getColorForType(pheroType, intensity)
					vector.FillRect(
						screen,
						float32(x)*float32(p.CellSize),
						float32(y)*float32(p.CellSize),
						float32(p.CellSize),
						float32(p.CellSize),
						c,
						false,
					)
				}
			}
		}
	}
}

func (p *PheromoneLayer) getColorForType(pheroType string, intensity uint8) color.RGBA {
	// Assign different colors to different pheromone types
	switch pheroType {
	case "food":
		return color.RGBA{0, intensity, 0, 128} // Green
	case "home":
		return color.RGBA{0, 0, intensity, 128} // Blue
	default:
		return color.RGBA{intensity, intensity, 0, 128} // Yellow for unknown types
	}
}

func (p *PheromoneLayer) initPheromoneType(pheroType string) {
	grid := make([][]float64, p.GridHeight)
	for i := range grid {
		grid[i] = make([]float64, p.GridWidth)
	}
	p.Grid[pheroType] = grid
}

func (p *PheromoneLayer) inGrid(x, y int) bool {
	return x >= 0 && x < p.GridWidth && y >= 0 && y < p.GridHeight
}
