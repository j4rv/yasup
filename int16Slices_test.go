package slices_test

import "testing"
import "github.com/j4rv/slices"

func Test_Int16Shuffle(t *testing.T) {
	shuffles := [][]int16{}
	for i := 0; i < 8; i++ {
		or := []int16{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		slices.Int16Shuffle(&or)
		shuffles = append(shuffles, or)
	}
	for i := range shuffles {
		for j := range shuffles {
			if i == j {
				continue
			}
			if slices.Int16Equals(shuffles[i], shuffles[j]) {
				// If there is any collision in 8 shuffles, the Shuffle function is probably broken
				t.Fail()
			}
		}
	}
}

func Test_Int16Equals(t *testing.T) {
	type TestCase struct {
		name string
		a, b []int16
		exp  bool
	}
	tcs := []TestCase{
		// nil checks
		{"Equals nil", nil, nil, true},
		{"Left nil, right empty", nil, []int16{}, false},
		{"Right nil, left empty", []int16{}, nil, false},
		{"Left nil, right not empty", nil, []int16{-32768}, false},
		{"Right nil, left not empty", []int16{-32768}, nil, false},

		{"Equals empty", []int16{}, []int16{}, true},
		{"Equals 1", []int16{-32768}, []int16{-32768}, true},
		{"Equals 2", []int16{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, []int16{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, true},
		{"Not equals 1", []int16{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, []int16{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, -32768}, false},
		{"Not equals 2", []int16{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, []int16{}, false},
		{"Not equals 3", []int16{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, []int16{-32768, -32768}, false},
	}
	for _, tc := range tcs {
		got := slices.Int16Equals(tc.a, tc.b)
		if got != tc.exp {
			t.Error(tc.name)
		}
	}
}
