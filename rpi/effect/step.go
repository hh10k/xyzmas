package effect

import (
	"time"

	"kurt.blackwell.id.au/xyzmas/color"
	"kurt.blackwell.id.au/xyzmas/maths"
)

func Step(interval time.Duration, amount int) Effect {
	startTime := time.Duration(0)
	return func(state State) State {
		if startTime == 0 {
			startTime = state.Time
		}

		t := ((state.Time - startTime) / interval)
		offset := int((t * time.Duration(amount)) % time.Duration(len(state.Colors)))
		c := make([]color.RGBA, len(state.Colors))
		for i := range c {
			c[i] = state.Colors[maths.Mod(i+offset, len(state.Colors))]
		}

		state.Colors = c
		return state
	}
}
