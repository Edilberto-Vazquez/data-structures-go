package graphs

// type nodes make(map[string][]string)

type graph struct {
	nodes        int
	adjacentList map[string][]string
}

func NewGraph() graph {
	return graph{nodes: 0, adjacentList: make(map[string][]string)}
}

func (g *graph) AddVertex(node string) {
	g.adjacentList[node] = []string{}
	g.nodes++
}

func (g *graph) AddEdge(node1 string, node2 string) {
	g.adjacentList[node1] = append(g.adjacentList[node1], node2)
}
