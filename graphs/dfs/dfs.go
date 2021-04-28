package dfs

// Edge - each edge is connected to another vertex with an associated weight
type Edge struct {
	vertex string
	weight int
}

// AdjacencyList - creates a maping from to a list of edges
type AdjacencyList map[string][]*Edge

// New - creates a new adjacency list useing a map
func New() *AdjacencyList {
	return &AdjacencyList{}
}

// AddEdge - creates a new edge with a float32 weight
func (al AdjacencyList) AddEdge(from, to string, edge int) {
	e := Edge{
		vertex: to,
		weight: edge,
	}
	al[from] = append(al[from], &e)
}

// GetEdges - returns a list of Edges at a given vertex
func (al AdjacencyList) GetEdges(vertex string) []Edge {
	edges := []Edge{}
	for _, v := range al[vertex] {
		edges = append(edges, *v)
	}
	return edges
}

// DFS - searches a graph using the depth first search traversal using a stack
func DFS(startLocation string, graph *AdjacencyList) []string {
	stack := []string{}
	visited := map[string]bool{}
	visitOrder := []string{}

	stack = append(stack, startLocation)
	for len(stack) != 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if !visited[node] {
			visitOrder = append(visitOrder, node)
			visited[node] = true
		}

		neighbors := graph.GetEdges(node)
		for _, neighbor := range neighbors {
			if !visited[neighbor.vertex] {
				stack = append(stack, neighbor.vertex)
			}
		}
	}
	return visitOrder
}
