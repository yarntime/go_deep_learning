package main

import (
	"fmt"
	"github.com/yarntime/go_deep_learning/linerunit"
)

type LinerActivator struct{}

func (l *LinerActivator) Activator(x float64) float64 {
	return x
}

var lineData = [][]float64{{5}, {3}, {8}, {1.4}, {10.1}}
var lineValue = []float64{5500, 2300, 7600, 1800, 11400}

func main() {
	l := linerunit.NewLinerUnit(1, &LinerActivator{})
	l.Fit(lineData, lineValue, 0.01, 50)
	predictedValue, _ := l.Predict([]float64{15})
	fmt.Printf("%f\n", predictedValue)
	fmt.Println(l.Weights)
	fmt.Println(l.Bias)
}
