package renderer

import (
	"kurt.blackwell.id.au/xyzmas/color"
	"kurt.blackwell.id.au/xyzmas/configuration"
)

type nullRenderer struct {
}

func (r *nullRenderer) Render(colors []color.RGBA) error {
	return nil
}

func newNullRenderer(config configuration.Configuration) (Renderer, error) {
	return &nullRenderer{}, nil
}
