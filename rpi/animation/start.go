package animation

import (
	"fmt"
	"math"
	"os"
	"time"

	"kurt.blackwell.id.au/xyzmas/color"
	"kurt.blackwell.id.au/xyzmas/configuration"
	"kurt.blackwell.id.au/xyzmas/effect"
	"kurt.blackwell.id.au/xyzmas/renderer"
	"kurt.blackwell.id.au/xyzmas/transition"
	"kurt.blackwell.id.au/xyzmas/value"
)

func Start(config configuration.Configuration, stop <-chan bool) error {
	ren, err := renderer.New(config)
	if err != nil {
		return err
	}

	var animation effect.Effect
	switch config.Mode {
	case "default", "random":
		animation = Random(config)
	case "rgb":
		animation = CycleRGB(config.LedCount)
	case "test":
		// animation = effect.Pixels(pixel.Switch(
		// 	color.RandomHclHues(0, 360, c, c, l, l),
		// 	ease.InOutCubic,
		// 	time.Second*1,
		// 	time.Second*1,
		// 	time.Second*2,
		// 	time.Second*3,
		// ))
		color1, color2 := color.RandomPair()

		animation = effect.HueShift(
			//value.InertialVelocity(value.InertialVelocityParams{
			// 	Acceleration: 1.0,
			// 	VelocityDrag: 0.95,
			// })
			// value.InertialAcceleration(value.InertialAccelerationParams{
			// 	AccelerationChange: 0.1,
			// 	VelocityDrag:       0.95,
			// 	AccelerationDrag:   0.95,
			// }),
			value.Linear(360/60),
			effect.Splotch(
				time.Second*5,
				time.Second*1,
				effect.Pipe(
					effect.Fill(color.SinWave(color1, color2, float64(math.Pi*2.0/float64(config.LedCount)))),
					effect.Slide(-50.0),
				),
				effect.Pipe(
					effect.Fill(color.SinWave(color1, color2, float64(math.Pi*2.0/float64(config.LedCount)))),
					effect.Slide(50.0),
				),
			))

	case "white":
		animation = effect.Fill(color.Constant(color.White))
	case "id":
		animation = Identification(config.LedCount, time.Duration(config.IdMode.StepIntervalMs)*time.Millisecond)
	default:
		return fmt.Errorf("unknown mode %s", config.Mode)
	}

	loop(ren, animation, config, stop)
	return nil
}

func loop(ren renderer.Renderer, animation effect.Effect, config configuration.Configuration, stop <-chan bool) {
	var frameDuration = time.Second / time.Duration(config.FramesPerSecond)

	now := time.Now()
	nextFrame := now
	startTime := now
	stopping := false
	black := make([]color.RGBA, config.LedCount)

	// For fading in and out when starting and stopping
	fadeState := transition.NewState(0)
	fadeState.FadeTo(1, 0, time.Second)

	// Combines stop/start fade and night time dimming
	rootFade := func(now time.Duration) float32 {
		var night = GetNightFactor(startTime.Add(now), config)
		var nightBrightness = (1.0 - night) + (night * config.NightBrightness)

		return fadeState.ValueAt(now) * nightBrightness
	}

	// Combines top level animation with dimming
	rootAnimation := effect.Transition(
		transition.Fade(), rootFade,
		effect.Fill(color.Constant(color.Black)),
		animation,
	)

	onStop := func() {
		fadeState.FadeTo(0, now.Sub(startTime), time.Second)
		stopping = true
	}

	onFrame := func(now time.Time) {

		state := rootAnimation(effect.State{
			Colors: black,
			Time:   now.Sub(startTime),
		})

		if err := ren.Render(state.Colors); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
			return
		}

		nextFrame = nextFrame.Add(frameDuration)
	}

	for !stopping || !fadeState.IsCompleteAt(nextFrame.Sub(startTime)) {
		delay := nextFrame.Sub(now)
		if delay < time.Millisecond {
			select {
			case <-stop:
				onStop()
			default:
				onFrame(nextFrame)
			}
		} else {
			select {
			case <-stop:
				onStop()
			case <-time.After(delay):
				onFrame(time.Now())
			}
		}

		now = time.Now()
		// If lagging skip frames
		if nextFrame.Before(now) {
			nextFrame = now
		}
	}

	// Clear before exiting
	ren.Render(black)
}
