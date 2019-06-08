package main

import (
	"fmt"
)

func main() {
	i := 1
	p := &i

	fmt.Println(p)
	fmt.Println(*p)
	fmt.Println(i)
	fmt.Println()

	*p = 555
	fmt.Println(p)
	fmt.Println(*p)
	fmt.Println(i)
	fmt.Println()

	i = 123
	fmt.Println(p)
	fmt.Println(*p)
	fmt.Println(i)
	fmt.Println()
}
