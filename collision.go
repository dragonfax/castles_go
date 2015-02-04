package main

/* Simple bounding box collision */

type Bounds struct {
	UpLeft   WindowPos
	LowRight WindowPos
}

func boundingSquare(centerPosition WindowPos, size int) Bounds {
	return Bounds{WindowPos{centerPosition.x - size/2, centerPosition.y - size/2}, WindowPos{centerPosition.x + size/2, centerPosition.y + size/2}}
}

func (tb Bounds) collidesWith(eb Bounds) bool {

	if tb.LowRight.x < eb.UpLeft.x || tb.UpLeft.x > eb.LowRight.x {
		return false
	}

	if tb.LowRight.y < eb.UpLeft.y || tb.UpLeft.y > eb.LowRight.y {
		return false
	}

	return true

}
