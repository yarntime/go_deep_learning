package main

import (
	"fmt"
	"github.com/yarntime/go_deep_learning/perceptron"
)

type SimpleActivator struct{}

func (s *SimpleActivator) Activator(x float64) float64 {
	if x > 0 {
		return 1.0
	} else {
		return 0.0
	}
}

var trainData = [][]float64{{1, 1}, {0, 0}, {1, 0}, {0, 1}}
var label = []float64{1, 0, 0, 0}

func main() {
	p := perceptron.NewPerceptron(2, &SimpleActivator{})
	p.Fit(trainData, label, 0.01, 100)
	predictedValue, _ := p.Predict([]float64{0, 1})
	fmt.Printf("%f\n", predictedValue)
	predictedValue, _ = p.Predict([]float64{1, 1})
	fmt.Printf("%f\n", predictedValue)
	predictedValue, _ = p.Predict([]float64{1, 0})
	fmt.Printf("%f\n", predictedValue)
	predictedValue, _ = p.Predict([]float64{0, 0})
	fmt.Printf("%f\n", predictedValue)
	fmt.Println(p.Weights)
	fmt.Println(p.Bias)
}
