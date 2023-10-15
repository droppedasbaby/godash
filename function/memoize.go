package function

import (
	"sync"
)

// CacheKeyer is an interface that represents any type that can be used as a key in a cache.
type CacheKeyer interface {
	// Key returns a string that can be used as a key in a cache.
	Key(...any) string
}

// Memoize returns a function that caches the results of f using the CacheKeyer provided for the key.
// The function f should return a value of type R, which will be cached.
func Memoize[CK CacheKeyer, R any](f func(...any) R, ck CK) func(args ...any) R {
	cache := make(map[string]R)
	mu := sync.RWMutex{}

	return func(args ...any) R {
		key := ck.Key(args...)

		mu.RLock()
		result, exists := cache[key]
		mu.RUnlock()

		if exists {
			return result
		}

		result = f(args...)

		mu.Lock()
		cache[key] = result
		mu.Unlock()

		return result
	}
}
