package color

func Fill(count int, colors RGBGenerator) []RGBA {
	c := make([]RGBA, count)
	for i := range c {
		c[i] = colors()
	}
	return c
}
