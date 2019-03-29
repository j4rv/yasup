package slices

import "math/rand"

//Uint64Shuffle will randomly swap the uint64 elements of a slice.
func Uint64Shuffle(sp *[]uint64) {
	slice := *sp
	rand.Shuffle(len(slice), func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})
}

//Uint64Equals compares two uint64 slices. Returns true if their elements are equal.
func Uint64Equals(a, b []uint64) bool {
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
