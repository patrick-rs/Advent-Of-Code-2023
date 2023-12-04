package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCalibrationNumber_ReturnsNumbersFromLine(t *testing.T) {
	input := "1a2"

	calibrationNumber := getCalibrationNumber(input)

	assert.Equal(t, 12, calibrationNumber)
}

func TestGetCalibrationNumber_ReturnsNumbersFromLineWithOneNumber(t *testing.T) {
	input := "J7J"

	calibrationNumber := getCalibrationNumber(input)

	assert.Equal(t, 7, calibrationNumber)
}

func TestGetCalibrationNumber_ReturnsNumbersFromLineWithManyNumbers(t *testing.T) {
	input := "iu23h4512n3j34uh8"

	calibrationNumber := getCalibrationNumber(input)

	assert.Equal(t, 28, calibrationNumber)
}
