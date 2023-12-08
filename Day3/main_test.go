package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckForSymbol_ReturnsTrue(t *testing.T) {
	input := `..........
..........
.....+....
.........`

	inputArray := [][]byte{}
	for _, l := range strings.Split(input, "\n") {
		inputArray = append(inputArray, []byte(l))
	}

	found := checkForSymbol(1, 4, 3, inputArray)

	assert.True(t, found)
}

func TestCheckForSymbol_ReturnsFalse(t *testing.T) {
	input := `..........
..........
..........
.........`

	inputArray := [][]byte{}
	for _, l := range strings.Split(input, "\n") {
		inputArray = append(inputArray, []byte(l))
	}

	found := checkForSymbol(1, 4, 2, inputArray)

	assert.True(t, found)
}

func TestCheckIfSymbol(t *testing.T) {
	input := 'a'
	assert.False(t, checkIfSymbol(input))
	input = 'z'
	assert.False(t, checkIfSymbol(input))
	input = '1'
	assert.False(t, checkIfSymbol(input))
	input = '9'
	assert.False(t, checkIfSymbol(input))
	input = '!'
	assert.True(t, checkIfSymbol(input))
	input = ')'
	assert.True(t, checkIfSymbol(input))
}
