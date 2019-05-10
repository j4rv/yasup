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
}
