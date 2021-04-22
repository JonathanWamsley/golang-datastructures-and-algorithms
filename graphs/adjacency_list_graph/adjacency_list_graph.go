package adjacency_list_graph

// Edge - each edge is connected to another vertex with an associated weight
type Edge struct {
	vertex int
	weight float32
}

// AdjacencyList - creates a maping from to a list of edges
type AdjacencyList map[int][]*Edge

// New - creates a new adjacency list useing a map
func New() *AdjacencyList {
	return &AdjacencyList{}
}

// AddVertex - creates a vertex without an associated edge
func (al AdjacencyList) AddVertex(vertex int) {
	al[vertex] = append(al[vertex], &Edge{})
}

// AddEdge - creates a new edge with a float32 weight
func (al AdjacencyList) AddEdge(from, to int, edge float32) {
	e := Edge{
		vertex: to,
		weight: edge,
	}
	if !al.IsVertex(to) {
		al.AddVertex(to)
	}
	al[from] = append(al[from], &e)
}

// RemoveEdge - removes an edge at a vertex
func (al AdjacencyList) RemoveEdge(atVertex, delVertex int) {
	if _, ok := al[atVertex]; ok {
		for index, v := range al[atVertex] {
			if v.vertex == delVertex {
				al[atVertex] = append(al[atVertex][:index], al[atVertex][index+1:]...)
				return
			}
		}
	}
}

// IsEdge - returns true if edge exists
func (al AdjacencyList) IsEdge(from, to int) bool {
	if _, ok := al[from]; ok {
		for _, v := range al[from] {
			if v.vertex == to {
				return true
			}
		}
	}
	return false
}

// IsVertex - returns true if vertex exists
func (al AdjacencyList) IsVertex(vertex int) bool {
	if _, ok := al[vertex]; ok {
		return true
	}
	return false
}
