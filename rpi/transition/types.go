package transition

import (
	"time"

	"kurt.blackwell.id.au/xyzmas/color"
)

type Transition func(a, b []color.RGBA, value float32) []color.RGBA
type Generator func() (Transition, time.Duration)
