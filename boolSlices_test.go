package slices_test

import "testing"
import "github.com/j4rv/slices"

func Test_BoolShuffle(t *testing.T) {
	shuffles := [][]bool{}
	for i := 0; i < 8; i++ {
		or := []bool{true, false, true, false, true, false, true, false, true, false, true, false}
		slices.BoolShuffle(&or)
		shuffles = append(shuffles, or)
	}
	for i := range shuffles {
		for j := range shuffles {
			if i == j {
				continue
			}
			if slices.BoolEquals(shuffles[i], shuffles[j]) {
				// If there is any collision in 8 shuffles, the Shuffle function is probably broken
				t.Fail()
			}
		}
	}
}

func Test_BoolEquals(t *testing.T) {
	type TestCase struct {
		name string
		a, b []bool
		exp  bool
	}
	tcs := []TestCase{
		// nil checks
		{"Equals nil", nil, nil, true},
		{"Left nil, right empty", nil, []bool{}, false},
		{"Right nil, left empty", []bool{}, nil, false},
		{"Left nil, right not empty", nil, []bool{false}, false},
		{"Right nil, left not empty", []bool{false}, nil, false},

		{"Equals empty", []bool{}, []bool{}, true},
		{"Equals 1", []bool{false}, []bool{false}, true},
		{"Equals 2", []bool{true, false, true, false, true, false, true, false, true, false, true, false}, []bool{true, false, true, false, true, false, true, false, true, false, true, false}, true},
		{"Not equals 1", []bool{true, false, true, false, true, false, true, false, true, false, true, false}, []bool{true, false, true, false, true, false, true, false, true, false, true, false, false}, false},
		{"Not equals 2", []bool{true, false, true, false, true, false, true, false, true, false, true, false}, []bool{}, false},
		{"Not equals 3", []bool{true, false, true, false, true, false, true, false, true, false, true, false}, []bool{false, false}, false},
	}
	for _, tc := range tcs {
		got := slices.BoolEquals(tc.a, tc.b)
		if got != tc.exp {
			t.Error(tc.name)
		}
	}
}
