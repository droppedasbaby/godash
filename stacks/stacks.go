package stacks

type Stack[T any] []T

// NewStack returns a new stack. FILO.
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

// Push pushes an element to the stack. FILO.
func (s *Stack[T]) Push(v T) {
	*s = append(*s, v)
}

// Pop pops an element from the stack. FILO.
func (s *Stack[T]) Pop() T {
	size := s.Size()
	v := (*s)[size-1]
	*s = (*s)[:size-1]
	return v
}

// Size returns the number of elements in the stack.
func (s *Stack[T]) Size() int {
	return len(*s)
}

// Peek returns the top element of the stack without removing it.
func (s *Stack[T]) Peek() T {
	size := s.Size()
	return (*s)[size-1]
}

// IsEmpty returns true if the stack is empty.
func (s *Stack[T]) IsEmpty() bool {
	size := s.Size()
	return size == 0
}
