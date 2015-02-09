package main

type Castle struct {
	health   int
	position BoardPos
}

func NewCastle() *Castle {
	this := new(Castle)
	this.health = 200
	this.position = BoardPos{BOARD_WIDTH_CELLS / 2, BOARD_HEIGHT_CELLS / 2}
	return this
}

func (this *Castle) eat() {
	if this.health > 0 {
		this.health -= 1
	}
	if this.health == 0 {
		eventSendC <- GameOverEvent{}
	}
}

func (this *Castle) draw() {
	wv := this.position.toWindowUpLeft()
	drawFilledRectangle(wv, CELL_SIZE, CELL_SIZE, gold(uint8(this.health)))
}
