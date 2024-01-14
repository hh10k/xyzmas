//go:build ws281x
// +build ws281x

package renderer

import (
	"fmt"
	"time"

	ws281x "github.com/rpi-ws281x/rpi-ws281x-go"
	"kurt.blackwell.id.au/xyzmas/color"
	"kurt.blackwell.id.au/xyzmas/configuration"
)

type ws281xRenderer struct {
	display *ws281x.WS2811
}

func (r ws281xRenderer) Render(colors []color.RGBA) error {
	leds := r.display.Leds(0)
	for i := 0; i < len(leds); i++ {
		c := colors[i]
		leds[i] = (uint32(c.R) << 16) + (uint32(c.G) << 8) + uint32(c.B)
	}

	if err := r.display.Render(); err != nil {
		return err
	}

	// Wait for pixels to be updated, otherwise if it is lagging the
	// next Render will have to wait and our timing will be messed up.
	return r.display.Wait()
}

func newWs281xRenderer(config configuration.Configuration) (Renderer, error) {
	hw := ws281x.HwDetect()
	fmt.Printf("Hardware Type    : %d\n", hw.Type)
	fmt.Printf("Hardware Version : 0x%08X\n", hw.Version)
	fmt.Printf("Periph base      : 0x%08X\n", hw.PeriphBase)
	fmt.Printf("Video core base  : 0x%08X\n", hw.VideocoreBase)
	fmt.Printf("Description      : %v\n", hw.Desc)

	options := ws281x.DefaultOptions
	options.Channels[0].GpioPin = config.GpioPin
	options.Channels[0].LedCount = config.LedCount
	options.Channels[0].Brightness = int(255 * config.Brightness)
	options.Channels[0].StripeType = ws281x.WS2811StripRGB
	fmt.Printf("Options          : %+v\n", options)

	display, err := ws281x.MakeWS2811(&options)
	if err != nil {
		return nil, err
	}

	// In practice, it only takes one retry after basic.target is ready
	retries := 5
	const retryInterval = 2 * time.Second
	for {
		err := display.Init()
		if err == nil {
			break
		}

		retries--
		if retries <= 0 {
			return nil, err
		}

		time.Sleep(retryInterval)
	}

	return &ws281xRenderer{display: display}, nil
}

func init() {
	renderers[Ws281xType] = newWs281xRenderer
	defaultRendererType = Ws281xType
}
