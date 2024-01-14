package color

func Constant(value RGBA) RGBGenerator {
	return func() RGBA {
		return value
	}
}

func Cycle(colors ...RGBGenerator) RGBGenerator {
	i := -1
	return func() RGBA {
		i = (i + 1) % len(colors)
		return colors[i]()
	}
}

func CycleRGBA(colors ...RGBA) RGBGenerator {
	i := -1
	return func() RGBA {
		i = (i + 1) % len(colors)
		return colors[i]
	}
}
