package main

import (
	"fmt"
)

type Student struct {
	name     string
	age      int
	nickName string
}

func main() {
	// var stdMap = make(map[string]Student)
	// stdMap["001"] = Student{
	// 	name:     "David Beckham",
	// 	age:      30,
	// 	nickName: "Dave",
	// }

	// stdMap["002"] = Student{
	// 	name:     "Peter Check",
	// 	age:      40,
	// 	nickName: "Pete",
	// }

	var stdMap = map[string]Student{
		"001": {name: "David Beckham", age: 30, nickName: "Dave"},
		"002": {name: "Peter Check", age: 33, nickName: "Pete"},
	}

	std001 := stdMap["001"]
	std001.name = "David Beckham updated"
	stdMap["001"] = std001

	fmt.Println(stdMap)

	delete(stdMap, "001")
	fmt.Println(stdMap)

	_, ok := stdMap["001"]

	fmt.Println(ok)
}
