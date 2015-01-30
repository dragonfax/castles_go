package main

import "fmt"

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

func (this *Wall) draw() {
	for _, p := range wallShapes[this.wType] {
		wx := this.position.x + p[0]
		wy := this.position.y + p[1]
		bv := boardToScreenPos(Vector{wx, wy})
		drawFilledRectangle(bv.x, bv.y, CELL_SIZE, CELL_SIZE, blue(100))
		drawRectangle(bv.x, bv.y, CELL_SIZE, CELL_SIZE, black())
	}
}

type Board struct {
	wallCells [BOARD_NUM_CELLS]uint8 // [x*y] = health of cell
}

func NewBoard() *Board {
	this := new(Board)
	return this
}

func (this *Board) draw() {
	for x := 0; x < BOARD_WIDTH_CELLS; x++ {
		for y := 0; y < BOARD_HEIGHT_CELLS; y++ {
			fmt.Printf("indexing %d,%d which is in cell %d and %d\n", x, y, x*BOARD_WIDTH_CELLS+y, BOARD_NUM_CELLS)
			color := this.wallCells[x*BOARD_WIDTH_CELLS+y]
			bv := boardToScreenPos(Vector{x, y})
			drawFilledRectangle(bv.x, bv.y, CELL_SIZE, CELL_SIZE, blue(color))
		}
	}
}

func (this *Board) isWallClear(wall Wall) bool {

	for _, p := range wallShapes[wall.wType] {
		wx := wall.position.x + p[0]
		wy := wall.position.y + p[1]
		if this.wallCells[wx*BOARD_WIDTH_CELLS+wy] != 0 {
			return false
		}
	}
	return true
}

func (this *Board) dropWall(wall Wall) bool {

	if this.isWallClear(wall) {
		for _, p := range wallShapes[wall.wType] {
			wx := wall.position.x + p[0]
			wy := wall.position.y + p[1]
			this.wallCells[wx*BOARD_WIDTH_CELLS+wy] = WALL_MAX_HEALTH
		}
		return true
	}

	return false
}
