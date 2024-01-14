package color

import "math"

const gamma = 2.2

var toGamma = makeToGamma()
var fromGamma = makeFromGamma()

// makeToGamma returns an array for converting a linear 0..255 RGBA value into a gamma-corrected 0..1023 one.
func makeToGamma() []int16 {
	x := make([]int16, 256)
	for i := 0; i < 256; i += 1 {
		x[i] = int16(math.Round(1023 * math.Pow(float64(i)/255.0, gamma)))
	}
	return x
}

// makeFromGamma returns an array for converting a gamma-corrected 0..511 RGBA value into a linear 0..255 one.
func makeFromGamma() []uint8 {
	x := make([]uint8, 1024)
	for i := 0; i < 1024; i += 1 {
		x[i] = uint8(math.Round(255 * math.Pow(float64(i)/1023.0, 1.0/gamma)))
	}
	return x
}

func LerpRGBLinear(a RGBA, b RGBA, t int16) RGBA {
	return NewRGBA(
		a.R+uint8(((int16(b.R)-int16(a.R))*t)>>8),
		a.G+uint8(((int16(b.G)-int16(a.G))*t)>>8),
		a.B+uint8(((int16(b.B)-int16(a.B))*t)>>8),
		a.A+uint8(((int16(b.A)-int16(a.A))*t)>>8),
	)
}

func LerpRGBGamma(c1 RGBA, c2 RGBA, t int32) RGBA {
	r1 := int32(toGamma[c1.R])
	r2 := int32(toGamma[c2.R])
	g1 := int32(toGamma[c1.G])
	g2 := int32(toGamma[c2.G])
	b1 := int32(toGamma[c1.B])
	b2 := int32(toGamma[c2.B])

	return NewRGBA(
		fromGamma[r1+((r2-r1)*t)>>8],
		fromGamma[g1+((g2-g1)*t)>>8],
		fromGamma[b1+((b2-b1)*t)>>8],
		c1.A+uint8(((int32(c2.A)-int32(c1.A))*t)>>8),
	)
}

func LerpRGB(a RGBA, b RGBA, t float32) RGBA {
	return LerpRGBGamma(a, b, int32(t*256))
}

func LerpRGBs(a []RGBA, b []RGBA, t float32) []RGBA {
	c := make([]RGBA, len(a))
	s := int32(t * 256)
	for i := range c {
		c[i] = LerpRGBGamma(a[i], b[i], s)
	}
	return c
}
