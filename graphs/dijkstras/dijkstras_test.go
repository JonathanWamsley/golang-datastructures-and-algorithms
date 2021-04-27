package dijkstras

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// test example from
// https://en.wikipedia.org/wiki/Breadth-first_search
func TestDijkstras(t *testing.T) {
	g := NewGraph()

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
	// g.AddEdge("Munchen", "Kassel", 502)

	g.AddEdge("Karlsruhue", "Augsburg", 250)
	g.AddEdge("Augsburg", "Karlsruhue", 250)

	g.AddEdge("Nurnberg", "Stuttgart", 183)
	g.AddEdge("Stuttgart", "Nurnberg", 183)

	g.AddEdge("Nurnberg", "Munchen", 167)
	// g.AddEdge("Munchen", "Nurnberg", 167)

	g.AddEdge("Augsburg", "Munchen", 84)
	// g.AddEdge("Munchen", "Augsburg", 84)

	path, err := Dijkstras("Frankfurt", "Munchen", g)
	assert.Nil(t, err)
	assert.EqualValues(t, 487, path.totalDistance)
	assert.EqualValues(t, []string{"Frankfurt", "Wurzburg", "Nurnberg", "Munchen"}, path.path)
}

func TestDijkstrasErrorNegative(t *testing.T) {
	assert.True(t, true)
	g := NewGraph()

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

	g.AddEdge("Wurzburg", "Nurnberg", -1003) // made negative
	g.AddEdge("Nurnberg", "Wurzburg", -1003) // made negative

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

	path, err := Dijkstras("Frankfurt", "Munchen", g)

	assert.NotNil(t, err)
	assert.EqualValues(t, errors.New("error: dijkstra's can not handle negative weights"), err)
	assert.EqualValues(t, Path{}.totalDistance, path.totalDistance)
	assert.EqualValues(t, Path{}.path, path.path)
}

func TestDijkstrasErrorNotConnected(t *testing.T) {
	g := NewGraph()

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

	// g.AddEdge("Kassel", "Munchen", 502) // removed connection to target location
	// g.AddEdge("Munchen", "Kassel", 502)

	g.AddEdge("Karlsruhue", "Augsburg", 250)
	g.AddEdge("Augsburg", "Karlsruhue", 250)

	g.AddEdge("Nurnberg", "Stuttgart", 183)
	g.AddEdge("Stuttgart", "Nurnberg", 183)

	// g.AddEdge("Nurnberg", "Munchen", 167) // removed connection to target location
	// g.AddEdge("Munchen", "Nurnberg", 167)

	// g.AddEdge("Augsburg", "Munchen", 84) // removed connection to target location
	// g.AddEdge("Munchen", "Augsburg", 84)

	startLocation := "Frankfurt"
	targetLocation := "Munchen"

	path, err := Dijkstras(startLocation, targetLocation, g)
	assert.NotNil(t, err)
	assert.EqualValues(t, fmt.Errorf("error: no connection exists from %v to %v", startLocation, targetLocation), err)
	assert.EqualValues(t, Path{}.totalDistance, path.totalDistance)
	assert.EqualValues(t, Path{}.path, path.path)

}
