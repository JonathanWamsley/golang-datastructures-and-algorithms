package linked_list_queue

import "errors"

// node - contains the data and pointer to next node
type node struct {
	Data interface{}
	next *node
}

// queue - the datastructure to easily remove from front and add to end
type queue struct {
	head *node
	tail *node
	len  int
}

// New - creates a new queue
func New() *queue {
	return &queue{}
}

// Enqueue - adds a node to the end
func (q *queue) Enqueue(d interface{}) *node {
	newTail := node{Data: d}
	if q.Empty() {
		q.head = &newTail
		q.tail = &newTail
		q.len++
		return &newTail
	}

	tail := q.tail
	tail.next = &newTail
	q.tail = &newTail
	q.len++
	return &newTail
}

// Dequeue - returns and remove the first node
func (q *queue) Dequeue() (interface{}, error) {
	if q.Empty() {
		return nil, errors.New("error: can not remove from empty queue")
	}
	head := q.head
	q.head = q.head.next
	q.len--
	return head.Data, nil
}

// Front - returns the first node
func (q *queue) Front() (interface{}, error) {
	if q.Empty() {
		return nil, errors.New("error: can not remove from empty queue")
	}
	return q.head.Data, nil
}

// Len - returns amount of nodes in queue
func (q *queue) Len() int {
	return q.len
}

// Empty - returns true if queue is empty
func (q *queue) Empty() bool {
	return q.len == 0
}
