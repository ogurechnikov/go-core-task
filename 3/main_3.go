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

	m.Remove("папа")
	fmt.Println("StringIntMap после удаления:", m)
}

func (s StringIntMap) Add(key string, value int) {
	s[key] = value
}

func (s StringIntMap) Remove(key string) {
	delete(s, key)
}
