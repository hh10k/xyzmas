package effect

import (
	"time"
)

func Cycle(interval time.Duration, effects ...Effect) Effect {
	startTime := time.Duration(0)
	return func(state State) State {
		if startTime == 0 {
			startTime = state.Time
		}

		t := int((state.Time - startTime) / interval)
		i := t % len(effects)
		return effects[i](state)
	}
}
