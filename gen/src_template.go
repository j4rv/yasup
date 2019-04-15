package main

import "text/template"

var srcTemplate = template.Must(template.New("").Parse(`
package slices

import (
	crypto "crypto/rand"
	"math/big"
	"math/rand"
)

//{{.TypeCased}}Insert will append elem at the position i
func {{.TypeCased}}Insert(elem {{.Type}}, sl *[]{{.Type}}, i int) {
	*sl = append(*sl, elem)
	copy((*sl)[i+1:], (*sl)[i:])
	(*sl)[i] = elem
}

//{{.TypeCased}}FastShuffle will randomly swap the {{.Type}} elements of a slice using math/rand (fast but not cryptographycally secure).
func {{.TypeCased}}FastShuffle(sp []{{.Type}}) {
	rand.Shuffle(len(sp), func(i, j int) {
		sp[i], sp[j] = sp[j], sp[i]
	})
}

//{{.TypeCased}}SecureShuffle will randomly swap the {{.Type}} elements of a slice using crypto/rand (resource intensive but cryptographycally secure).
func {{.TypeCased}}SecureShuffle(sp []{{.Type}}) {
	for i := int64(len(sp) - 1); i >= 0; i-- {
		bigRandI, err := crypto.Int(crypto.Reader, big.NewInt(i+1))
		if err != nil {
			panic(err)
		}
		randI := bigRandI.Int64()
		sp[i], sp[randI] = sp[randI], sp[i]
	}
}

//{{.TypeCased}}Equals compares two {{.Type}} slices. Returns true if their elements are equal.
func {{.TypeCased}}Equals(a, b []{{.Type}}) bool {
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
`))
