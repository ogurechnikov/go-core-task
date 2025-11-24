package main

import (
	"crypto/sha256"
	"fmt"
)

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
	fmt.Println("numDecimal", getType(numDecimal))
	fmt.Println("numOctal", getType(numOctal))
	fmt.Println("numHexadecimal", getType(numHexadecimal))
	fmt.Println("pi", getType(pi))
	fmt.Println("name", getType(name))
	fmt.Println("isActive", getType(isActive))
	fmt.Println("complexNum", getType(complexNum))

	fmt.Println(mergeToString())

	runes := []rune(mergeToString())
	fmt.Println(runes)

	hashslice := createHash(runes, "go-2024")
	fmt.Println(hashslice)
}

func getType(value any) string {
	return fmt.Sprintf("%T", value)
}

func mergeToString() string {
	data := []any{numDecimal, numOctal, numHexadecimal, pi, name, isActive, complexNum}

	var str string
	for _, item := range data {
		str += fmt.Sprintf("%v", item)
	}
	return str
}

func createHash(runes []rune, salt string) string {
	runeStr := string(runes)
	mid := len(runeStr) / 2
	hashStr := runeStr[:mid] + salt + runeStr[mid:]

	hash := sha256.New()
	hash.Write([]byte(hashStr))
	return fmt.Sprintf("%x", hash.Sum(nil))
}
