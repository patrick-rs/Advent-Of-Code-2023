package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	sumOfcalibrationValues := 0

	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		calibrationValue := getcalibrationValue(scanner.Text())
		sumOfcalibrationValues += calibrationValue
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Printf("Sum of calibration values: %d\n", sumOfcalibrationValues)
}

var numbersToDigitsMap map[string]string = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func getcalibrationValue(line string) int {

	var firstNumber string
	var secondNumber string

	for i := range line {
		foundDigit := ""

		if unicode.IsNumber(rune(line[i])) {
			foundDigit = string(line[i])
		} else {
			for num, digit := range numbersToDigitsMap {
				if strings.HasPrefix(line[i:], num) {
					foundDigit = digit
					break
				}
			}
		}

		if foundDigit != "" {
			if firstNumber == "" {
				firstNumber = foundDigit
			} else {
				secondNumber = foundDigit
			}
		}
	}

	if secondNumber == "" {
		num, err := strconv.ParseInt(fmt.Sprintf("%s%s", firstNumber, firstNumber), 0, 64)
		if err != nil {
			panic(err)
		}
		return int(num)
	}

	num, err := strconv.ParseInt(fmt.Sprintf("%s%s", firstNumber, secondNumber), 0, 64)
	if err != nil {
		panic(err)
	}
	return int(num)
}
