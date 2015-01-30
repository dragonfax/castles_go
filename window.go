package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

const CELL_SIZE = 20 // pixels across a square Cell

var renderer *sdl.Renderer
var window *sdl.Window

func windowToBoardPos(v WindowPos) BoardPos {
	return BoardPos{v.x / CELL_SIZE, v.y / CELL_SIZE}
}

func boardToWindowPos(sv BoardPos) WindowPos {
	return WindowPos{sv.x * CELL_SIZE, sv.y * CELL_SIZE}
}

func initWindow() {

	if sdl.Init(sdl.INIT_VIDEO) < 0 {
		panic(fmt.Sprintf(" SDLInitFailedException( Unable to initialize SDL: %v", sdl.GetError()))
	}

	var err error
	window, renderer, err = sdl.CreateWindowAndRenderer(640, 480, 0)
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
			switch event.(type) {
			case *sdl.QuitEvent:
				done = true
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

func drawFilledRectangle(v WindowPos, width, height int, color [4]uint8) {
	err := boxRGBA(renderer, v.x, v.y, v.x+width, v.y+height, color[0], color[1], color[2], color[3])
	if err != nil {
		panic(err)
	}
}

func drawRectangle(v WindowPos, width, height int, color [4]uint8) {
	err := rectangleRGBA(renderer, v.x, v.y, v.x+width, v.y+height, color[0], color[1], color[2], color[3])
	if err != nil {
		panic(err)
	}
}

func blue(r uint8) [4]uint8 {
	return [4]uint8{0, 0, r + 100, 255}
}

func black() [4]uint8 {
	return [4]uint8{0, 0, 0, 255}
}
