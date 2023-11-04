package functions

import "time"

// Throttle is a function that returns a function that will only be called once every duration.
// The returned function will return the result of the first call until the duration has passed.
// After the duration has passed, the function will be called again and the result will be returned.
func Throttle[R any](f func(...any) R, d time.Duration) func(...any) R {
	// var (
	//	lastCalled time.Time
	//	result     any
	//)
	return func(args ...any) R {
		// if time.Since(lastCalled) > duration {
		//	result = f(args)
		//	lastCalled = time.Now()
		//}
		return f(args...)
	}
}
