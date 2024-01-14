package color

import "github.com/lucasb-eyer/go-colorful"

func HueShift(rgb RGBA, shift float64) RGBA {
	col := colorful.Color{
		R: float64(rgb.R) / 255.0,
		G: float64(rgb.G) / 255.0,
		B: float64(rgb.B) / 255.0,
	}
	h, c, l := col.Hcl()
	return NewColorfulRGB(colorful.Hcl(h+shift, c, l))
}
