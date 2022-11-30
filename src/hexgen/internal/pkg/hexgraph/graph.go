package hexgraph

import (
	"github.com/3vilM33pl3/hexrepo/src/hexcloud/pkg/hexgrid"
	"github.com/dominikbraun/graph"
	"strconv"
)

type Hex struct {
	ID int64
	X  int64
	Y  int64
	Z  int64
}

type Map struct {
	Name       string
	HexGraph   *graph.Graph[string, Hex]
	RiverGraph *graph.Graph[string, Hex]
}

func NewMap(name string) Map {
	hexHash := func(h Hex) string {
		return strconv.FormatInt(h.ID, 10)
	}

	hexG := graph.New(hexHash)
	_ = hexG.AddVertex(Hex{ID: hexgrid.Pair(0, 0), X: 0, Y: 0, Z: 0})

	riverG := graph.New(hexHash, graph.Directed())
	_ = riverG.AddVertex(Hex{ID: hexgrid.Pair(0, 0), X: 0, Y: 0, Z: 0})

	return Map{name, &hexG, &riverG}
}
