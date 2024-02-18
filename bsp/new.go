package bsp

func New(options ...Option) *State {
	state := &State{
		branch: Branch{
			Direction: DirectionHorizontal,
			Ratio:     RatioEven,
		},
		keyMap: DefaultKeyMap(),
	}

	for _, option := range options {
		option(state)
	}

	return state
}
