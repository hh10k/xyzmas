package effect

func Pipe(effects ...Effect) Effect {
	return func(state State) State {
		for _, e := range effects {
			state = e(state)
		}
		return state
	}
}
