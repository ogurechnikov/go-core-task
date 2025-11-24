package main

import "fmt"

var (
	numDecimal     int       = 42
	numOctal       int       = 052
	numHexadecimal int       = 0x2A
	pi             float64   = 3.14
	name           string    = "Golang"
	isActive       bool      = true
	complexNum     complex64 = 1 + 2i
)

func main() {
	printType(numDecimal)
	printType(numOctal)
	printType(numHexadecimal)
	printType(pi)
	printType(name)
	printType(isActive)
	printType(complexNum)

	mergeToString()
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

func mergeToString() {
	data := []any{numDecimal, numOctal, numHexadecimal, pi, name, isActive, complexNum}

	var str string
	for _, item := range data {
		str += fmt.Sprintf("%v", item)
	}
	fmt.Println(str)
}
