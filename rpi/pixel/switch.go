package pixel

import (
	"time"

	"github.com/fogleman/ease"
	"kurt.blackwell.id.au/xyzmas/color"
	"kurt.blackwell.id.au/xyzmas/maths"
)

// Switch changes each pixels between the different colors
func Switch(
	colors color.RGBGenerator,
	easing ease.Function,
	fadeDurationMin time.Duration,
	fadeDurationMax time.Duration,
	idleDurationMin time.Duration,
	idleDurationMax time.Duration,
) PixelGenerator {
	params := &switchParams{
		colors:            colors,
		easing:            easing,
		fadeDurationMin:   fadeDurationMin,
		fadeDurationRange: int64(fadeDurationMax-fadeDurationMin) + 1,
		idleDurationMin:   idleDurationMin,
		idleDurationRange: int64(idleDurationMax-idleDurationMin) + 1,
	}
	return func() Pixel {
		pixel := switchPixel{
			state:  newSwitchState(0, colors(), params),
			params: params,
		}
		// Initial state, offset time randomly
		offset := time.Duration(maths.Rand.Int63n(int64(pixel.state.nextTime)))
		pixel.state.fromTime -= offset
		pixel.state.toTime -= offset
		pixel.state.nextTime -= offset
		return &pixel
	}
}

type switchParams struct {
	colors            color.RGBGenerator
	easing            ease.Function
	fadeDurationMin   time.Duration
	fadeDurationRange int64
	idleDurationMin   time.Duration
	idleDurationRange int64
}

type switchState struct {
	fromColor color.RGBA
	fromTime  time.Duration
	toColor   color.RGBA
	toTime    time.Duration
	nextTime  time.Duration
}

type switchPixel struct {
	state  switchState
	params *switchParams
}

func newSwitchState(fromTime time.Duration, fromColor color.RGBA, params *switchParams) switchState {
	toTime := fromTime + params.fadeDurationMin + time.Duration(maths.Rand.Int63n(params.fadeDurationRange))
	return switchState{
		fromTime:  fromTime,
		fromColor: fromColor,
		toTime:    toTime,
		toColor:   params.colors(),
		nextTime:  toTime + params.idleDurationMin + time.Duration(maths.Rand.Int63n(params.idleDurationRange)),
	}
}

func (pixel *switchPixel) Duration() time.Duration {
	return InfiniteDuration
}

func (pixel *switchPixel) Blend(now time.Duration, background color.RGBA) color.RGBA {
	for now >= pixel.state.nextTime {
		pixel.state = newSwitchState(pixel.state.nextTime, pixel.state.toColor, pixel.params)
	}

	if now >= pixel.state.toTime {
		return pixel.state.toColor
	}
	if now <= pixel.state.fromTime {
		return pixel.state.fromColor
	}

	t := float64(now-pixel.state.fromTime) / float64(pixel.state.toTime-pixel.state.fromTime)
	return color.LerpRGB(pixel.state.fromColor, pixel.state.toColor, float32(pixel.params.easing(t)))
}
