package main

import (
	"math/rand"
	"time"
)

const ENEMY_SIZE = CELL_SIZE / 5

type EnemySet map[*Enemy]bool

var enemySet EnemySet

func NewEnemySet() {
	enemySet = make(EnemySet)
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
	stopMoving bool
	eatTimer   *EatTimer
}

func NewEnemy() *Enemy {
	this := new(Enemy)

	this.moveToRandomEdgeOfMap()

	go this.moveLoop()
	go this.eventLoop()
	enemySet[this] = true
	this.eatTimer = NewEatTimer(time.Second / 5)
	return this
}

func (this *Enemy) moveToRandomEdgeOfMap() {

	pos := WindowPos{rand.Intn(WINDOW_WIDTH - 1), rand.Intn(WINDOW_HEIGHT - 1)}

	// which edge?
	switch rand.Intn(4) {
	case 0:
		pos.x = 0
	case 1:
		pos.x = WINDOW_WIDTH - 1
	case 2:
		pos.y = 0
	case 3:
		pos.y = WINDOW_HEIGHT - 1
	default:
		panic("my math is bad")
	}

	this.position = pos
}

func (this *Enemy) moveLoop() {
	moveTicker := time.NewTicker(time.Second / 20)
	for !this.stopMoving {
		this.move()
		<-moveTicker.C
	}
	delete(enemySet, this)
}

func (this *Enemy) draw() {
	drawFilledRectangle(this.position, ENEMY_SIZE, ENEMY_SIZE, red())
}

func (this *Enemy) inEatRange(wallPos BoardPos) bool {
	// if our boundries (plus a little) collide with those boundaries

	return this.position.dist(wallPos.toWindowCenter()) < CELL_SIZE
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

func (this *Enemy) bounds() Bounds {
	return boundingSquare(this.position, ENEMY_SIZE)
}

func (this *Enemy) checkCollisions(pos WindowPos) bool {

	// check against board
	if board.get(this.position.toBoard()) != 0 {
		return true
	}
	if board.get(this.position.toBoard()) != 0 {
		return true
	}

	// check against other enemies
	b := this.bounds()
	for e, _ := range enemySet {
		if e != this {
			eb := e.bounds()
			if b.collidesWith(eb) {
				return true
			}
		}
	}

	return false
}

func (this *Enemy) move() {

	wallPos := board.nearestWallPos(this.position.toBoard())
	castlePos := board.castle.position
	if !wallPos.isZero() && this.inEatRange(wallPos) {
		if this.eatTimer.timeToEat() {
			this.eatTimer.eating()
			board.eat(wallPos)
		}
	} else if this.inEatRange(castlePos) {
		if this.eatTimer.timeToEat() {
			this.eatTimer.eating()
			board.castle.eat()
		}
	} else {
		this.moveTowards(castlePos.toWindowCenter())
	}

}

func (this *Enemy) close() {
	this.stopMoving = true
	delete(enemySet, this)
}

func (this *Enemy) eventLoop() {
	eventC := GetEventReceiver()
	for !this.stopMoving {
		select {
		case event := <-eventC:
			switch event.(type) {
			case PeiceDroppedEvent:
				// collision and wallkill
				if board.get(this.position.toBoard()) != 0 {
					this.close()
				}
			}
		}
	}
	CloseEventReceiver(eventC)
}
