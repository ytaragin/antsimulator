# Ant Colony Simulator

A 2D ant colony simulation built with Go and [Ebiten](https://ebiten.org/) game engine.

## Features

- **Autonomous Ants**: Ants move randomly within their range, avoiding obstacles
- **Nest System**: Central nest that spawns ants over time
- **Pheromone System**: Multi-type pheromone layer with decay for trail-based behavior
  - Different pheromone types (food, home, etc.) with unique colors
  - Automatic decay over time
  - Efficient grid-based storage
- **Obstacle Avoidance**: Ants detect and navigate around blocks
- **Dynamic Environment**: Randomly generated obstacles and customizable world size

## Getting Started

### Prerequisites

- Go 1.25.4 or later
- Ebiten v2 dependencies (handled automatically by Go modules)

### Installation

```bash
git clone https://github.com/ytaragin/antsimulator.git
cd ants
go mod download
```

### Running the Simulation

```bash
go run main.go
```

The simulation window will open at 640x480 pixels with an internal game resolution of 320x240 (2x pixel scaling for retro aesthetic).

## Architecture

### Project Structure

```
ants/
├── main.go                    # Entry point
├── go.mod                     # Go module dependencies
└── simulator/
    ├── ant.go                 # Individual ant behavior
    ├── antclan.go            # Ant colony management
    ├── antsimulator.go       # Main simulation engine
    ├── block.go              # Obstacle blocks
    ├── nest.go               # Ant spawning nest
    ├── pheromonelayer.go     # Pheromone trail system
    └── rect.go               # Rectangle collision utilities
```

### Key Components

**AntSimulator**: Main game loop handling updates, rendering, and collision detection

**AntClan**: Manages the ant colony and nest

**Ant**: Individual ant with autonomous movement, target seeking, and collision avoidance

**PheromoneLayer**: Grid-based pheromone system supporting multiple types:
- Data structure: `map[string][][]float64` - each pheromone type has its own 2D grid
- Configurable decay rate and maximum strength
- Visual rendering with type-specific colors (green for food, blue for home)

**Nest**: Spawns ants at regular intervals

**Block**: Static obstacles that ants must navigate around

## Customization

### Adjust World Size

In `main.go`, modify the simulator dimensions:
```go
game := simulator.NewSimulator(320, 240) // width, height
```

### Pheromone Settings

In `antsimulator.go`, adjust cell size when creating the pheromone layer:
```go
g.pheromoneLayer = NewPheromoneLayer(g.width, g.height, 5.0) // cell size in pixels
```

In `pheromonelayer.go`, modify decay rate and strength:
```go
DecayRate:   0.01,      // Pheromone decay per frame
MaxStrength: 100.0,     // Maximum pheromone intensity
```

### Ant Behavior

In `ant.go`, tune movement parameters:
```go
Speed:       0.5,   // Movement speed in pixels per frame
TargetRange: 50.0,  // Maximum distance for picking new targets
```

### Nest Configuration

In `antclan.go`, adjust spawn settings:
```go
NewNest(x, y, 10, clan, 10, 60)
// Parameters: x, y, size, clan, totalAnts, framesPerAnt
```

## License

This project is open source and available under the MIT License.

## Acknowledgments

Built with [Ebiten](https://ebiten.org/) - A dead simple 2D game engine for Go
