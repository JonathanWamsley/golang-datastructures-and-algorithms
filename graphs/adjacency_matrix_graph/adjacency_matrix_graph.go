package adjacency_matrix_graph

import "errors"

type undirectedAdjMatrix [][]float32
type directedAdjMatrix [][]float32

// NewUndirectedGraph - creates a directed graph with a vertex amount
func NewUndirectedGraph(vertices int) *undirectedAdjMatrix {
	if vertices < 1 {
		return nil
	}
	m := make(undirectedAdjMatrix, vertices)
	for i := range m {
		m[i] = make([]float32, vertices)
	}
	return &m
}

// NewDirectedGraph - creates a directed graph with a vertex amount
func NewDirectedGraph(vertices int) *directedAdjMatrix {
	if vertices < 1 {
		return nil
	}
	m := make(directedAdjMatrix, vertices)
	for i := range m {
		m[i] = make([]float32, vertices)
	}
	return &m
}

// AddEdge - creates an undirected edge from vertex i to j
func (am *undirectedAdjMatrix) AddEdge(i, j int, value float32) error {
	if obErr := am.OutOfBoundsError(i, j); obErr != nil {
		return obErr
	}
	(*am)[i][j] = value
	(*am)[j][i] = value
	return nil
}

// AddEdge - creates a directed edge from 'from' to 'to' vertices
func (am *directedAdjMatrix) AddEdge(from, to int, value float32) error {
	if obErr := am.OutOfBoundsError(to, from); obErr != nil {
		return obErr
	}
	(*am)[from][to] = value
	return nil
}

// RemoveEdge - removes an undirected edge from vertex i to j
func (am *undirectedAdjMatrix) RemoveEdge(i, j int) error {
	if obErr := am.OutOfBoundsError(i, j); obErr != nil {
		return obErr
	}
	(*am)[i][j] = 0
	(*am)[j][i] = 0
	return nil
}

// RemoveEdge - removes an directed edge from 'from' to 'to' vertices
func (am *directedAdjMatrix) RemoveEdge(from, to int) error {
	if obErr := am.OutOfBoundsError(from, to); obErr != nil {
		return obErr
	}
	(*am)[from][to] = 0
	return nil
}

// IsEdge - returns true if vertex i, j is connected
func (am *undirectedAdjMatrix) IsEdge(i, j int) (bool, error) {
	if obErr := am.OutOfBoundsError(i, j); obErr != nil {
		return false, obErr
	}
	return (*am)[i][j] != 0, nil
}

// IsEdge - returns true if vertex 'from' is connected to 'to'
func (am *directedAdjMatrix) IsEdge(from, to int) (bool, error) {
	if obErr := am.OutOfBoundsError(from, to); obErr != nil {
		return false, obErr
	}
	return (*am)[from][to] != 0, nil
}

// OutOfBoundsError - returns an index out of bounds error
func (am *undirectedAdjMatrix) OutOfBoundsError(i, j int) error {
	if i >= len(*am) || j >= len(*am) {
		return errors.New("error: vertex out of bounds")
	}
	return nil
}

// OutOfBoundsError - returns an index out of bounds error
func (am *directedAdjMatrix) OutOfBoundsError(from, to int) error {
	if from >= len(*am) || to >= len(*am) {
		return errors.New("error: vertex out of bounds")
	}
	return nil
}
