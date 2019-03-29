package slices

import "math/rand"

//Uint8Shuffle will randomly swap the uint8 elements of a slice.
func Uint8Shuffle(sp *[]uint8) {
	slice := *sp
	rand.Shuffle(len(slice), func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})
}

//Uint8Equals compares two uint8 slices. Returns true if their elements are equal.
func Uint8Equals(a, b []uint8) bool {
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
