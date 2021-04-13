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
	assert.EqualValues(t, &node{data: 6, height: 1, root: n}, n)
	assert.Nil(t, a.root.left)
	assert.Nil(t, a.root.right)
	assert.EqualValues(t, 6, a.root.data)
	assert.EqualValues(t, 1, a.root.height)
}

// func TestInsertLeft(t *testing.T) {
// 	a := New()
// 	_ = a.Insert(6)
// 	_ = a.Insert(5)
// 	fmt.Println(a.root)
// 	fmt.Println(a.root.left)
// 	// assert.Nil(t, a.root.left.left)
// 	// assert.Nil(t, a.root.left.right)
// 	// assert.EqualValues(t, 1, a.root.left.data)
// 	// assert.EqualValues(t, 0, a.root.left.height)
// }

// func TestInsertLeft(t *testing.T) {
// 	n := node{data: 5}
// 	// node := n.Insert(2)
// 	node := n.Insert(3)
// 	_ = n.Insert(0)
// 	_ = n.Insert(8)
// 	_ = n.Insert(7)
// 	fmt.Println(n, node)
// }


// func TestRightRotation(t *testing.T) {
// 	n := node{data: 5}
// 	n.Insert(6)
// 	n.Insert(7)
// 	assert.EqualValues(t, 6, n.data)
// 	assert.NotNil(t, n.left)
// 	assert.NotNil(t, n.right)
// 	assert.EqualValues(t, 2, n.height)

// 	assert.EqualValues(t, 5, n.left.data)
// 	assert.Nil(t, n.left.left)
// 	assert.Nil(t, n.left.right)
// 	assert.EqualValues(t, 1, n.right.height)

// 	assert.EqualValues(t, 7, n.right.data)
// 	assert.Nil(t, n.right.left)
// 	assert.Nil(t, n.right.right)
// 	assert.EqualValues(t, 1, n.left.height)
// }

// func TestLeftRotation(t *testing.T) {
// 	n := node{data: 5}
// 	n.Insert(4)
// 	n.Insert(3)
// 	assert.EqualValues(t, 4, n.data)
// 	assert.NotNil(t, n.left)
// 	assert.NotNil(t, n.right)
// 	assert.EqualValues(t, 2, n.height)

// 	assert.EqualValues(t, 3, n.left.data)
// 	assert.Nil(t, n.left.left)
// 	assert.Nil(t, n.left.right)
// 	assert.EqualValues(t, 1, n.left.height)

// 	assert.EqualValues(t, 5, n.right.data)
// 	assert.Nil(t, n.right.left)
// 	assert.Nil(t, n.right.right)
// 	assert.EqualValues(t, 1, n.right.height)
// }

// https://www.geeksforgeeks.org/avl-tree-set-1-insertion/
// using link to go through a known working example to test
func TestInsertionExampleRightRotation(t *testing.T) {
	n := node{data: 13}
	n.Insert(10)
	n.Insert(15)
	n.Insert(5)
	n.Insert(11)
	n.Insert(16)
	n.Insert(4)
	n.Insert(8)

	assert.EqualValues(t, 13, n.data)
	assert.EqualValues(t, 10, n.left.data)
	assert.EqualValues(t, 15, n.right.data)
	assert.EqualValues(t, 5, n.left.left.data)
	assert.EqualValues(t, 11, n.left.right.data)
	assert.EqualValues(t, 16, n.right.right.data)

	//  result after right rotation
	n.Insert(3)

	// expected result after right rotation
	assert.EqualValues(t, 13, n.data)
	assert.EqualValues(t, 5, n.left.data)
	assert.EqualValues(t, 4, n.left.left.data)
	assert.EqualValues(t, 10, n.left.right.data)
	assert.EqualValues(t, 3, n.left.left.left.data)

	assert.EqualValues(t, 15, n.right.data)
	assert.EqualValues(t, 16, n.right.right.data)
}

func TestInsertionExampleLeftRotation(t *testing.T) {
	// setup
	n := node{data: 30}
	n.Insert(5)
	n.Insert(35)
	n.Insert(32)
	n.Insert(40)
	
	assert.EqualValues(t, 30, n.data)
	assert.EqualValues(t, 5, n.left.data)
	assert.EqualValues(t, 35, n.right.data)
	assert.EqualValues(t, 32, n.right.left.data)
	assert.EqualValues(t, 40, n.right.right.data)

	// insertion
	n.Insert(45)

	// expected result after left rotation
	assert.EqualValues(t, 35, n.data)
	assert.EqualValues(t, 30, n.left.data)
	assert.EqualValues(t, 40, n.right.data)
	assert.EqualValues(t, 5, n.left.left.data)
	assert.EqualValues(t, 32, n.left.right.data)
	assert.EqualValues(t, 45, n.right.right.data)
}

func TestInsertionExampleLeftRightRotation(t *testing.T) {
	// setup
	n := node{data: 13}
	n.Insert(10)
	n.Insert(15)
	n.Insert(5)
	n.Insert(11)
	n.Insert(16)
	n.Insert(4)
	n.Insert(6)

	assert.EqualValues(t, 13, n.data)
	assert.EqualValues(t, 10, n.left.data)
	assert.EqualValues(t, 15, n.right.data)
	assert.EqualValues(t, 5, n.left.left.data)
	assert.EqualValues(t, 11, n.left.right.data)
	assert.EqualValues(t, 16, n.right.right.data)
	assert.EqualValues(t, 4, n.left.left.left.data)
	assert.EqualValues(t, 6, n.left.left.right.data)

	// insertion
	n.Insert(7)

	// expected result after a left, right rotation
	assert.EqualValues(t, 13, n.data)
	assert.EqualValues(t, 6, n.left.data)
	assert.EqualValues(t, 15, n.right.data)
	assert.EqualValues(t, 5, n.left.left.data)
	assert.EqualValues(t, 10, n.left.right.data)
	assert.EqualValues(t, 16, n.right.right.data)
	assert.EqualValues(t, 4, n.left.left.left.data)
	assert.EqualValues(t, 7, n.left.right.left.data)
	assert.EqualValues(t, 11, n.left.right.right.data)
}

func TestInsertionExampleRightLeftRotation(t *testing.T) {
	// setup
	n := node{data: 5}
	n.Insert(2)
	n.Insert(7)
	n.Insert(1)
	n.Insert(4)
	n.Insert(6)
	n.Insert(9)
	n.Insert(3)
	n.Insert(16)

	assert.EqualValues(t, 5, n.data)
	assert.EqualValues(t, 2, n.left.data)
	assert.EqualValues(t, 7, n.right.data)
	assert.EqualValues(t, 1, n.left.left.data)
	assert.EqualValues(t, 4, n.left.right.data)
	assert.EqualValues(t, 6, n.right.left.data)
	assert.EqualValues(t, 9, n.right.right.data)
	assert.EqualValues(t, 3, n.left.right.left.data)
	assert.EqualValues(t, 16, n.right.right.right.data)

	// insertion
	n.Insert(15)

	// // expected result after a left, right rotation
	assert.EqualValues(t, 5, n.data)
	assert.EqualValues(t, 2, n.left.data)
	assert.EqualValues(t, 7, n.right.data)
	assert.EqualValues(t, 1, n.left.left.data)
	assert.EqualValues(t, 4, n.left.right.data)
	assert.EqualValues(t, 6, n.right.left.data)
	assert.EqualValues(t, 15, n.right.right.data)
	assert.EqualValues(t, 3, n.left.right.left.data)
	assert.EqualValues(t, 9, n.right.right.left.data)
	assert.EqualValues(t, 16, n.right.right.right.data)
}