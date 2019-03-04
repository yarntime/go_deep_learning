package types

type Node struct {
	LayerIndex int
	NodeIndex  int
	DownStream []*Connection
	UpStream   []*Connection
	Output     float64
	Delta      float64
}

func NewNode(layerIndex int, nodeIndex int) *Node {
	return &Node{
		LayerIndex: layerIndex,
		NodeIndex:  nodeIndex,
		DownStream: make([]*Connection, 0),
		UpStream:   make([]*Connection, 0),
	}
}

func (n *Node) AddDownStream(conn *Connection) {
	n.DownStream = append(n.DownStream, conn)
}

func (n *Node) AddUpStream(conn *Connection) {
	n.UpStream = append(n.UpStream, conn)
}

func (n *Node) SetOutput(output float64) {
	n.Output = output
}

func (n *Node) CalcOutPut() {
	result := 0.0
	for i := 0; i < len(n.UpStream); i++ {
		result += n.UpStream[i].UpStreamNode.Output * n.UpStream[i].Weight
	}
	n.Output = Sigmod(result)
}

func (n *Node) CalcHiddenLayerDelta() {
	result := 0.0
	for i := 0; i < len(n.DownStream); i++ {
		result += n.DownStream[i].DownStreamNode.Delta * n.DownStream[i].Weight
	}
	n.Delta = n.Output * (1 - n.Output) * result
}

func (n *Node) CalcOutPutLayerDelta(label float64) {
	n.Delta = n.Output * (1 - n.Output) * (label - n.Output)
}
