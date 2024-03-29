package main

import "fmt"

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))

	// Here Big is too large of a number but can be handled as a float64.
	// No compilation error is thrown here.
	fmt.Println(needFloat(Big))

	// The below line throws the following compilation error
	// constant 1267650600228229401496703205376 overflows int
	fmt.Println(float64(Big))

	const MaxUint = ^uint(0)
	const MaxInt = int(MaxUint >> 1)
	fmt.Println(MaxInt) // print 9223372036854775807

}
