package types

import "math"

func Sigmod(x float64) float64 {
	return 1 / (1 + math.Exp(-x))
}
