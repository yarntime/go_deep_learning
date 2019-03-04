package types

type Node struct {
	LayerIndex int
	NodeIndex  int
	DownStream []*Connection
	UpStream   []*Connection
	Output     float64
	Delta      float64
}

type Connection struct {
	UpStreamNode   *Node
	DownStreamNode *Node
	Weight         float64
	Gradient       float64
}

type Layer struct {
	LayerIndex int
	Nodes      []*Node
}

type Network struct {
	Connections []*Connection
	Layers      []*Layer
}
