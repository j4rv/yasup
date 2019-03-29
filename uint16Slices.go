package slices

import "math/rand"

//Uint16Shuffle will randomly swap the uint16 elements of a slice.
func Uint16Shuffle(sp *[]uint16) {
	slice := *sp
	rand.Shuffle(len(slice), func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})
}

//Uint16Equals compares two uint16 slices. Returns true if their elements are equal.
func Uint16Equals(a, b []uint16) bool {
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
