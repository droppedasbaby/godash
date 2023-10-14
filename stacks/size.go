package stacks

// Size returns the number of elements in the stack.
func Size[T any](s *Stack[T]) int {
	return len(*s)
}
