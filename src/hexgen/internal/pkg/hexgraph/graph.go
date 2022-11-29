package hexgraph

import (
	"github.com/dominikbraun/graph"
	"strconv"
)

type Hex struct {
	ID int
	X  int
	Y  int
	Z  int
}

func NewMap() graph.Graph[string, Hex] {
	hexHash := func(h Hex) string {
		return strconv.Itoa(h.ID)
	}

	g := graph.New(hexHash, graph.Directed(), graph.Acyclic())
	_ = g.AddVertex(Hex{ID: 1, X: 0, Y: 0, Z: 0})

	return g
}
