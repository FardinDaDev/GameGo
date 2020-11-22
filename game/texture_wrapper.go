package game

import (
	"fmt"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type TextureWrapper struct {
	texture *sdl.Texture
	width int32
	height int32
}

func LoadFromFile(renderer *sdl.Renderer, path string) (t *TextureWrapper) {
	loadedSurface, err := img.Load(path)

	if err != nil {
		fmt.Printf("Unable to load image %s! SDL_image Error: %s\n", path, err)
	}

	loadedSurface.SetColorKey(true, sdl.MapRGB(loadedSurface.Format, 0, 0xFF, 0xFF))

	newTexture, err := renderer.CreateTextureFromSurface(loadedSurface)
	if err != nil {
		fmt.Println("Error cannot create texture from surface:", err)
	}

	t = &TextureWrapper{
		texture: newTexture,
		height:  loadedSurface.H,
		width:   loadedSurface.W,
	}

	loadedSurface.Free()

	return t
}

func (t *TextureWrapper) Destroy() {
	t.texture.Destroy()
	t.width = 0
	t.height = 0
}

func (t *TextureWrapper) Render(renderer *sdl.Renderer, x, y int32) {
	dst := sdl.Rect{X: x, Y: y, W: t.width, H: t.height}
	renderer.Copy(t.texture, nil, &dst)
}