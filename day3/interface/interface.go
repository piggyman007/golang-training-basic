package main

import "fmt"

func main() {
	var i interface{} = 1

	fmt.Printf("%T - %v\n", i, i)

	i = "ok"
	fmt.Printf("%T - %v\n", i, i)

	t, ok := i.(string) // type assertion
	if ok {
		fmt.Println(t)
	} else {
		fmt.Println("Invalid type")
	}

	switch i.(type) { // type assertion
	case string:
		fmt.Println("String ", i)
	case int:
		fmt.Println("int ", i)
	}
}
