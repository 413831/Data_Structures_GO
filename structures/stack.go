package structures

import "errors"

var ErrEmptyStack = errors.New("empty stack")

type stack[T comparable] struct {
	size int64
	top  *node[T]
}

func NewStack[T comparable]() Stacker[T] {
	return &stack[T]{}
}

func (s *stack[T]) Pop() (T, error) {
	if s.size == 0 {
		var t T
		return t, ErrEmptyStack
	}

	topValue := s.top.value
	s.top = s.top.next
	s.size--

	return topValue, nil
}

func (s *stack[T]) Push(value T) {
	s.top = &node[T]{
		value: value,
		next:  s.top,
	}
	s.size++
}

func (s *stack[T]) Find(value T) bool {
	if s.size == 0 {
		return false
	}

	node := s.top

	for node.value != value {
		node = node.next

		if node == nil {
			return false
		}
	}

	return true
}

func (s stack[T]) Top() (T, error) {
	if s.size == 0 {
		var t T
		return t, ErrEmptyStack
	}

	return s.top.value, nil
}

func (s stack[T]) Bottom() (T, error) {
	if s.size == 0 {
		var t T
		return t, ErrEmptyStack
	}

	node := s.top
	bottom := node

	for node != nil {
		bottom = node
		node = node.next
	}

	return bottom.value, nil
}

func (s *stack[T]) Size() int64 {
	return s.size
}
