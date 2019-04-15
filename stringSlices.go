package slices

import (
	crypto "crypto/rand"
	"math/big"
	"math/rand"
)

//StringInsert will append elem at the position i
func StringInsert(elem string, sl *[]string, i int) {
	*sl = append(*sl, elem)
	copy((*sl)[i+1:], (*sl)[i:])
	(*sl)[i] = elem
}

//StringFastShuffle will randomly swap the string elements of a slice using math/rand (fast but not cryptographycally secure).
func StringFastShuffle(sp []string) {
	rand.Shuffle(len(sp), func(i, j int) {
		sp[i], sp[j] = sp[j], sp[i]
	})
}

//StringSecureShuffle will randomly swap the string elements of a slice using crypto/rand (resource intensive but cryptographycally secure).
func StringSecureShuffle(sp []string) {
	for i := int64(len(sp) - 1); i >= 0; i-- {
		bigRandI, err := crypto.Int(crypto.Reader, big.NewInt(i+1))
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
