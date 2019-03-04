package types

type Layer struct {
	LayerIndex int
	Nodes      []*Node
}

func NewLayer(layerIndex int, nodeCount int) *Layer {
	nodes := make([]*Node, 0)
	for i := 0; i < nodeCount; i++ {
		nodes = append(nodes, NewNode(layerIndex, i))
	}
	return &Layer{
		LayerIndex: layerIndex,
		Nodes:      nodes,
	}
}

func (l *Layer) SetOutPut(outputs []float64) {
	for i := 0; i < len(l.Nodes); i++ {
		l.Nodes[i].SetOutput(outputs[i])
	}
}

func (l *Layer) CalcOutPut() {
	for i := 0; i < len(l.Nodes); i++ {
		l.Nodes[i].CalcOutPut()
	}
}
