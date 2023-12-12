package main

import (
	"strconv"
	"strings"
	"unicode"
)

func main() {

}

func sumPartNumbers(input string) int {
	sum := 0
	rows := strings.Split(input, "\n")
	for y, line := range rows {
		xPos := -1
		length := -1
		yPos := -1
		num := ""
		for x, char := range line {
			if unicode.IsNumber(char) {
				if num == "" {
					xPos = x - 1
					yPos = y - 1
					length = 1
				}
				num += string(char)
				length += 1

				if x == len(line)-1 {
					if checkForSymbol(int64(yPos), int64(xPos), int64(length), rows) {
						number, err := strconv.ParseInt(num, 0, 64)
						if err != nil {
							panic(err)
						}
						sum += int(number)
						num = ""
					}
				}
			} else {
				if num != "" {
					if checkForSymbol(int64(yPos), int64(xPos), int64(length), rows) {
						number, err := strconv.ParseInt(num, 0, 64)
						if err != nil {
							panic(err)
						}
						sum += int(number)
						num = ""
					}

				}

			}
		}
	}
	return sum
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

	if int(xPos+length) >= len(rows[0]) {
		checkRightSide = false
		length -= 1
	}

	if yPos < 0 {
		for x := int64(0); x < length; x++ {

			if checkIfSymbol(rows[yPos][xPos+x]) {
				return true
			}
		}
	}

	if len(rows) == int(yPos+1) {
		return false
	}

	if checkRightSide {
		for i := 0; i < 3; i++ {
			if len(rows) == int(yPos)+i {
				break
			}
			if int(yPos)+i <= 0 {
				continue
			}
			if checkIfSymbol(rows[int(yPos)+i][xPos+length]) {
				return true
			}
		}
	}

	if checkLeftSide {
		for i := 0; i < 3; i++ {
			if len(rows) == int(yPos)+i {
				break
			}
			if int(yPos)+i <= 0 {
				continue
			}
			if checkIfSymbol(rows[int(yPos)+1][xPos+length]) {
				return true
			}
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
