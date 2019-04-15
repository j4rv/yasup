package slices_test

import "testing"
import "github.com/j4rv/slices"

func Test_ByteInsert(t *testing.T) {
	type testCase struct {
		name     string
		slice    []byte
		insertAt int
	}
	base := []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	tcs := []testCase{
		{"First", []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 0},
		{"Middle", []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, len(base) / 2},
		{"Last", []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, len(base)},
		{"Empty slice", []byte{}, 0},
		{"Nil slice", nil, 0},
	}
	for _, tc := range tcs {
		slices.ByteInsert(255, &tc.slice, tc.insertAt)
		if tc.slice[tc.insertAt] != 255 {
			t.Error(tc)
		}
	}
}

func Test_ByteFastShuffle(t *testing.T) {
	shuffles := [][]byte{}
	for i := 0; i < 8; i++ {
		or := []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		slices.ByteFastShuffle(or)
		shuffles = append(shuffles, or)
	}
	for i := range shuffles {
		for j := range shuffles {
			if i >= j {
				continue
			}
			if slices.ByteEquals(shuffles[i], shuffles[j]) {
				// If there is any collision in 8 shuffles, the Shuffle function is probably broken
				t.Fail()
			}
		}
	}
	// check that nil does not panic
	slices.ByteFastShuffle(nil)
}

func Test_ByteSecureShuffle(t *testing.T) {
	shuffles := [][]byte{}
	for i := 0; i < 8; i++ {
		or := []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		slices.ByteSecureShuffle(or)
		shuffles = append(shuffles, or)
	}
	for i := range shuffles {
		for j := range shuffles {
			if i >= j {
				continue
			}
			if slices.ByteEquals(shuffles[i], shuffles[j]) {
				// If there is any collision in 8 shuffles, the Shuffle function is probably broken
				t.Fail()
			}
		}
	}
	// check that nil does not panic
	slices.ByteSecureShuffle(nil)
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
		// golang treats empty and nil slices as the same thing in most cases, we'll do the same
		{"Left nil, right empty", nil, []byte{}, true},
		{"Right nil, left empty", []byte{}, nil, true},
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
			t.Error(tc)
		}
	}
}
