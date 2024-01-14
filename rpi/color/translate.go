package color

import (
	"math"

	"kurt.blackwell.id.au/xyzmas/maths"
)

func Translate(rgb []RGBA, distance float64) []RGBA {
	offset := int(math.Floor(distance))
	t := float32(distance) - float32(offset)

	c := make([]RGBA, len(rgb))
	for i := range c {
		prev := rgb[maths.Mod(i+offset, len(rgb))]
		next := rgb[maths.Mod(i+offset+1, len(rgb))]

		c[i] = LerpRGB(prev, next, t)
	}

	return c
}
