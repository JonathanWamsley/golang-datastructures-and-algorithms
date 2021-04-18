package avl_tree

// node - contains the data, height and children nodes
type node struct {
	data   int
	height int
	left   *node
	right  *node
}

// avlTree - represents a balanced binary tree
type avlTree struct {
	root *node
}

// New - creates a new avl tree
func New() *avlTree {
	return &avlTree{}
}

// Insert - adds data into the avl tree
func (a *avlTree) Insert(data int) *node {
	if a.root == nil {
		n := node{data: data, height: 1}
		a.root = &n
		return &n
	}
	return a.root.Insert(data)
}

// Insert - adds a node and performs a rotation if a tree high is not balanced
// reference: https://www.geeksforgeeks.org/avl-tree-set-1-insertion/
func (n *node) Insert(data int) *node {
	// perform normal bst insertion
	if n == nil {
		return &node{data: data, height: 1}
	} else if data <= n.data {

		n.left = n.left.Insert(data)
	} else if data > n.data {
		n.right = n.right.Insert(data)
	}

	// update height of ancestor node
	n.height = 1 + MaxHeight(n.left, n.right)

	// get the balance factor of the ancestor node to check whether this node became unbalanced
	balance := n.HeightDifference()

	if balance > 1 && data < n.left.data {
		n.rightRotation()
	}

	if balance < -1 && data > n.right.data {
		n.leftRotation()
	}

	if balance > 1 && data > n.left.data {
		n.left.leftRotation()
		n.rightRotation()
	}

	if balance < -1 && data < n.right.data {
		n.right.rightRotation()
		n.leftRotation()
	}

	return n
}

// MaxHeight - returns which ever child has a greater height
func MaxHeight(left, right *node) int {
	if left.Height() > right.Height() {
		return left.Height()
	}
	return right.Height()
}

// HeightDifference - returns height difference
func (n *node) HeightDifference() int {
	return n.left.Height() - n.right.Height()
}

/// Height - returns height of a node
func (n *node) Height() int {
	if n == nil {
		return 0
	}
	return n.height
}

func (n *node) rightRotation() {
	temp := &node{data: n.data, height: 1, left: n.left.right, right: n.right}

	*n = *n.left
	n.right = temp
}

func (n *node) leftRotation() {
	temp := &node{data: n.data, height: 1, left: n.left, right: n.right.left}

	*n = *n.right
	n.left = temp
}
