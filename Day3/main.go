package main

import (
	"fmt"
	"unicode"
)

func main() {

}

func checkIfSymbol(i byte) bool {
	fmt.Println(i)
	if unicode.IsDigit(rune(i)) {
		return false
	}
	if unicode.IsLetter(rune(i)) {
		return false
	}
	return true
}

func checkForSymbol(x1, y1, length int64, rows []string) bool {

	for x := int64(0); x < length; x++ {
		if checkIfSymbol(rows[y1][x]) {
			return true
		}
	}
	return false
}
