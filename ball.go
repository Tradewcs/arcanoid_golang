package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Ball struct {
	Position rl.Vector2
	Radius   float32
	Speed    rl.Vector2
}

func NewBall() Ball {
	return Ball{
		Position: rl.Vector2{X: screenWidth / 2 - 50, Y: screenHeight - 100},
		Radius:   5,
		Speed:    rl.Vector2{X: 4, Y: -4},
	}
}

func (b *Ball) Update(paddle *Paddle, bricks *[]Brick) {

	
	b.Position.X += b.Speed.X
	b.Position.Y += b.Speed.Y

	if b.Position.X < 0 || b.Position.X > screenWidth {
		b.Speed.X = -b.Speed.X
	}
	if b.Position.Y < 0 {
		b.Speed.Y = -b.Speed.Y
	}

	if b.Position.Y+b.Radius >= paddle.Position.Y && b.Position.X >= paddle.Position.X && b.Position.X <= paddle.Position.X+paddle.Width {
		b.Speed.Y = -b.Speed.Y
	}

	for i := 0; i < len(*bricks); i++ {
		brick := &(*bricks)[i]
		if !brick.Destroyed && b.Position.X >= brick.Position.X && b.Position.X <= brick.Position.X+brick.Width &&
			b.Position.Y >= brick.Position.Y && b.Position.Y <= brick.Position.Y+brick.Height {
			b.Speed.Y = -b.Speed.Y
			brick.Destroyed = true
		}
	}

	if b.Position.Y > screenHeight {
		*b = NewBall()
	}
}

func (b *Ball) Draw() {
	rl.DrawCircleV(b.Position, b.Radius, rl.White)
}
