package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetcalibrationValue_ReturnsNumbersFromLine(t *testing.T) {
	input := "1a2"

	calibrationValue := getcalibrationValue(input)

	assert.Equal(t, 12, calibrationValue)
}

func TestGetcalibrationValue_ReturnsNumbersFromLineWithOneNumber(t *testing.T) {
	input := "J7J"

	calibrationValue := getcalibrationValue(input)

	assert.Equal(t, 77, calibrationValue)
}

func TestGetcalibrationValue_ReturnsNumbersFromLineWithManyNumbers(t *testing.T) {
	input := "iu23h4512n3j34uh8"

	calibrationValue := getcalibrationValue(input)

	assert.Equal(t, 28, calibrationValue)
}

func TestReplaceSpelledOutNumbersWithNumbers_ReturnsOnlyDigits(t *testing.T) {
	input := "onetwothree"
	formattedString := replaceSpelledOutNumbersWithNumbers(input)
	assert.Equal(t, "123", formattedString)

	input = "fourfivesix"
	formattedString = replaceSpelledOutNumbersWithNumbers(input)
	assert.Equal(t, "456", formattedString)

	input = "seveneightnine"
	formattedString = replaceSpelledOutNumbersWithNumbers(input)
	assert.Equal(t, "789", formattedString)
}
