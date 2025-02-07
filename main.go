package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	screenWidth  = 800
	screenHeight = 600
)

func main() {
	rl.InitWindow(screenWidth, screenHeight, "Arkanoid in Go")
	rl.SetTargetFPS(60)

	game := NewGame()

	for !rl.WindowShouldClose() {
		game.Update()
		game.Draw()
	}

	rl.CloseWindow()
}
