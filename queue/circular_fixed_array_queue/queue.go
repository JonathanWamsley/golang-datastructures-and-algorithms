package circular_fixed_array_queue

import "errors"

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

// New - creates a new queue with the specified capacity
func New(capacity int) *queue {
	if capacity <= 0 {
		return nil
	}
	return &queue{
		items:    make([]item, capacity),
		capacity: capacity,
	}
}

// Enqueue - adds item to end of queue
func (q *queue) Enqueue(d interface{}) error {
	if q.Full() {
		return errors.New("error: queue capacity reached")
	}

	q.len++
	q.items[q.tail] = item{Data: d}

	q.tail = (q.tail + 1) % q.capacity
	return nil
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
	d := q.items[q.head].Data
	return d, nil
}

// Empty - returns true if queue is empty
func (q *queue) Empty() bool {
	return q.len == 0
}

// Full - returns true if queue is at max capacity
func (q *queue) Full() bool {
	return q.len == q.capacity
}

// Len - returns amount of nodes in queue
func (q *queue) Len() int {
	return q.len
}
