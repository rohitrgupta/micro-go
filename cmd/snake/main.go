package main

import (
	"bytes"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type Game struct {
	Screen   *ebiten.Image
	fontFace *text.GoTextFace
}

func NewGame() *Game {
	fontSource, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.PressStart2P_ttf))
	if err != nil {
		log.Fatal(err)
	}

	fontFace := &text.GoTextFace{
		Source: fontSource,
		Size:   8,
	}

	return &Game{
		fontFace: fontFace,
	}
}

func (g *Game) Update() error {
	return Update()
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.Screen = screen
	Draw()
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

var microGame = NewGame()

func main() {
	ebiten.SetWindowTitle("Snake using Micro-GO")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	if err := ebiten.RunGame(microGame); err != nil {
		log.Fatal(err)
	}
}
