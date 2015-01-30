package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

const CELL_SIZE = 20 // pixels across a square Cell

var renderer *sdl.Renderer

func screenToBoardPos(v Vector) Vector {
	return Vector{v.x / CELL_SIZE, v.y / CELL_SIZE}
}

func boardToScreenPos(sv Vector) Vector {
	return Vector{sv.x * CELL_SIZE, sv.y * CELL_SIZE}
}

func initWindow() {

	if sdl.Init(sdl.INIT_VIDEO) < 0 {
		panic(fmt.Sprintf(" SDLInitFailedException( Unable to initialize SDL: %v", sdl.GetError()))
	}

	window, renderer, err := sdl.CreateWindowAndRenderer(640, 480, nil)
	window.SetTitle("Castles")
	if err != nil {
		panic(fmt.Sprintf("SDLInitFailedException (Unable to create SDL screen: %v", sdl.GetError()))
	}

	renderer.SetDrawColor(0, 0, 0, 255)

	go windowEventsLoop()
}

func clearWindow() {
	renderer.Clear()
}

func flipWindow() {
	renderer.Present()
}

func windowEventsLoop() {
	eventReceiver := GetEventReceiver()
	for {
		select {
		case event := <-eventReceiver:
			switch e := event.(type) {
			case *sdl.QuitEvent:
				m.done = true
				/*case *sdl.WindowEvent:
				switch e.Event {
				case sdl.WINDOWEVENT_RESIZED:
					w := e.Data1
					h := e.Data2
					if w > 150 && h > 100 {
						QueueMain(func() {
							screen.resized(int(w), int(h))
						}, nil)
					}
				}*/
			}
		}
	}
}

func drawFilledRectangle(x, y, width, height int, color [4]uint8) {
	boxRGBA(renderer, x, y, x+width, y+height, color[0], color[1], color[2], color[3])
}

func drawRectangle(x, y, width, height int, color [4]uint8) {
	rectangleRGBA(renderer, x, y, x+width, y+height, color[0], color[1], color[2], color[3])
}

func blue(r int) [4]uint8 {
	return [4]uint8{0, 0, r + 100, 255}
}

func black() [4]uint8 {
	return [4]uint8{0, 0, 0, 255}
}
