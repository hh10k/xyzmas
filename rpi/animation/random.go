package animation

import (
	"fmt"
	"time"

	"kurt.blackwell.id.au/xyzmas/configuration"
	"kurt.blackwell.id.au/xyzmas/effect"
	"kurt.blackwell.id.au/xyzmas/maths"
	"kurt.blackwell.id.au/xyzmas/transition"
)

func Random(config configuration.Configuration) effect.Effect {
	anims := make([]*Animation, len(Animations))
	for i := range Animations {
		anims[i] = &Animations[i]
	}

	// Start at the end so that next generated number will shuffle the list
	i := len(anims)

	randomEffects := func(now time.Duration) (effect.Effect, time.Duration) {
		var night = GetNightFactor(time.Now(), config)

		for {
			// Go to next, re-shuffle list if required
			i++
			if i >= len(anims) {
				i = 0
				maths.Rand.Shuffle(len(anims), func(x int, y int) {
					anims[x], anims[y] = anims[y], anims[x]
				})
			}

			// Use this if not disallowed by night mode
			if anims[i].Policy.Night || night < maths.Rand.Float32() {
				break
			}
		}

		if config.Verbose {
			fmt.Println(anims[i].Name)
		}

		duration := anims[i].Policy.Duration
		config := effect.Context{
			LedCount:  config.LedCount,
			StartTime: now,
			Duration:  duration,
		}
		return anims[i].NewEffect(config), duration
	}

	randomTransitions := transition.Random()

	return effect.Sequence(randomEffects, randomTransitions)
}
