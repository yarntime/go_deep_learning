package types

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
