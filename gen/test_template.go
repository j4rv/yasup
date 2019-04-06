package main

import "text/template"

var testTemplate = template.Must(template.New("").Parse(`
package slices_test

import "testing"
import "github.com/j4rv/slices"

func Test_{{.TypeCased}}FastShuffle(t *testing.T) {
	shuffles := [][]{{.Type}}{}
	for i := 0; i < 8; i++ {
		or := []{{.Type}}{ {{.MultipleVals}} }
		slices.{{.TypeCased}}FastShuffle(or)
		shuffles = append(shuffles, or)
	}
	for i := range shuffles {
		for j := range shuffles {
			if i == j {
				continue
			}
			if slices.{{.TypeCased}}Equals(shuffles[i], shuffles[j]) {
				// If there is any collision in 8 shuffles, the Shuffle function is probably broken
				t.Fail()
			}
		}
	}
	// check that nil does not panic
	slices.{{.TypeCased}}FastShuffle(nil)
}

func Test_{{.TypeCased}}SecureShuffle(t *testing.T) {
	shuffles := [][]{{.Type}}{}
	for i := 0; i < 8; i++ {
		or := []{{.Type}}{ {{.MultipleVals}} }
		slices.{{.TypeCased}}SecureShuffle(or)
		shuffles = append(shuffles, or)
	}
	for i := range shuffles {
		for j := range shuffles {
			if i == j {
				continue
			}
			if slices.{{.TypeCased}}Equals(shuffles[i], shuffles[j]) {
				// If there is any collision in 8 shuffles, the Shuffle function is probably broken
				t.Fail()
			}
		}
	}
	// check that nil does not panic
	slices.{{.TypeCased}}SecureShuffle(nil)
}

func Test_{{.TypeCased}}Equals(t *testing.T) {
	type TestCase struct {
		name string
		a, b []{{.Type}}
		exp  bool
	}
	tcs := []TestCase{
		// nil checks
		TestCase{"Equals nil", nil, nil, true},
		// golang treats empty and nil slices as the same thing in most cases, we'll do the same
		TestCase{"Left nil, right empty", nil, []{{.Type}}{}, true},
		TestCase{"Right nil, left empty", []{{.Type}}{}, nil, true},
		TestCase{"Left nil, right not empty", nil, []{{.Type}}{ {{.SingleVal}} }, false},
		TestCase{"Right nil, left not empty", []{{.Type}}{ {{.SingleVal}} }, nil, false},

		TestCase{"Equals empty", []{{.Type}}{}, []{{.Type}}{}, true},
		TestCase{"Equals 1", []{{.Type}}{ {{.SingleVal}} }, []{{.Type}}{ {{.SingleVal}} }, true},
		TestCase{"Equals 2", []{{.Type}}{ {{.MultipleVals}} }, []{{.Type}}{ {{.MultipleVals}} }, true},
		TestCase{"Not equals 1", []{{.Type}}{ {{.MultipleVals}} }, []{{.Type}}{ {{.MultipleVals}}, {{.SingleVal}} }, false},
		TestCase{"Not equals 2", []{{.Type}}{ {{.MultipleVals}} }, []{{.Type}}{}, false},
		TestCase{"Not equals 3", []{{.Type}}{ {{.MultipleVals}} }, []{{.Type}}{ {{.SingleVal}}, {{.SingleVal}} }, false},
	}
	for _, tc := range tcs {
		got := slices.{{.TypeCased}}Equals(tc.a, tc.b)
		if got != tc.exp {
			t.Error(tc.name)
		}
	}
}
`))
