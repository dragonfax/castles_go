package main

import (
	"math/rand"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

type GameOverEvent struct{}

type Game struct {
	currentWall Wall
	drawTicker  *time.Ticker
	enemyTicker *time.Ticker
}

func NewGame() *Game {
	this := new(Game)
	this.drawTicker = time.NewTicker(time.Second / 60)
	this.enemyTicker = time.NewTicker(time.Second)
	NewEnemySet()
	NewBoard()
	this.pickRandomWall()
	return this
}

func (this *Game) run() {
	go this.inputLoop()
	go this.generateEnemyLoop()
	go this.drawLoop()
}

func (this *Game) draw() {
	clearWindow()
	board.draw()
	enemySet.draw()
	this.currentWall.draw()
	flipWindow()
}

func (this *Game) drawLoop() {
	waitC := make(WaitChannel)
	for {
		QueueMain(this.draw, waitC)
		<-waitC
		<-this.drawTicker.C
	}
}

func (this *Game) generateEnemy() {
	NewEnemy()
}

func (this *Game) generateEnemyLoop() {
	for {
		this.generateEnemy()
		<-this.enemyTicker.C
	}
}

func (this *Game) inputLoop() {
	eventC := GetEventReceiver()
	for {
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
}

func (this *Game) pickRandomWall() {
	wtype := WallType(rand.Int() % NUM_WALLTYPES)
	this.currentWall = Wall{this.currentWall.position, wtype}
}

func (this *Game) whenMouseMoves(event *sdl.MouseMotionEvent) {
	this.currentWall.position = WindowPos{int(event.X), int(event.Y)}.toBoard()
}

func (this *Game) whenMousePressed(event *sdl.MouseButtonEvent) {
	if board.dropWall(this.currentWall) {
		enemySet.wallKills(this.currentWall)
		this.pickRandomWall()
	}
}
