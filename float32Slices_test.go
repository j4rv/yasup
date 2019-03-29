package slices_test

import "testing"
import "github.com/j4rv/slices"

func Test_Float32Shuffle(t *testing.T) {
	shuffles := [][]float32{}
	for i := 0; i < 8; i++ {
		or := []float32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		slices.Float32Shuffle(&or)
		shuffles = append(shuffles, or)
	}
	for i := range shuffles {
		for j := range shuffles {
			if i == j {
				continue
			}
			if slices.Float32Equals(shuffles[i], shuffles[j]) {
				// If there is any collision in 8 shuffles, the Shuffle function is probably broken
				t.Fail()
			}
		}
	}
}

func Test_Float32Equals(t *testing.T) {
	type TestCase struct {
		name string
		a, b []float32
		exp  bool
	}
	tcs := []TestCase{
		// nil checks
		{"Equals nil", nil, nil, true},
		{"Left nil, right empty", nil, []float32{}, false},
		{"Right nil, left empty", []float32{}, nil, false},
		{"Left nil, right not empty", nil, []float32{-2147483648}, false},
		{"Right nil, left not empty", []float32{-2147483648}, nil, false},

		{"Equals empty", []float32{}, []float32{}, true},
		{"Equals 1", []float32{-2147483648}, []float32{-2147483648}, true},
		{"Equals 2", []float32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, []float32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, true},
		{"Not equals 1", []float32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, []float32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, -2147483648}, false},
		{"Not equals 2", []float32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, []float32{}, false},
		{"Not equals 3", []float32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, []float32{-2147483648, -2147483648}, false},
	}
	for _, tc := range tcs {
		got := slices.Float32Equals(tc.a, tc.b)
		if got != tc.exp {
			t.Error(tc.name)
		}
	}
}
