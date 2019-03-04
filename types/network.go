package types

import "errors"

type Network struct {
	Connections []*Connection
	Layers      []*Layer
}

func NewNetwork(layer []int) *Network {
	connections := make([]*Connection, 0)
	layers := make([]*Layer, 0)
	for i := 0; i < len(layer); i++ {
		layers = append(layers, NewLayer(i, layer[i]))
	}
	for i := 0; i < len(layer)-1; i++ {
		for j := 0; j < len(layers[i].Nodes); j++ {
			for k := 0; k < len(layers[i+1].Nodes); k++ {
				conn := NewConnection(layers[i].Nodes[j], layers[i+1].Nodes[k])
				connections = append(connections, conn)
				layers[i].Nodes[j].DownStream = append(layers[i].Nodes[j].DownStream, conn)
				layers[i+1].Nodes[k].UpStream = append(layers[i+1].Nodes[k].UpStream, conn)
			}
		}
	}
	return &Network{
		Connections: connections,
		Layers:      layers,
	}
}

func (n *Network) Train(labels [][]float64, data [][]float64, learningRate float64, iteration int) error {
	if len(data) == 0 {
		return errors.New("input data can not be null.")
	}
	if len(labels) == 0 {
		return errors.New("label data can not be null.")
	}
	if len(data) != len(labels) {
		return errors.New("input and label should have the same length.")
	}
	for i := 0; i < iteration; i++ {
		for j := 0; j < len(labels); j++ {
			n.trainOne(labels[j], data[j], learningRate)
		}
	}
	return nil
}

func (n *Network) trainOne(labels []float64, data []float64, learningRate float64) {
	n.Predict(data)
	n.calcDelta(labels)
	n.updateWeight(learningRate)
}

func (n *Network) calcDelta(label []float64) {
	outputNodes := n.Layers[len(n.Layers)-1].Nodes
	for i := 0; i < len(label); i++ {
		outputNodes[i].CalcOutPutLayerDelta(label[i])
	}
	for i := len(n.Layers) - 2; i >= 1; i-- {
		for j := 0; j < len(n.Layers[i].Nodes); j++ {
			n.Layers[i].Nodes[j].CalcHiddenLayerDelta()
		}
	}
}

func (n *Network) updateWeight(learningRate float64) {
	for i := 0; i < len(n.Layers)-1; i++ {
		for j := 0; j < len(n.Layers[i].Nodes); j++ {
			for k := 0; k < len(n.Layers[i].Nodes[j].DownStream); k++ {
				n.Layers[i].Nodes[j].DownStream[k].UpdateWeight(learningRate)
			}
		}
	}
}

func (n *Network) calcGradient() {
	for i := 0; i < len(n.Layers)-1; i++ {
		for j := 0; j < len(n.Layers[i].Nodes); j++ {
			for k := 0; k < len(n.Layers[i].Nodes[j].DownStream); k++ {
				n.Layers[i].Nodes[j].DownStream[k].CalcGradient()
			}
		}
	}
}

func (n *Network) Predict(data []float64) []float64 {
	n.Layers[0].SetOutPut(data)
	for i := 1; i < len(n.Layers); i++ {
		n.Layers[i].CalcOutPut()
	}
	result := make([]float64, 0)
	outputNodes := n.Layers[len(n.Layers)-1].Nodes
	for i := 0; i < len(outputNodes); i++ {
		result = append(result, outputNodes[i].Output)
	}
	return result
}
