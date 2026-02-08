package structures

import "errors"

var ErrEmptyQueue = errors.New("empty queue")

type queue[T comparable] struct {
	size   int64
	top    *node[T]
	bottom *node[T]
}

func NewQueue[T comparable]() Queuer[T] {
	return &queue[T]{}
}

func (q *queue[T]) Push(value T) {
	if q.size == 0 {
		q.bottom = &node[T]{
			value: value,
			next:  nil,
		}
		q.top = q.bottom
	} else {
		node := &node[T]{
			value: value,
			next:  nil,
		}
		q.bottom.next = node
		q.bottom = node
	}

	q.size++
}

func (q *queue[T]) Pop() (T, error) {
	if q.size == 0 {
		var t T
		return t, ErrEmptyQueue
	}

	topValue := q.top.value
	q.top = q.top.next

	if q.size == 1 {
		q.bottom = nil
	}

	q.size--

	return topValue, nil
}

func (q *queue[T]) Find(value T) bool {
	if q.size == 0 {
		return false
	}

	node := q.top

	for node != nil {
		if node.value == value {
			return true
		}

		node = node.next
	}

	return false
}

func (q *queue[T]) Top() (T, error) {
	if q.size == 0 {
		var t T
		return t, ErrEmptyQueue
	}

	return q.top.value, nil
}

func (q *queue[T]) Bottom() (T, error) {
	if q.size == 0 {
		var t T
		return t, ErrEmptyQueue
	}

	return q.bottom.value, nil
}

func (q *queue[T]) Size() int64 {
	return q.size
}
