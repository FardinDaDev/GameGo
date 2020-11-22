package main

import (
	"github.com/FardinDaDev/GoGame/game"
	"github.com/veandco/go-sdl2/sdl"
)

func main()  {
	var w game.Window

	w.Init("Test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 480, 640, false)

	for w.Running {
		w.EventHandler()
		w.Render()
		w.Update()
	}

	w.Clear()
}