package slices_test

import "testing"
import "github.com/j4rv/slices"

func Test_Int8Shuffle(t *testing.T) {
	shuffles := [][]int8{}
	for i := 0; i < 8; i++ {
		or := []int8{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		slices.Int8Shuffle(&or)
		shuffles = append(shuffles, or)
	}
	for i := range shuffles {
		for j := range shuffles {
			if i == j {
				continue
			}
			if slices.Int8Equals(shuffles[i], shuffles[j]) {
				// If there is any collision in 8 shuffles, the Shuffle function is probably broken
				t.Fail()
			}
		}
	}
}

func Test_Int8Equals(t *testing.T) {
	type TestCase struct {
		name string
		a, b []int8
		exp  bool
	}
	tcs := []TestCase{
		// nil checks
		{"Equals nil", nil, nil, true},
		{"Left nil, right empty", nil, []int8{}, false},
		{"Right nil, left empty", []int8{}, nil, false},
		{"Left nil, right not empty", nil, []int8{-128}, false},
		{"Right nil, left not empty", []int8{-128}, nil, false},

		{"Equals empty", []int8{}, []int8{}, true},
		{"Equals 1", []int8{-128}, []int8{-128}, true},
		{"Equals 2", []int8{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, []int8{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, true},
		{"Not equals 1", []int8{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, []int8{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, -128}, false},
		{"Not equals 2", []int8{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, []int8{}, false},
		{"Not equals 3", []int8{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, []int8{-128, -128}, false},
	}
	for _, tc := range tcs {
		got := slices.Int8Equals(tc.a, tc.b)
		if got != tc.exp {
			t.Error(tc.name)
		}
	}
}
