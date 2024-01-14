package effect

import (
	"time"

	"kurt.blackwell.id.au/xyzmas/color"
	"kurt.blackwell.id.au/xyzmas/maths"
	"kurt.blackwell.id.au/xyzmas/transition"
)

// Randomly assigns pixels to the given effects.  After holdDuration it will re-randomly assign pixels to effects
// and transition to the new arrangement.
func Splotch(
	holdDuration time.Duration,
	transitionDuration time.Duration,
	effects ...Effect,
) Effect {
	// Indexes into effects lists
	var fromEffects []int
	var toEffects []int
	var effectStates = make([]State, len(effects))
	var fade transition.State
	fade.From = 0
	fade.To = 1

	return func(state State) State {
		if fromEffects == nil || len(fromEffects) != len(state.Colors) {
			// Initial list
			fromEffects = make([]int, len(state.Colors))
			toEffects = make([]int, len(state.Colors))
			for i := range fromEffects {
				fromEffects[i] = maths.Rand.Intn(len(effects))
				toEffects[i] = maths.Rand.Intn(len(effects))
			}
			fade.FromTime = state.Time + holdDuration
			fade.ToTime = fade.FromTime + transitionDuration
		}

		// Start new fade period when needed
		for state.Time >= fade.ToTime {
			fade.FromTime = fade.ToTime + holdDuration
			fade.ToTime = fade.FromTime + transitionDuration

			fromEffects, toEffects = toEffects, fromEffects
			for i := range toEffects {
				toEffects[i] = maths.Rand.Intn(len(effects))
			}
		}
		t := fade.ValueAt(state.Time)

		// Update states for each of the effects
		for i := range effectStates {
			effectStates[i] = effects[i](state)
		}

		// Show the underlying effect, depending on which one is visible
		state.Colors = make([]color.RGBA, len(state.Colors))
		for i := range state.Colors {
			fromEffectIndex := fromEffects[i]
			toEffectIndex := toEffects[i]
			if fromEffectIndex == toEffectIndex {
				state.Colors[i] = effectStates[fromEffectIndex].Colors[i]
			} else {
				state.Colors[i] = color.LerpRGB(
					effectStates[fromEffectIndex].Colors[i],
					effectStates[toEffectIndex].Colors[i],
					t,
				)
			}
		}
		return state
	}
}
