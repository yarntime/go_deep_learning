package types

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

type Connection struct {
	UpStreamNode   *Node
	DownStreamNode *Node
	Weight         float64
	Gradient       float64
}

func NewConnection(upStreamNode *Node, downloadStreamNode *Node) *Connection {
	return &Connection{
		UpStreamNode:   upStreamNode,
		DownStreamNode: downloadStreamNode,
		Weight:         randOne(),
	}
}

func (c *Connection) CalcGradient() {
	c.Gradient = c.DownStreamNode.Delta * c.UpStreamNode.Output
}

func (c *Connection) UpdateWeight(rate float64) {
	c.CalcGradient()
	c.Weight = rate * c.Gradient
}

func randOne() float64 {
	flag := rand.Intn(2)
	number := rand.Float64()
	if flag == 0 {
		return -number
	}
	return number
}
