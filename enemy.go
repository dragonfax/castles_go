package main

import (
	"math/rand"
	"time"
)

const ENEMY_SIZE = CELL_SIZE / 5

type EnemySet map[*Enemy]bool

func (this EnemySet) wallKills(wall Wall) {
	for e, _ := range this {
		bpos := windowToBoardPos(e.position)
		if bpos == wall.position {
			e.wallkill()
		}
	}
}

func (this EnemySet) draw() {
	for e, _ := range this {
		e.draw()
	}
}

type EatTimer struct {
	when      time.Duration
	lastEaten time.Time
}

func NewEatTimer(when time.Duration) *EatTimer {
	this := new(EatTimer)
	this.lastEaten = time.Now()
	return this
}

func (this EatTimer) timeToEat() bool {
	return this.lastEaten.Add(this.when).Before(time.Now())
}

func (this *EatTimer) eating() {
	this.lastEaten = time.Now()
}

type Enemy struct {
	position   WindowPos
	direction  float32
	enemySet   EnemySet
	stopMoving bool
	board      *Board
	eatTimer   *EatTimer
}

func NewEnemy(enemySet EnemySet, board *Board) *Enemy {
	this := new(Enemy)

	this.moveToRandomEdgeOfMap()

	go this.moveLoop()
	this.enemySet = enemySet
	this.enemySet[this] = true
	this.board = board
	return this
}

func (this *Enemy) moveToRandomEdgeOfMap() {

	pos := WindowPos{rand.Intn(640), rand.Intn(480)}

	// which edge?
	switch rand.Intn(5) {
	case 0:
		pos.x = 0
	case 1:
		pos.x = 640
	case 2:
		pos.y = 0
	case 3:
		pos.y = 480
	default:
		panic("my math is bad")
	}

	this.position = pos
}

func (this *Enemy) close() {
	delete(this.enemySet, this)
}

func (this *Enemy) moveLoop() {
	moveTicker := time.NewTicker(time.Second / 5)
	for !this.stopMoving {
		this.move()
		<-moveTicker.C
	}
	delete(this.enemySet, this)
}

func (this *Enemy) draw() {
	drawFilledRectangle(this.position, ENEMY_SIZE, ENEMY_SIZE, red())
}

func (this *Enemy) inEatRange(wallPos BoardPos) {
	// if our boundries (plus a little) collide with those boundaries

	return this.eatCollision().collidsWidth(this.board.wallCollision(wallPos))
}

func (this *Enemy) move() {

	wallPos := this.board.nearestWalPos(windowToBoardPos(this.position))
	if this.inEatRange(wallPos) {
		if this.eatTimer.timeToEat() {
			this.eatTimer.eating()
			this.board.eat(wallPos)
		}
	} else {
		this.moveTowards(wallPos)
	}

}

func (this *Enemy) wallkill() {
	this.stopMoving = true
	delete(this.enemySet, this)
}
