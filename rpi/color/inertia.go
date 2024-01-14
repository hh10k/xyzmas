package color

import (
	"math"

	colorful "github.com/lucasb-eyer/go-colorful"
	"kurt.blackwell.id.au/xyzmas/maths"
)

func InertiaHclHues(
	cMin float64, cMax float64,
	lMin float64, lMax float64,
) RGBGenerator {
	hueAcc := 0.1
	cAcc := 0.002
	lAcc := 0.002

	hueT := maths.Rand.Float64() * 360
	hueV := 0.0
	cT := maths.Rand.Float64() * 2 * math.Pi
	cV := 0.0
	cHalf := cMax - cMin
	cMiddle := cMin + cHalf
	lT := maths.Rand.Float64() * 2 * math.Pi
	lV := 0.0
	lHalf := (lMax - lMin) / 2
	lMiddle := lMin + lHalf

	return func() RGBA {
		hueV = (hueV + (maths.Rand.Float64()*2-1)*hueAcc) * 0.95
		cV = (cV + (maths.Rand.Float64()*2-1)*cAcc) * 0.95
		lV = (lV + (maths.Rand.Float64()*2-1)*lAcc) * 0.95
		hueT += hueV
		cT += cV
		lT += lV

		return NewColorfulRGB(colorful.Hsl(
			math.Mod(hueT, 360),
			cMiddle+math.Sin(cT)*cHalf,
			lMiddle+math.Sin(lT)*lHalf,
		))
	}
}
