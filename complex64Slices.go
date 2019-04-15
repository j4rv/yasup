package slices

import (
	crypto "crypto/rand"
	"math/big"
	"math/rand"
)

//Complex64Insert will append elem at the position i
func Complex64Insert(elem complex64, sl *[]complex64, i int) {
	*sl = append(*sl, elem)
	copy((*sl)[i+1:], (*sl)[i:])
	(*sl)[i] = elem
}

//Complex64FastShuffle will randomly swap the complex64 elements of a slice using math/rand (fast but not cryptographycally secure).
func Complex64FastShuffle(sp []complex64) {
	rand.Shuffle(len(sp), func(i, j int) {
		sp[i], sp[j] = sp[j], sp[i]
	})
}

//Complex64SecureShuffle will randomly swap the complex64 elements of a slice using crypto/rand (resource intensive but cryptographycally secure).
func Complex64SecureShuffle(sp []complex64) {
	for i := int64(len(sp) - 1); i >= 0; i-- {
		bigRandI, err := crypto.Int(crypto.Reader, big.NewInt(i+1))
		if err != nil {
			panic(err)
		}
		randI := bigRandI.Int64()
		sp[i], sp[randI] = sp[randI], sp[i]
	}
}

//Complex64Equals compares two complex64 slices. Returns true if their elements are equal.
func Complex64Equals(a, b []complex64) bool {
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
