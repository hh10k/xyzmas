package effect

import (
	"time"

	"kurt.blackwell.id.au/xyzmas/color"
	"kurt.blackwell.id.au/xyzmas/value"
)

func Slide(pixelsPerSec float64) Effect {
	pixelsSpeed := -pixelsPerSec / float64(time.Second)
	return func(state State) State {
		state.Colors = color.Translate(state.Colors, float64(state.Time)*pixelsSpeed)
		return state
	}
}

func SlideValue(fn value.Function) Effect {
	return func(state State) State {
		state.Colors = color.Translate(state.Colors, fn(state.Time))
		return state
	}
}
