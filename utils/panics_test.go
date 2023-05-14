package utils_test

import (
	"testing"

	"github.com/GrewalAS/godash/utils"
)

func TestPanics(t *testing.T) {
	t.Parallel()

	t.Run("panics", func(t *testing.T) {
		t.Parallel()
		if !utils.Panics(func() { panic("oh no") }) {
			t.Errorf("Panics() = false, want true")
		}
	})

	t.Run("does not panic", func(t *testing.T) {
		t.Parallel()
		if utils.Panics(func() {}) {
			t.Errorf("Panics() = true, want false")
		}
	})
}
