package renderer

import (
	"fmt"
	"strings"

	"github.com/logrusorgru/aurora/v3"
	"kurt.blackwell.id.au/xyzmas/color"
	"kurt.blackwell.id.au/xyzmas/configuration"
)

const DOT = "O"

type consoleRenderer struct {
}

func (r *consoleRenderer) Render(colors []color.RGBA) error {
	var buffer strings.Builder

	for _, c := range colors {
		if c.R == c.G && c.G == c.B {
			// Grey range
			if c.R == 0 {
				buffer.WriteString(aurora.Black(DOT).String())
			} else {
				g := uint8(24 * int(c.R) / 256)
				buffer.WriteString(aurora.Gray(g, DOT).String())
			}
		} else {
			// RGB range
			r := uint8(5 * int(c.R) / 255)
			g := uint8(5 * int(c.G) / 255)
			b := uint8(5 * int(c.B) / 255)
			buffer.WriteString(aurora.Index(16+36*r+6*g+b, DOT).String())
		}
	}

	// Move back to the start of the line after rendering
	// By doing this after rendering we will avoid overwriting other console output
	buffer.WriteString("\x1B[2A\n")

	fmt.Println(buffer.String())

	return nil
}

func newConsoleRenderer(config configuration.Configuration) (Renderer, error) {
	return &consoleRenderer{}, nil
}
