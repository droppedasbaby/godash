package stacks

// Push adds an element to the top of the stack.
func Push[T any](s *Stack[T], elem T) {
	*s = append(*s, elem)
}
