package main

import (
	vectors3d "drs_fabian/nyx/vectors"
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

func Init(renderer *sdl.Renderer, window *sdl.Window) {

	vector := vectors3d.Cartesian(100, 400, 300)
	renderer.SetDrawColor(255, 0, 0, 255)

	running := true
	for running {
		renderer.DrawLine(0, 0, int32(vector.X), int32(vector.Y))
		renderer.Present()
		sdl.Delay(1000 / 60)
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				fmt.Println("Quit")
				running = false
				break
			}
		}
	}
}

func main() {

	/* IF CANNOT CREATE WINDOW PAINC */
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	/* CREATE WINDOW */
	window, err := sdl.CreateWindow("Nyx - Galaxy Simulation", sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED, 800, 500, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, 0)
	if err != nil {
		panic(err)
	}
	defer renderer.Destroy()

	Init(renderer, window)
}
