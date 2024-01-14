package effect

import (
	"time"

	"kurt.blackwell.id.au/xyzmas/color"
	"kurt.blackwell.id.au/xyzmas/pixel"
)

func Flash(
	colors color.RGBGenerator,
	fadeDurationMin time.Duration,
	fadeDurationMax time.Duration,
	blankDurationMin time.Duration,
	blankDurationMax time.Duration) Effect {
	return Pixels(pixel.Flash(colors, fadeDurationMin, fadeDurationMax, blankDurationMin, blankDurationMax))
}
