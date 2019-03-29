package slices

import "math/rand"

//IntShuffle will randomly swap the int elements of a slice.
func IntShuffle(sp *[]int) {
	slice := *sp
	rand.Shuffle(len(slice), func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})
}

//IntEquals compares two int slices. Returns true if their elements are equal.
func IntEquals(a, b []int) bool {
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
