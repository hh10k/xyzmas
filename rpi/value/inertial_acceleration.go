package value

import (
	"time"

	"kurt.blackwell.id.au/xyzmas/maths"
)

type InertialAccelerationParams struct {
	InitialValue, InitialVelocity, InitialAcceleration float64
	AccelerationChange                                 float64
	VelocityDrag, AccelerationDrag                     float64
}

func InertialAcceleration(params InertialAccelerationParams) Function {
	prevValue := params.InitialValue
	nextValue := params.InitialValue
	velocity := params.InitialVelocity
	acceleration := params.InitialAcceleration
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
			acceleration = (acceleration + (maths.Rand.Float64()*2-1)*params.AccelerationChange) * params.AccelerationDrag
			velocity = (velocity + acceleration) * params.VelocityDrag
			prevValue = nextValue
			nextValue += velocity
		}

		return prevValue + (nextValue-prevValue)*float64(t-prevT)/float64(inertiaTickInterval)
	}
}
