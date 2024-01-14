package renderer

import "kurt.blackwell.id.au/xyzmas/color"

type Renderer interface {
	Render(colors []color.RGBA) error
}
