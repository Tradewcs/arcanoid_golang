package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Brick struct {
	Position  rl.Vector2
	Width     float32
	Height    float32
	HitsLeft  int
	Color     rl.Color
	Destroyed bool
}

func GenerateBricks(level int) []Brick {
	var bricks []Brick
	rows := 5 + level
	cols := 10
	brickWidth := float32(70)
	brickHeight := float32(20)

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			hits := 3 
			color := rl.Red 

			bricks = append(bricks, Brick{
				Position: rl.Vector2{
					X: float32(col) * (brickWidth + 5),
					Y: float32(row) * (brickHeight + 5),
				},
				Width:    brickWidth,
				Height:   brickHeight,
				HitsLeft: hits,
				Color:    color,
			})
		}
	}

	return bricks
}

func (b *Brick) Hit() {
	b.HitsLeft--

	if b.HitsLeft == 2 {
		b.Color = rl.Blue
	} else if b.HitsLeft == 1 {
		b.Color = rl.Green
	} else {
		b.Destroyed = true
	}
}

func (b *Brick) Draw() {
	if !b.Destroyed {
		rl.DrawRectangleV(b.Position, rl.Vector2{X: b.Width, Y: b.Height}, b.Color)
	}
}
