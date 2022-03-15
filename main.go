package main

import (
	"fmt"
	"heuristic/heuristic"
	"time"
)

type tracker struct {
	totalDistance int
	route         *route
	lastRoute     *route
}

type route struct {
	node *heuristic.Node
	next *route
}

func main() {
	graph := make(map[string]*heuristic.Node)
	heuristic.BuildPoints(graph)
	heuristic.BuildEdges(graph)

	// for graphNode, val := range graph {
	// 	fmt.Printf("%s\t%p\t%v\n", graphNode, graph[graphNode], val.Edges)
	// }
	heuristicSearch(graph, "s1", "s8")
}

func heuristicSearch(graph map[string]*heuristic.Node, from, to string) {
	var fromNode *heuristic.Node = graph[from]
	var toNode *heuristic.Node = graph[to]
	startRoute := &route{
		node: graph[from],
	}
	t := tracker{
		route:     startRoute,
		lastRoute: startRoute,
	}
	var distanceTraveled int

	routeElem := startRoute
	for t.lastRoute.node.NodeName != toNode.NodeName {

		fromNode, distanceTraveled = calculateNodeDistance(&t, fromNode, fromNode.Edges)
		t.totalDistance += distanceTraveled
		newRoute := &route{
			node: fromNode,
		}
		for routeElem.next != nil {
			routeElem = routeElem.next
		}
		routeElem.next = newRoute
		t.lastRoute = newRoute
		fmt.Println("total", t.totalDistance)
		printTraveled(t)
		time.Sleep(5 * time.Second)
	}
}

func printTraveled(t tracker) {
	routeElem := t.route
	for routeElem.next != nil {
		fmt.Printf("%s > ", routeElem.node.NodeName)
		routeElem = routeElem.next
	}
	fmt.Print("\n")

}

func calculateNodeDistance(t *tracker, from *heuristic.Node, edges []heuristic.Edge) (*heuristic.Node, int) {

	var idx, min, aux int
	for i, edge := range edges {
		aux = t.totalDistance + edge.Distance + edge.Node.Weight
		if i == 0 {
			idx = i
			min = aux
		}
		if aux < min {
			idx = i
			min = aux
		}
	}

	fmt.Println(t.totalDistance, min)

	return edges[idx].Node, edges[idx].Distance
}
