package slices

import (
	crypto "crypto/rand"
	"math/big"
	"math/rand"
)

// TODO: Test that a nil slice does not panic in FastShuffle and SecureShuffle

//Uint32FastShuffle will randomly swap the uint32 elements of a slice using math/big (fast but not cryptographycally secure).
func Uint32FastShuffle(sp []uint32) {
	rand.Shuffle(len(sp), func(i, j int) {
		sp[i], sp[j] = sp[j], sp[i]
	})
}

//Uint32SecureShuffle will randomly swap the uint32 elements of a slice using crypto/rand (resource intensive but cryptographycally secure).
func Uint32SecureShuffle(sp []uint32) {
	for i := len(sp) - 1; i > 0; i-- {
		bigRandI, err := crypto.Int(crypto.Reader, big.NewInt(int64(i)))
		if err != nil {
			panic(err)
		}
		randI := bigRandI.Int64()
		sp[i], sp[randI] = sp[randI], sp[i]
	}
}

//Uint32Equals compares two uint32 slices. Returns true if their elements are equal.
func Uint32Equals(a, b []uint32) bool {
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
