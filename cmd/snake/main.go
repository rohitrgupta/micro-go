package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	Screen *ebiten.Image
}

func (g *Game) Update() error {
	return Update()
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.Screen = screen
	Draw()
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth * int(global_scale), screenHeight * int(global_scale)
}

var game = &Game{}

func main() {
	ebiten.SetWindowSize(screenWidth*int(global_scale), screenHeight*int(global_scale))
	ebiten.SetWindowTitle("Snake using Micro-GO")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
