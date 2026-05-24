package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Point struct {
	X float32
	Y float32
	c int
}

const bodyColor = 8
const headColor = 9
const foodColor = 10

var cellSize float32 = 8
var gridWidth = int(float32(screenWidth) / cellSize)
var gridHeight = int(float32(screenHeight) / cellSize)

var snake = []Point{}
var food = Point{X: 10, Y: 10, c: foodColor}
var direction = Point{X: 1, Y: 0}
var updateSpeed = time.Second / 4
var lastUpdateTime time.Time

func init() {
	fmt.Println("Initializing snake game size", gridWidth, gridHeight)
	snake = append(snake, Point{X: 5, Y: 5, c: bodyColor})
	snake = append(snake, Point{X: 3, Y: 5, c: headColor})
	lastUpdateTime = time.Now()
	food = Point{X: float32(rand.Intn(gridWidth)), Y: float32(rand.Intn(gridHeight)), c: foodColor}
}

func Update() error {

	if btn(0) && direction.Y != 1 {
		direction = Point{X: 0, Y: -1}
	} else if btn(1) && direction.Y != -1 {
		direction = Point{X: 0, Y: 1}
	} else if btn(2) && direction.X != 1 {
		direction = Point{X: -1, Y: 0}
	} else if btn(3) && direction.X != -1 {
		direction = Point{X: 1, Y: 0}
	}

	if time.Since(lastUpdateTime) < updateSpeed {
		return nil
	}

	lastUpdateTime = time.Now()
	for i := len(snake) - 1; i > 0; i-- {
		snake[i] = snake[i-1]
		snake[i].c = bodyColor
	}
	snake[0].X += direction.X
	snake[0].Y += direction.Y
	snake[0].c = headColor

	if snake[0].X == food.X && snake[0].Y == food.Y {
		food = Point{X: float32(rand.Intn(gridWidth)), Y: float32(rand.Intn(gridHeight)), c: foodColor}
		snake = append(snake, Point{X: snake[len(snake)-1].X, Y: snake[len(snake)-1].Y, c: bodyColor})
	}

	if snake[0].X < 0 ||
		snake[0].X >= float32(gridWidth) ||
		snake[0].Y < 0 ||
		snake[0].Y >= float32(gridHeight) {

		snake = snake[:0]
		snake = append(snake, Point{X: 3, Y: 5, c: headColor})
		direction = Point{X: 1, Y: 0}
		return nil
	}
	// fmt.Println(snake)
	return nil
}

func Draw() {
	for _, point := range snake {
		rect(point.X*cellSize, point.Y*cellSize, cellSize, cellSize, point.c)
	}
	rect(food.X*cellSize, food.Y*cellSize, cellSize, cellSize, food.c)
}
