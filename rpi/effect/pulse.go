package effect

import (
	"time"

	"kurt.blackwell.id.au/xyzmas/color"
	"kurt.blackwell.id.au/xyzmas/pixel"
)

func Pulse(
	colors color.RGBGenerator,
	durationMin time.Duration,
	durationMax time.Duration,
	exponent float64) Effect {
	return Pixels(pixel.Pulse(colors, durationMin, durationMax, exponent))
}
