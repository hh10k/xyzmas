package transition

import (
	"math"

	"kurt.blackwell.id.au/xyzmas/color"
)

// Stack stacks pixels from the right down to the left of the string
//
// spacing is a value 0..1 for how far the pixels are spaced apart as they move in
func Stack(spacing float32) Transition {
	return func(base, stacked []color.RGBA, t float32) []color.RGBA {
		c := make([]color.RGBA, len(base))

		// len(c) 5, spacing 1 => length 21
		//                      #####|----|----|----|----|
		// #####|----|----|----|----|
		// len(c) 5, spacing 0 => length 9
		//      #####|||||
		// #####|||||
		// len(c) 4, spacing 0 => length 7
		//     ####||||
		// ####||||
		// len(c) 4, spacing 1 => length 13
		//              ####|---|---|---|
		// ####|---|---|---|

		count := float32(len(c))
		countBetween := spacing * (count - 1)
		countStride := countBetween + 1
		length := count + countBetween*(count-1)
		tx := int(math.Floor(float64((t*length - count - 1) / countBetween)))

		for i := range c {
			if i <= tx {
				c[i] = stacked[i]
				continue
			}
			ix := float32(i) - count + t*length
			ci := int(math.Floor(float64(ix / countStride)))
			if ci < 0 || (ix-float32(ci)*countStride) >= 1 {
				c[i] = base[i]
			} else {
				c[i] = stacked[ci]
			}
		}

		return c
	}
}
