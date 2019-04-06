package slices_test

import "testing"
import "github.com/j4rv/slices"

func Test_Uint16FastShuffle(t *testing.T) {
	shuffles := [][]uint16{}
	for i := 0; i < 8; i++ {
		or := []uint16{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		slices.Uint16FastShuffle(or)
		shuffles = append(shuffles, or)
	}
	for i := range shuffles {
		for j := range shuffles {
			if i == j {
				continue
			}
			if slices.Uint16Equals(shuffles[i], shuffles[j]) {
				// If there is any collision in 8 shuffles, the Shuffle function is probably broken
				t.Fail()
			}
		}
	}
	// check that nil does not panic
	slices.Uint16FastShuffle(nil)
}

func Test_Uint16SecureShuffle(t *testing.T) {
	shuffles := [][]uint16{}
	for i := 0; i < 8; i++ {
		or := []uint16{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		slices.Uint16SecureShuffle(or)
		shuffles = append(shuffles, or)
	}
	for i := range shuffles {
		for j := range shuffles {
			if i == j {
				continue
			}
			if slices.Uint16Equals(shuffles[i], shuffles[j]) {
				// If there is any collision in 8 shuffles, the Shuffle function is probably broken
				t.Fail()
			}
		}
	}
	// check that nil does not panic
	slices.Uint16SecureShuffle(nil)
}

func Test_Uint16Equals(t *testing.T) {
	type TestCase struct {
		name string
		a, b []uint16
		exp  bool
	}
	tcs := []TestCase{
		// nil checks
		{"Equals nil", nil, nil, true},
		// golang treats empty and nil slices as the same thing in most cases, we'll do the same
		{"Left nil, right empty", nil, []uint16{}, true},
		{"Right nil, left empty", []uint16{}, nil, true},
		{"Left nil, right not empty", nil, []uint16{65535}, false},
		{"Right nil, left not empty", []uint16{65535}, nil, false},

		{"Equals empty", []uint16{}, []uint16{}, true},
		{"Equals 1", []uint16{65535}, []uint16{65535}, true},
		{"Equals 2", []uint16{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, []uint16{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, true},
		{"Not equals 1", []uint16{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, []uint16{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 65535}, false},
		{"Not equals 2", []uint16{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, []uint16{}, false},
		{"Not equals 3", []uint16{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, []uint16{65535, 65535}, false},
	}
	for _, tc := range tcs {
		got := slices.Uint16Equals(tc.a, tc.b)
		if got != tc.exp {
			t.Error(tc.name)
		}
	}
}
