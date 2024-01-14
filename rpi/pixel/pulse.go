package pixel

import (
	"math"
	"time"

	"kurt.blackwell.id.au/xyzmas/color"
	"kurt.blackwell.id.au/xyzmas/maths"
)

func Pulse(
	colors color.RGBGenerator,
	durationMin time.Duration,
	durationMax time.Duration,
	exponent float64) PixelGenerator {
	var durationRange = int64(durationMax-durationMin) + 1
	return func() Pixel {
		return &pulsePixel{
			duration: durationMin + time.Duration(maths.Rand.Int63n(durationRange)),
			color:    colors(),
			exponent: exponent,
		}
	}
}

type pulsePixel struct {
	duration time.Duration
	color    color.RGBA
	exponent float64
}

func (pixel *pulsePixel) Duration() time.Duration {
	return pixel.duration
}

func (pixel *pulsePixel) Blend(now time.Duration, background color.RGBA) color.RGBA {
	tLinear := 0.5 - 0.5*math.Cos((2*math.Pi*float64(now))/float64(pixel.duration))
	tExp := math.Pow(tLinear, pixel.exponent)
	return color.LerpRGB(background, pixel.color, float32(tExp))
}
