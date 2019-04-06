package slices_test

import "testing"
import "github.com/j4rv/slices"

// TODO: Test that a nil slice does not panic in FastShuffle and SecureShuffle

func Test_Uint64FastShuffle(t *testing.T) {
	shuffles := [][]uint64{}
	for i := 0; i < 8; i++ {
		or := []uint64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		slices.Uint64FastShuffle(or)
		shuffles = append(shuffles, or)
	}
	for i := range shuffles {
		for j := range shuffles {
			if i == j {
				continue
			}
			if slices.Uint64Equals(shuffles[i], shuffles[j]) {
				// If there is any collision in 8 shuffles, the Shuffle function is probably broken
				t.Fail()
			}
		}
	}
	// check that nil does not panic
	slices.Uint64FastShuffle(nil)
}

func Test_Uint64SecureShuffle(t *testing.T) {
	shuffles := [][]uint64{}
	for i := 0; i < 8; i++ {
		or := []uint64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		slices.Uint64SecureShuffle(or)
		shuffles = append(shuffles, or)
	}
	for i := range shuffles {
		for j := range shuffles {
			if i == j {
				continue
			}
			if slices.Uint64Equals(shuffles[i], shuffles[j]) {
				// If there is any collision in 8 shuffles, the Shuffle function is probably broken
				t.Fail()
			}
		}
	}
	// check that nil does not panic
	slices.Uint64SecureShuffle(nil)
}

func Test_Uint64Equals(t *testing.T) {
	type TestCase struct {
		name string
		a, b []uint64
		exp  bool
	}
	tcs := []TestCase{
		// nil checks
		{"Equals nil", nil, nil, true},
		// golang treats empty and nil slices as the same thing in most cases, we'll do the same
		{"Left nil, right empty", nil, []uint64{}, true},
		{"Right nil, left empty", []uint64{}, nil, true},
		{"Left nil, right not empty", nil, []uint64{18446744073709551615}, false},
		{"Right nil, left not empty", []uint64{18446744073709551615}, nil, false},

		{"Equals empty", []uint64{}, []uint64{}, true},
		{"Equals 1", []uint64{18446744073709551615}, []uint64{18446744073709551615}, true},
		{"Equals 2", []uint64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, []uint64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, true},
		{"Not equals 1", []uint64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, []uint64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 18446744073709551615}, false},
		{"Not equals 2", []uint64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, []uint64{}, false},
		{"Not equals 3", []uint64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, []uint64{18446744073709551615, 18446744073709551615}, false},
	}
	for _, tc := range tcs {
		got := slices.Uint64Equals(tc.a, tc.b)
		if got != tc.exp {
			t.Error(tc.name)
		}
	}
}
