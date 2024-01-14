package renderer

import (
	"kurt.blackwell.id.au/xyzmas/color"
	"kurt.blackwell.id.au/xyzmas/configuration"
)

type limiterRenderer struct {
	renderer Renderer
	config   configuration.Configuration
	limit    uint64
}

func (r *limiterRenderer) Render(colors []color.RGBA) error {
	var rgbTotal uint64 = 0
	for _, rgb := range colors {
		rgbTotal += uint64(rgb.R) + uint64(rgb.G) + uint64(rgb.B)
	}

	rgbTotal = uint64(float64(rgbTotal) * r.config.Brightness)
	if rgbTotal > r.limit {
		// Too bright

		c := make([]color.RGBA, len(colors))
		t := int16(256 * r.limit / rgbTotal)
		for i, x := range colors {
			c[i] = color.NewRGBA(
				uint8((int16(x.R)*t)>>8),
				uint8((int16(x.G)*t)>>8),
				uint8((int16(x.B)*t)>>8),
				255,
			)
		}
		colors = c
	}

	if err := r.renderer.Render(colors); err != nil {
		return err
	}

	return nil
}

func NewLimiterRenderer(renderer Renderer, config configuration.Configuration) (Renderer, error) {
	return &limiterRenderer{
		renderer: renderer,
		config:   config,
		limit:    uint64(config.AmpLimit * 50 * 255),
	}, nil
}
