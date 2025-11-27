package main

import (
	"fmt"
)

type StringIntMap map[string]int

func main() {

	m := StringIntMap{}

	m.Add("мама", 1)
	m.Add("папа", 2)
	m.Add("дедушка", 3)
	fmt.Println("StringIntMap после добавления:", m)
}

func (s StringIntMap) Add(key string, value int) {
	s[key] = value
}
