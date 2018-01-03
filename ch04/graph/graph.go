package main

import "fmt"

var graph = make(map[string]map[string]bool)

func addEdge(from, to string) {
	edges := graph[from]
	if edges == nil {
		edges = make(map[string]bool)
		graph[from] = edges
	}
	edges[to] = true
}

func hasEdge(from, to string) bool {
	return graph[from][to]
}

func main() {
	addEdge("A", "b")
	addEdge("X", "y")
	fmt.Printf("%v\n", graph)
	fmt.Printf("%t, %t %t\n", hasEdge("A", "b"), hasEdge("X", "y"), hasEdge("O", "p"))
}
