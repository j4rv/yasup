package slices_test

import "testing"
import "github.com/j4rv/slices"

func Test_Uint8FastShuffle(t *testing.T) {
	shuffles := [][]uint8{}
	for i := 0; i < 8; i++ {
		or := []uint8{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		slices.Uint8FastShuffle(or)
		shuffles = append(shuffles, or)
	}
	for i := range shuffles {
		for j := range shuffles {
			if i == j {
				continue
			}
			if slices.Uint8Equals(shuffles[i], shuffles[j]) {
				// If there is any collision in 8 shuffles, the Shuffle function is probably broken
				t.Fail()
			}
		}
	}
}

func Test_Uint8SecureShuffle(t *testing.T) {
	shuffles := [][]uint8{}
	for i := 0; i < 8; i++ {
		or := []uint8{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		slices.Uint8SecureShuffle(or)
		shuffles = append(shuffles, or)
	}
	for i := range shuffles {
		for j := range shuffles {
			if i == j {
				continue
			}
			if slices.Uint8Equals(shuffles[i], shuffles[j]) {
				// If there is any collision in 8 shuffles, the Shuffle function is probably broken
				t.Fail()
			}
		}
	}
}

func Test_Uint8Equals(t *testing.T) {
	type TestCase struct {
		name string
		a, b []uint8
		exp  bool
	}
	tcs := []TestCase{
		// nil checks
		{"Equals nil", nil, nil, true},
		// golang treats empty and nil slices as the same thing in most cases, we'll do the same
		{"Left nil, right empty", nil, []uint8{}, true},
		{"Right nil, left empty", []uint8{}, nil, true},
		{"Left nil, right not empty", nil, []uint8{255}, false},
		{"Right nil, left not empty", []uint8{255}, nil, false},

		{"Equals empty", []uint8{}, []uint8{}, true},
		{"Equals 1", []uint8{255}, []uint8{255}, true},
		{"Equals 2", []uint8{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, []uint8{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, true},
		{"Not equals 1", []uint8{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, []uint8{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 255}, false},
		{"Not equals 2", []uint8{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, []uint8{}, false},
		{"Not equals 3", []uint8{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, []uint8{255, 255}, false},
	}
	for _, tc := range tcs {
		got := slices.Uint8Equals(tc.a, tc.b)
		if got != tc.exp {
			t.Error(tc.name)
		}
	}
}
