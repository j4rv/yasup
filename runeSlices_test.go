package slices_test

import "testing"
import "github.com/j4rv/slices"

func Test_RuneShuffle(t *testing.T) {
	shuffles := [][]rune{}
	for i := 0; i < 8; i++ {
		or := []rune{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		slices.RuneShuffle(&or)
		shuffles = append(shuffles, or)
	}
	for i := range shuffles {
		for j := range shuffles {
			if i == j {
				continue
			}
			if slices.RuneEquals(shuffles[i], shuffles[j]) {
				// If there is any collision in 8 shuffles, the Shuffle function is probably broken
				t.Fail()
			}
		}
	}
}

func Test_RuneEquals(t *testing.T) {
	type TestCase struct {
		name string
		a, b []rune
		exp  bool
	}
	tcs := []TestCase{
		// nil checks
		{"Equals nil", nil, nil, true},
		{"Left nil, right empty", nil, []rune{}, false},
		{"Right nil, left empty", []rune{}, nil, false},
		{"Left nil, right not empty", nil, []rune{2147483647}, false},
		{"Right nil, left not empty", []rune{2147483647}, nil, false},

		{"Equals empty", []rune{}, []rune{}, true},
		{"Equals 1", []rune{2147483647}, []rune{2147483647}, true},
		{"Equals 2", []rune{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, []rune{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, true},
		{"Not equals 1", []rune{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, []rune{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 2147483647}, false},
		{"Not equals 2", []rune{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, []rune{}, false},
		{"Not equals 3", []rune{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, []rune{2147483647, 2147483647}, false},
	}
	for _, tc := range tcs {
		got := slices.RuneEquals(tc.a, tc.b)
		if got != tc.exp {
			t.Error(tc.name)
		}
	}
}
