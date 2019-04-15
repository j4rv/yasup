package slices_test

import "testing"
import "github.com/j4rv/slices"

func Test_Complex128Insert(t *testing.T) {
	type testCase struct {
		name     string
		slice    []complex128
		insertAt int
	}
	base := []complex128{0 + 0i, 1 + 2i, -2 + 7.5i, 3 + 42.1i, 4 - 74.6i, -5 + 4i, 6 - 88i, 7 - 0i, 8 + 100i, 9 + 99i}
	tcs := []testCase{
		{"First", []complex128{0 + 0i, 1 + 2i, -2 + 7.5i, 3 + 42.1i, 4 - 74.6i, -5 + 4i, 6 - 88i, 7 - 0i, 8 + 100i, 9 + 99i}, 0},
		{"Middle", []complex128{0 + 0i, 1 + 2i, -2 + 7.5i, 3 + 42.1i, 4 - 74.6i, -5 + 4i, 6 - 88i, 7 - 0i, 8 + 100i, 9 + 99i}, len(base) / 2},
		{"Last", []complex128{0 + 0i, 1 + 2i, -2 + 7.5i, 3 + 42.1i, 4 - 74.6i, -5 + 4i, 6 - 88i, 7 - 0i, 8 + 100i, 9 + 99i}, len(base)},
		{"Empty slice", []complex128{}, 0},
		{"Nil slice", nil, 0},
	}
	for _, tc := range tcs {
		slices.Complex128Insert(-52.6084+155.80287i, &tc.slice, tc.insertAt)
		if tc.slice[tc.insertAt] != -52.6084+155.80287i {
			t.Error(tc)
		}
	}
}

func Test_Complex128FastShuffle(t *testing.T) {
	shuffles := [][]complex128{}
	for i := 0; i < 8; i++ {
		or := []complex128{0 + 0i, 1 + 2i, -2 + 7.5i, 3 + 42.1i, 4 - 74.6i, -5 + 4i, 6 - 88i, 7 - 0i, 8 + 100i, 9 + 99i}
		slices.Complex128FastShuffle(or)
		shuffles = append(shuffles, or)
	}
	for i := range shuffles {
		for j := range shuffles {
			if i >= j {
				continue
			}
			if slices.Complex128Equals(shuffles[i], shuffles[j]) {
				// If there is any collision in 8 shuffles, the Shuffle function is probably broken
				t.Fail()
			}
		}
	}
	// check that nil does not panic
	slices.Complex128FastShuffle(nil)
}

func Test_Complex128SecureShuffle(t *testing.T) {
	shuffles := [][]complex128{}
	for i := 0; i < 8; i++ {
		or := []complex128{0 + 0i, 1 + 2i, -2 + 7.5i, 3 + 42.1i, 4 - 74.6i, -5 + 4i, 6 - 88i, 7 - 0i, 8 + 100i, 9 + 99i}
		slices.Complex128SecureShuffle(or)
		shuffles = append(shuffles, or)
	}
	for i := range shuffles {
		for j := range shuffles {
			if i >= j {
				continue
			}
			if slices.Complex128Equals(shuffles[i], shuffles[j]) {
				// If there is any collision in 8 shuffles, the Shuffle function is probably broken
				t.Fail()
			}
		}
	}
	// check that nil does not panic
	slices.Complex128SecureShuffle(nil)
}

func Test_Complex128Equals(t *testing.T) {
	type TestCase struct {
		name string
		a, b []complex128
		exp  bool
	}
	tcs := []TestCase{
		// nil checks
		{"Equals nil", nil, nil, true},
		// golang treats empty and nil slices as the same thing in most cases, we'll do the same
		{"Left nil, right empty", nil, []complex128{}, true},
		{"Right nil, left empty", []complex128{}, nil, true},
		{"Left nil, right not empty", nil, []complex128{-52.6084 + 155.80287i}, false},
		{"Right nil, left not empty", []complex128{-52.6084 + 155.80287i}, nil, false},

		{"Equals empty", []complex128{}, []complex128{}, true},
		{"Equals 1", []complex128{-52.6084 + 155.80287i}, []complex128{-52.6084 + 155.80287i}, true},
		{"Equals 2", []complex128{0 + 0i, 1 + 2i, -2 + 7.5i, 3 + 42.1i, 4 - 74.6i, -5 + 4i, 6 - 88i, 7 - 0i, 8 + 100i, 9 + 99i}, []complex128{0 + 0i, 1 + 2i, -2 + 7.5i, 3 + 42.1i, 4 - 74.6i, -5 + 4i, 6 - 88i, 7 - 0i, 8 + 100i, 9 + 99i}, true},
		{"Not equals 1", []complex128{0 + 0i, 1 + 2i, -2 + 7.5i, 3 + 42.1i, 4 - 74.6i, -5 + 4i, 6 - 88i, 7 - 0i, 8 + 100i, 9 + 99i}, []complex128{0 + 0i, 1 + 2i, -2 + 7.5i, 3 + 42.1i, 4 - 74.6i, -5 + 4i, 6 - 88i, 7 - 0i, 8 + 100i, 9 + 99i, -52.6084 + 155.80287i}, false},
		{"Not equals 2", []complex128{0 + 0i, 1 + 2i, -2 + 7.5i, 3 + 42.1i, 4 - 74.6i, -5 + 4i, 6 - 88i, 7 - 0i, 8 + 100i, 9 + 99i}, []complex128{}, false},
		{"Not equals 3", []complex128{0 + 0i, 1 + 2i, -2 + 7.5i, 3 + 42.1i, 4 - 74.6i, -5 + 4i, 6 - 88i, 7 - 0i, 8 + 100i, 9 + 99i}, []complex128{-52.6084 + 155.80287i, -52.6084 + 155.80287i}, false},
	}
	for _, tc := range tcs {
		got := slices.Complex128Equals(tc.a, tc.b)
		if got != tc.exp {
			t.Error(tc)
		}
	}
}
