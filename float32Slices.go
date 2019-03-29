package slices

import "math/rand"

//Float32Shuffle will randomly swap the float32 elements of a slice.
func Float32Shuffle(sp *[]float32) {
	slice := *sp
	rand.Shuffle(len(slice), func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})
}

//Float32Equals compares two float32 slices. Returns true if their elements are equal.
func Float32Equals(a, b []float32) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false // one of them is nil, the other is not.
	}
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
