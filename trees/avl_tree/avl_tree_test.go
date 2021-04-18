package avl_tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	a := New()
	assert.Nil(t, a.root)
}

func TestInsertEmpty(t *testing.T) {
	a := New()
	n := a.Insert(6)
	assert.EqualValues(t, &node{data: 6, height: 1}, n)
	assert.Nil(t, a.root.left)
	assert.Nil(t, a.root.right)
	assert.EqualValues(t, 6, a.root.data)
	assert.EqualValues(t, 1, a.root.height)
}

// https://www.geeksforgeeks.org/avl-tree-set-1-insertion/
// using link to go through a known working example to test
func TestInsertionExampleRightRotation(t *testing.T) {
	a := New()
	a.Insert(13)
	a.Insert(10)
	a.Insert(15)
	a.Insert(5)
	a.Insert(11)
	a.Insert(16)
	a.Insert(4)
	a.Insert(8)

	assert.EqualValues(t, 13, a.root.data)
	assert.EqualValues(t, 10, a.root.left.data)
	assert.EqualValues(t, 15, a.root.right.data)
	assert.EqualValues(t, 5, a.root.left.left.data)
	assert.EqualValues(t, 11, a.root.left.right.data)
	assert.EqualValues(t, 16, a.root.right.right.data)

	//  result after right rotation
	a.Insert(3)

	// expected result after right rotation
	assert.EqualValues(t, 13, a.root.data)
	assert.EqualValues(t, 5, a.root.left.data)
	assert.EqualValues(t, 4, a.root.left.left.data)
	assert.EqualValues(t, 10, a.root.left.right.data)
	assert.EqualValues(t, 3, a.root.left.left.left.data)

	assert.EqualValues(t, 15, a.root.right.data)
	assert.EqualValues(t, 16, a.root.right.right.data)
}

func TestInsertionExampleLeftRotation(t *testing.T) {
	// setup
	a := New()
	a.Insert(30)
	a.Insert(5)
	a.Insert(35)
	a.Insert(32)
	a.Insert(40)

	assert.EqualValues(t, 30, a.root.data)
	assert.EqualValues(t, 5, a.root.left.data)
	assert.EqualValues(t, 35, a.root.right.data)
	assert.EqualValues(t, 32, a.root.right.left.data)
	assert.EqualValues(t, 40, a.root.right.right.data)

	// insertion
	a.Insert(45)

	// expected result after left rotation
	assert.EqualValues(t, 35, a.root.data)
	assert.EqualValues(t, 30, a.root.left.data)
	assert.EqualValues(t, 40, a.root.right.data)
	assert.EqualValues(t, 5, a.root.left.left.data)
	assert.EqualValues(t, 32, a.root.left.right.data)
	assert.EqualValues(t, 45, a.root.right.right.data)
}

func TestInsertionExampleLeftRightRotation(t *testing.T) {
	// 	// setup
	a := New()
	a.Insert(13)
	a.Insert(10)
	a.Insert(15)
	a.Insert(5)
	a.Insert(11)
	a.Insert(16)
	a.Insert(4)
	a.Insert(6)

	assert.EqualValues(t, 13, a.root.data)
	assert.EqualValues(t, 10, a.root.left.data)
	assert.EqualValues(t, 15, a.root.right.data)
	assert.EqualValues(t, 5, a.root.left.left.data)
	assert.EqualValues(t, 11, a.root.left.right.data)
	assert.EqualValues(t, 16, a.root.right.right.data)
	assert.EqualValues(t, 4, a.root.left.left.left.data)
	assert.EqualValues(t, 6, a.root.left.left.right.data)

	// insertion
	a.Insert(7)
	// expected result after a left, right rotation
	assert.EqualValues(t, 13, a.root.data)
	assert.EqualValues(t, 6, a.root.left.data)
	assert.EqualValues(t, 15, a.root.right.data)
	assert.EqualValues(t, 5, a.root.left.left.data)
	assert.EqualValues(t, 10, a.root.left.right.data)
	assert.EqualValues(t, 16, a.root.right.right.data)
	assert.EqualValues(t, 4, a.root.left.left.left.data)
	assert.EqualValues(t, 7, a.root.left.right.left.data)
	assert.EqualValues(t, 11, a.root.left.right.right.data)
}

func TestInsertionExampleRightLeftRotation(t *testing.T) {
	// setup
	a := New()
	a.Insert(5)
	a.Insert(2)
	a.Insert(7)
	a.Insert(1)
	a.Insert(4)
	a.Insert(6)
	a.Insert(9)
	a.Insert(3)
	a.Insert(16)

	assert.EqualValues(t, 5, a.root.data)
	assert.EqualValues(t, 2, a.root.left.data)
	assert.EqualValues(t, 7, a.root.right.data)
	assert.EqualValues(t, 1, a.root.left.left.data)
	assert.EqualValues(t, 4, a.root.left.right.data)
	assert.EqualValues(t, 6, a.root.right.left.data)
	assert.EqualValues(t, 9, a.root.right.right.data)
	assert.EqualValues(t, 3, a.root.left.right.left.data)
	assert.EqualValues(t, 16, a.root.right.right.right.data)

	// insertion
	a.Insert(15)

	// // expected result after a left, right rotation
	assert.EqualValues(t, 5, a.root.data)
	assert.EqualValues(t, 2, a.root.left.data)
	assert.EqualValues(t, 7, a.root.right.data)
	assert.EqualValues(t, 1, a.root.left.left.data)
	assert.EqualValues(t, 4, a.root.left.right.data)
	assert.EqualValues(t, 6, a.root.right.left.data)
	assert.EqualValues(t, 15, a.root.right.right.data)
	assert.EqualValues(t, 3, a.root.left.right.left.data)
	assert.EqualValues(t, 9, a.root.right.right.left.data)
	assert.EqualValues(t, 16, a.root.right.right.right.data)
}
