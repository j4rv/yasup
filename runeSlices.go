package slices

import (
	crypto "crypto/rand"
	"math/big"
	"math/rand"
)

//RuneInsert will append elem at the position i
func RuneInsert(elem rune, sl *[]rune, i int) {
	*sl = append(*sl, elem)
	copy((*sl)[i+1:], (*sl)[i:])
	(*sl)[i] = elem
}

//RuneFastShuffle will randomly swap the rune elements of a slice using math/rand (fast but not cryptographycally secure).
func RuneFastShuffle(sp []rune) {
	rand.Shuffle(len(sp), func(i, j int) {
		sp[i], sp[j] = sp[j], sp[i]
	})
}

//RuneSecureShuffle will randomly swap the rune elements of a slice using crypto/rand (resource intensive but cryptographycally secure).
func RuneSecureShuffle(sp []rune) {
	for i := int64(len(sp) - 1); i >= 0; i-- {
		bigRandI, err := crypto.Int(crypto.Reader, big.NewInt(i+1))
		if err != nil {
			panic(err)
		}
		randI := bigRandI.Int64()
		sp[i], sp[randI] = sp[randI], sp[i]
	}
}

//RuneEquals compares two rune slices. Returns true if their elements are equal.
func RuneEquals(a, b []rune) bool {
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
