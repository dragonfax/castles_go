package main

import (
	"math/rand"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

type Game struct {
	currentWall Wall
	board       Board
	enemySet    EnemySet
	drawTicker  *time.Ticker
	enemyTicker *time.Ticker
}

func NewGame() *Game {
	this = new(Game)
	this.drawTicker = NewTicker()
	this.enemyTicker = time.Ticker(time.Second)
	this.enemySet = make(EnemySet)
	this.board = NewBoard()
	this.currentWall = this.pickRandomWall()
	return this
}

func (this *Game) run() {
	go this.inputLoop()
	go this.generateEnemyLoop()
	go this.drawLoop()
}

func (this *Game) draw() {
	windowClear()
	this.board.draw()
	this.enemySet.draw()
	this.currentWall.draw()
	windowFlip()
}

func (this *Game) drawLoop() {
	waitC := make(WaitChannel)
	for {
		QueueMain(this.draw, waitC)
		<-waitC
		<-this.drawTicker
	}
}

func (this *Game) generateEnemy() {
	e := NewEnemy(this.enemySet)
}

func (this *Game) generateEnemyLoop() {
	for {
		this.generateEnemy()
		<-enemyTicker
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
	this.currentWall = NewWall(wtype, this.currentWall.position)
}

func (this *Game) whenMouseMoves(event *sdl.MouseMotionEvent) {
	/*
		move wall position (on screen)
	*/
}

func (this *Game) whenMousePressed(event *sdl.MouseButtonEvent) {
	if this.board.dropWall(this.currentWall) {
		this.enemySet.wallKills(this.currentWall)
		this.pickRandomWallt()
	}
}
