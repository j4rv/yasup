package slices

import "math/rand"

//Int32Shuffle will randomly swap the int32 elements of a slice.
func Int32Shuffle(sp *[]int32) {
	slice := *sp
	rand.Shuffle(len(slice), func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})
}

//Int32Equals compares two int32 slices. Returns true if their elements are equal.
func Int32Equals(a, b []int32) bool {
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
