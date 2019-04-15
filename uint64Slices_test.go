package slices_test

import "testing"
import "github.com/j4rv/slices"

func Test_Uint64Insert(t *testing.T) {
	type testCase struct {
		name     string
		slice    []uint64
		insertAt int
	}
	base := []uint64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	tcs := []testCase{
		{"First", []uint64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 0},
		{"Middle", []uint64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, len(base) / 2},
		{"Last", []uint64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, len(base)},
		{"Empty slice", []uint64{}, 0},
		{"Nil slice", nil, 0},
	}
	for _, tc := range tcs {
		slices.Uint64Insert(18446744073709551615, &tc.slice, tc.insertAt)
		if tc.slice[tc.insertAt] != 18446744073709551615 {
			t.Error(tc)
		}
	}
}

func Test_Uint64FastShuffle(t *testing.T) {
	shuffles := [][]uint64{}
	for i := 0; i < 8; i++ {
		or := []uint64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		slices.Uint64FastShuffle(or)
		shuffles = append(shuffles, or)
	}
	for i := range shuffles {
		for j := range shuffles {
			if i >= j {
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
			if i >= j {
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
			t.Error(tc)
		}
	}
}
