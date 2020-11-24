package game

import (
	"fmt"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
	"os"
)

//Global Variables (messy)
var (
	degrees      float64          = 0
	flipType     sdl.RendererFlip = sdl.FLIP_NONE
	arrowTexture *TextureWrapper
	font         *ttf.Font
	textTexture  *TextureWrapper

	gButton [TOTAL_BUTTONS]MouseButton
	gButtonSpriteSheetTexture *TextureWrapper
	BUTTON_WIDTH, BUTTON_HEIGHT int32 = 300, 200
	gSpriteClips [TOTAL_BUTTONS]sdl.Rect
)

//Constants (Global Variable)
const (
	TOTAL_BUTTONS int = 4
	SCREEN_WIDTH = 1920 //TODO: Removing SCREEN_WIDTH & SCREEN_HEIGHT...
	SCREEN_HEIGHT = 1080
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

func (w *Window) NewWindow(title string, x int32, y int32, width int32, height int32, fullscreen bool) {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to initialise SDL2: %s\n", w.err)
		os.Exit(1)
	}

	if err := ttf.Init(); err != nil {
		fmt.Printf("Failed to initialize TTF: %s\n", err)
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

	w.Renderer, w.err = sdl.CreateRenderer(w.Window, -1, sdl.RENDERER_ACCELERATED | sdl.RENDERER_PRESENTVSYNC)

	if w.err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create renderer: %s\n", w.err)
		os.Exit(2)
	}

	if err := img.Init(img.INIT_JPG | img.INIT_PNG); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load img: %s\n", w.err)
		os.Exit(2)
	}
	sdl.SetHint(sdl.HINT_RENDER_SCALE_QUALITY, "1")

	w.Running = true
}

func (w *Window) LoadTexture(path string) *sdl.Texture {
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

func (w *Window) LoadMedia() (err error) {
	//arrowTexture = LoadFromFile(w.Renderer, "./img/preview.png")

	//font, err := ttf.OpenFont("./fonts/Roboto-Medium.ttf", 28)
	//
	//if err != nil {
	//	fmt.Println("Couldn't load the font...", err)
	//}
	//
	//textTexture = Init()
	//textColor := sdl.Color{R: 0, G: 0, B: 0}
	//textTexture.LoadFromRenderedText(w.Renderer, font, "Swaggerboi69", textColor)


	//gButtonSpriteSheetTexture = NewTextureWrapper().LoadFromFile(w.Renderer, "./img/button.png")

	gButtonSpriteSheetTexture = NewTextureWrapper().LoadFromFile(w.Renderer, "./img/button.png")

	for i := 0; i < BUTTON_SPRITE_TOTAL; i++ {
		gSpriteClips[i] = sdl.Rect{X: 0, Y: int32(i * 200), W: BUTTON_WIDTH, H: BUTTON_HEIGHT}
	}

	gButton[0] = NewMouseButton(0, 0)
	gButton[1] = NewMouseButton(SCREEN_WIDTH - BUTTON_WIDTH, 0)
	gButton[2] = NewMouseButton(0, SCREEN_HEIGHT - BUTTON_HEIGHT)
	gButton[3] = NewMouseButton(SCREEN_WIDTH - BUTTON_WIDTH, SCREEN_HEIGHT - BUTTON_HEIGHT)

	fmt.Println(gButton)

	return nil
}

func (w *Window) Render() {
	w.Renderer.SetDrawColor(0xFF, 0xFF, 0xFF, 0xFF)
	w.Renderer.Clear()

	for i := 0; i < TOTAL_BUTTONS; i++ {
		gButton[i].Render(w.Renderer, gButtonSpriteSheetTexture)
	}

	//arrowTexture.RenderEx(w.Renderer, (screenWidth - arrowTexture.width) / 2, (screenHeight - arrowTexture.height) / 2, nil, degrees, nil, flipType)

	//if degrees != 0 {
	//	textColor := sdl.Color{R: 0, G: 0, B: 0}
	//	str := fmt.Sprintf("the row degress is %f", degrees)
	//	textTexture.LoadFromRenderedText(w.Renderer, font, str, textColor)
	//}
	//
	//textTexture.Render(w.Renderer, 0, 0, nil)

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

		case *sdl.MouseMotionEvent:
			for i := 0; i < TOTAL_BUTTONS; i++ {
				gButton[i].HandleEvent(t.X, t.Y, t.Type)
			}
			break
		case *sdl.MouseButtonEvent:
			for i := 0; i < TOTAL_BUTTONS; i++ {
				gButton[i].HandleEvent(t.X, t.Y, t.Type)
			}
			break
		}
	}
}

func (w *Window) Update() {
	//fmt.Println("Update...")

}

func (w *Window) Clear() {
	//arrowTexture.Destroy()
	gButtonSpriteSheetTexture.Destroy()
	//textTexture.Destroy()

	w.Renderer.Destroy()
	w.Window.Destroy()

	ttf.Quit()
	img.Quit()
	sdl.Quit()
}
