package main

import (
	"fmt"
	"testing"
)

func TestDist(t *testing.T) {

	fmt.Printf("%d\n", 0^2)

	dist := (Vector{0, 0}).dist(Vector{1, 0})
	if dist != 1 {
		t.Fatalf("failure dist was %f", dist)
	}

	dist = (Vector{1, 0}).dist(Vector{0, 0})
	if dist != 1 {
		t.Fatalf("failure dist was %f", dist)
	}

	dist = (Vector{0, 0}).dist(Vector{0, 0})
	if dist != 0 {
		t.Fatalf("failure dist was %f", dist)
	}

	dist = (Vector{0, 0}).dist(Vector{1, 1})
	if !(dist > 1 && dist < 2) {
		t.Fatalf("failure dist was %f", dist)
	}
}
