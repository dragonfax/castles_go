package main

import "fmt"

const BOARD_WIDTH_CELLS = 40
const BOARD_HEIGHT_CELLS = 20
const BOARD_NUM_CELLS = BOARD_WIDTH_CELLS * BOARD_HEIGHT_CELLS

type Board struct {
	wallCells [BOARD_NUM_CELLS]uint8 // [x*y] = health of cell
	castle    *Castle
}

var board Board

func NewBoard() {
	board = Board{}
	board.castle = NewCastle()
	go board.eventLoop()
}

func (this *Board) get(p BoardPos) uint8 {
	if p.x >= BOARD_WIDTH_CELLS {
		panic(fmt.Sprintf("requested an x coordinate too wide %d", p.x))
	}
	if p.y >= BOARD_HEIGHT_CELLS {
		panic(fmt.Sprintf("requested a y coordinate too high %d", p.y))
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
	if health > 0 {
		if health >= 5 {
			health = health - 5
		} else {
			health = 0
		}
		this.set(p, health)
	}
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
	this.castle.draw()
}

func (this *Board) nearestWallPos(enemyPos BoardPos) BoardPos {
	lowestDist := 9999.0
	closestWallPos := BoardPos{}
	for x := 0; x < BOARD_WIDTH_CELLS; x++ {
		for y := 0; y < BOARD_HEIGHT_CELLS; y++ {
			bp := BoardPos{x, y}
			if this.get(bp) != 0 {
				dist := bp.dist(enemyPos)
				if dist < lowestDist {
					lowestDist = dist
					closestWallPos = bp
				}
			}
		}
	}
	return closestWallPos
}

func (this *Board) eventLoop() {
	eventC := GetEventReceiver()
	for {
		select {
		case event := <-eventC:
			switch event.(type) {
			case PeiceDroppedEvent:
				NewPeice()
			}
		}
	}
}
