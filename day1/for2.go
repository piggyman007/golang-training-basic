package main

import "fmt"

func main() {
	names := []string{"game", "bank", "samui", "dog", "jiew"}
	for i, name := range names {
		fmt.Printf("%v: %v\n", i, name)
	}
}
