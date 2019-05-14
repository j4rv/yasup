package main

import (
	"fmt"

	"github.com/j4rv/yasup"
)

func main() {
	ints := []int{4, 5, 6}
	fmt.Println(ints, "starting slice")

	yasup.IntInsert(&ints, 9, 1)
	fmt.Println(ints, "'9' inserted at position 1")
	yasup.IntDelete(&ints, 2)
	fmt.Println(ints, "deleted element at position 2")

	for i := 0; i < 2; i++ {
		yasup.IntPush(&ints, i)
		fmt.Println(ints, "pushed", i)
	}

	yasup.IntFastShuffle(ints)
	fmt.Println(ints, "shuffled (fast)")
	yasup.IntSecureShuffle(ints)
	fmt.Println(ints, "shuffled (secure)")

	for i := 0; i < 6; i++ {
		pop, err := yasup.IntPop(&ints)
		if err != nil {
			fmt.Println("oops, popped too much!. Got err:", err)
		} else {
			fmt.Println("popped:", pop, "current slice:", ints)
		}
	}

	fmt.Println("equals (1):", yasup.IntEquals([]int{4, 5, 6}, []int{4, 5, 6}))
	fmt.Println("equals (2):", yasup.IntEquals([]int{4, 5, 6}, []int{}))

	fmt.Println("Returns errors instead of panicking")
	fmt.Println("yasup.IntDelete(&ints, 99) returned: ", yasup.IntDelete(&ints, 99))

	ints = []int{0, 1, 2, 0, 2, 1}
	fmt.Println("new slice", ints)
	fmt.Println("index of 0:", yasup.IntIndex(ints, 0))
	fmt.Println("last index of 0:", yasup.IntLastIndex(ints, 0))
	fmt.Println("index of 1:", yasup.IntIndex(ints, 1))
	fmt.Println("last index of 1:", yasup.IntLastIndex(ints, 1))

	ints = []int{0, 1, 2, 3, 0, 1, 2, 3, 0, 1, 2, 3, 0, 1, 2, 3, 0, 1, 2, 3}
	fmt.Println("new slice", ints)
	fmt.Println("replaced 0 with 9 one time. Ocurrences:", yasup.IntReplace(ints, 0, 9, 1), "result:", ints)
	fmt.Println("replaced 1 with 9 zero times. Ocurrences:", yasup.IntReplace(ints, 1, 9, 0), "result:", ints)
	fmt.Println("replaced 2 with 9 ten times. Ocurrences:", yasup.IntReplace(ints, 2, 9, 10), "result:", ints)
	fmt.Println("replaced all 3s with 9s. Ocurrences:", yasup.IntReplaceAll(ints, 3, 9), "result:", ints)

}
