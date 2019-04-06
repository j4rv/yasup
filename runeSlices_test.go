package slices_test

import "testing"
import "github.com/j4rv/slices"

// TODO: Test that a nil slice does not panic in FastShuffle and SecureShuffle

func Test_RuneFastShuffle(t *testing.T) {
	shuffles := [][]rune{}
	for i := 0; i < 8; i++ {
		or := []rune{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		slices.RuneFastShuffle(or)
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
	// check that nil does not panic
	slices.RuneFastShuffle(nil)
}

func Test_RuneSecureShuffle(t *testing.T) {
	shuffles := [][]rune{}
	for i := 0; i < 8; i++ {
		or := []rune{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		slices.RuneSecureShuffle(or)
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
	// check that nil does not panic
	slices.RuneSecureShuffle(nil)
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
		// golang treats empty and nil slices as the same thing in most cases, we'll do the same
		{"Left nil, right empty", nil, []rune{}, true},
		{"Right nil, left empty", []rune{}, nil, true},
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
