package transition

import (
	"kurt.blackwell.id.au/xyzmas/color"
)

func Fade() Transition {
	return func(a, b []color.RGBA, value float32) []color.RGBA {
		return color.LerpRGBs(a, b, value)
	}
}
