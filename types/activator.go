package types

import "math"

func Sigmod(x float64) float64 {
	return 1 / (1 + math.Exp(-x))
}

func Relu(x float64) float64 {
	if x < 0 {
		return 0
	}
	return x
}
