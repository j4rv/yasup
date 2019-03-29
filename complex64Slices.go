package slices

import "math/rand"

//Complex64Shuffle will randomly swap the complex64 elements of a slice.
func Complex64Shuffle(sp *[]complex64) {
	slice := *sp
	rand.Shuffle(len(slice), func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})
}

//Complex64Equals compares two complex64 slices. Returns true if their elements are equal.
func Complex64Equals(a, b []complex64) bool {
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
