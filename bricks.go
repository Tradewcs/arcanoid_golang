package main

import rl "github.com/gen2brain/raylib-go/raylib"


type Brick struct {
	Position  rl.Vector2
	Width     float32
	Height    float32
	Destroyed bool
}

func GenerateBricks() []Brick {
	bricks := []Brick{}
	rows := 5
	cols := 10
	brickWidth := float32(75)
	brickHeight := float32(20)

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			bricks = append(bricks, Brick{
				Position: rl.Vector2{X: float32(col) * (brickWidth + 5), Y: float32(row) * (brickHeight + 5)},
				Width:    brickWidth,
				Height:   brickHeight,
			})
		}
	}

	return bricks
}

func (b *Brick) Draw() {
	if !b.Destroyed {
		rl.DrawRectangleV(b.Position, rl.Vector2{X: b.Width, Y: b.Height}, rl.Red)
	}
}
