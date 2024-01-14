package renderer

import (
	"fmt"

	"kurt.blackwell.id.au/xyzmas/configuration"
)

type RendererType string

const (
	DefaultType RendererType = "default"
	ConsoleType RendererType = "console"
	NullType    RendererType = "null"
	Ws281xType  RendererType = "ws281x"
)

var renderers = make(map[RendererType]func(config configuration.Configuration) (Renderer, error))
var defaultRendererType = ConsoleType

func init() {
	renderers[ConsoleType] = newConsoleRenderer
	renderers[NullType] = newNullRenderer
}

func New(config configuration.Configuration) (Renderer, error) {
	name := RendererType(config.Renderer)
	if name == DefaultType {
		name = defaultRendererType
	}

	factory := renderers[name]
	if factory == nil {
		return nil, fmt.Errorf("No renderer '%s' found", name)
	}

	ren, err := factory(config)
	if err == nil && config.Benchmark {
		ren, err = NewBenchmarkRenderer(ren, config)
	}
	if err == nil && config.AmpLimit != 0 {
		ren, err = NewLimiterRenderer(ren, config)
	}

	return ren, err
}
