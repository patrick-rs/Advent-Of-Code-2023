package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	var sumOfValidGameNumbers int64
	var sumOfGamePower int64

	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		gameNumber := determineIfValidGame(scanner.Text())
		sumOfValidGameNumbers += gameNumber

		gamePower := calculateGamePower(scanner.Text())
		sumOfGamePower += gamePower
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Printf("Sum of valid game nubers: %d\n", sumOfValidGameNumbers)
	fmt.Printf("Sum of game power: %d\n", sumOfGamePower)
}

func determineIfValidGame(line string) int64 {
	gameNumber, games := getGameNumberAndGames(line)

	for _, game := range games {
		numberOfRed, numberOfGreen, numberOfBlue := getNumberOfRGBCubes(game)
		if numberOfRed > 12 || numberOfGreen > 13 || numberOfBlue > 14 {
			return 0
		}
	}

	return gameNumber
}

func calculateGamePower(game string) int64 {
	_, games := getGameNumberAndGames(game)
	var maximumRed, maximumGreen, maximumBlue int64

	for _, game := range games {
		numberOfRed, numberOfGreen, numberOfBlue := getNumberOfRGBCubes(game)

		if numberOfRed > maximumRed {
			maximumRed = numberOfRed
		}

		if numberOfGreen > maximumGreen {
			maximumGreen = numberOfGreen
		}

		if numberOfBlue > maximumBlue {
			maximumBlue = numberOfBlue
		}
	}

	return maximumRed * maximumGreen * maximumBlue
}

func getGameNumberAndGames(line string) (int64, []string) {
	lineSplit := strings.Split(line, ":")

	re := regexp.MustCompile(`(\d+)$`)

	matches := re.FindAll([]byte(lineSplit[0]), 1)

	gameNumber, err := strconv.ParseInt(string(matches[0]), 0, 64)

	if err != nil {
		panic(err)
	}

	games := strings.Split(lineSplit[1], ";")

	return gameNumber, games
}

func getNumberOfRGBCubes(game string) (int64, int64, int64) {
	var numberOfRed int64
	var numberOfGreen int64
	var numberOfBlue int64
	var err error

	redRegex := regexp.MustCompile(`(\d+) red`)

	redMatches := redRegex.FindStringSubmatch(game)

	if len(redMatches) == 2 {
		numberOfRed, err = strconv.ParseInt(string(redMatches[1]), 0, 64)
		if err != nil {
			panic(err)
		}
	}

	greenRegex := regexp.MustCompile(`(\d+) green`)

	greenMatches := greenRegex.FindStringSubmatch(game)

	if len(greenMatches) == 2 {
		numberOfGreen, err = strconv.ParseInt(string(greenMatches[1]), 0, 64)
		if err != nil {
			panic(err)
		}
	}

	blueRegex := regexp.MustCompile(`(\d+) blue`)

	blueMatches := blueRegex.FindStringSubmatch(game)

	if len(blueMatches) == 2 {
		numberOfBlue, err = strconv.ParseInt(string(blueMatches[1]), 0, 64)
		if err != nil {
			panic(err)
		}
	}

	return numberOfRed, numberOfGreen, numberOfBlue
}
