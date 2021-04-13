package avl_tree

type node struct {
	data int
	height int
	left *node
	right *node

	root *node
}

type avlTree struct {
	root *node
}

func New() *avlTree {
	return &avlTree{}
}

func (a *avlTree) Insert(data int) *node {
	if a.root == nil {
		n := node{data: data, height: 1}
		n.root = &n
		a.root = &n
		return &n
	}
	return a.root.Insert(data)
}

// https://www.geeksforgeeks.org/avl-tree-set-1-insertion/
func (n *node) Insert(data int) *node {
	// perform normal bst insertion
	if n == nil {
		return &node{data: data, height: 1}
	} else if data <= n.data {

		n.left = n.left.Insert(data)
	} else if data > n.data {
		n.right = n.right.Insert(data)
	} else {
		return n
	}

	// update height of ancestor node
	n.height = 1 + max(n.left, n.right)

	// get the balance factor of the ancestor node to check whether this node became unbalanced
	balance := n.Balanced()

	if balance > 1 && data < n.left.data {
		n.RightRotation()
	}

	if balance < -1 && data > n.right.data {
		n.LeftRotation()
	}

	if balance > 1 && data > n.left.data {
		n.left.LeftRotation()
		n.RightRotation()
	}

	if balance < -1 && data < n.right.data {
		n.right.RightRotation()
		n.LeftRotation()
	}

	return n
}

func max(left, right *node) int {
	if left.Height() > right.Height() {
		return left.Height()
	}
	return right.Height()
}

func (n *node) Balanced() int {
	if n == nil {
		return 0
	}
	return n.left.Height() - n.right.Height()
}

func (n *node) Height() int {
	if n == nil {
		return 0
	}
	return n.height
}

func (n *node) RightRotation() {	
	temp := &node{data: n.data, height: 1, left: n.left.right, right: n.right}


	*n = *n.left
	n.right = temp

}

func (n *node) LeftRotation() {
	temp := &node{data: n.data, height: 1, left: n.left, right: n.right.left}

	*n = *n.right
	n.left = temp
}
