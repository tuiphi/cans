package bsp

import "math"

type Ratio float64

const (
	RatioEqual     Ratio = 1. / 2
	RatioTwoThirds Ratio = 2. / 3
	RatioOneThird  Ratio = 1. / 3
	RatioGolden    Ratio = 1 / math.Phi
)
