package day2

import (
	"adventofcode/helpers"
	"fmt"
	"strings"
)

const rock = 0
const paper = 1
const scissors = 2

const win = "win"
const lose = "lose"
const draw = "draw"

// Day2 Puzzle
func Day2(filepath string) {
	fmt.Println("Day 2 - Rock Paper Scissors, it's a fix!")

	originalScore := 0
	maxFitchingScore := 0
	for _, l := range helpers.ReadLines(filepath) {
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

func scoreRound(me, you int) int {
	// Initial score is based on value of the option you selected
	scoreFromSelected := me + 1

	var score int
	// Draw = 3 points
	if me == you {
		score = scoreFromSelected + 3
		// Win (each RPS in cycle always beats the one before) = 6 points
	} else if me == (you+1)%3 {
		score = scoreFromSelected + 6
		// Loss = 0 points
	} else {
		score = scoreFromSelected
	}

	return score
}

func selectRpsToMatchResult(opponent int, result string) int {
	var rps int
	switch result {
	case win:
		rps = (opponent + 1) % 3
	case lose:
		// -1 % 3 != 2 with go's mod operator add 3 so we always have a positive value to mod
		rps = ((opponent - 1) + 3) % 3
	case draw:
		rps = opponent
	}
	return rps
}

func translateToRps(letter string) int {
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
