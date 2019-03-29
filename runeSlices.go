package slices

import "math/rand"

//RuneShuffle will randomly swap the rune elements of a slice.
func RuneShuffle(sp *[]rune) {
	slice := *sp
	rand.Shuffle(len(slice), func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})
}

//RuneEquals compares two rune slices. Returns true if their elements are equal.
func RuneEquals(a, b []rune) bool {
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
