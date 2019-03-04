package types

type Connection struct {
	UpStreamNode   *Node
	DownStreamNode *Node
	Weight         float64
	Gradient       float64
}

type Network struct {
	Connections []*Connection
	Layers      []*Layer
}
