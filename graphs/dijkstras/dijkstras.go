package dijkstras

import (
	"container/heap"
	"errors"
	"fmt"
	"math"
)

// ====================================================================
// An adjacency list is used to represent a graph

// Edge - each edge is connected to another vertex with an associated weight
type Edge struct {
	vertex string
	weight int
}

// AdjacencyList - creates a maping from to a list of edges
type AdjacencyList map[string][]*Edge

// NewGraph - creates graph using an adjacency list
func NewGraph() *AdjacencyList {
	return &AdjacencyList{}
}

// AddEdge - creates a new edge directed edge as an int
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

// ====================================================================
// Heap was implemented using the builtin container heap.
// More information can be found at
// https://golang.org/pkg/container/heap/#example__priorityQueue

// Path - stores a path direction and its total distance
type Path struct {
	path          []string // stores the pathway to get from a startLocation to a given vertex
	totalDistance int      // the priority target
	index         int      // he index of the item in the heap
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Path

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].totalDistance < pq[j].totalDistance
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	path := x.(*Path)
	path.index = n
	*pq = append(*pq, path)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	path := old[n-1]
	old[n-1] = nil  // avoid memory leak
	path.index = -1 // for safety
	*pq = old[0 : n-1]
	return path
}

// NewDistanceMap - initialized a graphs vertex's distance to inf
func NewDistanceMap(g *AdjacencyList) map[string]float64 {
	d := make(map[string]float64, len(*g))
	for k := range *g {
		d[k] = math.Inf(1)
	}
	return d
}

// Dijkstras - finds the shortest path in a graph, returning the path to get their and its distance
func Dijkstras(startLocation, targetLocation string, g *AdjacencyList) (Path, error) {
	// creation of auxiliary data structures
	pq := make(PriorityQueue, 1)
	distance := NewDistanceMap(g)
	bestPaths := Path{}

	// initialization of auxiliary data structures
	distance[startLocation] = 0
	pq[0] = &Path{path: []string{startLocation}}

	for pq.Len() != 0 {
		minPathway := heap.Pop(&pq).(*Path)

		currentPath := minPathway.path
		lastLocation := currentPath[len(currentPath)-1]
		currentDistance := int(minPathway.totalDistance)
		if len(bestPaths.path) != 0 && currentDistance > bestPaths.totalDistance {
			break
		}

		neighbors := g.GetEdges(lastLocation)
		for _, neighbor := range neighbors {
			if neighbor.weight < 0 {
				return Path{}, errors.New("error: dijkstra's can not handle negative weights")
			}

			distanceToNeighbor := float64(currentDistance + neighbor.weight)
			if distanceToNeighbor < distance[neighbor.vertex] {
				distance[neighbor.vertex] = distanceToNeighbor

				updatedPath := append(currentPath, neighbor.vertex)
				item := &Path{
					path:          updatedPath,
					totalDistance: int(distanceToNeighbor),
				}
				heap.Push(&pq, item)
				// fmt.Println("traveled from", lastLocation, "to", neighbor.vertex, "at", updatedPath, distanceToNeighbor)
				if neighbor.vertex == targetLocation {
					if len(bestPaths.path) == 0 || item.totalDistance < bestPaths.totalDistance {
						bestPaths.path = updatedPath
						bestPaths.totalDistance = item.totalDistance
					}
				}
			}
		}

	}
	if len(bestPaths.path) == 0 {
		return Path{}, fmt.Errorf(fmt.Sprintf("error: no connection exists from %v to %v", startLocation, targetLocation))
	}
	return bestPaths, nil
}
