package structures

import "errors"

var (
	ErrEmptyDoubleLinkedList             = errors.New("empty list")
	ErrIndexOutOfRangeInDoubleLinkedList = errors.New("index out of range in double linked list")
)

type doubleLinkedList[T comparable] struct {
	first *doubleNode[T]
	size  int64
}

func NewDoubleLinkedList[T comparable]() DoubleLinkedLister[T] {
	return &doubleLinkedList[T]{}
}

func (dl *doubleLinkedList[T]) Size() int64 {
	return dl.size
}

func (dl doubleLinkedList[T]) Find(value T) bool {
	for current := dl.first; current != nil; current = current.next {
		if current.value == value {
			return true
		}
	}

	return false
}

func (dl doubleLinkedList[T]) PushAt(index int, value T) error {
	if index < 0 || index > int(dl.Size()) {
		return ErrIndexOutOfRangeInDoubleLinkedList
	}

	newNode := &doubleNode[T]{value: value}

	if dl.first == nil || index == 0 {
		newNode.next = dl.first

		if dl.first != nil {
			dl.first.prev = newNode
		}

		dl.first = newNode
		dl.size++

		return nil
	}

	current := dl.first

	for i := 0; i < index-1; i++ {
		current = current.next
	}

	newNode.next = current.next
	newNode.prev = current
	if current.next != nil {
		current.next.prev = newNode
	}

	current.next = newNode
	dl.size++

	return nil
}

func (dl doubleLinkedList[T]) PopAt(index int) (T, error) {
	if dl.Size() == 0 {
		var t T
		return t, ErrEmptyDoubleLinkedList
	}

	if index < 0 || index > int(dl.Size()) {
		var t T
		return t, ErrIndexOutOfRangeInDoubleLinkedList
	}

	var value T

	if index == 0 {
		value = dl.first.value
		dl.first = dl.first.next
		if dl.first != nil {
			dl.first.prev = nil
		}

		dl.size--
		return value, nil
	}

	current := dl.first

	for i := 0; i < index-1; i++ {
		current = current.next
	}

	value = current.next.value

	if current.next.next != nil {
		current.next.next.prev = current
	}
	current.next = current.next.next
	dl.size--

	return value, nil
}

func (dl doubleLinkedList[T]) First() (T, error) {
	if dl.Size() == 0 {
		var t T
		return t, ErrEmptyDoubleLinkedList
	}

	return dl.first.value, nil
}

func (dl doubleLinkedList[T]) Last() (T, error) {
	if dl.Size() == 0 {
		var t T
		return t, ErrEmptyDoubleLinkedList
	}

	current := dl.first

	for current.next != nil {
		current = current.next
	}

	return current.value, nil
}

func (dl doubleLinkedList[T]) GetAt(index int) (T, error) {
	if dl.Size() == 0 {
		var t T
		return t, ErrEmptyDoubleLinkedList
	}

	if index < 0 || index > int(dl.Size()) {
		var t T
		return t, ErrIndexOutOfRangeInDoubleLinkedList
	}

	current := dl.first

	for i := 0; i < index-1; i++ {
		current = current.next
	}

	return current.value, nil
}
