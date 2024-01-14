package effect

import (
	"time"

	"kurt.blackwell.id.au/xyzmas/color"
	"kurt.blackwell.id.au/xyzmas/maths"
	"kurt.blackwell.id.au/xyzmas/pixel"
)

type pixelState struct {
	startTime time.Duration
	pixel     pixel.Pixel
}

func Pixels(pixelGenerator pixel.PixelGenerator) Effect {
	var pixels []pixelState

	return func(state State) State {
		// (Re)create pixels if necessary
		if pixels == nil || len(pixels) != len(state.Colors) {
			pixels = make([]pixelState, len(state.Colors))
			for i := range pixels {
				p := pixelState{
					startTime: state.Time,
					pixel:     pixelGenerator(),
				}
				duration := p.pixel.Duration()
				if duration != pixel.InfiniteDuration {
					// Start at random time during pixel animation
					p.startTime -= time.Duration(maths.Rand.Int63n(int64(duration + 1)))
				}
				pixels[i] = p
			}
		}

		// Update pixels
		c := make([]color.RGBA, len(state.Colors))

		for i, p := range pixels {
			for {
				duration := p.pixel.Duration()
				if duration == pixel.InfiniteDuration ||
					p.startTime+p.pixel.Duration() > state.Time {
					break
				}
				p = pixelState{
					startTime: p.startTime + p.pixel.Duration(),
					pixel:     pixelGenerator(),
				}
				pixels[i] = p
			}

			c[i] = p.pixel.Blend(state.Time-p.startTime, state.Colors[i])
		}

		state.Colors = c
		return state
	}
}
