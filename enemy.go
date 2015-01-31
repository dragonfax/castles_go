package main

import (
	"math"
	"math/rand"
	"time"
)

const ENEMY_SIZE = CELL_SIZE / 5

type EnemySet map[*Enemy]bool

func (this EnemySet) wallKills(wall Wall) {
	for e, _ := range this {
		bpos := e.position.toBoard()
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
	this.eatTimer = NewEatTimer(time.Second / 5)
	return this
}

func (this *Enemy) moveToRandomEdgeOfMap() {

	pos := WindowPos{rand.Intn(640), rand.Intn(480)}

	// which edge?
	switch rand.Intn(4) {
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
	moveTicker := time.NewTicker(time.Second / 10)
	for !this.stopMoving {
		this.move()
		<-moveTicker.C
	}
	delete(this.enemySet, this)
}

func (this *Enemy) draw() {
	drawFilledRectangle(this.position, ENEMY_SIZE, ENEMY_SIZE, red())
}

var EAT_RANGE = math.Sqrt(((CELL_SIZE / 2) ^ 2) * 2)

func (this *Enemy) inEatRange(wallPos BoardPos) bool {
	// if our boundries (plus a little) collide with those boundaries

	return this.position.dist(wallPos.toWindowCenter()) < EAT_RANGE
}

func (this *Enemy) moveTowards(wallPos WindowPos) {

	newPos := this.position

	if wallPos.x < this.position.x {
		newPos.x -= 1
	} else if wallPos.x > this.position.x {
		newPos.x += 1
	}

	if wallPos.y < this.position.y {
		newPos.y -= 1
	} else if wallPos.y > this.position.y {
		newPos.y += 1
	}

	if !this.checkCollisions(newPos) {
		this.position = newPos
	}
}

func (this *Enemy) collidesWith(e *Enemy) bool {

	tb := this.bounds()
	eb := e.bounds()

	if tb.LowRight.x < eb.UpLeft.x || tb.UpLeft.x > eb.LowRight.x {
		return false
	}

	if tb.LowRight.y < eb.UpLeft.y || tb.UpLeft.y > eb.LowRight.y {
		return false
	}

	return true

}

func (this *Enemy) bounds() Bounds {
	upLeftCorner := this.position
	lowRightCorner := WindowPos{this.position.x + ENEMY_SIZE, this.position.y + ENEMY_SIZE}
	return Bounds{upLeftCorner, lowRightCorner}
}

func (this *Enemy) checkCollisions(pos WindowPos) bool {

	b := this.bounds()

	// check against board
	if this.board.get(b.UpLeft.toBoard()) != 0 {
		return true
	}
	if this.board.get(b.LowRight.toBoard()) != 0 {
		return true
	}

	// check against other enemies
	for e, _ := range this.enemySet {
		if e != this {
			if this.collidesWith(e) {
				return true
			}
		}
	}

	return false
}

func (this *Enemy) move() {

	wallPos := this.board.nearestWallPos(this.position.toBoard())
	if this.inEatRange(wallPos) {
		if this.eatTimer.timeToEat() {
			this.eatTimer.eating()
			this.board.eat(wallPos)
		}
	} else {
		this.moveTowards(wallPos.toWindowCenter())
	}

}

func (this *Enemy) wallkill() {
	this.stopMoving = true
	delete(this.enemySet, this)
}
