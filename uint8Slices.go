package slices

import (
	crypto "crypto/rand"
	"math/big"
	"math/rand"
)

//Uint8Insert will append elem at the position i
func Uint8Insert(elem uint8, sl *[]uint8, i int) {
	*sl = append(*sl, elem)
	copy((*sl)[i+1:], (*sl)[i:])
	(*sl)[i] = elem
}

//Uint8FastShuffle will randomly swap the uint8 elements of a slice using math/rand (fast but not cryptographycally secure).
func Uint8FastShuffle(sp []uint8) {
	rand.Shuffle(len(sp), func(i, j int) {
		sp[i], sp[j] = sp[j], sp[i]
	})
}

//Uint8SecureShuffle will randomly swap the uint8 elements of a slice using crypto/rand (resource intensive but cryptographycally secure).
func Uint8SecureShuffle(sp []uint8) {
	for i := int64(len(sp) - 1); i >= 0; i-- {
		bigRandI, err := crypto.Int(crypto.Reader, big.NewInt(i+1))
		if err != nil {
			panic(err)
		}
		randI := bigRandI.Int64()
		sp[i], sp[randI] = sp[randI], sp[i]
	}
}

//Uint8Equals compares two uint8 slices. Returns true if their elements are equal.
func Uint8Equals(a, b []uint8) bool {
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
