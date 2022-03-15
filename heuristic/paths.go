package heuristic

type Edge struct {
	Distance int
	Node     *Node
}

func BuildEdges(graph map[string]*Node) {
	buildEdge([...]*Node{graph["s1"], graph["s2"]}, 4)
	buildEdge([...]*Node{graph["s1"], graph["s3"]}, 10)

	buildEdge([...]*Node{graph["s2"], graph["s4"]}, 3)
	buildEdge([...]*Node{graph["s2"], graph["s7"]}, 8)
	buildEdge([...]*Node{graph["s2"], graph["s10"]}, 1)
	buildEdge([...]*Node{graph["s2"], graph["s11"]}, 1)

	buildEdge([...]*Node{graph["s3"], graph["s5"]}, 8)
	buildEdge([...]*Node{graph["s3"], graph["s10"]}, 8)
	buildEdge([...]*Node{graph["s3"], graph["s12"]}, 8)
	buildEdge([...]*Node{graph["s3"], graph["s13"]}, 2)

	buildEdge([...]*Node{graph["s4"], graph["s7"]}, 3)
	buildEdge([...]*Node{graph["s4"], graph["s8"]}, 8)
	buildEdge([...]*Node{graph["s4"], graph["s10"]}, 8)

	buildEdge([...]*Node{graph["s5"], graph["s8"]}, 7)
	buildEdge([...]*Node{graph["s5"], graph["s9"]}, 7)

	buildEdge([...]*Node{graph["s7"], graph["s8"]}, 8)

	buildEdge([...]*Node{graph["s8"], graph["s9"]}, 8)

	buildEdge([...]*Node{graph["s10"], graph["s12"]}, 1)
	buildEdge([...]*Node{graph["s11"], graph["s13"]}, 2)
}

func buildEdge(nodes [2]*Node, distance int) {
	nodes[0].Edges = append(nodes[0].Edges, Edge{
		Distance: distance,
		Node:     nodes[1],
	})

	nodes[1].Edges = append(nodes[1].Edges, Edge{
		Distance: distance,
		Node:     nodes[0],
	})
}
