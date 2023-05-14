package utils

// Panics returns true if the provided function panics.
// Source: https://www.reddit.com/r/golang/comments/2tj8b3/comment/cnzy3kc
//
//nolint:nonamedreturns // Can't find a way to make this function work without named returns.
func Panics(doesItPanic func()) (panics bool) {
	defer func() {
		if r := recover(); r != nil {
			panics = true
		}
	}()

	doesItPanic()

	return panics
}
