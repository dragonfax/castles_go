package main

import "time"

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

type Enemy struct {
	position   WindowPos
	direction  float32
	enemySet   EnemySet
	stopMoving bool
	target     BoardPos
	board      *Board
}

func NewEnemy(enemySet EnemySet, board *Board) *Enemy {
	this := new(Enemy)

	this.moveToEdgeOfMap()

	go this.moveLoop()
	this.enemySet = enemySet
	this.enemySet[this] = true
	this.board = board
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
	drawFilledRectangle(this.position, ENEMY_SIZE, ENEMY_SIZE, red())
}

func (this *Enemy) move() {

	if this.board.get(this.target) == 0 {
		// choose a new target
		this.target = this.board.getRandomWallPosition()
	}

	nextStep := WindowPos{}

	targetWindowPos := boardToWindowPos(this.target)

	if targetWindowPos.x > this.position.x {
		nextStep.x = this.position.x + ENEMY_SIZE
	} else {
		nextStep.x = this.position.x - ENEMY_SIZE
	}

	if targetWindowPos.y > this.position.y {
		nextStep.y = this.position.y + ENEMY_SIZE
	} else {
		nextStep.y = this.position.y - ENEMY_SIZE
	}

	bp := windowToBoardPos(nextStep)

	if this.board.get(bp) != 0 {
		// wall in the way, eat it.
		this.board.eat(bp)
	} else {
		// empty space, move there
		this.position = nextStep
	}

}

func (this *Enemy) wallkill() {
	this.stopMoving = true
	delete(this.enemySet, this)
}
