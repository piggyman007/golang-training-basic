package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	words := strings.Split(s, " ")
	count := make(map[string]int)
	for _, word := range words {
		decoratedWord := strings.ToUpper(word)
		decoratedWord = strings.Replace(decoratedWord, ",", "", -1)
		decoratedWord = strings.Replace(decoratedWord, "?", "", -1)
		if _, ok := count[decoratedWord]; ok {
			count[decoratedWord]++
		} else {
			count[decoratedWord] = 1
		}
	}

	return count
}

func main() {
	fmt.Println(WordCount("i am A boy, are you a boy"))
	fmt.Println(WordCount("If it looks like a duck, swims like a duck,"))
}
