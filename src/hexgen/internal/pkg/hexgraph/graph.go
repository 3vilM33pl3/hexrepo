package hexgraph

import (
	"github.com/3vilM33pl3/hexrepo/src/hexcloud/pkg/hexgrid"
	"github.com/dominikbraun/graph"
	"github.com/dominikbraun/graph/draw"
	"os"
)

type Hex struct {
	ID int64
	X  int64
	Y  int64
	Z  int64
}

type Map struct {
	Name       string
	HexGraph   graph.Graph[int64, Hex]
	RiverGraph graph.Graph[int64, Hex]
}

func NewMap(name string) Map {
	hexHash := func(h Hex) int64 {
		return h.ID
	}

	hexG := graph.New(hexHash)
	riverG := graph.New(hexHash, graph.Directed())

	return Map{name, hexG, riverG}
}

func (m *Map) Generate(size int) {

	m.HexGraph.AddVertex(Hex{ID: hexgrid.Pair(0, 0), X: 0, Y: 0, Z: 0})
	m.RiverGraph.AddVertex(Hex{ID: hexgrid.Pair(0, 0), X: 0, Y: 0, Z: 0})

	for i := 1; i <= size; i++ {
		hexes := hexgrid.Ring(hexgrid.NewHexXYZ(0, 0, 0), int64(i))

		for _, hex := range hexes {
			m.HexGraph.AddVertex(Hex{ID: hexgrid.Pair(hex.X, hex.Y), X: hex.X, Y: hex.Y, Z: hex.Z})
		}

		for _, hex := range hexes {
			neighbours := hexgrid.Ring(hexgrid.NewHexXYZ(hex.X, hex.Y, -hex.X-hex.Y), 1)
			for _, neighbour := range neighbours {
				nb, err := m.HexGraph.Vertex(hexgrid.Pair(neighbour.X, neighbour.Y))
				if err != nil {
					continue
				}
				m.HexGraph.AddEdge(hexgrid.Pair(hex.X, hex.Y), nb.ID)
			}
		}
	}

}

func (m *Map) DOT() {
	file, _ := os.Create("./mygraph.gv")
	_ = draw.DOT(m.HexGraph, file)
}
