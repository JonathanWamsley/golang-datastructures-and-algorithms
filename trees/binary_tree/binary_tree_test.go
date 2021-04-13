package binary_tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	b := New()
	assert.EqualValues(t, 1, b.root.Data)
	assert.EqualValues(t, 2, b.root.left.Data)
	assert.EqualValues(t, 3, b.root.right.Data)
	assert.EqualValues(t, 4, b.root.left.left.Data)
	assert.EqualValues(t, 5, b.root.left.right.Data)
	assert.EqualValues(t, 6, b.root.right.left.Data)
	assert.EqualValues(t, 7, b.root.right.right.Data)
}

func TestPreOrderTraversalRec(t *testing.T) {
	b := New()
	s := PreOrderTraversalRec(b.root)
	assert.EqualValues(t, "[ 1 2 4 5 3 6 7 ]", s)
}

func TestInOrderTraversalRec(t *testing.T) {
	b := New()
	s := InOrderTraversalRec(b.root)
	assert.EqualValues(t, "[ 4 2 5 1 6 3 7 ]", s)
}

func TestPostOrderTraversalRec(t *testing.T) {
	b := New()
	s := PostOrderTraversalRec(b.root)
	assert.EqualValues(t, "[ 4 5 2 6 7 3 1 ]", s)
}

func TestPreOrderTraversalIter(t *testing.T) {
	b := New()
	s := PreOrderTraversalIter(b.root)
	assert.EqualValues(t, "[ 1 2 4 5 3 6 7 ]", s)
}

func TestInPreOrderTraversalIter(t *testing.T) {
	b := New()
	s := InOrderTraversalIter(b.root)
	assert.EqualValues(t, "[ 4 2 5 1 6 3 7 ]", s)
}

func TestPostPreOrderTraversalIter(t *testing.T) {
	b := New()
	s := PostOrderTraversalIter(b.root)
	assert.EqualValues(t, "[ 4 5 2 6 7 3 1 ]", s)
}
