package transition

import (
	"time"

	"kurt.blackwell.id.au/xyzmas/maths"
)

type predefined struct {
	Transition func() Transition
	Duration   time.Duration
}

var transitions = [...]predefined{
	{Wipe, time.Second},
	{ReverseWipe, time.Second},
	{Dissolve, time.Second},
	{Fade, time.Second},
}

func Random() Generator {
	return func() (Transition, time.Duration) {
		next := transitions[maths.Rand.Intn(len(transitions))]
		return next.Transition(), next.Duration
	}
}

func Cycle() Generator {
	i := -1
	return func() (Transition, time.Duration) {
		i = (i + 1) % len(transitions)
		next := transitions[i]
		return next.Transition(), next.Duration
	}
}
