package slices

import "math/rand"

//Complex128Shuffle will randomly swap the complex128 elements of a slice.
func Complex128Shuffle(sp *[]complex128) {
	slice := *sp
	rand.Shuffle(len(slice), func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})
}

//Complex128Equals compares two complex128 slices. Returns true if their elements are equal.
func Complex128Equals(a, b []complex128) bool {
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
