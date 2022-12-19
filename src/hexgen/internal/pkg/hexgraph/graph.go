package hexgraph

import (
	"fmt"
	"github.com/3vilM33pl3/hexrepo/src/hexcloud/pkg/hexgrid"
	"gonum.org/v1/gonum/graph/encoding/dot"
	"gonum.org/v1/gonum/graph/simple"
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
	HexGraph   *simple.UndirectedGraph
	RiverGraph *simple.DirectedGraph
}

func NewMap(name string) Map {

	hexG := simple.NewUndirectedGraph()
	hexR := simple.NewDirectedGraph()

	return Map{name, hexG, hexR}
}

func (m *Map) Generate(size int) {

	m.HexGraph.AddNode(simple.Node(hexgrid.Pair(0, 0)))
	m.RiverGraph.AddNode(simple.Node(hexgrid.Pair(0, 0)))


	for i := 1; i <= size; i++ {
		hexes := hexgrid.Ring(hexgrid.NewHexXYZ(0, 0, 0), int64(i))

		var node simple.Node
		for _, hex := range hexes {
			node = simple.Node(hexgrid.Pair(hex.X, hex.Y))
			m.HexGraph.AddNode(node)
		}

		for _, hex := range hexes {
			neighbours := hexgrid.Ring(hexgrid.NewHexXYZ(hex.X, hex.Y, -hex.X-hex.Y), 1)
			for _, neighbour := range neighbours {
				nid := hexgrid.Pair(neighbour.X, neighbour.Y)

				// Check if node exists
				if m.HexGraph.Node(nid) == nil {
					continue
				}

				m.HexGraph.SetEdge(simple.Edge{simple.Node(hexgrid.Pair(hex.X, hex.Y)), m.HexGraph.Node(nid)})

			}
		}
	}

}

func (m *Map) DOT() {
	result, _ := dot.Marshal(m.HexGraph, "", "", "  ")
	file, err := os.Create("mygraph.gv")
	if err != nil {
		fmt.Println(err)
		return
	}
	file.WriteString(string(result))
}
