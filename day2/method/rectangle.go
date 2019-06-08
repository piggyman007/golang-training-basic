package main

import (
	"fmt"
)

type Arearer interface {
	area() float64
}

type Rectangle struct {
	width  float64
	length float64
}

func (rec Rectangle) area() float64 {
	return rec.width * rec.length
}

type Triangle struct {
	base   float64
	height float64
}

func (tri Triangle) area() float64 {
	return tri.base * tri.height / 2
}

func printArea(a Arearer) {
	fmt.Println(a.area())
}

func main() {
	rec := Rectangle{3, 4}
	printArea(rec)

	tri := Triangle{10, 10}
	printArea(tri)
}
