package effect

import (
	"time"

	"github.com/fogleman/ease"
)

// Applies an easing function to an effect's animation
func Timing(
	context Context,
	easing ease.Function,
	effects ...Effect,
) Effect {
	return func(state State) State {
		originalTime := state.Time

		d := float64(context.Duration)
		t := float64(state.Time-context.StartTime) / d
		state.Time = time.Duration(easing(t) * d)

		for _, e := range effects {
			state = e(state)
		}

		state.Time = originalTime
		return state
	}
}
