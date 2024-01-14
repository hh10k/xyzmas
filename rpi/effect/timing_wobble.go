package effect

import (
	"math"
	"time"
)

func TimingWobble(
	frequency time.Duration,
	amplitude time.Duration,
	pow float64,
	effects ...Effect,
) Effect {
	return func(state State) State {
		originalTime := state.Time

		t := math.Sin(float64(state.Time) * math.Pi / float64(frequency))
		s := math.Signbit(t)
		t = (1 - math.Pow(1-math.Abs(t), pow))
		if s {
			t = -t
		}
		state.Time = state.Time + time.Duration(t*float64(amplitude))

		for _, e := range effects {
			state = e(state)
		}

		state.Time = originalTime
		return state
	}
}
