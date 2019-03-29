package slices

import "math/rand"

//Int8Shuffle will randomly swap the int8 elements of a slice.
func Int8Shuffle(sp *[]int8) {
	slice := *sp
	rand.Shuffle(len(slice), func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})
}

//Int8Equals compares two int8 slices. Returns true if their elements are equal.
func Int8Equals(a, b []int8) bool {
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
