package structures

type Sizer[T comparable] interface {
	Size() int64
}

type Finder[T comparable] interface {
	Find(value T) bool
}

type Stacker[T comparable] interface {
	Sizer[T]
	Finder[T]
	Pop() (T, error)
	Push(value T)
	Top() (T, error)
	Bottom() (T, error)
}

type Queuer[T comparable] interface {
	Sizer[T]
	Finder[T]
	Pop() (T, error)
	Push(value T)
	Top() (T, error)
	Bottom() (T, error)
}

type SimpleLinkedLister[T comparable] interface {
	Sizer[T]
	Finder[T]
	PushAt(index int, value T) error
	PopAt(index int) (T, error)
	First() (T, error)
	Last() (T, error)
	GetAt(index int) (T, error)
}

type DoubleLinkedLister[T comparable] interface {
	Sizer[T]
	Finder[T]
	PushAt(index int, value T) error
	PopAt(index int) (T, error)
	First() (T, error)
	Last() (T, error)
	GetAt(index int) (T, error)
}
