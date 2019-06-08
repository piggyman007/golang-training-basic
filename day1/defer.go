package main

import (
	"fmt"
)

func main() {
	fmt.Println("start")

	for i := 0; i < 10; i++ {
		defer func() {
			fmt.Println(i)
		}()
	}

	fmt.Println("end")
}
