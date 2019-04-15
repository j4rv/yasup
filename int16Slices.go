package slices

import (
	crypto "crypto/rand"
	"math/big"
	"math/rand"
)

//Int16Insert will append elem at the position i
func Int16Insert(elem int16, sl *[]int16, i int) {
	*sl = append(*sl, elem)
	copy((*sl)[i+1:], (*sl)[i:])
	(*sl)[i] = elem
}

//Int16FastShuffle will randomly swap the int16 elements of a slice using math/rand (fast but not cryptographycally secure).
func Int16FastShuffle(sp []int16) {
	rand.Shuffle(len(sp), func(i, j int) {
		sp[i], sp[j] = sp[j], sp[i]
	})
}

//Int16SecureShuffle will randomly swap the int16 elements of a slice using crypto/rand (resource intensive but cryptographycally secure).
func Int16SecureShuffle(sp []int16) {
	for i := int64(len(sp) - 1); i >= 0; i-- {
		bigRandI, err := crypto.Int(crypto.Reader, big.NewInt(i+1))
		if err != nil {
			panic(err)
		}
		randI := bigRandI.Int64()
		sp[i], sp[randI] = sp[randI], sp[i]
	}
}

//Int16Equals compares two int16 slices. Returns true if their elements are equal.
func Int16Equals(a, b []int16) bool {
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
