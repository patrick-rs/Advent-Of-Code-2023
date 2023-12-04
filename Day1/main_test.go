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
