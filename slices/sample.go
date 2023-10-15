package slices

import (
	"math/rand"
	"time"
)

// Sample returns a random element from the slice.
func Sample[T any](s []T) T {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return s[r.Intn(len(s))]
}

// SampleWithSeed returns a random element from the slice with the provided seed.
func SampleWithSeed[T any](s []T, seed int64) T {
	r := rand.New(rand.NewSource(seed))
	return s[r.Intn(len(s))]
}
