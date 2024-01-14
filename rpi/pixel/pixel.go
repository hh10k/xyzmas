package pixel

import (
	"time"

	"kurt.blackwell.id.au/xyzmas/color"
)

const InfiniteDuration time.Duration = -1

type PixelGenerator func() Pixel

type Pixel interface {
	Duration() time.Duration
	Blend(now time.Duration, background color.RGBA) color.RGBA
}
