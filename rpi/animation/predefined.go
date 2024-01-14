package animation

import (
	"math"
	"time"

	"github.com/fogleman/ease"
	"github.com/lucasb-eyer/go-colorful"
	"kurt.blackwell.id.au/xyzmas/color"
	"kurt.blackwell.id.au/xyzmas/effect"
	"kurt.blackwell.id.au/xyzmas/maths"
	"kurt.blackwell.id.au/xyzmas/pixel"
	"kurt.blackwell.id.au/xyzmas/transition"
	"kurt.blackwell.id.au/xyzmas/value"
)

var long = AnimationPolicy{
	Duration: 50 * time.Second,
	Night:    true,
}
var longDaytime = AnimationPolicy{
	Duration: 50 * time.Second,
	Night:    false,
}
var medium = AnimationPolicy{
	Duration: 30 * time.Second,
	Night:    true,
}
var mediumDaytime = AnimationPolicy{
	Duration: 30 * time.Second,
	Night:    false,
}
var short = AnimationPolicy{
	Duration: 15 * time.Second,
	Night:    true,
}
var shortDaytime = AnimationPolicy{
	Duration: 15 * time.Second,
	Night:    false,
}
var sparkle = AnimationPolicy{
	Duration: 5 * time.Second,
	Night:    false,
}

type AnimationPolicy struct {
	Duration time.Duration
	Night    bool
}

type Animation struct {
	Name      string
	Policy    AnimationPolicy
	NewEffect func(context effect.Context) effect.Effect
}

func NewAnimation(
	name string,
	policy AnimationPolicy,
	newEffect func(context effect.Context) effect.Effect,
) Animation {
	return Animation{
		Name:      name,
		Policy:    policy,
		NewEffect: newEffect,
	}
}

