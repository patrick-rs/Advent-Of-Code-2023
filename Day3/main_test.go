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

	found := checkForSymbol(1, 4, 3, strings.Split(input, "\n"))
	assert.True(t, found)
}

func TestCheckForSymbol_ReturnsFalse(t *testing.T) {
	input := `..........
..........
..........
.........`

	found := checkForSymbol(1, 4, 2, strings.Split(input, "\n"))

	assert.True(t, found)
}

func TestCheckIfSymbol(t *testing.T) {
	assert.False(t, checkIfSymbol('a'))
	assert.False(t, checkIfSymbol('z'))
	assert.False(t, checkIfSymbol('1'))
	assert.False(t, checkIfSymbol('9'))
	assert.True(t, checkIfSymbol('!'))
	assert.True(t, checkIfSymbol(')'))
}
