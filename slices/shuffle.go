package slices

import (
	"math/rand"
	"time"
)

func shuffleWithSeed[T any](s []T, seed int64) []T {
	r := rand.New(rand.NewSource(seed))
	for i := range s {
		j := r.Intn(i + 1)
		s[i], s[j] = s[j], s[i]
	}
	return s
}

// Shuffle shuffles the slice in place.
func Shuffle[T any](s []T) []T {
	return shuffleWithSeed(s, time.Now().UnixNano())
}

// ShuffleWithSeed shuffles the slice with the provided seed in place.
func ShuffleWithSeed[T any](s []T, seed int64) []T {
	return shuffleWithSeed(s, seed)
}
