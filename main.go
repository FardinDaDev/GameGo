package main

import (
	"github.com/FardinDaDev/GoGame/game"
	"github.com/FardinDaDev/GoGame/logger"
	"github.com/veandco/go-sdl2/sdl"
)

func main()  {
	var w game.Window

	w.NewWindow("Test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 1920, 1080, true)

	if w.LoadMedia() != nil {
		logger.GameDir.Println("Cannot load media...")
	}

	for w.Running {
		w.EventHandler()
		w.Render()
		w.Update()
	}

	w.Clear()
}