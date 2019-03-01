package perceptron

import (
	"github.com/yarntime/go_deep_learning/interfaces"
	"github.com/pkg/errors"
	"fmt"
)

type Perceptron struct {
	Number int
	Activator interfaces.Activator
	Weights []float64
	Bias    float64
}

func NewPerceptron(number int, activator interfaces.Activator) *Perceptron {
	return &Perceptron{
		Number: number,
		Activator: activator,
		Weights: make([]float64, number),
		Bias: 0.0,
	}
}

func (p *Perceptron) Fit(input [][]float64, label []float64, learningRate float64, iteration int) error {
	if len(input) == 0 {
		return errors.New("input data can not be null.")
	}
	if len(label) == 0 {
		return errors.New("label data can not be null.")
	}
	if len(input) != len(label) {
		return errors.New("input and label should have the same length.")
	}
	for i := 0; i < iteration; i++ {
		err := p.fitOne(input, label, learningRate)
		if err != nil {
			return err
		}
	}
	return nil
}

func (p *Perceptron) fitOne(input [][]float64, label []float64, learningRate float64) error {
	for i := 0; i < len(input); i++ {
		predictedValue, err := p.Predict(input[i])
		if err != nil {
			return err
		}
		p.updateWeights(predictedValue, label[i], learningRate, input[i])
	}
	return nil
}

func (p *Perceptron) updateWeights(predictedValue float64, labelValue float64, learningRate float64, input []float64) {
	delta := labelValue - predictedValue
	for i := 0; i < p.Number; i++ {
		p.Weights[i] += delta * learningRate * input[i]
	}
	p.Bias += delta * learningRate
}

func (p *Perceptron) Predict(input []float64) (float64, error) {
	if p.Number != len(input) {
		return 0.0, errors.New(fmt.Sprintf("Wrong input data length. perceptron number(%d), input data length(%d)", p.Number, len(input)))
	}
	result := 0.0
	for i := 0; i < p.Number; i++ {
		result += input[i] * p.Weights[i]
	}
	return p.Activator.Activator(result + p.Bias), nil
}
