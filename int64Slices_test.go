package slices_test

import "testing"
import "github.com/j4rv/slices"

// TODO: Test that a nil slice does not panic in FastShuffle and SecureShuffle

func Test_Int64FastShuffle(t *testing.T) {
	shuffles := [][]int64{}
	for i := 0; i < 8; i++ {
		or := []int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		slices.Int64FastShuffle(or)
		shuffles = append(shuffles, or)
	}
	for i := range shuffles {
		for j := range shuffles {
			if i == j {
				continue
			}
			if slices.Int64Equals(shuffles[i], shuffles[j]) {
				// If there is any collision in 8 shuffles, the Shuffle function is probably broken
				t.Fail()
			}
		}
	}
	// check that nil does not panic
	slices.Int64FastShuffle(nil)
}

func Test_Int64SecureShuffle(t *testing.T) {
	shuffles := [][]int64{}
	for i := 0; i < 8; i++ {
		or := []int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		slices.Int64SecureShuffle(or)
		shuffles = append(shuffles, or)
	}
	for i := range shuffles {
		for j := range shuffles {
			if i == j {
				continue
			}
			if slices.Int64Equals(shuffles[i], shuffles[j]) {
				// If there is any collision in 8 shuffles, the Shuffle function is probably broken
				t.Fail()
			}
		}
	}
	// check that nil does not panic
	slices.Int64SecureShuffle(nil)
}

func Test_Int64Equals(t *testing.T) {
	type TestCase struct {
		name string
		a, b []int64
		exp  bool
	}
	tcs := []TestCase{
		// nil checks
		{"Equals nil", nil, nil, true},
		// golang treats empty and nil slices as the same thing in most cases, we'll do the same
		{"Left nil, right empty", nil, []int64{}, true},
		{"Right nil, left empty", []int64{}, nil, true},
		{"Left nil, right not empty", nil, []int64{-9223372036854775807}, false},
		{"Right nil, left not empty", []int64{-9223372036854775807}, nil, false},

		{"Equals empty", []int64{}, []int64{}, true},
		{"Equals 1", []int64{-9223372036854775807}, []int64{-9223372036854775807}, true},
		{"Equals 2", []int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, []int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, true},
		{"Not equals 1", []int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, []int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, -9223372036854775807}, false},
		{"Not equals 2", []int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, []int64{}, false},
		{"Not equals 3", []int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, []int64{-9223372036854775807, -9223372036854775807}, false},
	}
	for _, tc := range tcs {
		got := slices.Int64Equals(tc.a, tc.b)
		if got != tc.exp {
			t.Error(tc.name)
		}
	}
}
