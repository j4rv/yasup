package slices

import (
	crypto "crypto/rand"
	"math/big"
	"math/rand"
)

//BoolInsert will append elem at the position i
func BoolInsert(elem bool, sl *[]bool, i int) {
	*sl = append(*sl, elem)
	copy((*sl)[i+1:], (*sl)[i:])
	(*sl)[i] = elem
}

//BoolFastShuffle will randomly swap the bool elements of a slice using math/rand (fast but not cryptographycally secure).
func BoolFastShuffle(sp []bool) {
	rand.Shuffle(len(sp), func(i, j int) {
		sp[i], sp[j] = sp[j], sp[i]
	})
}

//BoolSecureShuffle will randomly swap the bool elements of a slice using crypto/rand (resource intensive but cryptographycally secure).
func BoolSecureShuffle(sp []bool) {
	for i := int64(len(sp) - 1); i >= 0; i-- {
		bigRandI, err := crypto.Int(crypto.Reader, big.NewInt(i+1))
		if err != nil {
			panic(err)
		}
		randI := bigRandI.Int64()
		sp[i], sp[randI] = sp[randI], sp[i]
	}
}

//BoolEquals compares two bool slices. Returns true if their elements are equal.
func BoolEquals(a, b []bool) bool {
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
