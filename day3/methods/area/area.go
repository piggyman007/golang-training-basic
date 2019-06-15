package main

import "fmt"

type H interface {
	Height() float64
}

type Arearer interface {
	area() float64
}

type triangle struct {
	height float64
	base   float64
}

func (tri *triangle) area() float64 {
	return 0.5 * tri.base * tri.height
}

func (tri triangle) Height() float64 {
	return tri.height
}

func printArea(a Arearer) {
	fmt.Println(a.area())
}

func printH(h H) {
	fmt.Println(h.Height())
}

func main() {
	rec := &triangle{3, 4}
	printArea(rec)

	tri := &triangle{height: 3, base: 4}
	printArea(tri)
	printH(tri)

	tri2 := new(triangle)
	fmt.Printf("%v\n", tri2)

	tri3 := &triangle{}
	fmt.Printf("%v\n", tri3)
}
