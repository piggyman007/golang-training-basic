package main

import (
	"fmt"
)

func main() {
	primes := [6]int{1, 2, 3, 4, 5, 6}
	sum(primes)

	ages := [6]int{12, 15, 16, 17, 15}
	sum(ages)

	xxx(ages)
}

func sum(nums [6]int) {
	s := 0
	for _, n := range nums {
		s += n
	}

	fmt.Println(s)
}

func xxx(x interface{}) {
	myInt, ok := x.([6]int)
	if ok {
		fmt.Println("OK => ", myInt)
	} else {
		fmt.Println(("Not OK"))
	}
}
