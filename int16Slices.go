package slices

import "math/rand"

//Int16Shuffle will randomly swap the int16 elements of a slice.
func Int16Shuffle(sp *[]int16) {
	slice := *sp
	rand.Shuffle(len(slice), func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})
}

//Int16Equals compares two int16 slices. Returns true if their elements are equal.
func Int16Equals(a, b []int16) bool {
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
