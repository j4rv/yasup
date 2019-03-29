package slices

import "math/rand"

//Float64Shuffle will randomly swap the float64 elements of a slice.
func Float64Shuffle(sp *[]float64) {
	slice := *sp
	rand.Shuffle(len(slice), func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})
}

//Float64Equals compares two float64 slices. Returns true if their elements are equal.
func Float64Equals(a, b []float64) bool {
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
