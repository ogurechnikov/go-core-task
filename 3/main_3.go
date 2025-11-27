package main

import (
	"fmt"
	"maps"
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

	cm := m.Copy()
	fmt.Println("Копия StringIntMap:", cm)

	m.Add("бабушка", 4)
	fmt.Println("Оригинал(проверка):", m)
	fmt.Println("Копия(проверка):", cm)
}

func (s StringIntMap) Add(key string, value int) {
	s[key] = value
}

func (s StringIntMap) Remove(key string) {
	delete(s, key)
}

func (s StringIntMap) Copy() StringIntMap {
	cm := make(StringIntMap)
	// Подкапотный способ
	// for key, value := range s {
	// 	cm[key] = value
	// }

	// Современный способ
	maps.Copy(cm, s)
	return cm
}
