package slices_test

import "testing"
import "github.com/j4rv/slices"

func Test_UintShuffle(t *testing.T) {
	shuffles := [][]uint{}
	for i := 0; i < 8; i++ {
		or := []uint{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		slices.UintShuffle(&or)
		shuffles = append(shuffles, or)
	}
	for i := range shuffles {
		for j := range shuffles {
			if i == j {
				continue
			}
			if slices.UintEquals(shuffles[i], shuffles[j]) {
				// If there is any collision in 8 shuffles, the Shuffle function is probably broken
				t.Fail()
			}
		}
	}
}

func Test_UintEquals(t *testing.T) {
	type TestCase struct {
		name string
		a, b []uint
		exp  bool
	}
	tcs := []TestCase{
		// nil checks
		{"Equals nil", nil, nil, true},
		{"Left nil, right empty", nil, []uint{}, false},
		{"Right nil, left empty", []uint{}, nil, false},
		{"Left nil, right not empty", nil, []uint{4294967295}, false},
		{"Right nil, left not empty", []uint{4294967295}, nil, false},

		{"Equals empty", []uint{}, []uint{}, true},
		{"Equals 1", []uint{4294967295}, []uint{4294967295}, true},
		{"Equals 2", []uint{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, []uint{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, true},
		{"Not equals 1", []uint{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, []uint{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 4294967295}, false},
		{"Not equals 2", []uint{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, []uint{}, false},
		{"Not equals 3", []uint{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, []uint{4294967295, 4294967295}, false},
	}
	for _, tc := range tcs {
		got := slices.UintEquals(tc.a, tc.b)
		if got != tc.exp {
			t.Error(tc.name)
		}
	}
}
