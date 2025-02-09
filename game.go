package main

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	Paddle    Paddle
	Ball      Ball
	Bricks    []Brick
	Points    int
	Level     int
	Speed     float32
	BricksHit int
}

func NewGame() *Game {
	game := &Game{
		Paddle: NewPaddle(),
		Ball:   NewBall(),
		Points: 0,
		Level:  1,
		Speed:  5,
	}
	game.GenerateLevel()
	return game
}

func (g *Game) Update() {
	g.Paddle.Update()
	g.Ball.Update(&g.Paddle, &g.Bricks, g)
}

func (g *Game) Draw(bg rl.Texture2D) {
	rl.BeginDrawing()
	rl.ClearBackground(rl.White)

	bgRect := rl.NewRectangle(0, 0, float32(screenWidth), float32(screenHeight))
	rl.DrawTexturePro(bg, rl.NewRectangle(0, 0, float32(bg.Width), float32(bg.Height)), bgRect, rl.Vector2{}, 0, rl.White)

	g.Paddle.Draw()
	g.Ball.Draw()
	for _, brick := range g.Bricks {
		brick.Draw()
	}
	rl.DrawText(fmt.Sprintf("Points: %d", g.Points), 10, 25, 20, rl.White)
	rl.DrawText(fmt.Sprintf("Level: %d", g.Level), 10, 50, 20, rl.White)

	rl.EndDrawing()
}

func (g *Game) GenerateLevel() {
	g.Bricks = GenerateBricks(g.Level)
	g.Ball = NewBall()
	g.Paddle = NewPaddle()
	g.Speed += float32(g.Level) * 0.5
}
