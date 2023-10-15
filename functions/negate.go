package functions

// Negate returns a function that negates the result of the given function f.
func Negate(f func(...any) bool) func(...any) bool {
	return func(args ...any) bool {
		return !f(args...)
	}
}
