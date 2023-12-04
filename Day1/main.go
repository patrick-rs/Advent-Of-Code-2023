package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

func getcalibrationValue(line string) int {
	var firstNumber rune
	var secondNumber rune

	for _, r := range line {
		if unicode.IsNumber(r) {
			if firstNumber == 0 {
				firstNumber = r
			} else {
				secondNumber = r
			}
		}
	}

	if secondNumber == 0 {
		num, err := strconv.ParseInt(fmt.Sprintf("%c%c", firstNumber, firstNumber), 0, 64)
		if err != nil {
			panic(err)
		}
		return int(num)
	}

	num, err := strconv.ParseInt(fmt.Sprintf("%c%c", firstNumber, secondNumber), 0, 64)
	if err != nil {
		panic(err)
	}
	return int(num)
}
