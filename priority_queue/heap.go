package priority_queue

import (
	"errors"
)

type heap []int

// New - creates a new heap
func New() *heap {
	return &heap{}
}

// Insert - adds an int to the heap
func (h *heap) Insert(data int) int {
	*h = append(*h, data)

	h.percolateUp(h.Len() - 1)
	return data
}

// Delete - removes an int from the heap
func (h *heap) Delete() (int, error) {
	if len(*h) == 0 {
		return -1, errors.New("error: can not delete, priority queue is empty")
	}
	data := (*h)[0]

	if len(*h) == 1 {
		*h = heap{}
		return data, nil
	}

	(*h)[0] = (*h)[len(*h)-1]
	(*h) = (*h)[:len(*h)-1]
	h.percolateDown(0)
	return data, nil
}

// Max - gets the max int from the heap
func (h *heap) Max() (interface{}, error) {
	if h.Len() == 0 {
		return nil, errors.New("error: priority queue is empty")
	}
	return (*h)[0], nil
}

// Len - returns amount of ints in heap
func (h *heap) Len() int {
	return len(*h)
}

func (h *heap) percolateUp(i int) {
	parent, err := h.parent(i)
	if err != nil {
		return
	}

	if (*h)[i] > (*h)[parent] {
		(*h)[i], (*h)[parent] = (*h)[parent], (*h)[i]
		h.percolateUp(parent)
	}
}

func (h *heap) percolateDown(i int) {
	current := (*h)[i]
	leftMax, _ := h.leftChild(i)
	rightMax, _ := h.rightChild(i)
	max := i

	if leftMax != -1 && (*h)[leftMax] > current {
		max = leftMax
	}
	if rightMax != -1 && (*h)[rightMax] > current && (*h)[rightMax] > (*h)[leftMax] {
		max = rightMax
	}
	if max != i {
		(*h)[i], (*h)[max] = (*h)[max], (*h)[i]
		h.percolateDown(max)
	}

}

func (h *heap) parent(i int) (int, error) {
	if i <= 0 || i >= h.Len() {
		return -1, errors.New("error: index out of bounds")
	}
	return (i - 1) / 2, nil
}

func (h *heap) leftChild(i int) (int, error) {
	lc := (i * 2) + 1
	if lc >= h.Len() {
		return -1, errors.New("error: index out of bounds")
	}
	return lc, nil
}

func (h *heap) rightChild(i int) (int, error) {
	rc := (i * 2) + 2
	if rc >= h.Len() {
		return -1, errors.New("error: index out of bounds")
	}
	return rc, nil
}
