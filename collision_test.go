package main

import "testing"

func TestCollidesWith(t *testing.T) {

	if !boundingSquare(WindowPos{10, 10}, 2).collidesWith(boundingSquare(WindowPos{10, 10}, 2)) {
		t.Fatal("should have been true")
	}

	if boundingSquare(WindowPos{10, 10}, 2).collidesWith(boundingSquare(WindowPos{1, 1}, 2)) {
		t.Fatal("should have been false")
	}

	if !boundingSquare(WindowPos{10, 10}, 2).collidesWith(boundingSquare(WindowPos{9, 9}, 2)) {
		t.Fatal("should have been true")
	}

}
