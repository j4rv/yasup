package slices_test

import "testing"
import "github.com/j4rv/slices"

// TODO: Test that a nil slice does not panic in FastShuffle and SecureShuffle

func Test_IntFastShuffle(t *testing.T) {
	shuffles := [][]int{}
	for i := 0; i < 8; i++ {
		or := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		slices.IntFastShuffle(or)
		shuffles = append(shuffles, or)
	}
	for i := range shuffles {
		for j := range shuffles {
			if i == j {
				continue
			}
			if slices.IntEquals(shuffles[i], shuffles[j]) {
				// If there is any collision in 8 shuffles, the Shuffle function is probably broken
				t.Fail()
			}
		}
	}
	// check that nil does not panic
	slices.IntFastShuffle(nil)
}

func Test_IntSecureShuffle(t *testing.T) {
	shuffles := [][]int{}
	for i := 0; i < 8; i++ {
		or := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		slices.IntSecureShuffle(or)
		shuffles = append(shuffles, or)
	}
	for i := range shuffles {
		for j := range shuffles {
			if i == j {
				continue
			}
			if slices.IntEquals(shuffles[i], shuffles[j]) {
				// If there is any collision in 8 shuffles, the Shuffle function is probably broken
				t.Fail()
			}
		}
	}
	// check that nil does not panic
	slices.IntSecureShuffle(nil)
}

func Test_IntEquals(t *testing.T) {
	type TestCase struct {
		name string
		a, b []int
		exp  bool
	}
	tcs := []TestCase{
		// nil checks
		{"Equals nil", nil, nil, true},
		// golang treats empty and nil slices as the same thing in most cases, we'll do the same
		{"Left nil, right empty", nil, []int{}, true},
		{"Right nil, left empty", []int{}, nil, true},
		{"Left nil, right not empty", nil, []int{-2147483648}, false},
		{"Right nil, left not empty", []int{-2147483648}, nil, false},

		{"Equals empty", []int{}, []int{}, true},
		{"Equals 1", []int{-2147483648}, []int{-2147483648}, true},
		{"Equals 2", []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, true},
		{"Not equals 1", []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, -2147483648}, false},
		{"Not equals 2", []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, []int{}, false},
		{"Not equals 3", []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, []int{-2147483648, -2147483648}, false},
	}
	for _, tc := range tcs {
		got := slices.IntEquals(tc.a, tc.b)
		if got != tc.exp {
			t.Error(tc.name)
		}
	}
}
