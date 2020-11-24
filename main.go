package main

import (
	"fmt"
	"github.com/FardinDaDev/GoGame/game"
	"github.com/veandco/go-sdl2/sdl"
)

func main()  {
	var w game.Window

	w.Init("Test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 640, 480, false)

	if w.LoadMedia() != nil {
		fmt.Println("Cannot load media")
	}

	for w.Running {
		w.EventHandler()
		w.Render()
		w.Update()
	}

	w.Clear()
}