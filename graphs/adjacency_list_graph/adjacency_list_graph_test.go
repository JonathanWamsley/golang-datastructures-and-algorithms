package adjacency_list_graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddVertex(t *testing.T) {
	al := New()
	al.AddVertex(1)
	assert.True(t, al.IsVertex(1))
	assert.False(t, al.IsVertex(2))

	al.AddEdge(1, 2, 2.0)
	assert.True(t, al.IsVertex(1))
	assert.True(t, al.IsVertex(2))
}

func TestAddEdge(t *testing.T) {
	al := New()
	al.AddEdge(1, 2, 1.0)
	al.AddEdge(2, 3, 2.0)
	assert.True(t, al.IsEdge(1, 2))
	assert.False(t, al.IsEdge(2, 1))
	assert.True(t, al.IsEdge(2, 3))
	assert.False(t, al.IsEdge(3, 2))
}

func TestRemoveEdge(t *testing.T) {
	al := New()
	al.AddEdge(1, 2, 1.0)
	assert.True(t, al.IsEdge(1, 2))
	al.RemoveEdge(1, 2)
	assert.False(t, al.IsEdge(1, 2))
}
