package slices_test

import "testing"
import "github.com/j4rv/slices"

func Test_StringFastShuffle(t *testing.T) {
	shuffles := [][]string{}
	for i := 0; i < 8; i++ {
		or := []string{"0", "1", "2", "3", "4", "5", "6", "7", "lorem", "ipsum"}
		slices.StringFastShuffle(or)
		shuffles = append(shuffles, or)
	}
	for i := range shuffles {
		for j := range shuffles {
			if i == j {
				continue
			}
			if slices.StringEquals(shuffles[i], shuffles[j]) {
				// If there is any collision in 8 shuffles, the Shuffle function is probably broken
				t.Fail()
			}
		}
	}
}

func Test_StringSecureShuffle(t *testing.T) {
	shuffles := [][]string{}
	for i := 0; i < 8; i++ {
		or := []string{"0", "1", "2", "3", "4", "5", "6", "7", "lorem", "ipsum"}
		slices.StringSecureShuffle(or)
		shuffles = append(shuffles, or)
	}
	for i := range shuffles {
		for j := range shuffles {
			if i == j {
				continue
			}
			if slices.StringEquals(shuffles[i], shuffles[j]) {
				// If there is any collision in 8 shuffles, the Shuffle function is probably broken
				t.Fail()
			}
		}
	}
}

func Test_StringEquals(t *testing.T) {
	type TestCase struct {
		name string
		a, b []string
		exp  bool
	}
	tcs := []TestCase{
		// nil checks
		{"Equals nil", nil, nil, true},
		// golang treats empty and nil slices as the same thing in most cases, we'll do the same
		{"Left nil, right empty", nil, []string{}, true},
		{"Right nil, left empty", []string{}, nil, true},
		{"Left nil, right not empty", nil, []string{"foobar"}, false},
		{"Right nil, left not empty", []string{"foobar"}, nil, false},

		{"Equals empty", []string{}, []string{}, true},
		{"Equals 1", []string{"foobar"}, []string{"foobar"}, true},
		{"Equals 2", []string{"0", "1", "2", "3", "4", "5", "6", "7", "lorem", "ipsum"}, []string{"0", "1", "2", "3", "4", "5", "6", "7", "lorem", "ipsum"}, true},
		{"Not equals 1", []string{"0", "1", "2", "3", "4", "5", "6", "7", "lorem", "ipsum"}, []string{"0", "1", "2", "3", "4", "5", "6", "7", "lorem", "ipsum", "foobar"}, false},
		{"Not equals 2", []string{"0", "1", "2", "3", "4", "5", "6", "7", "lorem", "ipsum"}, []string{}, false},
		{"Not equals 3", []string{"0", "1", "2", "3", "4", "5", "6", "7", "lorem", "ipsum"}, []string{"foobar", "foobar"}, false},
	}
	for _, tc := range tcs {
		got := slices.StringEquals(tc.a, tc.b)
		if got != tc.exp {
			t.Error(tc.name)
		}
	}
}
