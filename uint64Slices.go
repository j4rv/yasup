package slices

import (
	crypto "crypto/rand"
	"math/big"
	"math/rand"
)

//Uint64FastShuffle will randomly swap the uint64 elements of a slice using math/rand (fast but not cryptographycally secure).
func Uint64FastShuffle(sp []uint64) {
	rand.Shuffle(len(sp), func(i, j int) {
		sp[i], sp[j] = sp[j], sp[i]
	})
}

//Uint64SecureShuffle will randomly swap the uint64 elements of a slice using crypto/rand (resource intensive but cryptographycally secure).
func Uint64SecureShuffle(sp []uint64) {
	for i := len(sp) - 1; i > 0; i-- {
		bigRandI, err := crypto.Int(crypto.Reader, big.NewInt(int64(i)))
		if err != nil {
			panic(err)
		}
		randI := bigRandI.Int64()
		sp[i], sp[randI] = sp[randI], sp[i]
	}
}

//Uint64Equals compares two uint64 slices. Returns true if their elements are equal.
func Uint64Equals(a, b []uint64) bool {
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
