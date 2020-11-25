package game

import (
	"fmt"
	"github.com/FardinDaDev/GoGame/logger"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type TextureWrapper struct {
	texture *sdl.Texture
	width   int32
	height  int32
}

func NewTextureWrapper() *TextureWrapper {
	return &TextureWrapper{
		texture: nil,
		height:  0,
		width:   0,
	}
}

func (t *TextureWrapper) LoadFromFile(renderer *sdl.Renderer, path string) ( *TextureWrapper) {
	loadedSurface, err := img.Load(path)

	if err != nil {
		fmt.Printf("Unable to load image %s! SDL_image Error: %s\n", path, err)
	}

	loadedSurface.SetColorKey(true, sdl.MapRGB(loadedSurface.Format, 0, 0xFF, 0xFF))

	newTexture, err := renderer.CreateTextureFromSurface(loadedSurface)
	if err != nil {
		logger.GameDir.Println("Error cannot create texture from surface:", err)
	}

	t = &TextureWrapper{
		texture: newTexture,
		height: loadedSurface.H,
		width: loadedSurface.W,
	}

	defer loadedSurface.Free()

	return t
}

func (t *TextureWrapper) SetBlendMode(blending sdl.BlendMode) {
	t.texture.SetBlendMode(blending)
}

func (t *TextureWrapper) SetAlpha(alpha uint8) {
	t.texture.SetAlphaMod(alpha)
}

func (t *TextureWrapper) SetColor(red, green, blue uint8) {
	t.texture.SetColorMod(red, green, blue)
}


func (t *TextureWrapper) LoadFromRenderedText(renderer *sdl.Renderer, font *ttf.Font, textureText string, textureColor sdl.Color) *TextureWrapper {

	textSurface, err := font.RenderUTF8Solid(textureText, textureColor)

	if err != nil {
		logger.GameDir.Println("Couldn't render SDL ttf text", err)
	}

	newTexture, err := renderer.CreateTextureFromSurface(textSurface)

	if err != nil {
		logger.GameDir.Println("Couldn't renderer text from surface", err)
	}

	t = &TextureWrapper{
		texture: newTexture,
		width: textSurface.W,
		height: textSurface.H,
	}

	defer textSurface.Free()

	return t
}

func (t *TextureWrapper) Destroy() {
	t.texture.Destroy()
	t.width = 0
	t.height = 0
}

func (t *TextureWrapper) Render(renderer *sdl.Renderer, x, y int32, clip *sdl.Rect) {
	dst := sdl.Rect{X: x, Y: y, W: t.width, H: t.height}

	if !clip.Empty() {
		dst.W = clip.W
		dst.H = clip.H
	}

	renderer.Copy(t.texture, clip, &dst)
}

func (t *TextureWrapper) RenderEx(renderer *sdl.Renderer, x, y int32, clip *sdl.Rect, angle float64, center *sdl.Point, flip sdl.RendererFlip) {
	dst := sdl.Rect{X: x, Y: y, W: t.width, H: t.height}

	if !clip.Empty() {
		dst.W = clip.W
		dst.H = clip.H
	}

	renderer.CopyEx(t.texture, clip, &dst, angle, center, flip)
}
