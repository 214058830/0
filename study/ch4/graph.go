package graph

import "fmt"

// 字符串到map映射
var graph = make(map[string]map[string]bool)

func addRdge(from, to string) {
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
