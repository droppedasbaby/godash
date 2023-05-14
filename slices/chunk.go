package slices

// Chunk returns a slice of slices of size n. If the slice cannot be evenly divided into chunks of size n, the last
// chunk will be smaller than size n.
func Chunk[T any](slice []T, n int) (cs [][]T) {
	cs = [][]T{}
	for i := 0; i < len(slice); i += n {
		min := i + n
		if len(slice) < min {
			min = len(slice)
		}
		cs = append(cs, slice[i:min])
	}
	return
}
