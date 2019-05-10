// Code generated by yasupGen; DO NOT EDIT.

package yasup

import (
	crypto "crypto/rand"
	"math/big"
	"math/rand"
)

var zeroValueUint uint

//UintContains will return true if elem is present in the slice and false otherwise.
func UintContains(sl []uint, elem uint) bool {
	for i := range sl {
		if sl[i] == elem {
			return true
		}
	}
	return false
}

//UintInsert will append elem at the position i. Might return ErrIndexOutOfBounds.
func UintInsert(sl *[]uint, elem uint, i int) error {
	if i < 0 || i > len(*sl) {
		return ErrIndexOutOfBounds
	}
	*sl = append(*sl, elem)
	copy((*sl)[i+1:], (*sl)[i:])
	(*sl)[i] = elem
	return nil
}

//UintDelete delete the element at the position i. Might return ErrIndexOutOfBounds.
func UintDelete(sl *[]uint, i int) error {
	if i < 0 || i >= len(*sl) {
		return ErrIndexOutOfBounds
	}
	*sl = append((*sl)[:i], (*sl)[i+1:]...)
	return nil
}

//UintPush is equivalent to UintInsert with index len(*sl)
func UintPush(sl *[]uint, elem uint) {
	UintInsert(sl, elem, len(*sl))
}

//UintFrontPush is equivalent to UintInsert with index 0
func UintFrontPush(sl *[]uint, elem uint) {
	UintInsert(sl, elem, 0)
}

//UintPop is equivalent to getting and removing the last element of the slice. Might return ErrEmptySlice.
func UintPop(sl *[]uint) (uint, error) {
	if len(*sl) == 0 {
		return zeroValueUint, ErrEmptySlice
	}
	last := len(*sl) - 1
	ret := (*sl)[last]
	UintDelete(sl, last)
	return ret, nil
}

//UintPop is equivalent to getting and removing the first element of the slice. Might return ErrEmptySlice.
func UintFrontPop(sl *[]uint) (uint, error) {
	if len(*sl) == 0 {
		return zeroValueUint, ErrEmptySlice
	}
	ret := (*sl)[0]
	UintDelete(sl, 0)
	return ret, nil
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

//UintFastShuffle will randomly swap the uint elements of a slice using math/rand (fast but not cryptographycally secure).
func UintFastShuffle(sp []uint) {
	rand.Shuffle(len(sp), func(i, j int) {
		sp[i], sp[j] = sp[j], sp[i]
	})
}

//UintSecureShuffle will randomly swap the uint elements of a slice using crypto/rand (resource intensive but cryptographycally secure).
func UintSecureShuffle(sp []uint) error {
	var i int64
	size := int64(len(sp)) - 1
	for i = 0; i < size+1; i++ {
		bigRandI, err := crypto.Int(crypto.Reader, big.NewInt(size))
		if err != nil {
			return err
		}
		randI := bigRandI.Int64()
		sp[size-i], sp[randI] = sp[randI], sp[size-i]
	}
	return nil
}
