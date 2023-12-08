package main

import (
	"unicode"
)

func main() {

}

func checkIfSymbol(i byte) bool {
	if unicode.IsDigit(rune(i)) {
		return false
	}
	if unicode.IsLetter(rune(i)) {
		return false
	}
	if i == '.' {
		return false
	}
	return true
}

func checkForSymbol(yPos, xPos, length int64, rows []string) bool {

	checkRightSide := true
	checkLeftSide := true

	if xPos == -1 {
		checkLeftSide = false
		xPos += 1
		length -= 1
	}

	if int(xPos+length) == len(rows[0]) {
		checkRightSide = false
		length -= 1
	}

	for x := int64(0); x < length; x++ {
		if checkIfSymbol(rows[yPos][xPos+x]) {
			return true
		}
	}

	if len(rows) == int(yPos+1) {
		return false
	}

	if checkRightSide {
		if checkIfSymbol(rows[yPos+1][xPos]) {
			return true
		}
	}

	if checkLeftSide {
		if checkIfSymbol(rows[yPos+1][xPos+length]) {
			return true
		}
	}

	if len(rows) == int(yPos+2) {
		return false
	}

	for x := int64(0); x < length; x++ {
		if checkIfSymbol(rows[yPos+2][xPos+x]) {
			return true
		}
	}

	return false
}
