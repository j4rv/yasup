package slices

import (
	crypto "crypto/rand"
	"math/big"
	"math/rand"
)

//StringFastShuffle will randomly swap the string elements of a slice using math/rand (fast but not cryptographycally secure).
func StringFastShuffle(sp []string) {
	rand.Shuffle(len(sp), func(i, j int) {
		sp[i], sp[j] = sp[j], sp[i]
	})
}

//StringSecureShuffle will randomly swap the string elements of a slice using crypto/rand (resource intensive but cryptographycally secure).
func StringSecureShuffle(sp []string) {
	for i := len(sp) - 1; i > 0; i-- {
		bigRandI, err := crypto.Int(crypto.Reader, big.NewInt(int64(i)))
		if err != nil {
			panic(err)
		}
		randI := bigRandI.Int64()
		sp[i], sp[randI] = sp[randI], sp[i]
	}
}

//StringEquals compares two string slices. Returns true if their elements are equal.
func StringEquals(a, b []string) bool {
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
