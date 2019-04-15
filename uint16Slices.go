package slices

import (
	crypto "crypto/rand"
	"math/big"
	"math/rand"
)

//Uint16Insert will append elem at the position i
func Uint16Insert(elem uint16, sl *[]uint16, i int) {
	*sl = append(*sl, elem)
	copy((*sl)[i+1:], (*sl)[i:])
	(*sl)[i] = elem
}

//Uint16FastShuffle will randomly swap the uint16 elements of a slice using math/rand (fast but not cryptographycally secure).
func Uint16FastShuffle(sp []uint16) {
	rand.Shuffle(len(sp), func(i, j int) {
		sp[i], sp[j] = sp[j], sp[i]
	})
}

//Uint16SecureShuffle will randomly swap the uint16 elements of a slice using crypto/rand (resource intensive but cryptographycally secure).
func Uint16SecureShuffle(sp []uint16) {
	for i := int64(len(sp) - 1); i >= 0; i-- {
		bigRandI, err := crypto.Int(crypto.Reader, big.NewInt(i+1))
		if err != nil {
			panic(err)
		}
		randI := bigRandI.Int64()
		sp[i], sp[randI] = sp[randI], sp[i]
	}
}

//Uint16Equals compares two uint16 slices. Returns true if their elements are equal.
func Uint16Equals(a, b []uint16) bool {
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
