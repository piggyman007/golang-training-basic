package main

import "fmt"

type Sayer interface {
	SayName()
	SetAge(x int)
}

////////////////////////////////////////////////////////////////////////////////////////

type Cat struct {
	name string
	age  int
}

func (c Cat) SetAge(age int) {
}

func (c Cat) SayName() {
	fmt.Println(c.name)
}

////////////////////////////////////////////////////////////////////////////////////////
type Person struct {
	id    int
	name  string
	age   int
	books map[string]string
}

// func (p Person) String() string {
// 	// if call fmt.Println(person), will call this function
// 	// see >>> stringer
// 	return fmt.Sprintf("my name is %s, age is %d, my id is %d\n", p.name, p.age, p.id)
// }

func (p Person) SayName() {
	fmt.Println(p.name)
}

func (p Person) SayID() {
	fmt.Println(p.id)
}

func (p Person) SayAge() {
	fmt.Println(p.age)
}

func (p *Person) SetAge(age int) {
	p.age = age
	p.books["gobooks"] = "ok"
	fmt.Println("books in method => ", p.books)
}

////////////////////////////////////////////////////////////////////////////////////////

func AnnounceName(s Sayer) {
	s.SayName()
}

////////////////////////////////////////////////////////////////////////////////////////

func main() {
	// การ new แบบ pointer, e.g., a := &Person{} จะมองเห็นทั้ง
	// (r Reciever) xx และ
	// (r *Reciever) xx

	// การ new แบบ normal, e.g., a := Person{} จะมองเห็นเฉพาะ
	// (r Reciever) xx
	// จะมองไม่เห็น (r *Reciever) xx
	aood := &Person{id: 101, name: "Alongkorn", age: 15, books: map[string]string{"gobooks": "perfect go"}}

	// aood.SayName()
	// aood.SayAge()
	// aood.SayID()

	aood.SetAge(123)
	aood.SayAge()

	// struct & slice will keep address (change in method will cause book in main() change also)
	fmt.Println("book in main  => ", aood.books)

	c := Cat{name: "Meaw"}

	AnnounceName(c)
	AnnounceName(aood)

}
