package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumPartNumbersTopLeft(t *testing.T) {
	input := `.#..
..3.
....`

	assert.Equal(t, 3, sumPartNumbers(input))
}

func TestSumPartNumbersLeftSide(t *testing.T) {
	input := `.#..
3...
....`

	assert.Equal(t, 3, sumPartNumbers(input))
}

func TestSumPartNumbersRightSide(t *testing.T) {
	input := `..#.
...3
....`

	assert.Equal(t, 3, sumPartNumbers(input))
}

func TestSumPartNumbersBelowSide(t *testing.T) {
	input := `....
...3
..#.`

	assert.Equal(t, 3, sumPartNumbers(input))
}

func TestSumPartNumbersRightSideReturns0(t *testing.T) {
	input := `....
...3
....`

	assert.Equal(t, 0, sumPartNumbers(input))
}

func TestSumPartNumbersFirstNumberSide(t *testing.T) {
	input := `1...
.$..
....`

	assert.Equal(t, 1, sumPartNumbers(input))
}

func TestSumPartNumbers(t *testing.T) {
	input := `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`

	assert.Equal(t, 4361, sumPartNumbers(input))
}

func TestCheckForSymbol_ReturnsTrue(t *testing.T) {
	input := `..........
..........
......+...
.........`

	found := checkForSymbol(2, 4, 3, strings.Split(input, "\n"))
	assert.True(t, found)
}

func TestCheckForSymbol_ReturnsFalse(t *testing.T) {
	input := `..........
..........
..........
.........`

	found := checkForSymbol(1, 4, 2, strings.Split(input, "\n"))

	assert.False(t, found)
}

func TestCheckForSymbol_ReturnsTrueWhenAtBottom(t *testing.T) {
	input := `..........
..........
..........
...._....`

	found := checkForSymbol(2, 4, 2, strings.Split(input, "\n"))

	assert.True(t, found)
}

func TestCheckForSymbol_ReturnsFalseWhenAtBottom(t *testing.T) {
	input := `..........
..........
..........
.........`

	found := checkForSymbol(3, 4, 2, strings.Split(input, "\n"))

	assert.False(t, found)
}

func TestCheckForSymbol_ReturnsFalseWhenAtSecondToLastRow(t *testing.T) {
	input := `..........
..........
..........
.........`

	found := checkForSymbol(3, 3, 2, strings.Split(input, "\n"))

	assert.False(t, found)
}

func TestCheckForSymbol_ReturnsTrueWhenStartingOnLeftEdge(t *testing.T) {
	input := `..........
..........
&.........
.........`

	found := checkForSymbol(1, 0, 2, strings.Split(input, "\n"))

	assert.True(t, found)
}

func TestCheckForSymbol_ReturnsTrueWhenStartingOnRightEdge(t *testing.T) {
	input := `..........
.........&
..........
.........`

	found := checkForSymbol(0, 8, 2, strings.Split(input, "\n"))

	assert.True(t, found)
}

func TestCheckIfSymbol(t *testing.T) {
	assert.False(t, checkIfSymbol('a'))
	assert.False(t, checkIfSymbol('z'))
	assert.False(t, checkIfSymbol('1'))
	assert.False(t, checkIfSymbol('9'))
	assert.True(t, checkIfSymbol('!'))
	assert.True(t, checkIfSymbol('+'))
	assert.True(t, checkIfSymbol(')'))
}
