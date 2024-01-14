package color

import (
	"math"

	colorful "github.com/lucasb-eyer/go-colorful"
	"kurt.blackwell.id.au/xyzmas/maths"
)

const hueStepLarge = 120
const hueStepSmall = 80

const (
	tripletSaturatedHuesStepLarge int = iota
	tripletSaturatedHueStepSmall
	tripletDesaturatedHueStepLarge
	tripletDesaturatedHueStepSmall
	tripletSingleHueSaturations
	tripletSingleHueRandSL
	tripletCount
)

func RandomTriplet() (color1, color2, color3 RGBA) {
	switch maths.Rand.Intn(tripletCount) {
	case tripletSaturatedHuesStepLarge:
		hue1 := maths.Rand.Float64() * 360
		hue2 := math.Mod(hue1+hueStepLarge, 360)
		hue3 := math.Mod(hue2+hueStepLarge, 360)
		color1 = NewColorfulRGB(colorful.Hsv(hue1, 1, 1))
		color2 = NewColorfulRGB(colorful.Hsv(hue2, 1, 1))
		color3 = NewColorfulRGB(colorful.Hsv(hue3, 1, 1))
	case tripletSaturatedHueStepSmall:
		hue1 := maths.Rand.Float64() * 360
		hue2 := math.Mod(hue1+hueStepSmall, 360)
		hue3 := math.Mod(hue2+hueStepSmall, 360)
		if maths.Rand.Intn(2) == 0 {
			hue3, hue1 = hue1, hue3
		}
		color1 = NewColorfulRGB(colorful.Hsv(hue1, 1, 0.6))
		color2 = NewColorfulRGB(colorful.Hsv(hue2, 1, 0.8))
		color3 = NewColorfulRGB(colorful.Hsv(hue3, 1, 1))
	case tripletDesaturatedHueStepLarge:
		hue1 := maths.Rand.Float64() * 360
		hue2 := math.Mod(hue1+hueStepLarge, 360)
		hue3 := math.Mod(hue2+hueStepLarge, 360)
		color1 = NewColorfulRGB(colorful.Hsv(hue1, 0.7, 1))
		color2 = NewColorfulRGB(colorful.Hsv(hue2, 0.7, 1))
		color3 = NewColorfulRGB(colorful.Hsv(hue3, 0.7, 1))
	case tripletDesaturatedHueStepSmall:
		hue1 := maths.Rand.Float64() * 360
		hue2 := math.Mod(hue1+hueStepSmall, 360)
		hue3 := math.Mod(hue2+hueStepSmall, 360)
		if maths.Rand.Intn(2) == 0 {
			hue3, hue1 = hue1, hue3
		}
		color1 = NewColorfulRGB(colorful.Hsv(hue1, 0.7, 0.6))
		color2 = NewColorfulRGB(colorful.Hsv(hue2, 0.7, 0.8))
		color3 = NewColorfulRGB(colorful.Hsv(hue3, 0.7, 1))
	case tripletSingleHueSaturations:
		hue := maths.Rand.Float64() * 360
		color1 = NewColorfulRGB(colorful.Hsv(hue, 0.4, 1))
		color2 = NewColorfulRGB(colorful.Hsv(hue, 0.7, 1))
		color3 = NewColorfulRGB(colorful.Hsv(hue, 1, 1))
	case tripletSingleHueRandSL:
		hue := maths.Rand.Float64() * 360
		c := 0.6 + maths.Rand.Float64()*0.2
		l := 0.6
		color1 = NewColorfulRGB(colorful.Hsl(hue, c, l))
		color2 = NewColorfulRGB(colorful.Hsl(math.Mod(hue+120, 360), c, l))
		color3 = NewColorfulRGB(colorful.Hsl(math.Mod(hue+240, 360), c, l))
	}

	if maths.Rand.Intn(2) == 0 {
		color3, color1 = color1, color3
	}

	return
}
