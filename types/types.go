package types

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
