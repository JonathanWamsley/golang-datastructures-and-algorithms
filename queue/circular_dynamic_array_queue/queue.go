package circular_dynamic_array_queue

import (
	"errors"
)

const DefaultCapacity = 16

// item - holds the data
type item struct {
	Data interface{}
}

// queue - queue implementing FIFO operations
type queue struct {
	items    []item
	head     int
	tail     int
	len      int
	capacity int
}

// New - creates a new queue that grows in size
func New() *queue {
	return &queue{
		items:    make([]item, DefaultCapacity),
		capacity: DefaultCapacity,
	}
}

// Enqueue - adds item to end of queue
func (q *queue) Enqueue(data interface{}) (interface{}, error) {
	if q.Full() {
		q.Double()
	}
	d := item{Data: data}
	q.items[q.tail] = d
	q.tail = (q.tail + 1) % q.capacity
	q.len++
	return d, nil
}

// Dequeue - removes front item in queue
func (q *queue) Dequeue() (interface{}, error) {
	if q.Empty() {
		return nil, errors.New("error: can not remove from empty queue")
	}

	data := q.items[q.head].Data
	q.items[q.head] = item{}
	q.head = (q.head + 1) % q.capacity
	q.len--

	return data, nil
}

// Front - returns the first element
func (q *queue) Front() (interface{}, error) {
	if q.Empty() {
		return nil, errors.New("error: can not get element from empty queue")
	}
	return q.items[q.head].Data, nil
}

// Empty - returns true if queue is empty
func (q *queue) Empty() bool {
	return q.len == 0
}

// Full - returns true if queue is at max capacity
func (q *queue) Full() bool {
	return q.len == q.capacity
}

// Double - doubles the size of the array
func (q *queue) Double() {
	q.items = append(q.items, make([]item, q.len)...)
	q.capacity = len(q.items)
	q.head = 0
	q.tail = q.len
}

// Shrink - removes half the array when 3/4 unoccupied
func (q *queue) Shrink() {
	// TODO: implement me
}

// Len - returns amount of nodes in queue
func (q *queue) Len() int {
	return q.len
}
