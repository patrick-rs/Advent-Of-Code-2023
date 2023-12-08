package main

import (
	"fmt"
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

	found := checkForSymbol(1, 4, 3, 6, inputArray)

	fmt.Println(strings.Split(input, "\n"))

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

	found := checkForSymbol(1, 4, 3, 6, inputArray)

	fmt.Println(strings.Split(input, "\n"))

	assert.True(t, found)
}
