package bfs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBFS(t *testing.T) {
	assert.True(t, true)
	g := New()

	g.AddEdge("Frankfurt", "Mannhelm", 85)
	g.AddEdge("Mannhelm", "Frankfurt", 85)

	g.AddEdge("Frankfurt", "Wurzburg", 217)
	g.AddEdge("Wurzburg", "Frankfurt", 217)

	g.AddEdge("Frankfurt", "Kassel", 173)
	g.AddEdge("Kassel", "Frankfurt", 173)

	g.AddEdge("Mannhelm", "Karlsruhue", 80)
	g.AddEdge("Karlsruhue", "Mannhelm", 80)

	g.AddEdge("Wurzburg", "Erfurt", 186)
	g.AddEdge("Erfurt", "Wurzburg", 186)

	g.AddEdge("Wurzburg", "Nurnberg", 103)
	g.AddEdge("Nurnberg", "Wurzburg", 103)

	g.AddEdge("Kassel", "Munchen", 502)
	g.AddEdge("Munchen", "Kassel", 502)

	g.AddEdge("Karlsruhue", "Augsburg", 250)
	g.AddEdge("Augsburg", "Karlsruhue", 250)

	g.AddEdge("Nurnberg", "Stuttgart", 183)
	g.AddEdge("Stuttgart", "Nurnberg", 183)

	g.AddEdge("Nurnberg", "Munchen", 167)
	g.AddEdge("Munchen", "Nurnberg", 167)

	g.AddEdge("Augsburg", "Munchen", 84)
	g.AddEdge("Munchen", "Augsburg", 84)

	visitedOrder := BFS("Frankfurt", g)
	assert.EqualValues(t, []string{"Frankfurt", "Mannhelm", "Wurzburg", "Kassel", "Karlsruhue", "Erfurt", "Nurnberg", "Munchen", "Augsburg", "Stuttgart"}, visitedOrder)

}
