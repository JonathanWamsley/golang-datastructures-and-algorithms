package dfs

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
