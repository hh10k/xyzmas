package effect

import (
	"kurt.blackwell.id.au/xyzmas/color"
	"kurt.blackwell.id.au/xyzmas/value"
)

func HueShift(hueShift value.Function, effect Effect) Effect {
	return func(state State) State {
		value := hueShift(state.Time)

		state = effect(state)
		colors := make([]color.RGBA, len(state.Colors))
		for i, c := range state.Colors {
			colors[i] = color.HueShift(c, value)
		}

		state.Colors = colors
		return state
	}
}
