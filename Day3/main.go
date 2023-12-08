package main

import (
	"fmt"
	"unicode"
)

func main() {

}

func checkIfSymbol(i rune) bool {
	fmt.Println(i)
	if unicode.IsDigit(i) {
		return false
	}
	if unicode.IsLetter(i) {
		return false
	}
	return true
}

func checkForSymbol(x1, y1, length int64, rows [][]byte) bool {

	return false
}
