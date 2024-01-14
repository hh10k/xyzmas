package transition

import (
	"kurt.blackwell.id.au/xyzmas/color"
	"kurt.blackwell.id.au/xyzmas/maths"
)

func Dissolve() Transition {
	var pixelsLevels []uint8
	return func(a, b []color.RGBA, value float32) []color.RGBA {
		if pixelsLevels == nil || len(pixelsLevels) != len(a) {
			pixelsLevels = make([]uint8, len(a))
			for i := range pixelsLevels {
				pixelsLevels[i] = uint8(maths.Rand.Intn(256))
			}
		}

		c := make([]color.RGBA, len(a))

		level := uint8(255 * value)
		for i := range c {
			if level < pixelsLevels[i] {
				c[i] = a[i]
			} else {
				c[i] = b[i]
			}
		}

		return c
	}
}
