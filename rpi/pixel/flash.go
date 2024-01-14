package pixel

import (
	"time"

	"kurt.blackwell.id.au/xyzmas/color"
	"kurt.blackwell.id.au/xyzmas/maths"
)

func Flash(
	colors color.RGBGenerator,
	fadeDurationMin time.Duration,
	fadeDurationMax time.Duration,
	blankDurationMin time.Duration,
	blankDurationMax time.Duration) PixelGenerator {
	var fadeDurationRange = int64(fadeDurationMax-fadeDurationMin) + 1
	var blankDurationRange = int64(blankDurationMax-blankDurationMin) + 1
	return func() Pixel {
		fadeDuration := fadeDurationMin + time.Duration(maths.Rand.Int63n(fadeDurationRange))
		blankDuration := blankDurationMin + time.Duration(maths.Rand.Int63n(blankDurationRange))
		return &flashPixel{
			duration:      fadeDuration,
			durationTotal: fadeDuration + blankDuration,
			color:         colors(),
		}
	}
}

type flashPixel struct {
	// How long the fade goes for
	duration time.Duration
	// Fade + blank duration
	durationTotal time.Duration
	// The colour
	color color.RGBA
}

func (pixel *flashPixel) Duration() time.Duration {
	return pixel.durationTotal
}

func (pixel *flashPixel) Blend(now time.Duration, background color.RGBA) color.RGBA {
	if now >= pixel.duration {
		return background
	}

	t := float64(now) / float64(pixel.duration)
	t = 1 - t
	t = t * t
	t = t * t
	t = t * t
	t = t * t
	return color.LerpRGB(background, pixel.color, float32(t))
}
