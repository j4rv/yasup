package slices

import (
	crypto "crypto/rand"
	"math/big"
	"math/rand"
)

//ByteInsert will append elem at the position i
func ByteInsert(elem byte, sl *[]byte, i int) {
	*sl = append(*sl, elem)
	copy((*sl)[i+1:], (*sl)[i:])
	(*sl)[i] = elem
}

//ByteFastShuffle will randomly swap the byte elements of a slice using math/rand (fast but not cryptographycally secure).
func ByteFastShuffle(sp []byte) {
	rand.Shuffle(len(sp), func(i, j int) {
		sp[i], sp[j] = sp[j], sp[i]
	})
}

//ByteSecureShuffle will randomly swap the byte elements of a slice using crypto/rand (resource intensive but cryptographycally secure).
func ByteSecureShuffle(sp []byte) {
	for i := int64(len(sp) - 1); i >= 0; i-- {
		bigRandI, err := crypto.Int(crypto.Reader, big.NewInt(i+1))
		if err != nil {
			panic(err)
		}
		randI := bigRandI.Int64()
		sp[i], sp[randI] = sp[randI], sp[i]
	}
}

//ByteEquals compares two byte slices. Returns true if their elements are equal.
func ByteEquals(a, b []byte) bool {
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
