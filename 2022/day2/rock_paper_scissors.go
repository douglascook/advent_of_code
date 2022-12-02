package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const rock = "rock"
const paper = "paper"
const scissors = "scissors"
const win = "win"
const lose = "lose"
const draw = "draw"

func main() {
	fmt.Println("Day 1 - Counting calories!")

	originalScore := 0
	maxFitchingScore := 0
	for _, l := range readLines("input.txt") {
		opponent, me, _ := strings.Cut(l, " ")
		rpsMe := translateToRps(me)
		rpsOpponent := translateToRps(opponent)

		// Part 1 translated second letter as RPS, calculate the score using that translation
		originalScore = originalScore + scoreRound(rpsMe, rpsOpponent)

		// Part 2 the second letter tells you to win, lose or draw
		result := translateToResult(me)
		rpsMe = selectRpsToMatchResult(rpsOpponent, result)
		maxFitchingScore = maxFitchingScore + scoreRound(rpsMe, rpsOpponent)
	}
	fmt.Println("Total score with original translation =", originalScore)
	fmt.Println("Total score with max fitching results =", maxFitchingScore)

}

func scoreRound(me, you string) int {
	var score int

	// Initial score from what you selected
	switch me {
	case rock:
		score = 1
	case paper:
		score = 2
	case scissors:
		score = 3
	}

	// Draw = 3 points
	if me == you {
		score += 3
		// Win = 6 points
	} else if (me == rock && you == scissors) || (me == scissors && you == paper) || (me == paper && you == rock) {
		score += 6
	}
	// Loss = 0 points

	return score
}

func selectRpsToMatchResult(opponent, result string) string {
	var rps string
	if result == win {
		switch opponent {
		case rock:
			rps = paper
		case paper:
			rps = scissors
		case scissors:
			rps = rock
		}
	} else if result == lose {
		switch opponent {
		case rock:
			rps = scissors
		case paper:
			rps = rock
		case scissors:
			rps = paper
		}
	} else {
		// Draw you choose the same as your opponent
		rps = opponent
	}
	return rps
}

func translateToRps(letter string) string {
	if letter == "A" || letter == "X" {
		return rock
	}
	if letter == "B" || letter == "Y" {
		return paper
	}
	if letter == "C" || letter == "Z" {
		return scissors
	}
	panic("Unexpected letter")
}

func translateToResult(letter string) string {
	switch letter {
	case "X":
		return lose
	case "Y":
		return draw
	case "Z":
		return win
	default:
		panic("Unexpected letter")
	}
}

// TODO learn how to use modules properly and put this into shared helpers
func readLines(filepath string) []string {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}
