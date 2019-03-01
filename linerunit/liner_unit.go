package linerunit

import (
	"github.com/yarntime/go_deep_learning/interfaces"
	"github.com/yarntime/go_deep_learning/perceptron"
)

type LinerUnit struct {
	*perceptron.Perceptron
}

func NewLinerUnit(number int, activator interfaces.Activator) *LinerUnit {
	return &LinerUnit{
		perceptron.NewPerceptron(number, activator),
	}
}
