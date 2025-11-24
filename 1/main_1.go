package main

import "fmt"

var numDecimal int = 42
var numOctal int = 052
var numHexadecimal int = 0x2A
var pi float64 = 3.14
var name string = "Golang"
var isActive bool = true
var complexNum complex64 = 1 + 2i

func main() {
	printType(numDecimal)
	printType(numOctal)
	printType(numHexadecimal)
	printType(pi)
	printType(name)
	printType(isActive)
	printType(complexNum)
}

func printType(value any) {
	switch v := value.(type) {
	case int:
		fmt.Printf("Тип: int, Значение: %d\n", v)
	case float64:
		fmt.Printf("Тип: float64, Значение: %f\n", v)
	case string:
		fmt.Printf("Тип: string, Значение: %s\n", v)
	case bool:
		fmt.Printf("Тип: bool, Значение: %t\n", v)
	case complex64:
		fmt.Printf("Тип: complex64, Значение: %v\n", v)
	default:
		fmt.Printf("Тип: %T, Значение: %v\n", v, v)
	}
}
