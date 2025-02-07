package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Game struct {
	Paddle Paddle
	Ball   Ball
	Bricks []Brick
}

func NewGame() *Game {
	game := &Game{
		Paddle: NewPaddle(),
		Ball:   NewBall(),
		Bricks: GenerateBricks(),
	}
	return game
}

func (g *Game) Update() {
	g.Paddle.Update()
	g.Ball.Update(&g.Paddle, &g.Bricks)
}

func (g *Game) Draw() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.Black)

	g.Paddle.Draw()
	g.Ball.Draw()
	for _, brick := range g.Bricks {
		brick.Draw()
	}

	rl.EndDrawing()
}
