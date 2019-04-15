package slices

import (
	crypto "crypto/rand"
	"math/big"
	"math/rand"
)

//Complex128Insert will append elem at the position i
func Complex128Insert(elem complex128, sl *[]complex128, i int) {
	*sl = append(*sl, elem)
	copy((*sl)[i+1:], (*sl)[i:])
	(*sl)[i] = elem
}

//Complex128FastShuffle will randomly swap the complex128 elements of a slice using math/rand (fast but not cryptographycally secure).
func Complex128FastShuffle(sp []complex128) {
	rand.Shuffle(len(sp), func(i, j int) {
		sp[i], sp[j] = sp[j], sp[i]
	})
}

//Complex128SecureShuffle will randomly swap the complex128 elements of a slice using crypto/rand (resource intensive but cryptographycally secure).
func Complex128SecureShuffle(sp []complex128) {
	for i := int64(len(sp) - 1); i >= 0; i-- {
		bigRandI, err := crypto.Int(crypto.Reader, big.NewInt(i+1))
		if err != nil {
			panic(err)
		}
		randI := bigRandI.Int64()
		sp[i], sp[randI] = sp[randI], sp[i]
	}
}

//Complex128Equals compares two complex128 slices. Returns true if their elements are equal.
func Complex128Equals(a, b []complex128) bool {
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
