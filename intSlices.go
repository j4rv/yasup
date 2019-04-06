package slices

import (
	crypto "crypto/rand"
	"math/big"
	"math/rand"
)

//IntFastShuffle will randomly swap the int elements of a slice using math/rand (fast but not cryptographycally secure).
func IntFastShuffle(sp []int) {
	rand.Shuffle(len(sp), func(i, j int) {
		sp[i], sp[j] = sp[j], sp[i]
	})
}

//IntSecureShuffle will randomly swap the int elements of a slice using crypto/rand (resource intensive but cryptographycally secure).
func IntSecureShuffle(sp []int) {
	for i := len(sp) - 1; i > 0; i-- {
		bigRandI, err := crypto.Int(crypto.Reader, big.NewInt(int64(i)))
		if err != nil {
			panic(err)
		}
		randI := bigRandI.Int64()
		sp[i], sp[randI] = sp[randI], sp[i]
	}
}

//IntEquals compares two int slices. Returns true if their elements are equal.
func IntEquals(a, b []int) bool {
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
