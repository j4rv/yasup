package slices

import "math/rand"

//Uint32Shuffle will randomly swap the uint32 elements of a slice.
func Uint32Shuffle(sp *[]uint32) {
	slice := *sp
	rand.Shuffle(len(slice), func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})
}

//Uint32Equals compares two uint32 slices. Returns true if their elements are equal.
func Uint32Equals(a, b []uint32) bool {
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
