package main

type Vector struct {
	x, y int
}

func (this Vector) dist(v Vector) int {
	return math.sqrt((this.x-v.x)^2, (this.y-v.y)^2)
}

type WindowPos Vector

func (this WindowPos) toBoard() BoardPos {
	return BoardPos{this.x / CELL_SIZE, this.y / CELL_SIZE}
}

func (this WindowPos) dist(v WindowPos) int {
	return this.dist(Vector{v.x, v.y})
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

func (this BoardPos) dist(v BoardPos) int {
	return this.dist(Vector{v.x, v.y})
}

type Bounds struct {
	UpLeft   WindowPos
	LowRight WindowPos
}
