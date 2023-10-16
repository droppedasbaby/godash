package slices

import (
	"math/rand"
	"time"
)

func sampleWithSeed[T any](s []T, seed int64) T {
	r := rand.New(rand.NewSource(seed))
	return s[r.Intn(len(s))]
}

// Sample returns a random element from the slice.
func Sample[T any](s []T) T {
	return sampleWithSeed(s, time.Now().UnixNano())
}

// SampleWithSeed returns a random element from the slice with the provided seed.
func SampleWithSeed[T any](s []T, seed int64) T {
	return sampleWithSeed(s, seed)
}
