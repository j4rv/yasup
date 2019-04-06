package slices_test

import "testing"
import "github.com/j4rv/slices"

// TODO: Test that a nil slice does not panic in FastShuffle and SecureShuffle

func Test_Int32FastShuffle(t *testing.T) {
	shuffles := [][]int32{}
	for i := 0; i < 8; i++ {
		or := []int32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		slices.Int32FastShuffle(or)
		shuffles = append(shuffles, or)
	}
	for i := range shuffles {
		for j := range shuffles {
			if i == j {
				continue
			}
			if slices.Int32Equals(shuffles[i], shuffles[j]) {
				// If there is any collision in 8 shuffles, the Shuffle function is probably broken
				t.Fail()
			}
		}
	}
	// check that nil does not panic
	slices.Int32FastShuffle(nil)
}

func Test_Int32SecureShuffle(t *testing.T) {
	shuffles := [][]int32{}
	for i := 0; i < 8; i++ {
		or := []int32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		slices.Int32SecureShuffle(or)
		shuffles = append(shuffles, or)
	}
	for i := range shuffles {
		for j := range shuffles {
			if i == j {
				continue
			}
			if slices.Int32Equals(shuffles[i], shuffles[j]) {
				// If there is any collision in 8 shuffles, the Shuffle function is probably broken
				t.Fail()
			}
		}
	}
	// check that nil does not panic
	slices.Int32SecureShuffle(nil)
}

func Test_Int32Equals(t *testing.T) {
	type TestCase struct {
		name string
		a, b []int32
		exp  bool
	}
	tcs := []TestCase{
		// nil checks
		{"Equals nil", nil, nil, true},
		// golang treats empty and nil slices as the same thing in most cases, we'll do the same
		{"Left nil, right empty", nil, []int32{}, true},
		{"Right nil, left empty", []int32{}, nil, true},
		{"Left nil, right not empty", nil, []int32{-2147483648}, false},
		{"Right nil, left not empty", []int32{-2147483648}, nil, false},

		{"Equals empty", []int32{}, []int32{}, true},
		{"Equals 1", []int32{-2147483648}, []int32{-2147483648}, true},
		{"Equals 2", []int32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, []int32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, true},
		{"Not equals 1", []int32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, []int32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, -2147483648}, false},
		{"Not equals 2", []int32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, []int32{}, false},
		{"Not equals 3", []int32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, []int32{-2147483648, -2147483648}, false},
	}
	for _, tc := range tcs {
		got := slices.Int32Equals(tc.a, tc.b)
		if got != tc.exp {
			t.Error(tc.name)
		}
	}
}
