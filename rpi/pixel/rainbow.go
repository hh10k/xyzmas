package pixel

import (
	"math"
	"time"

	"github.com/lucasb-eyer/go-colorful"
	"kurt.blackwell.id.au/xyzmas/color"
	"kurt.blackwell.id.au/xyzmas/maths"
)

func RainbowPulse(
	durationMin time.Duration,
	durationMax time.Duration,
	exponent float64,
	cycleSpeed float64) PixelGenerator {
	var durationRange = int64(durationMax-durationMin) + 1

	return func() Pixel {
		return &rainbowPixel{
			duration:       durationMin + time.Duration(maths.Rand.Int63n(durationRange)),
			colorHueOffset: maths.Rand.Float64() * 360,
			colorC:         0.6 + maths.Rand.Float64()*0.3,
			colorL:         0.4 + maths.Rand.Float64()*0.3,
			exponent:       exponent,
			cycleSpeed:     cycleSpeed * 360,
		}
	}
}

type rainbowPixel struct {
	duration       time.Duration
	exponent       float64
	cycleSpeed     float64
	colorHueOffset float64
	colorC         float64
	colorL         float64
}

func (pixel *rainbowPixel) Duration() time.Duration {
	return pixel.duration
}

func (pixel *rainbowPixel) Blend(now time.Duration, background color.RGBA) color.RGBA {
	tLinear := 0.5 - 0.5*math.Cos((2*math.Pi*float64(now))/float64(pixel.duration))
	tExp := math.Pow(tLinear, pixel.exponent)
	hue := math.Mod(pixel.colorHueOffset+(float64(now)/float64(pixel.duration))*pixel.cycleSpeed, 360)
	return color.LerpRGB(background, color.NewColorfulRGB(colorful.Hsv(hue, pixel.colorC, pixel.colorL)), float32(tExp))
}
