package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDetermineIfValidGame_ValidGames(t *testing.T) {
	input := "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"
	gameNumber := determineIfValidGame(input)
	assert.Equal(t, 1, gameNumber)

	input = "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue"
	gameNumber = determineIfValidGame(input)
	assert.Equal(t, 2, gameNumber)

	input = "Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green"
	gameNumber = determineIfValidGame(input)
	assert.Equal(t, 5, gameNumber)
}

func TestDetermineIfValidGame_InvalidGames(t *testing.T) {
	input := "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red"
	gameNumber := determineIfValidGame(input)
	assert.Equal(t, 0, gameNumber)

	input = "Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red"
	gameNumber = determineIfValidGame(input)
	assert.Equal(t, 0, gameNumber)
}
