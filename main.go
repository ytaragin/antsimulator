package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/ytaragin/ants/simulator"
)

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	game := simulator.NewSimulator(640, 480)
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
