package effect

import (
	"kurt.blackwell.id.au/xyzmas/color"
	"kurt.blackwell.id.au/xyzmas/maths"
)

func Sparkle(colors color.RGBGenerator, ratio float32) Effect {
	p := int(256 * ratio)
	return func(state State) State {
		c := make([]color.RGBA, len(state.Colors))
		for i, rgb := range state.Colors {
			if maths.Rand.Intn(256) <= p {
				rgb = colors()
			}
			c[i] = rgb
		}

		state.Colors = c
		return state
	}
}
