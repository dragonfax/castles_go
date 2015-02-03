package main

import "math"

type Vector struct {
	x, y int
}

func (this Vector) dist(v Vector) float64 {
	return math.Sqrt(math.Pow(float64(this.x-v.x), 2) + math.Pow(float64(this.y-v.y), 2))
}

type WindowPos Vector

func (this WindowPos) toBoard() BoardPos {
	return BoardPos{this.x / CELL_SIZE, this.y / CELL_SIZE}
}

func (this WindowPos) dist(v WindowPos) float64 {
	return Vector{this.x, this.y}.dist(Vector{v.x, v.y})
}

type BoardPos Vector

func (this BoardPos) toWindowCenter() WindowPos {
	return WindowPos{this.x*CELL_SIZE + CELL_SIZE/2, this.y*CELL_SIZE + CELL_SIZE/2}
}

func (this BoardPos) toWindowUpLeft() WindowPos {
	return WindowPos{this.x * CELL_SIZE, this.y * CELL_SIZE}
}

func (this BoardPos) toWindowLowRight() WindowPos {
	return WindowPos{this.x*CELL_SIZE + CELL_SIZE, this.y*CELL_SIZE + CELL_SIZE}
}

func (this BoardPos) toWindowBounds() Bounds {
	return Bounds{this.toWindowUpLeft(), this.toWindowLowRight()}
}

func (this BoardPos) dist(v BoardPos) float64 {
	return Vector{this.x, this.y}.dist(Vector{v.x, v.y})
}

func (this BoardPos) isZero() bool {
	return this.x == 0 && this.y == 0
}
