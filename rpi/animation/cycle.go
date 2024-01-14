package animation

import (
	"time"

	"kurt.blackwell.id.au/xyzmas/color"
	"kurt.blackwell.id.au/xyzmas/effect"
)

func CycleRGB(count int) effect.Effect {
	return effect.Cycle(
		3*time.Second,
		effect.Fill(color.Constant(color.Red)),
		effect.Fill(color.Constant(color.Green)),
		effect.Fill(color.Constant(color.Blue)),
	)
}
