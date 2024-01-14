package color

import (
	"math"

	colorful "github.com/lucasb-eyer/go-colorful"
	"kurt.blackwell.id.au/xyzmas/maths"
)

const (
	pairOpposite int = iota
	pairTetradic
	pairCount
)

func RandomHuePair() (hue1, hue2 float64) {
	hue1 = maths.Rand.Float64() * 360
	hue2 = hue1
	switch maths.Rand.Intn(pairCount) {
	case pairTetradic:
		hue2 = math.Mod(hue1+(360.0/3.0), 360)
	case pairOpposite:
		hue2 = math.Mod(hue1+(360.0/2.0), 360)
	}

	if maths.Rand.Intn(2) == 0 {
		hue2, hue1 = hue1, hue2
	}

	return
}

func RandomPair() (color1, color2 RGBA) {
	hue1, hue2 := RandomHuePair()

	s := 0.7 + maths.Rand.Float64()*0.3

	color1 = NewColorfulRGB(colorful.Hsv(hue1, s, 1))
	color2 = NewColorfulRGB(colorful.Hsv(hue2, s, 1))

	return
}

// RandomChromaLuminance returns a random chroma and luminance that are 'nice'
func RandomChromaLuminance() (c, l float64) {
	// Slightly more likely to be chroma of 1
	c = math.Min(1, 0.6+maths.Rand.Float64()*0.5)
	// Slightly more likely to be luminance of 0.5, up to 0.8
	l = math.Max(0.5, 0.4+maths.Rand.Float64()*0.4)
	return
}
