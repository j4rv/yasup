package slices_test

import "testing"
import "github.com/j4rv/slices"

// TODO: Test that a nil slice does not panic in FastShuffle and SecureShuffle

func Test_Int16FastShuffle(t *testing.T) {
	shuffles := [][]int16{}
	for i := 0; i < 8; i++ {
		or := []int16{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		slices.Int16FastShuffle(or)
		shuffles = append(shuffles, or)
	}
	for i := range shuffles {
		for j := range shuffles {
			if i == j {
				continue
			}
			if slices.Int16Equals(shuffles[i], shuffles[j]) {
				// If there is any collision in 8 shuffles, the Shuffle function is probably broken
				t.Fail()
			}
		}
	}
	// check that nil does not panic
	slices.Int16FastShuffle(nil)
}

func Test_Int16SecureShuffle(t *testing.T) {
	shuffles := [][]int16{}
	for i := 0; i < 8; i++ {
		or := []int16{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		slices.Int16SecureShuffle(or)
		shuffles = append(shuffles, or)
	}
	for i := range shuffles {
		for j := range shuffles {
			if i == j {
				continue
			}
			if slices.Int16Equals(shuffles[i], shuffles[j]) {
				// If there is any collision in 8 shuffles, the Shuffle function is probably broken
				t.Fail()
			}
		}
	}
	// check that nil does not panic
	slices.Int16SecureShuffle(nil)
}

func Test_Int16Equals(t *testing.T) {
	type TestCase struct {
		name string
		a, b []int16
		exp  bool
	}
	tcs := []TestCase{
		// nil checks
		{"Equals nil", nil, nil, true},
		// golang treats empty and nil slices as the same thing in most cases, we'll do the same
		{"Left nil, right empty", nil, []int16{}, true},
		{"Right nil, left empty", []int16{}, nil, true},
		{"Left nil, right not empty", nil, []int16{-32768}, false},
		{"Right nil, left not empty", []int16{-32768}, nil, false},

		{"Equals empty", []int16{}, []int16{}, true},
		{"Equals 1", []int16{-32768}, []int16{-32768}, true},
		{"Equals 2", []int16{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, []int16{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, true},
		{"Not equals 1", []int16{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, []int16{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, -32768}, false},
		{"Not equals 2", []int16{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, []int16{}, false},
		{"Not equals 3", []int16{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, []int16{-32768, -32768}, false},
	}
	for _, tc := range tcs {
		got := slices.Int16Equals(tc.a, tc.b)
		if got != tc.exp {
			t.Error(tc.name)
		}
	}
}
