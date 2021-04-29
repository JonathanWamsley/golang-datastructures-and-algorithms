package connected_components

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// reference https://www.baeldung.com/cs/graph-connected-components
func TestFindComponents(t *testing.T) {
	g := New()
	// 3 connected components
	g.AddEdge("v1", "v2", 1)
	g.AddEdge("v2", "v1", 1)

	g.AddEdge("v3", "v2", 1)
	g.AddEdge("v2", "v3", 1)

	g.AddEdge("v3", "v1", 1)
	g.AddEdge("v1", "v3", 1)

	// 2 connected components
	g.AddEdge("v4", "v5", 1)
	g.AddEdge("v5", "v4", 1)

	// 2 connected components

	g.AddEdge("v7", "v6", 1)
	g.AddEdge("v6", "v7", 1)

	// 1 connected component
	g.AddVertex("v8")

	componentAmount, components := FindComponents(g)
	assert.EqualValues(t, 4, componentAmount)

	connectionOne := components["v1"]
	connectionTwo := components["v2"]
	connectionThree := components["v3"]
	assert.True(t, connectionOne == connectionTwo && connectionTwo == connectionThree)

	connectionFour := components["v4"]
	connectionFive := components["v5"]
	assert.True(t, connectionFour == connectionFive)

	connectionSix := components["v6"]
	connectionSeven := components["v7"]
	assert.True(t, connectionSix == connectionSeven)

	connectionEight := components["v8"]
	assert.True(t, connectionOne != connectionFour && connectionFour != connectionSix && connectionSix != connectionEight)
}
