package effect

import (
	"time"

	"kurt.blackwell.id.au/xyzmas/transition"
)

type Generator func(now time.Duration) (effect Effect, duration time.Duration)

func Sequence(effects Generator, transitions transition.Generator) Effect {
	var current Effect
	var currentDuration time.Duration
	var trans Effect
	var transState transition.State
	next, nextDuration := transitions()

	return func(state State) State {
		if current == nil {
			// First
			current, currentDuration = effects(state.Time)
			trans = current
			transState = transition.NewState(1)
			transState.FadeTo(1, state.Time, 0)
		}

		// The current effect will be shown for at least currentDuration, including fade at either end
		// It will be shown for no less than 2 * nextDuration due to fading
		nextStartTime := transState.FromTime + currentDuration - nextDuration
		if nextStartTime < transState.ToTime {
			nextStartTime = transState.ToTime
		}

		if state.Time >= nextStartTime {
			prev := current
			current, currentDuration = effects(state.Time)
			transState = transition.NewState(0)
			transState.FadeTo(1, nextStartTime, nextDuration)
			trans = Transition(next, transState.ValueAt, prev, current)

			next, nextDuration = transitions()
		}

		return trans(state)
	}
}
