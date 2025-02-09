package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Paddle struct {
	Position rl.Vector2
	Width    float32
	Height   float32
	Speed    float32
}

func NewPaddle() Paddle {
	return Paddle{
		Position: rl.Vector2{X: 350, Y: 550},
		Width:    100,
		Height:   10,
		Speed:    5,
	}
}

func (p *Paddle) Update() {
	if rl.IsKeyDown(rl.KeyLeft) && p.Position.X > 0 {
		p.Position.X -= p.Speed
	}
	if rl.IsKeyDown(rl.KeyRight) && p.Position.X+p.Width < 800 {
		p.Position.X += p.Speed
	}
}

func (p *Paddle) Draw() {
	rl.DrawRectangleV(p.Position, rl.Vector2{X: p.Width, Y: p.Height}, rl.White)
}
