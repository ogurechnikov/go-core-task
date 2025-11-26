package main

import (
	"fmt"
)

var StringIntMap = map[string]int{}

func main() {
	Add("мама", 1)
	Add("папа", 2)
	fmt.Println(StringIntMap)
}

func Add(key string, value int) {
	StringIntMap[key] = value
}
