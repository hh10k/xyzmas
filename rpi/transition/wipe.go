package transition

import (
	"math"

	"kurt.blackwell.id.au/xyzmas/color"
)

func Wipe() Transition {
	return func(a, b []color.RGBA, value float32) []color.RGBA {
		count := len(a)
		c := make([]color.RGBA, count)

		x := value * float32(count)
		index := int(math.Floor(float64(x)))

		for i := 0; i < index; i++ {
			c[i] = b[i]
		}
		c[index] = color.LerpRGB(a[index], b[index], x-float32(index))
		for i := index + 1; i < count; i++ {
			c[i] = a[i]
		}

		return c
	}
}

func ReverseWipe() Transition {
	wipe := Wipe()
	return func(a, b []color.RGBA, value float32) []color.RGBA {
		return wipe(b, a, 1-value)
	}
}
