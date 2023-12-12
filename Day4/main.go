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
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	total := 0
	for scanner.Scan() {
		game := strings.Split(scanner.Text(), ":")[1]
		cards := strings.Split(game, "|")

		winningNumbers := grabNumbersFromLine(cards[0])
		numbersWeHave := grabNumbersFromLine(cards[1])

		total += calculateCardScore(numbersWeHave, winningNumbers)
	}

	fmt.Printf("Total score: %d\n", total)
}

func calculateCardScore(numbersWeHave, winningNumbers []int) int {
	score := 0

	winNums := map[int]bool{}
	for _, num := range winningNumbers {
		winNums[num] = true
	}

	for _, numWeHave := range numbersWeHave {
		if winNums[numWeHave] {
			if score == 0 {
				score = 1
			} else {
				score = score * 2
			}
		}
	}

	return score
}

func grabNumbersFromLine(line string) []int {
	re := regexp.MustCompile(`(\d+)`)

	matches := re.FindAll([]byte(line), -1)

	nums := []int{}

	for _, match := range matches {
		n, err := strconv.ParseInt(string(match), 0, 64)

		if err != nil {
			panic(err)
		}

		nums = append(nums, int(n))
	}
	return nums
}
