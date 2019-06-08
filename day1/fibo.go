package main

import (
	"fmt"
)

func fibonacci() func() int {
	x, y := 0, 1
	return func() int {
		x, y = y, x+y
		return x
	}
}

func adder() (func() int, func() int) {
	sum := 0
	return func() int {
			sum++
			return sum
		},
		func() int {
			return sum
		}
}

func main() {
	f := fibonacci()
	fmt.Printf("%v ", 0)
	for i := 0; i < 10; i++ {
		fmt.Printf("%v ", f())
	}
	fmt.Println()

	inc, cur := adder()
	fmt.Println(cur())

	fmt.Println(inc())
	fmt.Println(inc())

	fmt.Println(cur())
}
