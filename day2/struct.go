package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{X: 1, Y: 2}
	fmt.Printf("%T\n", v)   // print type => main.Vertex
	fmt.Printf("% #v\n", v) // print type and value (with space) => main.Vertex{X: 1, Y: 2}
	fmt.Printf("%#v\n", v)  // print type and value => main.Vertex{X:1, Y:2}
	fmt.Printf("%+v\n", v)  // print value and field name => {X:1 Y:2}

	p := &v
	(*p).X = 1e9
	p.Y = 123
	fmt.Printf("% #v\n", *p)
	fmt.Printf("% #v\n", v)
}
