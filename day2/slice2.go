package main

import (
	"fmt"
)

func main() {
	// s := []int{2, 3, 5, 7, 11, 13}
	// printSlice(s)

	// s = s[:0] // Slice the slice to give it zero length.
	// printSlice(s)

	// s = s[:4] // Extend its length.
	// printSlice(s)

	// s = s[2:] // Drop its first two values.
	// printSlice(s)

	// b := make([]int, 0, 5) // len(b)=0, cap(b)=5

	// b = b[:cap(b)] // len(b)=5, cap(b)=5
	// printSlice(b)
	// b = b[1:] // len(b)=4, cap(b)=4
	// printSlice(b)

	var s []int
	printSlice(s)

	s = append(s, 0) // append works on nil slices.
	printSlice(s)

	s = append(s, 1) // The slice grows as needed.
	printSlice(s)

	new := []int{2, 3, 4, 5}
	s = append(s, new...) // We can add more than one element at a time.
	printSlice(s)

	test(new...)
}

func test(input ...int) {
	fmt.Println(input[0])
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
