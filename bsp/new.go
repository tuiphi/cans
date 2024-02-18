package bsp

type Option func(*State)

func WithBranch(branch Branch) Option {
	return func(state *State) {
		state.branch = branch
	}
}

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
