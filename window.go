
import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

const CELL_SIZE = 20 // pixels across a square Cell

var renderer *sdl.Renderer

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

