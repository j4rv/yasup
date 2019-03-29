package slices_test

import "testing"
import "github.com/j4rv/slices"

func Test_ByteShuffle(t *testing.T) {
	shuffles := [][]byte{}
	for i := 0; i < 8; i++ {
		or := []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		slices.ByteShuffle(&or)
		shuffles = append(shuffles, or)
	}
	for i := range shuffles {
		for j := range shuffles {
			if i == j {
				continue
			}
			if slices.ByteEquals(shuffles[i], shuffles[j]) {
				// If there is any collision in 8 shuffles, the Shuffle function is probably broken
				t.Fail()
			}
		}
	}
}

func Test_ByteEquals(t *testing.T) {
	type TestCase struct {
		name string
		a, b []byte
		exp  bool
	}
	tcs := []TestCase{
		// nil checks
		{"Equals nil", nil, nil, true},
		{"Left nil, right empty", nil, []byte{}, false},
		{"Right nil, left empty", []byte{}, nil, false},
		{"Left nil, right not empty", nil, []byte{255}, false},
		{"Right nil, left not empty", []byte{255}, nil, false},

		{"Equals empty", []byte{}, []byte{}, true},
		{"Equals 1", []byte{255}, []byte{255}, true},
		{"Equals 2", []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, true},
		{"Not equals 1", []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 255}, false},
		{"Not equals 2", []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, []byte{}, false},
		{"Not equals 3", []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, []byte{255, 255}, false},
	}
	for _, tc := range tcs {
		got := slices.ByteEquals(tc.a, tc.b)
		if got != tc.exp {
			t.Error(tc.name)
		}
	}
}
