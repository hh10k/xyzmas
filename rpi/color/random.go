package color

import (
	"math"

	colorful "github.com/lucasb-eyer/go-colorful"
	"kurt.blackwell.id.au/xyzmas/maths"
)

func RandomRGBsOf(colors ...RGBA) RGBGenerator {
	return func() RGBA {
		i := maths.Rand.Intn(len(colors))
		return colors[i]
	}
}

func RandomRGBs(
	min RGBA,
	max RGBA,
) RGBGenerator {
	rangeR := int(max.R) - int(min.R)
	if rangeR < 0 {
		rangeR = -rangeR
		min.R = max.R
	}
	rangeG := int(max.G) - int(min.G)
	if rangeG < 0 {
		rangeG = -rangeG
		min.G = max.G
	}
	rangeB := int(max.B) - int(min.B)
	if rangeB < 0 {
		rangeB = -rangeB
		min.B = max.B
	}

	rangeR++
	rangeG++
	rangeB++
	return func() RGBA {
		return NewRGB(
			min.R+uint8(maths.Rand.Intn(rangeR)),
			min.G+uint8(maths.Rand.Intn(rangeG)),
			min.B+uint8(maths.Rand.Intn(rangeB)),
		)
	}
}

func RandomHslHues(
	hueMin float64, hueMax float64,
	saturationMin float64, saturationMax float64,
	luminanceMin float64, luminanceMax float64) RGBGenerator {

	hueRange := hueMax - hueMin
	saturationRange := saturationMax - saturationMin
	luminanceRange := luminanceMax - luminanceMin

	return func() RGBA {
		return NewColorfulRGB(colorful.Hsl(
			math.Mod(hueMin+maths.Rand.Float64()*hueRange, 360),
			saturationMin+maths.Rand.Float64()*saturationRange,
			luminanceMin+maths.Rand.Float64()*luminanceRange,
		))
	}
}

func RandomHclHues(
	hueMin float64, hueMax float64,
	cMin float64, cMax float64,
	lMin float64, lMax float64) RGBGenerator {

	hueRange := hueMax - hueMin
	cRange := cMax - cMin
	lRange := lMax - lMin

	return func() RGBA {
		return NewColorfulRGB(colorful.Hsl(
			math.Mod(hueMin+maths.Rand.Float64()*hueRange, 360),
			cMin+maths.Rand.Float64()*cRange,
			lMin+maths.Rand.Float64()*lRange,
		))
	}
}

func CycleHclHues(
	hueStep float64,
	cMin float64, cMax float64,
	lMin float64, lMax float64) RGBGenerator {

	if maths.Rand.Intn(2) == 0 {
		hueStep = -hueStep
	}

	hue := maths.Rand.Float64() * 360
	cRange := cMax - cMin
	lRange := lMax - lMin

	return func() RGBA {
		c := colorful.Hcl(
			hue,
			cMin+maths.Rand.Float64()*cRange,
			lMin+maths.Rand.Float64()*lRange,
		)

		hue += hueStep

		return NewColorfulRGB(c)
	}
}

func CycleHsvHues(
	hueStep float64,
	saturationMin float64, saturationMax float64,
	valueMin float64, valueMax float64) RGBGenerator {

	if maths.Rand.Intn(2) == 0 {
		hueStep = -hueStep
	}

	hue := maths.Rand.Float64() * 360
	saturationRange := saturationMax - saturationMin
	valueRange := valueMax - valueMin

	return func() RGBA {
		c := colorful.Hsv(
			hue,
			saturationMin+maths.Rand.Float64()*saturationRange,
			valueMin+maths.Rand.Float64()*valueRange,
		)

		hue = math.Mod(hue+hueStep, 360)

		return NewColorfulRGB(c)
	}
}
