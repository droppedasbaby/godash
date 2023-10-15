package function

import (
	"fmt"
	"github.com/GrewalAS/godash"
	"github.com/GrewalAS/godash/utils"
	"sync"
)

// CacheKeyer is an interface that represents any type that can be used as a key in a cache.
type CacheKeyer interface {
	Key() string
}

// Memoize returns a function that caches the results of f using the CacheKeyer provided for the key.
// The function f should return a value of type R, which will be cached.
func Memoize[R any, CK CacheKeyer](f func(CK) R) func(CK) R {
	cache := make(map[string]R)
	mu := sync.RWMutex{}

	return func(ck CK) R {
		key := ck.Key()

		mu.RLock()
		result, exists := cache[key]
		mu.RUnlock()
		if exists {
			return result
		}

		result = f(ck)

		mu.Lock()
		cache[key] = result
		mu.Unlock()

		return result
	}
}

// MemoizeResult returns a function that caches the results of f using the CacheKeyer provided for the key.
// The function f should return a utils.Result[R], which will be cached.
func MemoizeResult[R any, CK CacheKeyer](f func(CK) utils.Result[R]) func(CK) utils.Result[R] {
	cache := make(map[string]utils.Result[R])
	mu := sync.RWMutex{}

	return func(ck CK) utils.Result[R] {
		key := ck.Key()

		mu.RLock()
		result, exists := cache[key]
		mu.RUnlock()

		if exists {
			return result
		}

		result = f(ck)

		mu.Lock()
		cache[key] = result
		mu.Unlock()

		return result
	}
}

// comparableKeyer is a custom type that wraps a Comparable type and implements the CacheKeyer interface for memoization.
type comparableKeyer[C comparable] struct {
	v C
}

// Key converts the comparable value to a string that will be used as a cache key.
func (ck comparableKeyer[C]) Key() string {
	return fmt.Sprintf("%v", ck.v)
}

// MemoizeComparable caches the results of a function that takes a comparable argument and returns a value of the same
// type.
func MemoizeComparable[C comparable, R any](f func(C) R) func(C) R {
	memoized := Memoize(func(ck comparableKeyer[C]) R {
		return f(ck.v)
	})

	return func(c C) R {
		return memoized(comparableKeyer[C]{v: c})
	}
}

// hashableKeyer is a custom type that wraps a Hashable type and implements the CacheKeyer interface for memoization.
type hashableKeyer[H godash.Hashable] struct {
	v H
}

// Key uses the Hash method from the Hashable interface to generate a string key for caching.
func (ck hashableKeyer[H]) Key() string {
	return ck.v.Hash()
}

// MemoizeHashable caches the results of a function that takes a hashable argument and returns a value of the same type.
func MemoizeHashable[H godash.Hashable, R any](f func(H) R) func(H) R {
	memoized := Memoize(func(ck hashableKeyer[H]) R {
		return f(ck.v)
	})

	return func(h H) R {
		return memoized(hashableKeyer[H]{v: h})
	}
}
