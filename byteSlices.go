package slices

import (
	crypto "crypto/rand"
	"math/big"
	"math/rand"
)

// TODO: Test that a nil slice does not panic in FastShuffle and SecureShuffle

//ByteFastShuffle will randomly swap the byte elements of a slice using math/big (fast but not cryptographycally secure).
func ByteFastShuffle(sp []byte) {
	rand.Shuffle(len(sp), func(i, j int) {
		sp[i], sp[j] = sp[j], sp[i]
	})
}

//ByteSecureShuffle will randomly swap the byte elements of a slice using crypto/rand (resource intensive but cryptographycally secure).
func ByteSecureShuffle(sp []byte) {
	for i := len(sp) - 1; i > 0; i-- {
		bigRandI, err := crypto.Int(crypto.Reader, big.NewInt(int64(i)))
		if err != nil {
			panic(err)
		}
		randI := bigRandI.Int64()
		sp[i], sp[randI] = sp[randI], sp[i]
	}
}

//ByteEquals compares two byte slices. Returns true if their elements are equal.
func ByteEquals(a, b []byte) bool {
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
