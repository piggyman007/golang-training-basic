package main

import "fmt"

type person struct {
	name string
}

func (p person) Walk() {
	fmt.Printf("%v is walking\n", p.name)
}

func (p person) Eat() {
	fmt.Printf("%v is eating\n", p.name)
}

func (p person) Greeting() {
	fmt.Printf("hello %v\n", p.name)
}

func (p person) Name() string {
	return p.name
}

func (p *person) SetName(n string) {
	p.name = n
}

func main() {
	p := person{name: "Alongkorn"}
	p.Walk()
	p.Eat()
	p.Greeting()

	(&p).SetName("Dave")
	fmt.Println(p.Name())
	p.Greeting()
}
