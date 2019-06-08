package main

import (
	"fmt"
	_ "sample/golang-training-2019-06-01/day2/lib1"
)

func init() {
	fmt.Println("Init")
}

func main() {
	primes := [6]int{2, 3, 4, 7, 11}
	var s []int = primes[1:4]
	fmt.Printf("%T => %v\n", s, s)
	fmt.Printf("%T => %v\n", primes, primes)

	s[0] = 10
	primes[3] = 111
	fmt.Printf("%T => %v\n", s, s)
	fmt.Printf("%T => %v\n", primes, primes)

	ss := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
	}
	fmt.Println(ss)
}
