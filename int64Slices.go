package slices

import (
	crypto "crypto/rand"
	"math/big"
	"math/rand"
)

// TODO: Test that a nil slice does not panic in FastShuffle and SecureShuffle

//Int64FastShuffle will randomly swap the int64 elements of a slice using math/big (fast but not cryptographycally secure).
func Int64FastShuffle(sp []int64) {
	rand.Shuffle(len(sp), func(i, j int) {
		sp[i], sp[j] = sp[j], sp[i]
	})
}

//Int64SecureShuffle will randomly swap the int64 elements of a slice using crypto/rand (resource intensive but cryptographycally secure).
func Int64SecureShuffle(sp []int64) {
	for i := len(sp) - 1; i > 0; i-- {
		bigRandI, err := crypto.Int(crypto.Reader, big.NewInt(int64(i)))
		if err != nil {
			panic(err)
		}
		randI := bigRandI.Int64()
		sp[i], sp[randI] = sp[randI], sp[i]
	}
}

//Int64Equals compares two int64 slices. Returns true if their elements are equal.
func Int64Equals(a, b []int64) bool {
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
