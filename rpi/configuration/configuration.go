package configuration

import (
	"encoding/json"
	"flag"
	"io"
	"os"
)

type IdModeConfiguration struct {
	StepIntervalMs int `json:"stepIntervalMs"`
}

type Configuration struct {
	ConfigPath string `json:"-"`

	Verbose         bool    `json:"verbose"`
	Renderer        string  `json:"renderer"`
	Mode            string  `json:"mode"`
	Benchmark       bool    `json:"benchmark"`
	GpioPin         int     `json:"gpioPin"`
	LedCount        int     `json:"ledCount"`
	Brightness      float64 `json:"brightness"`
	AmpLimit        float64 `json:"ampLimit"`
	FramesPerSecond int     `json:"framesPerSecond"`

	// Mode = id
	IdMode IdModeConfiguration `json:"idMode"`

	NightEnabled    bool    `json:"nightEnabled"`
	NightBrightness float32 `json:"nightBrightness"`
	SunsetStart     int     `json:"sunsetStart"`
	SunsetEnd       int     `json:"sunsetEnd"`
	SunriseStart    int     `json:"sunriseStart"`
	SunriseEnd      int     `json:"sunriseEnd"`
}

// Create configuration with default values
func New() Configuration {
	return Configuration{
		Verbose:         false,
		Renderer:        "default",
		Mode:            "default",
		GpioPin:         18,
		LedCount:        50,
		Brightness:      1,
		AmpLimit:        0,
		FramesPerSecond: 60,
	}
}

func Load() (Configuration, error) {
	config := New()

	// Parse command line to get optional config path
	flags := flag.NewFlagSet("xyzmas", flag.ExitOnError)
	config.AddFlags(flags)
	flags.Parse(os.Args[1:])

	// Read config from that path, if not the default.
	err := config.Read()
	if err != nil {
		return config, err
	}

	// Re-parse args to override read config
	flags = flag.NewFlagSet("xyzmas", flag.ExitOnError)
	config.AddFlags(flags)
	flags.Parse(os.Args[1:])

	return config, nil
}

func (config *Configuration) AddFlags(flags *flag.FlagSet) {
	flags.StringVar(&config.ConfigPath, "config", "config.json", "Path to configuration file")
	flags.StringVar(&config.Renderer, "renderer", config.Renderer, "Name of renderer: 'ws281x' (default when available), 'console', or 'null'")
	flags.StringVar(&config.Mode, "mode", config.Mode, "Name of mode: 'random', 'rgb' or 'id'")
	flags.BoolVar(&config.Benchmark, "benchmark", config.Benchmark, "Whether to print the effective refresh rate")
	flags.BoolVar(&config.Verbose, "verbose", config.Verbose, "Print what effect is starting")
	flags.IntVar(&config.GpioPin, "gpio", config.GpioPin, "GPIO pin")
	flags.IntVar(&config.LedCount, "count", config.LedCount, "Number of LEDs")
	flags.Float64Var(&config.Brightness, "brightness", config.Brightness, "Brightness (0..1)")
	flags.Float64Var(&config.AmpLimit, "limit", config.AmpLimit, "Limit, in amps")
	flags.IntVar(&config.FramesPerSecond, "fps", config.FramesPerSecond, "Frames per second")
	flags.BoolVar(&config.NightEnabled, "night", config.NightEnabled, "Whether to adjust brightness and disabled intense effects after sunset")
}

func (config *Configuration) Read() error {
	configFile, err := os.Open(config.ConfigPath)
	if err != nil {
		return err
	}

	configJson, err := io.ReadAll(configFile)
	if err != nil {
		return err
	}

	return json.Unmarshal(configJson, config)
}
