package color

import (
	"fmt"
	"testing"
)

func TestLerpRGBGamma(t *testing.T) {
	var tests = []struct {
		c1, c2 RGBA
		t      int32
		want   RGBA
	}{
		{NewRGBA(0, 63, 127, 0), NewRGBA(255, 255, 255, 255), 0, NewRGBA(0, 63, 127, 0)},
		{NewRGBA(0, 0, 0, 0), NewRGBA(63, 127, 255, 255), 256, NewRGBA(63, 127, 255, 255)},
		{NewRGBA(0, 0, 0, 0), NewRGBA(255, 255, 255, 255), 128, NewRGBA(186, 186, 186, 127)},
	}

	// fmt.Printf("%+v\n", toGamma)
	// fmt.Printf("%+v\n", fromGamma)

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			ans := LerpRGBGamma(tt.c1, tt.c2, tt.t)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}
