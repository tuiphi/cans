package window

import "github.com/tuiphy/soda"

type Window interface {
	State() soda.State
	ViewOnly() bool
}

type ExtendedState interface {
	soda.State

	SetActive(active bool)
}
