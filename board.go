package main

const BOARD_WIDTH_CELLS = 100
const BOARD_HEIGHT_CELLS = 80
const BOARD_NUM_CELLS = BOARD_WIDTH_CELLS * BOARD_HEIGHT_CELLS

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

type Direction int

const (
	Up Direction = iota
	Right
	Left
	Down
)

type Board struct {
	wallCells [BOARD_NUM_CELLS]int // [x*y] = health of cell
}

func NewBoard() *Board {
	this := new(Board)
}

func (this *Board) dropWall(wall Wall) bool {
	is the space clear of walls
		then drop it
		does it collide with any enemies
			then wall kill them.
		clear out the currentWall
		pick a new random wall
}


type Wall struct {
	position    Vector
	orientation Direction
	wType       WallType
	health      int
}


