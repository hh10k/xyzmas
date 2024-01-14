package renderer

import (
	"fmt"
	"time"

	"kurt.blackwell.id.au/xyzmas/color"
	"kurt.blackwell.id.au/xyzmas/configuration"
)

const BENCHMARK_INTERVAL = time.Second

type benchmarkRenderer struct {
	renderer Renderer
	config   configuration.Configuration
	lastTime time.Time

	frameCount int
	rgbCount   uint64
	rgbMax     uint64
}

func (r *benchmarkRenderer) Render(colors []color.RGBA) error {
	if err := r.renderer.Render(colors); err != nil {
		return err
	}

	now := time.Now()
	r.frameCount++

	var rgbTotal uint64 = 0
	for _, rgb := range colors {
		rgbTotal += uint64(rgb.R) + uint64(rgb.G) + uint64(rgb.B)
	}
	r.rgbCount += rgbTotal
	if rgbTotal > r.rgbMax {
		r.rgbMax = rgbTotal
	}

	interval := now.Sub(r.lastTime)
	if interval >= BENCHMARK_INTERVAL {
		// 60mA per RGB LED for 5V ws2811 strips
		var amps float32 = float32(0.06 * r.config.Brightness / (3 * 255))

		fmt.Printf(
			"%f fps, %fA max, %fA avg\n",
			float32(time.Duration(r.frameCount)*time.Second)/float32(interval),
			amps*float32(r.rgbMax),
			amps*float32(r.rgbCount)/float32(r.frameCount),
		)

		r.lastTime = now
		r.frameCount = 0
		r.rgbCount = 0
		r.rgbMax = 0
	}

	return nil
}

func NewBenchmarkRenderer(renderer Renderer, config configuration.Configuration) (Renderer, error) {
	return &benchmarkRenderer{
		renderer:   renderer,
		config:     config,
		lastTime:   time.Now(),
		frameCount: 0,
		rgbCount:   0,
		rgbMax:     0,
	}, nil
}
