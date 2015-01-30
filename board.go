package main

const BOARD_WIDTH_CELLS = 100
const BOARD_HEIGHT_CELLS = 80
const BOARD_NUM_CELLS = BOARD_WIDTH_CELLS * BOARD_HEIGHT_CELLS

const WALL_MAX_HEALTH = 100

type WallType int

const (
	LL WallType = iota
	RL
	Dot
	Straight
	RZig
	LZig
	T
)
const NUM_WALLTYPES = int(T) + 1

var wallShapes = [][][]int{

	// Left L
	[][]int{
		[]int{1, 0},
		[]int{1, 1},
		[]int{1, 2},
		[]int{0, 2},
	},

	// Right L
	[][]int{
		[]int{0, 0},
		[]int{0, 1},
		[]int{0, 2},
		[]int{1, 2},
	},

	// Dot
	[][]int{
		[]int{0, 0},
	},

	// Straight
	[][]int{
		[]int{0, 0},
		[]int{0, 1},
		[]int{0, 2},
		[]int{0, 3},
	},

	// RZig
	[][]int{
		[]int{0, 0},
		[]int{0, 1},
		[]int{1, 1},
		[]int{1, 2},
	},

	// LZig
	[][]int{
		[]int{1, 0},
		[]int{1, 1},
		[]int{0, 1},
		[]int{0, 2},
	},

	// T
	[][]int{
		[]int{0, 0},
		[]int{0, 1},
		[]int{1, 1},
		[]int{0, 2},
	},
}

type Wall struct {
	position Vector
	wType    WallType
}

func (this Wall) draw() {
	screenPos := boardToScreenPos(this.position)
	color := this.wallCells[wx*wy]
	drawFilledRectangle(screenPos.x, screenPos.y, CELL_SIZE, CELL_SIZE, blue(100))
	drawRectangle(screenPos.x, screenPos.y, CELL_SIZE, CELL_SIZE, black())
}

func (this *Wall) draw() {
	for _, p := range wallShapes[this.wType] {
		wx := this.x + p[0]
		wy := this.y + p[1]
		drawFilledRectangle(wx, wy, CELL_SIZE, CELL_SIZE, blue(100))
		drawRectangle(wx, wy, CELL_SIZE, CELL_SIZE, black())
	}
}

type Board struct {
	wallCells [BOARD_NUM_CELLS]int // [x*y] = health of cell
}

func NewBoard() *Board {
	this := new(Board)
	return this
}

func (this *Board) draw() {
	for _, p := range wallShapes[wall.wType] {
		wx := wall.x + p[0]
		wy := wall.y + p[1]
		color := this.wallCells[wx*wy]
		drawFilledRectangle(wx, wy, CELL_SIZE, CELL_SIZE, blue(color))
	}
}

func (this *Board) isWallClear(wall Wall) {

	for _, p := range wallShapes[wall.wType] {
		wx := wall.x + p[0]
		wy := wall.y + p[1]
		if this.wallCells[wx*wy] != 0 {
			return false
		}
	}
	return true
}

func (this *Board) dropWall(wall Wall) bool {

	if this.isWallClear(wall) {
		for _, p := range wallShapes[wall.wType] {
			wx := wall.x + p[0]
			wy := wall.y + p[1]
			this.wallCells[wx*wy] = WALL_MAX_HEALTH
		}
		return true
	}

	return false
}
