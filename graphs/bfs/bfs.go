package bfs

import (
	"errors"
)

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

func (al AdjacencyList) Edges(vertex string) []Edge {
	edges := []Edge{}
	for _, v := range al[vertex] {
		edges = append(edges, *v)
	}
	return edges
}

// ====================================================================
type Queue struct {
	items    []string // place a, place b, place c
	distance []int
}

// NewQueue - creates a new queue
func NewQueue() *Queue {
	return &Queue{}
}

// Enqueue - adds items to the end
func (q *Queue) Enqueue(item string, distance int) {
	q.items = append(q.items, item)
	q.distance = append(q.distance, distance)
}

// Dequeue - pops an item from the front
func (q *Queue) Dequeue() (string, error) {
	if q.IsEmpty() {
		return "", errors.New("error: queue empty")
	}
	item := q.items[0]
	q.items = q.items[1:]

	q.distance = q.distance[1:]
	return item, nil
}

// IsEmpty - returns true if no items aer in the queue
func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

// BFS - searches a graph using the breadth first search traversal using a queue
func BFS(start string, g *AdjacencyList) []string {
	q := NewQueue()
	visited := make(map[string]bool)
	visitedOrder := []string{}

	q.Enqueue(start, 0)
	visited[start] = true

	for !q.IsEmpty() {
		node, _ := q.Dequeue()
		visitedOrder = append(visitedOrder, node)
		visited[node] = true
		neighbors := g.Edges(node)

		for _, neighbor := range neighbors {
			if !visited[neighbor.vertex] {
				q.Enqueue(neighbor.vertex, neighbor.weight)
				visited[neighbor.vertex] = true
			}
		}
	}
	return visitedOrder
}
