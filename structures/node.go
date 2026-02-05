package structures

type node[T comparable] struct {
	value T
	next  *node[T]
}
