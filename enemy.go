package main

import "time"

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

type Enemy struct {
	position   WindowPos
	direction  float32
	enemySet   EnemySet
	stopMoving bool
	target     BoardPos
	board      *Board
}

func NewEnemy(enemySet EnemySet) *Enemy {
	this := new(Enemy)

	this.moveToEdgeOfMap()

	go this.moveLoop()
	this.enemySet = enemySet
	this.enemySet[this] = true
	return this
}

func (this *Enemy) moveToEdgeOfMap() {
	this.position = WindowPos{0, 0}
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
}

func (this *Enemy) move() {
	/*
		verify there is a wall in direction
		if so
			move towards wall
		if not
			find a new random direction towards a wall.
		if next to wall
			eat wall
		if not moving and not eating a wall and not next to a wall
			choose a new diretion towards a wall.
	*/
}

func (this *Enemy) wallkill() {
	this.stopMoving = true
	delete(this.enemySet, this)
}
