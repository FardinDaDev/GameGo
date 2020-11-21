package game

import (
	"fmt"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"os"
)

type Window struct {
	Window   *sdl.Window
	surface  *sdl.Surface
	Renderer *sdl.Renderer
	texture  *sdl.Texture
	event    sdl.Event
	err      error
	Running  bool
}

func (w *Window) Init(title string, x int32, y int32, width int32, height int32, fullscreen bool) {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to initialise: %s\n", w.err)
		os.Exit(1)
	}

	var flags uint32 = sdl.WINDOW_SHOWN

	if fullscreen {
		flags = sdl.WINDOW_FULLSCREEN_DESKTOP
	}

	w.Window, w.err = sdl.CreateWindow(title, x, y,
		width, height, flags)

	if w.err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create window: %s\n", w.err)
		os.Exit(2)
	}

	w.Renderer, w.err = sdl.CreateRenderer(w.Window, -1, sdl.RENDERER_ACCELERATED)

	if w.err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create renderer: %s\n", w.err)
		os.Exit(2)
	}

	w.Running = true
}

func (w *Window) LoadTexture(path string) *sdl.Texture {
	img.Init(img.INIT_JPG | img.INIT_PNG)
	sdl.SetHint(sdl.HINT_RENDER_SCALE_QUALITY, "1")

	surfaceImg, err := img.Load(path)

	if err != nil {
		fmt.Printf("Unable to load the img: %s\n", err)
	}

	textureImg, err := w.Renderer.CreateTextureFromSurface(surfaceImg)

	if err != nil {
		fmt.Printf("Unable to load the texture: %s\n", err)
	}

	defer surfaceImg.Free()

	return textureImg
}


func (w *Window) Render() {
	w.Renderer.SetDrawColor(255, 255, 55, 255)
	w.Renderer.Clear()

	w.Renderer.Present()
}

func (w *Window) EventHandler() {
	for w.event = sdl.PollEvent(); w.event != nil; w.event = sdl.PollEvent() {
		switch t := w.event.(type) {
		case *sdl.QuitEvent:
			fmt.Println("Quit...")
			w.Running = false
		case *sdl.KeyboardEvent:
			if t.Keysym.Sym == sdl.K_ESCAPE {
				fmt.Println("Quit...")
				w.Running = false
			}
		}
	}
}

func (w *Window) Update() {
	//
}

func (w *Window) Clear() {
	w.Renderer.Destroy()
	w.Window.Destroy()
	sdl.Quit()
}
