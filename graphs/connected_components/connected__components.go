package connected_components

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

func (al AdjacencyList) AddVertex(vertex string) {
	al[vertex] = []*Edge{}
}

// GetEdges - returns a list of Edges at a given vertex
func (al AdjacencyList) GetEdges(vertex string) []Edge {
	edges := []Edge{}
	for _, v := range al[vertex] {
		edges = append(edges, *v)
	}
	return edges
}

// FindComponents - finds connected components in a graph and labels them together
func FindComponents(graph *AdjacencyList) (int, map[string]int) {
	visited := map[string]bool{}
	components := map[string]int{}

	index := 0
	for vertex := range *graph {
		if !visited[vertex] {
			index++
			FindComponentsDFS(index, vertex, visited, components, *graph)

		}
	}
	return index, components
}

// FindComponentsDFS - traversal method for finding all connected components
func FindComponentsDFS(index int, vertex string, visited map[string]bool, components map[string]int, graph AdjacencyList) {
	if visited[vertex] {
		return
	}

	visited[vertex] = true
	components[vertex] = index
	neighbors := graph.GetEdges(vertex)
	for _, neighbor := range neighbors {
		if !visited[neighbor.vertex] {
			FindComponentsDFS(index, neighbor.vertex, visited, components, graph)
		}
	}
}
