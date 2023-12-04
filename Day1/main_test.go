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

func TestReplaceSpelledOutNumbersAndGetCalibrationValues(t *testing.T) {
	input := "two1nine"
	calibrationValue := getcalibrationValue(input)
	assert.Equal(t, 29, calibrationValue)

	input = "eightwothree"
	calibrationValue = getcalibrationValue(input)
	assert.Equal(t, 83, calibrationValue)

	input = "abcone2threexyz"
	calibrationValue = getcalibrationValue(input)
	assert.Equal(t, 13, calibrationValue)

	input = "xtwone3four"
	calibrationValue = getcalibrationValue(input)
	assert.Equal(t, 24, calibrationValue)

	input = "4nineeightseven2"
	calibrationValue = getcalibrationValue(input)
	assert.Equal(t, 42, calibrationValue)

	input = "zoneight234"
	calibrationValue = getcalibrationValue(input)
	assert.Equal(t, 14, calibrationValue)

	input = "7pqrstsixteen"
	calibrationValue = getcalibrationValue(input)
	assert.Equal(t, 76, calibrationValue)
}
