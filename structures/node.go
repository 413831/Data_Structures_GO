package structures

type node[T comparable] struct {
	value T
	next  *node[T]
}

type doubleNode[T comparable] struct {
	value T
	prev  *doubleNode[T]
	next  *doubleNode[T]
}
