package animation

import (
	"fmt"
	"time"

	"kurt.blackwell.id.au/xyzmas/color"
	"kurt.blackwell.id.au/xyzmas/effect"
	"kurt.blackwell.id.au/xyzmas/maths"
)

type lightCodes struct {
	/** Length of binary sequence. */
	codeLength int
	/** For each light index, return the code it will repeat. */
	codesByIndex []int
	/** Maps a code to the actual light index.  The code map appear multiple times with different rotations. */
	indexByCodes []int
}

// How many times we need to repeat the code pattern to get back to where we start.
// Due to the odd/even nature, this must be a multiple of 2 to get back to where we started.
const patternRepeats = 2

var colors = [3]color.RGBA{color.Red, color.Green, color.Blue}
var lowHigh = [3][2]int{
	{1, 2},
	{0, 2},
	{0, 1},
}

func Identification(count int, interval time.Duration) effect.Effect {
	lightCodes := generateLightCodes(count)
	fmt.Printf("Generated id codes with length %d\n", lightCodes.codeLength)

	patternLength := lightCodes.codeLength * patternRepeats
	lightColors := make([]color.RGBA, patternLength*count)
	for i := 0; i < count; i++ {
		getColorPattern(
			lightColors[i*patternLength:(i+1)*patternLength],
			lightCodes.codesByIndex[i],
			lightCodes.codeLength,
		)
	}

	return func(state effect.State) effect.State {
		t := int(state.Time/interval) % patternLength
		c := make([]color.RGBA, count)

		for i := range c {
			c[i] = lightColors[i*patternLength+t]
		}

		state.Colors = c
		return state
	}
}

// Fill the dest with the repeating pattern for the code and code length
// The dest should be patternRepeats * length long.
func getColorPattern(dest []color.RGBA, code int, length int) {
	d := 0
	c := 0

	// Generate at least a length worth first so that it 'settles' into the final repeating pattern.
	shift := length + maths.Rand.Intn(len(dest))

	i := 1
	for d < len(dest) {
		c = lowHigh[c][(code>>(length-i))&1]

		if shift > 0 {
			shift--
		} else {
			dest[d] = colors[c]
			d++
		}

		i++
		if i > length {
			i = 1
		}
	}
}

func generateLightCodes(count int) lightCodes {
	// Try longer lengths until we have enough.
	for codeLength := 2; ; codeLength++ {
		codesByIndex := make([]int, 0)
		indexByCodes := make([]int, 1<<codeLength)

		// The first code (0) comes for free. Since it won't be used by any others
		// we can use use indexByCodes[code] != 0 to test wether that index can be used.

		for code := 1; code < len(indexByCodes); code++ {
			if indexByCodes[code] != 0 {
				// Code already used
				continue
			}

			// Rotate number to record all versions
			n := code
			for rotateI := 0; rotateI < codeLength; rotateI++ {
				indexByCodes[n] = len(codesByIndex)
				n = (n >> 1) + (n % 2 << (codeLength - 1))
			}

			codesByIndex = append(codesByIndex, code)
			if len(codesByIndex) >= count {
				// Enough codes have been generated.
				return lightCodes{
					codeLength,
					codesByIndex,
					indexByCodes,
				}
			}
		}
	}
}
