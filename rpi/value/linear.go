package value

import "time"

func Linear(perSecond float64) Function {
	first := true
	var start time.Duration

	return func(t time.Duration) float64 {
		if first {
			start = t
			first = false
		}
		return (t - start).Seconds() * perSecond
	}
}