var Animations = []Animation{
	NewAnimation("white", short,
		func(context effect.Context) effect.Effect {
			return effect.Fill(color.Constant(color.White))
		}),
	NewAnimation("sliding-random", long,
		func(context effect.Context) effect.Effect {
			return effect.Pipe(
				effect.Fill(color.RandomRGBs(color.Black, color.White)),
				effect.Slide(1.0),
			)
		}),
	NewAnimation("sliding-rainbow", long,
		func(context effect.Context) effect.Effect {
			c, l := color.RandomChromaLuminance()
			return effect.Pipe(
				effect.Fill(color.CycleHclHues(360/float64(context.LedCount), c, c, l, l)),
				effect.Slide(50.0),
			)
		}),
	NewAnimation("sliding-duo-lava", long,
		func(context effect.Context) effect.Effect {
			color1, color2 := color.RandomPair()

			return effect.Splotch(
				time.Second*5,
				time.Second*1,
				effect.Pipe(
					effect.Fill(color.SinWave(color1, color2, float64(math.Pi*2.0/float64(context.LedCount)))),
					effect.Slide(-50.0),
				),
				effect.Pipe(
					effect.Fill(color.SinWave(color1, color2, float64(math.Pi*2.0/float64(context.LedCount)))),
					effect.Slide(50.0),
				),
			)
		}),
	NewAnimation("sliding-duo-lava-hue", long,
		func(context effect.Context) effect.Effect {
			color1, color2 := color.RandomPair()

			return effect.HueShift(
				// value.InertialVelocity(value.InertialVelocityParams{
				// 	Acceleration: 1.0,
				// 	VelocityDrag: 0.95,
				// }),
				value.Linear(360/60),
				effect.Splotch(
					time.Second*5,
					time.Second*1,
					effect.Pipe(
						effect.Fill(color.SinWave(color1, color2, float64(math.Pi*2.0/float64(context.LedCount)))),
						effect.Slide(-50.0),
					),
					effect.Pipe(
						effect.Fill(color.SinWave(color1, color2, float64(math.Pi*2.0/float64(context.LedCount)))),
						effect.Slide(50.0),
					),
				))
		}),
	NewAnimation("pulse-random", medium,
		func(context effect.Context) effect.Effect {
			return effect.Pulse(
				color.RandomHclHues(0, 360, 1, 1, 0.5, 0.5),
				2*time.Second, 8*time.Second,
				20)
		}),
	NewAnimation("pulse-cycle", medium,
		func(context effect.Context) effect.Effect {
			return effect.Pulse(
				color.CycleHclHues(1, 1, 1, 0.5, 0.5),
				2*time.Second, 8*time.Second,
				20)
		}),
	NewAnimation("pulse-rainbow", medium,
		func(context effect.Context) effect.Effect {
			return effect.Pixels(pixel.RainbowPulse(
				3*time.Second, 6*time.Second,
				3,
				5))
		}),
	NewAnimation("pulse-rainbow-inverse", medium,
		func(context effect.Context) effect.Effect {
			return effect.Pipe(
				effect.Fill(color.Constant(color.White)),
				effect.Pixels(pixel.RainbowPulse(
					5*time.Second, 10*time.Second,
					0.5,
					5)),
			)
		}),
	NewAnimation("switch-colourful", medium,
		func(context effect.Context) effect.Effect {
			c, l := color.RandomChromaLuminance()
			return effect.Pixels(pixel.Switch(
				color.RandomHclHues(0, 360, c, c, l, l),
				ease.InOutCubic,
				time.Second*1,
				time.Second*1,
				time.Second*2,
				time.Second*3,
			))
		}),
	NewAnimation("switch-cycle-wobble", long,
		func(context effect.Context) effect.Effect {
			c, l := color.RandomChromaLuminance()
			return effect.TimingWobble(
				time.Second*2,
				time.Millisecond*500,
				1,
				effect.Pixels(pixel.Switch(
					color.CycleHclHues(360/3/float64(context.LedCount), c, c, l, l),
					ease.InOutCubic,
					time.Second*1,
					time.Second*1,
					time.Millisecond*1000,
					time.Millisecond*1300,
				)),
			)
		}),
	NewAnimation("switch-cycle-inertia", long,
		func(context effect.Context) effect.Effect {
			return effect.Pixels(pixel.Switch(
				color.InertiaHclHues(0.5, 0.9, 0.5, 0.9),
				ease.InOutCubic,
				time.Second*1,
				time.Second*1,
				time.Millisecond*1000,
				time.Millisecond*1300,
			))
		}),
	// NewAnimation("flash-white-on-color", normal,
	// 	func(context effect.Context) effect.Effect {
	// 		return effect.Pipe(
	// 			effect.Fill(
	// 				color.Constant(color.NewColorfulRGB(colorful.Hsv(
	// 					maths.Rand.Float64()*360,
	// 					1,
	// 					0.5,
	// 				))),
	// 			),
	// 			effect.Flash(
	// 				color.Constant(color.White),
	// 				10000*time.Millisecond, 10000*time.Millisecond,
	// 				0*time.Second, 0*time.Second),
	// 		)
	// 	}),
	NewAnimation("flash", medium,
		func(context effect.Context) effect.Effect {
			hue1 := maths.Rand.Float64() * 360
			hue2 := hue1
			switch maths.Rand.Intn(2) {
			case 1:
				// Complementary colours
				hue2 = math.Mod(hue1+(360.0/3.0), 360)
			case 2:
				// Opposite colours
				hue2 = math.Mod(hue1+(360.0/2.0), 360)
			}

			return effect.Pipe(
				effect.Fill(
					color.Constant(color.NewColorfulRGB(colorful.Hsv(hue1, 1, 0.5))),
				),
				effect.Pulse(
					color.Constant(color.NewColorfulRGB(colorful.Hsv(hue2, 0.7, 0.8))),
					6*time.Second, 8*time.Second,
					20),
				effect.Flash(
					color.Constant(color.NewColorfulRGB(colorful.Hsv(hue2, 0.2, 1))),
					10000*time.Millisecond, 10000*time.Millisecond,
					0*time.Second, 0*time.Second),
			)
		}),
	NewAnimation("hues", long,
		func(context effect.Context) effect.Effect {
			hue := maths.Rand.Float64() * 360
			return effect.Pipe(
				effect.Fill(color.RandomHslHues(
					hue, hue+20,
					0.7, 1,
					0.4, 0.7,
				)),
				effect.Slide(2.0),
			)
		}),
	NewAnimation("gold", long,
		func(context effect.Context) effect.Effect {
			return effect.Pipe(
				effect.Fill(color.RandomHslHues(
					35, 45,
					0.7, 1,
					0.4, 0.7,
				)),
				effect.Slide(2.0),
				effect.Pulse(
					color.Constant(color.Black),
					2*time.Second, 8*time.Second,
					20),
			)
		}),
	NewAnimation("sparkle-white", sparkle,
		func(context effect.Context) effect.Effect {
			return effect.Sparkle(color.Constant(color.White), 0.025)
		}),
	NewAnimation("sparkle-rainbow", sparkle,
		func(context effect.Context) effect.Effect {
			return effect.Sparkle(color.RandomHslHues(0, 360, 1, 1, 0.5, 0.5), 0.05)
		}),
	NewAnimation("step-binary", mediumDaytime,
		func(context effect.Context) effect.Effect {
			hue := maths.Rand.Float64() * 360
			c := 0.6 + maths.Rand.Float64()*0.2
			l := 0.6
			rgb1 := color.NewColorfulRGB(colorful.Hsl(hue, c, l))
			rgb2 := color.NewColorfulRGB(colorful.Hsl(math.Mod(hue+180, 360), c, l))
			return effect.Timing(
				context, ease.InCubic,
				effect.Pipe(
					effect.Fill(
						color.Cycle(
							color.Constant(rgb1),
							color.Constant(rgb2),
						),
					),
					effect.Step(long.Duration/20, 1),
				),
			)
		}),
	// TODO: One like this, but the colours are slowly changing
	NewAnimation("cycle-triple", shortDaytime,
		func(context effect.Context) effect.Effect {
			color1, color2, color3 := color.RandomTriplet()
			return effect.Timing(
				context, ease.InCubic,
				effect.Cycle(
					long.Duration/40,
					effect.Fill(color.Constant(color1)),
					effect.Fill(color.Constant(color2)),
					effect.Fill(color.Constant(color3)),
				),
			)
		}),
	NewAnimation("splotch-triple", medium,
		func(context effect.Context) effect.Effect {
			color1, color2, color3 := color.RandomTriplet()

			return effect.Splotch(
				time.Second*1,
				time.Millisecond*500,
				effect.Fill(color.Constant(color1)),
				effect.Fill(color.Constant(color2)),
				effect.Fill(color.Constant(color3)),
			)
		}),
	NewAnimation("coen-pirate-ish", shortDaytime,
		func(context effect.Context) effect.Effect {
			return effect.Timing(
				context, ease.InQuart,
				effect.Pulse(
					color.CycleRGBA(
						color.Red,
						color.White,
						color.Magenta,
					),
					time.Second, 2*time.Second,
					20),
			)
		}),
	NewAnimation("pulse-along", mediumDaytime,
		func(context effect.Context) effect.Effect {
			first := true
			transitionDuration := 2 * time.Second
			return effect.Sequence(
				func(start time.Duration) (effect.Effect, time.Duration) {
					duration := transitionDuration * 2
					if first {
						first = false
						duration = transitionDuration
						return effect.Fill(color.Constant(color.Black)), duration
					}
					return effect.Transition(
						transition.Fade(),
						func(now time.Duration) float32 {
							t := min(1, float32(now-start)/float32(transitionDuration*13/10))
							return 1 - t*t
						},
						effect.Fill(color.Constant(color.Black)),
						effect.Fill(color.CycleHclHues(1, 1, 1, 0.5, 0.5)),
					), duration
				},
				func() (transition.Transition, time.Duration) {
					if maths.Rand.Float32() < 0.5 {
						return transition.Wipe(), transitionDuration
					} else {
						return transition.ReverseWipe(), transitionDuration
					}
				})
		},
	),
	// NewAnimation("stack", longDaytime,
	// 	func(context effect.Context) effect.Effect {
	// 		i := 0
	// 		transitionDuration := 120 * time.Second
	// 		return effect.Sequence(
	// 			func(start time.Duration) (effect.Effect, time.Duration) {
	// 				duration := transitionDuration * 2
	// 				if i == 0 {
	// 					i++
	// 					duration = transitionDuration
	// 					return effect.Fill(color.Constant(color.Black)), duration
	// 				}
	// 				i++
	// 				return effect.Transition(
	// 					transition.Fade(),
	// 					func(now time.Duration) float32 {
	// 						t := min(1, float32(now-start)/float32(transitionDuration))
	// 						return 1 - t
	// 					},
	// 					effect.Fill(color.Constant(color.Black)),
	// 					effect.Fill(color.RandomHclHues(0, 360, 1, 1, 0.5, 0.5)),
	// 				), duration
	// 			},
	// 			func() (transition.Transition, time.Duration) {
	// 				return transition.Stack(0.05), transitionDuration
	// 			})
	// 	},
	// ),
}
