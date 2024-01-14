package value

import (
	"time"

	"kurt.blackwell.id.au/xyzmas/maths"
)

type InertialVelocityParams struct {
	InitialValue, InitialVelocity float64
	Acceleration                  float64
	VelocityDrag                  float64
}

func InertialVelocity(params InertialVelocityParams) Function {
	prevValue := params.InitialValue
	nextValue := params.InitialValue
	velocity := params.InitialVelocity
	var prevT time.Duration

	return func(t time.Duration) float64 {
		var steps int
		if t < prevT {
			// Time reversed
			prevT = t
			steps = 0
		} else {
			// Step forward so that t >= prevT
			fullSteps := (t - prevT) / inertiaTickInterval
			prevT += inertiaTickInterval * fullSteps
			steps = min(inertiaMaxSteps, int(fullSteps))
		}

		for i := 0; i < steps; i++ {
			velocity = (velocity + (maths.Rand.Float64()*2-1)*params.Acceleration) * params.VelocityDrag
			prevValue = nextValue
			nextValue += velocity
		}

		return prevValue + (nextValue-prevValue)*float64(t-prevT)/float64(inertiaTickInterval)
	}
}
