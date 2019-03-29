package slices

import "math/rand"

//Int64Shuffle will randomly swap the int64 elements of a slice.
func Int64Shuffle(sp *[]int64) {
	slice := *sp
	rand.Shuffle(len(slice), func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})
}

//Int64Equals compares two int64 slices. Returns true if their elements are equal.
func Int64Equals(a, b []int64) bool {
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
