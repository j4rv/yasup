// Code generated by yasupGen; DO NOT EDIT.

package yasup_test

import (
	yasup "github.com/j4rv/yasup"
	"testing"
)

func Test_Float64Insert(t *testing.T) {
	type testCase struct {
		name     string
		slice    []float64
		insertAt int
	}
	base := []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	tcs := []testCase{
		{"First", []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 0},
		{"Middle", []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, len(base) / 2},
		{"Last", []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, len(base)},
		{"Empty slice", []float64{}, 0},
		{"Nil slice", nil, 0},
	}
	for _, tc := range tcs {
		yasup.Float64Insert(&tc.slice, -9223372036854775807, tc.insertAt)
		if tc.slice[tc.insertAt] != -9223372036854775807 {
			t.Error(tc)
		}
	}
}

func Test_Float64FastShuffle(t *testing.T) {
	shuffles := [][]float64{}
	for i := 0; i < 8; i++ {
		or := []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		yasup.Float64FastShuffle(or)
		shuffles = append(shuffles, or)
	}
	for i := range shuffles {
		for j := range shuffles {
			if i >= j {
				continue
			}
			if yasup.Float64Equals(shuffles[i], shuffles[j]) {
				// If there is any collision in 8 shuffles, the Shuffle function is probably broken
				t.Fail()
			}
		}
	}
	// check that nil does not panic
	yasup.Float64FastShuffle(nil)
}

func Test_Float64SecureShuffle(t *testing.T) {
	shuffles := [][]float64{}
	for i := 0; i < 8; i++ {
		or := []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		yasup.Float64SecureShuffle(or)
		shuffles = append(shuffles, or)
	}
	for i := range shuffles {
		for j := range shuffles {
			if i >= j {
				continue
			}
			if yasup.Float64Equals(shuffles[i], shuffles[j]) {
				// If there is any collision in 8 shuffles, the Shuffle function is probably broken
				t.Fail()
			}
		}
	}
	// check that nil does not panic
	yasup.Float64SecureShuffle(nil)
}

func Test_Float64Equals(t *testing.T) {
	type TestCase struct {
		name string
		a, b []float64
		exp  bool
	}
	tcs := []TestCase{
		// nil checks
		{"Equals nil", nil, nil, true},
		// golang treats empty and nil slices as the same thing in most cases, we'll do the same
		{"Left nil, right empty", nil, []float64{}, true},
		{"Right nil, left empty", []float64{}, nil, true},
		{"Left nil, right not empty", nil, []float64{-9223372036854775807}, false},
		{"Right nil, left not empty", []float64{-9223372036854775807}, nil, false},

		{"Equals empty", []float64{}, []float64{}, true},
		{"Equals 1", []float64{-9223372036854775807}, []float64{-9223372036854775807}, true},
		{"Equals 2", []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, true},
		{"Not equals 1", []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, -9223372036854775807}, false},
		{"Not equals 2", []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, []float64{}, false},
		{"Not equals 3", []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, []float64{-9223372036854775807, -9223372036854775807}, false},
	}
	for _, tc := range tcs {
		got := yasup.Float64Equals(tc.a, tc.b)
		if got != tc.exp {
			t.Error(tc)
		}
	}
}
