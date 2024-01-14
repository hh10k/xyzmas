package color

import "github.com/lucasb-eyer/go-colorful"

type RGBA struct {
	R uint8
	G uint8
	B uint8
	A uint8
}

type RGBGenerator func() RGBA

func NewRGB(r uint8, g uint8, b uint8) RGBA {
	return RGBA{r, g, b, 255}
}

func NewRGBA(r uint8, g uint8, b uint8, a uint8) RGBA {
	return RGBA{r, g, b, a}
}

func NewColorfulRGB(c colorful.Color) RGBA {
	r, g, b := c.Clamped().RGB255()
	return RGBA{r, g, b, 255}
}

var Black = NewRGB(0x00, 0x00, 0x00)
var White = NewRGB(0xff, 0xff, 0xff)
var Red = NewRGB(0xff, 0x00, 0x00)
var Green = NewRGB(0x00, 0xff, 0x00)
var Blue = NewRGB(0x00, 0x00, 0xff)
var Orange = NewRGB(0xff, 0xff, 0x00)
var Cyan = NewRGB(0x00, 0xff, 0xff)
var Magenta = NewRGB(0xff, 0x00, 0xff)
