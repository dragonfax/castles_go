package main

import (
	"math/rand"

	"github.com/veandco/go-sdl2/sdl"
)

var peice *Peice

type PeiceType int

const (
	LL PeiceType = iota
	RL
	Dot
	Straight
	RZig
	LZig
	T
)
const NUM_PEICETYPES = int(T) + 1

var peiceShapes = [][][]int{

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

type PeiceDroppedEvent struct{}

type Peice struct {
	position   BoardPos
	wType      PeiceType
	stopMoving bool
}

const WALL_MAX_HEALTH = 100

func NewPeice() {
	wtype := PeiceType(rand.Int() % NUM_PEICETYPES)
	peice = &Peice{BoardPos{}, wtype, false}
	go peice.eventLoop()
}

func (this *Peice) draw() {
	for _, p := range peiceShapes[this.wType] {
		bx := this.position.x + p[0]
		by := this.position.y + p[1]
		wv := BoardPos{bx, by}.toWindowUpLeft()
		drawFilledRectangle(wv, CELL_SIZE, CELL_SIZE, blue(255))
		drawRectangle(wv, CELL_SIZE, CELL_SIZE, black())
	}
}

func (this *Peice) drop() {

	if this.isClear() {
		for _, p := range this.enumerateWallPositions() {
			board.set(p, WALL_MAX_HEALTH)
		}
		eventSendC <- PeiceDroppedEvent{}
		this.close()
	}
}

func (this *Peice) close() {
	this.stopMoving = true
}

func (this *Peice) enumerateWallPositions() []BoardPos {

	peicePoints := peiceShapes[this.wType]
	positions := make([]BoardPos, len(peicePoints))
	for i, p := range peicePoints {
		positions[i] = BoardPos{this.position.x + p[0], this.position.y + p[1]}
	}

	return positions
}

func (this *Peice) isClear() bool {

	for _, p := range peiceShapes[this.wType] {
		wx := this.position.x + p[0]
		wy := this.position.y + p[1]
		w := BoardPos{wx, wy}
		if board.get(w) != 0 {
			return false
		}
	}
	return true
}

func (this *Peice) eventLoop() {
	eventC := GetEventReceiver()
	for !this.stopMoving {
		select {
		case event := <-eventC:
			switch e := event.(type) {
			case *sdl.MouseMotionEvent:
				this.whenMouseMoves(e)
			case *sdl.MouseButtonEvent:
				this.whenMousePressed(e)
			}
		}
	}
	CloseEventReceiver(eventC)
}

func (this *Peice) whenMouseMoves(event *sdl.MouseMotionEvent) {
	this.position = WindowPos{int(event.X), int(event.Y)}.toBoard()
}

func (this *Peice) whenMousePressed(event *sdl.MouseButtonEvent) {
	this.drop()
}
