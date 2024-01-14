package effect

import (
	"time"

	"kurt.blackwell.id.au/xyzmas/transition"
)

func Transition(transition transition.Transition, getValue func(now time.Duration) float32, first Effect, second Effect) Effect {
	return func(state State) State {
		value := getValue(state.Time)

		if value < 0.001 {
			return first(state)
		}
		if value > 0.999 {
			return second(state)
		}

		firstState := first(state)
		secondState := second(state)
		state.Colors = transition(firstState.Colors, secondState.Colors, value)
		return state
	}
}
