package effect

import (
	"kurt.blackwell.id.au/xyzmas/color"
)

func Fill(colors color.RGBGenerator) Effect {
	var c []color.RGBA
	return func(state State) State {
		if c == nil || len(c) != len(state.Colors) {
			c = color.Fill(len(state.Colors), colors)
		}

		state.Colors = c
		return state
	}
}
