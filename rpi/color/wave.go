package color

import (
	"math"
)

func SinWave(color1 RGBA, color2 RGBA, step float64) RGBGenerator {
	var t float64
	return func() RGBA {
		c := LerpRGB(color1, color2, 0.5+0.5*float32(math.Sin(t)))
		t += step
		return c
	}
}
