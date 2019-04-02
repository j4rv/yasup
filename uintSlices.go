package slices

import (
	crypto "crypto/rand"
	"math/big"
	"math/rand"
)

// TODO: Test that a nil slice does not panic in FastShuffle and SecureShuffle

//UintFastShuffle will randomly swap the uint elements of a slice using math/big (fast but not cryptographycally secure).
func UintFastShuffle(sp []uint) {
	rand.Shuffle(len(sp), func(i, j int) {
		sp[i], sp[j] = sp[j], sp[i]
	})
}

//UintSecureShuffle will randomly swap the uint elements of a slice using crypto/rand (resource intensive but cryptographycally secure).
func UintSecureShuffle(sp []uint) {
	for i := len(sp) - 1; i > 0; i-- {
		bigRandI, err := crypto.Int(crypto.Reader, big.NewInt(int64(i)))
		if err != nil {
			panic(err)
		}
		randI := bigRandI.Int64()
		sp[i], sp[randI] = sp[randI], sp[i]
	}
}

//UintEquals compares two uint slices. Returns true if their elements are equal.
func UintEquals(a, b []uint) bool {
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
