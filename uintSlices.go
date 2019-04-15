package slices

import (
	crypto "crypto/rand"
	"math/big"
	"math/rand"
)

//UintInsert will append elem at the position i
func UintInsert(elem uint, sl *[]uint, i int) {
	*sl = append(*sl, elem)
	copy((*sl)[i+1:], (*sl)[i:])
	(*sl)[i] = elem
}

//UintFastShuffle will randomly swap the uint elements of a slice using math/rand (fast but not cryptographycally secure).
func UintFastShuffle(sp []uint) {
	rand.Shuffle(len(sp), func(i, j int) {
		sp[i], sp[j] = sp[j], sp[i]
	})
}

//UintSecureShuffle will randomly swap the uint elements of a slice using crypto/rand (resource intensive but cryptographycally secure).
func UintSecureShuffle(sp []uint) {
	for i := int64(len(sp) - 1); i >= 0; i-- {
		bigRandI, err := crypto.Int(crypto.Reader, big.NewInt(i+1))
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
