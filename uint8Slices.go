package slices

import (
	crypto "crypto/rand"
	"math/big"
	"math/rand"
)

// TODO: Test that a nil slice does not panic in FastShuffle and SecureShuffle

//Uint8FastShuffle will randomly swap the uint8 elements of a slice using math/big (fast but not cryptographycally secure).
func Uint8FastShuffle(sp []uint8) {
	rand.Shuffle(len(sp), func(i, j int) {
		sp[i], sp[j] = sp[j], sp[i]
	})
}

//Uint8SecureShuffle will randomly swap the uint8 elements of a slice using crypto/rand (resource intensive but cryptographycally secure).
func Uint8SecureShuffle(sp []uint8) {
	for i := len(sp) - 1; i > 0; i-- {
		bigRandI, err := crypto.Int(crypto.Reader, big.NewInt(int64(i)))
		if err != nil {
			panic(err)
		}
		randI := bigRandI.Int64()
		sp[i], sp[randI] = sp[randI], sp[i]
	}
}

//Uint8Equals compares two uint8 slices. Returns true if their elements are equal.
func Uint8Equals(a, b []uint8) bool {
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
