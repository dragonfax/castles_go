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
	position BoardPos
	wType    WallType
}

func (this *Wall) draw() {
	for _, p := range wallShapes[this.wType] {
		bx := this.position.x + p[0]
		by := this.position.y + p[1]
		wv := BoardPos{bx, by}.toWindowUpLeft()
		drawFilledRectangle(wv, CELL_SIZE, CELL_SIZE, blue(255))
		drawRectangle(wv, CELL_SIZE, CELL_SIZE, black())
	}
}

type Board struct {
	wallCells [BOARD_NUM_CELLS]uint8 // [x*y] = health of cell
}

func NewBoard() *Board {
	this := new(Board)
	return this
}

func (this *Board) get(p BoardPos) uint8 {
	if p.x >= BOARD_WIDTH_CELLS {
		panic("requested a coorinate too wide")
	}
	if p.y >= BOARD_HEIGHT_CELLS {
		panic("requested a coorinate too high")
	}
	return this.wallCells[p.x*BOARD_HEIGHT_CELLS+p.y]
}

func (this *Board) set(b BoardPos, value uint8) {
	if b.x >= BOARD_WIDTH_CELLS {
		panic("requested a coorinate too wide")
	}
	if b.y >= BOARD_HEIGHT_CELLS {
		panic("requested a coorinate too high")
	}
	this.wallCells[b.x*BOARD_HEIGHT_CELLS+b.y] = value
}

func (this *Board) eat(p BoardPos) {
	health := this.get(p)
	if health >= 5 {
		health = health - 5
	} else {
		health = 0
	}
	this.set(p, health)
}

func (this *Board) getRandomWallPosition() BoardPos {
	p := BoardPos{}
	for x := 0; x < BOARD_WIDTH_CELLS; x++ {
		for y := 0; y < BOARD_HEIGHT_CELLS; y++ {
			p.x = x
			p.y = y
			if this.get(p) != 0 {
				return p
			}
		}
	}
	return p
}

func (this *Board) draw() {
	p := BoardPos{}
	for x := 0; x < BOARD_WIDTH_CELLS; x++ {
		for y := 0; y < BOARD_HEIGHT_CELLS; y++ {
			p.x = x
			p.y = y
			health := this.get(p)
			wv := p.toWindowUpLeft()
			drawFilledRectangle(wv, CELL_SIZE, CELL_SIZE, blue(uint8(float32(health)*(255/WALL_MAX_HEALTH))))
		}
	}
}

func (this *Board) isWallClear(wall Wall) bool {

	for _, p := range wallShapes[wall.wType] {
		wx := wall.position.x + p[0]
		wy := wall.position.y + p[1]
		w := BoardPos{wx, wy}
		if this.get(w) != 0 {
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
			w := BoardPos{wx, wy}
			this.set(w, WALL_MAX_HEALTH)
		}
		return true
	}

	return false
}

func (this *Board) nearestWallPos(enemyPos BoardPos) BoardPos {
	lowestDist := 9999
	closestWallPos := Vector{BOARD_WIDTH_CELLS / 2, BOARD_HEIGHT_CELLS / 2}
	for x := 0; x < BOARD_WIDTH_CELLS; x++ {
		for y := 0; y < BOARD_HEIGHT_CELLS; y++ {
			bp := BoardPos{x, y}.dist(enemyPos)
			if this.get(pb) != 0 {
				dist := bp.dist(enemyPos)
				if dist < lowestDist {
					lowestDist = dist
					closestWallPos = bp
				}
			}
		}
	}
	return bp
}
