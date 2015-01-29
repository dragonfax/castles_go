package main

type WallType int

const (
	LL WallType = iota
	RL
	Dot
	Straight
	RZig
	LZig
	T
)
const NUM_WALLTYPES = int(T) + 1

type Direction int

const (
	Up Direction = iota
	Right
	Left
	Down
)

const BOARD_WIDTH_CELLS = 100
const BOARD_HEIGHT_CELLS = 80
const BOARD_NUM_CELLS = BOARD_WIDTH_CELLS * BOARD_HEIGHT_CELLS

type Board struct {
	wallCells [BOARD_NUM_CELLS]int
}

func NewBoard() *Board {
	this := new(Board)
}

func (this *Board) dropWallAt(x, y float32, wall Wall) {
	is the space clear of walls
		then drop it
		does it collide with any enemies
			then wall kill them.
		clear out the currentWall
		pick a new random wall
}


type Wall struct {
	position    Vector
	orientation Direction
	wType       WallType
	health      int
}

var enemyList = make(map[*Enemy]bool)

type Enemy struct {
	position  Vector
	direction float32
}

func NewEnemy() *Enemy {
	this := new(Enemy)
	choose a location along map edge
	go this.moveLoop()
	return this
}

func (this *Enemy) moveLoop() {
	for {
		move()
		wait on enemy ticker
		if enemy killed
			stop moving and exit goroutine
			remove fro all structures
				especially drawing.
	}
}

func (this *Enemy) move() {
	verify there is a wall in direction
	if so
		move towards wall
	if not
		find a new random direction towards a wall.
	if next to wall
		eat wall
	if not moving and not eating a wall and not next to a wall
		choose a new diretion towards a wall.
}

func (this *Enemy) wallkill() {
	destroy this enemy
}



func mainThreadLoop() {
	for {
		handle main thread queue events
		read events from devices
	}
}

type Game struct {
	currentWall Wall
	board Board
}

func NewGame() *Game {
	this = new(Game)
	return this
}

func (this *Game) setup() {
	init screen and devices
	this.board = NewBoard()
	this.currentWall = this.pickRandomWall()
}

func (this *Game) run() {
	this.setup()
	go handleEventsLoop()
	go this.generateEnemyLoop()
	go this.drawLoop()
}

func (this *Game) draw() {
	draw board
	draw all emenites
	draw currentWall
}

func (this *Game) drawLoop() {
	for {
		this.draw()
		wait on render fps limiter
	}
}

func (this *Game) generateEnemy() {
	NewEnemy()
}

func (this *Game) generateEnemyLoop() {
	for {
		this.generateEnemy
		wait on enemy generation ticker
	}
}

func (this *Game) handleEvents() {
	read events
	track mouse
	draw current wall at normalized mouse position
	if mouse cliekd
		this.board.drawWallAt(mouse location)
}

func (this *Game) handleEventsLoop() {
	for {
		this.handleEvents()
		wait for input ticker
	}
}

func (this *Game) pickRandomWall() {
	wtype := WallType(rand.Int() % NUM_WALLTYPES)
	this.currentWall = NewWall(wtype)
}


func main() {
	go NewGame().run()
	mainThreadLoop()
}
