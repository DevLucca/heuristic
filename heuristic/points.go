package heuristic

type Node struct {
	Weight   int
	NodeName string
	Edges    []Edge
}

func BuildPoints(graph map[string]*Node) {
	graph["s1"] = &Node{
		NodeName: "s1",
		Weight:   50,
	}

	graph["s2"] = &Node{
		NodeName: "s2",
		Weight:   40,
	}

	graph["s3"] = &Node{
		NodeName: "s3",
		Weight:   40,
	}

	graph["s4"] = &Node{
		NodeName: "s4",
		Weight:   30,
	}

	graph["s5"] = &Node{
		NodeName: "s5",
		Weight:   20,
	}

	graph["s7"] = &Node{
		NodeName: "s7",
		Weight:   20,
	}

	graph["s8"] = &Node{
		NodeName: "s8",
		Weight:   2,
	}

	graph["s9"] = &Node{
		NodeName: "s9",
		Weight:   10,
	}

	graph["s10"] = &Node{
		NodeName: "s10",
		Weight:   25,
	}

	graph["s11"] = &Node{
		NodeName: "s11",
		Weight:   45,
	}

	graph["s12"] = &Node{
		NodeName: "s12",
		Weight:   15,
	}

	graph["s13"] = &Node{
		NodeName: "s13",
		Weight:   45,
	}
}
