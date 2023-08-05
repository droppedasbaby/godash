package stacks

// Pop removes and returns the top element of the stack.
// If the stack is empty, it panics like accessing an index out of range.
func Pop[T any](s *Stack[T]) T {
	top := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return top
}
