package stack

type Stack[T any] interface {
	Push(value T)
}

type linkedStack[T any] struct {
	head *node[T]
	size int
}

type node[T any] struct {
	value T
	next  *node[T]
}

func NewStack[T any]() Stack[T] {
	return &linkedStack[T]{}
}

func (s *linkedStack[T]) Push(value T) {
	newNode := &node[T]{value: value}
	if s.head != nil {
		newNode.next = s.head
	}
	s.head = newNode
	s.size++
}
