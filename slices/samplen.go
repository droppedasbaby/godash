package slices

import (
	"math/rand"
	"time"
)

func sampleNWithSeed[T any](s []T, n int, seed int64) (res []T) {
	r := rand.New(rand.NewSource(seed))
	res = make([]T, n)
	for i := 0; i < n; i++ {
		res[i] = s[r.Intn(len(s))]
	}
	return
}

// SampleN returns n random elements from the slice.
func SampleN[T any](s []T, n int) []T {
	return sampleNWithSeed(s, n, time.Now().UnixNano())
}

// SampleNWithSeed returns n random elements from the slice with the provided seed.
func SampleNWithSeed[T any](s []T, n int, seed int64) []T {
	return sampleNWithSeed(s, n, seed)
}
