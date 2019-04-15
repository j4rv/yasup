package slices

import (
	crypto "crypto/rand"
	"math/big"
	"math/rand"
)

//Float32Insert will append elem at the position i
func Float32Insert(elem float32, sl *[]float32, i int) {
	*sl = append(*sl, elem)
	copy((*sl)[i+1:], (*sl)[i:])
	(*sl)[i] = elem
}

//Float32FastShuffle will randomly swap the float32 elements of a slice using math/rand (fast but not cryptographycally secure).
func Float32FastShuffle(sp []float32) {
	rand.Shuffle(len(sp), func(i, j int) {
		sp[i], sp[j] = sp[j], sp[i]
	})
}

//Float32SecureShuffle will randomly swap the float32 elements of a slice using crypto/rand (resource intensive but cryptographycally secure).
func Float32SecureShuffle(sp []float32) {
	for i := int64(len(sp) - 1); i >= 0; i-- {
		bigRandI, err := crypto.Int(crypto.Reader, big.NewInt(i+1))
		if err != nil {
			panic(err)
		}
		randI := bigRandI.Int64()
		sp[i], sp[randI] = sp[randI], sp[i]
	}
}

//Float32Equals compares two float32 slices. Returns true if their elements are equal.
func Float32Equals(a, b []float32) bool {
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
