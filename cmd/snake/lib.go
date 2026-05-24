package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var global_scale float32 = 4.0

var screenWidth = 128
var screenHeight = 128

var originalPico8Palette = []color.Color{
	color.RGBA{R: 0, G: 0, B: 0, A: 255},       // 0 black
	color.RGBA{R: 29, G: 43, B: 83, A: 255},    // 1 dark-blue
	color.RGBA{R: 126, G: 37, B: 83, A: 255},   // 2 dark-purple
	color.RGBA{R: 0, G: 135, B: 81, A: 255},    // 3 dark-green
	color.RGBA{R: 171, G: 82, B: 54, A: 255},   // 4 brown
	color.RGBA{R: 95, G: 87, B: 79, A: 255},    // 5 dark-gray
	color.RGBA{R: 194, G: 195, B: 199, A: 255}, // 6 light-gray
	color.RGBA{R: 255, G: 241, B: 232, A: 255}, // 7 white
	color.RGBA{R: 255, G: 0, B: 77, A: 255},    // 8 red
	color.RGBA{R: 255, G: 163, B: 0, A: 255},   // 9 orange
	color.RGBA{R: 255, G: 236, B: 39, A: 255},  // 10 yellow
	color.RGBA{R: 0, G: 228, B: 54, A: 255},    // 11 green
	color.RGBA{R: 41, G: 173, B: 255, A: 255},  // 12 blue
	color.RGBA{R: 131, G: 118, B: 156, A: 255}, // 13 indigo
	color.RGBA{R: 255, G: 119, B: 168, A: 255}, // 14 pink
	color.RGBA{R: 255, G: 204, B: 170, A: 255}, // 15 peach
}

// rect draws a rectangle on the screen with the given coordinates, width, height, and color index.
func rect(x, y, w, h float32, c int) {
	if c < 0 || c >= len(originalPico8Palette) {
		return
	}
	vector.FillRect(game.Screen, x*global_scale, y*global_scale, w*global_scale, h*global_scale, originalPico8Palette[c], false)
}

var buttonMapping = map[int]ebiten.Key{
	0: ebiten.KeyArrowUp,
	1: ebiten.KeyArrowDown,
	2: ebiten.KeyArrowLeft,
	3: ebiten.KeyArrowRight,
}

var buttonHistory = make(map[int]bool)

func btn(btn int) bool {
	if btn < 0 || btn >= len(buttonMapping) {
		return false
	}
	if ebiten.IsKeyPressed(buttonMapping[btn]) {
		return true
	}
	return false
}
