package adjacency_matrix_graph

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUndirectedAdjacencyGraph(t *testing.T) {
	u := NewUndirectedGraph(3)
	assert.EqualValues(t, 3, len(*u))
}

func TestNewUndirectedAdjacencyGraphNil(t *testing.T) {
	u := NewUndirectedGraph(0)
	assert.Nil(t, u)
}

func TestNewdirectedAdjacencyGraph(t *testing.T) {
	d := NewDirectedGraph(3)
	assert.EqualValues(t, 3, len(*d))
}

func TestNewdirectedAdjacencyGraphNil(t *testing.T) {
	d := NewDirectedGraph(0)
	assert.Nil(t, d)
}

func TestUndirectedAddEdge(t *testing.T) {
	d := NewUndirectedGraph(3)
	err1 := d.AddEdge(0, 1, 1.5)
	err2 := d.AddEdge(1, 2, 2.0)
	assert.Nil(t, err1)
	assert.Nil(t, err2)

	b1, err1 := d.IsEdge(0, 1)
	b2, err2 := d.IsEdge(1, 0)
	b3, err3 := d.IsEdge(1, 2)
	b4, err4 := d.IsEdge(2, 1)
	assert.True(t, b1)
	assert.True(t, b2)
	assert.True(t, b3)
	assert.True(t, b4)
	assert.Nil(t, err1)
	assert.Nil(t, err2)
	assert.Nil(t, err3)
	assert.Nil(t, err4)

}

func TestDirectedAddEdge(t *testing.T) {
	d := NewDirectedGraph(3)
	err1 := d.AddEdge(0, 1, 1.5)
	err2 := d.AddEdge(1, 2, 2.0)
	assert.Nil(t, err1)
	assert.Nil(t, err2)

	b1, err1 := d.IsEdge(0, 1)
	b2, err2 := d.IsEdge(1, 0)
	b3, err3 := d.IsEdge(1, 2)
	b4, err4 := d.IsEdge(2, 1)
	assert.True(t, b1)
	assert.False(t, b2)
	assert.True(t, b3)
	assert.False(t, b4)
	assert.Nil(t, err1)
	assert.Nil(t, err2)
	assert.Nil(t, err3)
	assert.Nil(t, err4)
}

func TestUndirectedRemove(t *testing.T) {
	u := NewUndirectedGraph(3)
	err1 := u.AddEdge(0, 1, 1.5)
	err2 := u.RemoveEdge(1, 2)
	assert.Nil(t, err1)
	assert.Nil(t, err2)
}

func TestDirectedRemove(t *testing.T) {
	d := NewDirectedGraph(3)
	err1 := d.AddEdge(0, 1, 1.5)
	err2 := d.RemoveEdge(1, 2)
	assert.Nil(t, err1)
	assert.Nil(t, err2)
}

func TestUndirectedOutOfBounds(t *testing.T) {
	u := NewUndirectedGraph(2)
	err1 := u.AddEdge(1, 2, 1.5)
	assert.EqualValues(t, errors.New("error: vertex out of bounds"), err1)

	_, err2 := u.IsEdge(1, 2)
	assert.EqualValues(t, errors.New("error: vertex out of bounds"), err2)

	err3 := u.RemoveEdge(1, 2)
	assert.EqualValues(t, errors.New("error: vertex out of bounds"), err3)

}

func TestDirectedOutOfBounds(t *testing.T) {
	d := NewDirectedGraph(2)
	err1 := d.AddEdge(1, 2, 1.5)
	assert.EqualValues(t, errors.New("error: vertex out of bounds"), err1)

	_, err2 := d.IsEdge(1, 2)
	assert.EqualValues(t, errors.New("error: vertex out of bounds"), err2)

	err3 := d.RemoveEdge(1, 2)
	assert.EqualValues(t, errors.New("error: vertex out of bounds"), err3)
}
