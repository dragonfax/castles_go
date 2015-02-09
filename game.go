package main

import "time"

type GameOverEvent struct{}

type Game struct {
	drawTicker  *time.Ticker
	enemyTicker *time.Ticker
}

func NewGame() *Game {
	this := new(Game)
	this.drawTicker = time.NewTicker(time.Second / 60)
	this.enemyTicker = time.NewTicker(time.Second)
	NewEnemySet()
	NewBoard()
	NewPeice()
	return this
}

func (this *Game) run() {
	go this.generateEnemyLoop()
	go this.drawLoop()
}

func (this *Game) draw() {
	clearWindow()
	board.draw()
	enemySet.draw()
	peice.draw()
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
