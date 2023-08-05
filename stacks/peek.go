package stacks

// Peek returns the top element of the stack without removing it.
// If the stack is empty, it panics like accessing an index out of range.
func Peek[T any](s *Stack[T]) T {
	return (*s)[len(*s)-1]
}
