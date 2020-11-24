package game

import "github.com/veandco/go-sdl2/sdl"

const (
	BUTTON_SPRITE_MOUSE_OUT = iota
	BUTTON_SPRITE_MOUSE_OVER_MOTION
	BUTTON_SPRITE_MOUSE_DOWN
	BUTTON_SPRITE_MOUSE_UP
	BUTTON_SPRITE_TOTAL
)

type MouseButton struct {
	position sdl.Point
	currentSprite int
}

func NewMouseButton(x, y int32) MouseButton {
	return MouseButton{
		position: sdl.Point{X: x, Y: y},
		currentSprite: BUTTON_SPRITE_MOUSE_OUT,
	}
}

func (m *MouseButton) SetPosition(x, y int32) {
	m.position = sdl.Point{X: x, Y: y}
}

func (m *MouseButton) HandleEvent(x, y int32, t uint32) {
	inside := true

	switch {
	case x < m.position.X:
		inside = false
		break
	case x > (m.position.X + BUTTON_WIDTH):
		inside = false
		break
	case y < m.position.Y:
		inside = false
		break
	case y > (m.position.Y + BUTTON_HEIGHT):
		inside = false
		break
	}

	if !inside {
		m.currentSprite = BUTTON_SPRITE_MOUSE_OUT
	} else {
		switch t {
		case sdl.MOUSEMOTION:
			m.currentSprite = BUTTON_SPRITE_MOUSE_OVER_MOTION
			break
		case sdl.MOUSEBUTTONDOWN:
			m.currentSprite = BUTTON_SPRITE_MOUSE_DOWN
			break
		case sdl.MOUSEBUTTONUP:
			m.currentSprite = BUTTON_SPRITE_MOUSE_UP
			break
		}
	}
}

func (m *MouseButton) Render(renderer *sdl.Renderer, wrapper *TextureWrapper) {
	wrapper.Render(renderer, m.position.X, m.position.Y, &gSpriteClips[m.currentSprite])
}
