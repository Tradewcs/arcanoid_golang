package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	screenWidth  = 745
	screenHeight = 600
)

func main() {
	rl.InitWindow(screenWidth, screenHeight, "Arkanoid in Go")
	rl.SetTargetFPS(60)

	game := NewGame()
	bg := rl.LoadTexture("assets/background.png")
	defer rl.UnloadTexture(bg)

	for !rl.WindowShouldClose() {
		game.Update()
		game.Draw(bg)
	}

	rl.CloseWindow()
}
