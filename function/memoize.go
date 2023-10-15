package function

//import (
//	"fmt"
//	"github.com/GrewalAS/godash"
//	"sync"
//)
//
//type CacheKeyer interface {
//	Key() string
//}
//
//// Memoize returns a function that caches the results of f using the CacheKeyer provided for the key.
//func Memoize[R any, T CacheKeyer](f func(T) R) func(T) R {
//	cache := make(map[string]R)
//	mu := sync.RWMutex{}
//
//	return func(c T) R {
//		key := c.Key()
//
//		mu.RLock()
//		result, exists := cache[key]
//		mu.RUnlock()
//		if exists {
//			return result
//		}
//		result = f(c)
//
//		mu.Lock()
//		cache[key] = result
//		mu.Unlock()
//
//		return result
//	}
//}
//
//// comparableKeyer is a custom type that wraps a Comparable type and implements the CacheKeyer interface.
//type comparableKeyer[C comparable] struct {
//	v C
//}
//
//// Key returns the string representation of the ComparableKeyer value using the fmt.Sprintf function.
//func (ck comparableKeyer[C]) Key() string {
//	return fmt.Sprintf("%v", ck.v)
//}
//
//// MemoizeComparable returns a function that caches the results of f using a type that implements Comparable for the key.
//func MemoizeComparable[C comparable, R any](f func(C) R) func(comparableKeyer[C]) R {
//	return Memoize(func(keyer comparableKeyer[C]) R {
//		return f(keyer.v)
//	})
//}
//
//// Define a function that operates on an int.
//func myFunction(i int) int {
//	// Your logic here.
//	return i * 2
//}
//
//func all() {
//	memoizedFunc := MemoizeComparable(myFunction)
//	_ := memoizedFunc(1)
//}
//
//// hashableKeyer is a custom type that wraps a Hashable type and implements the CacheKeyer interface.
//type hashableKeyer[H godash.Hashable] struct {
//	v H
//}
//
//// Key returns the string representation of the HashableKeyer value using the Hash method.
//func (ck hashableKeyer[H]) Key() string {
//	return ck.v.Hash()
//}
//
//// MemoizeHashable returns a function that caches the results of f using a type that implements Hashable for the key.
//func MemoizeHashable[H godash.Hashable, R any](f func(H) R) func(keyer hashableKeyer[H]) R {
//	return Memoize(func(keyer hashableKeyer[H]) R {
//		return f(keyer.v)
//	})
//}
