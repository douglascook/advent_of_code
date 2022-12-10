package day10

import (
	"adventofcode/helpers"
	"fmt"
	"strings"
)

func Day10(filepath string) {
	fmt.Println("Day 10 - Cathode Ray Tube Signal Stuff!")

	instructions := helpers.ReadLines(filepath)
	registerOverTime := getRegisterOverTime(instructions)

	score := signalStrengthScore(registerOverTime, []int{20, 60, 100, 140, 180, 220})
	fmt.Println(score)

	writeToCRT(registerOverTime)
}

func getRegisterOverTime(instructions []string) []int {
	cycle := 0
	// Pad with fake cycle zero to make things simpler
	signals := []int{0}

	// Register starts at 1
	x := 1
	for _, instruction := range instructions {
		// No-op takes one cycle and leaves register unmodified
		if instruction == "noop" {
			cycle++
			signals = append(signals, x)
		} else {
			// Addx starts
			addx := parseAddx(instruction)
			cycle++
			signals = append(signals, x)
			// First cycle register is unchanged
			cycle++
			signals = append(signals, x)
			// Second cycle register is updated
			x += addx
		}
	}
	return signals
}

func signalStrengthScore(registerOverTime []int, cycles []int) int {
	score := 0
	for _, c := range cycles {
		fmt.Println("Cycle", c, "register", registerOverTime[c])
		score += c * registerOverTime[c]
		fmt.Println("Score", score)
	}
	return score
}

func writeToCRT(registerOverTime []int) {
	printLine(registerOverTime[1:41])
	printLine(registerOverTime[41:81])
	printLine(registerOverTime[81:121])
	printLine(registerOverTime[121:161])
	printLine(registerOverTime[161:201])
	printLine(registerOverTime[201:241])
}

// TODO why does this not work with outer loop and the obvious mod logic???
func printLine(registerOverTime []int) {
	for col, r := range registerOverTime {
		if col <= r+1 && col >= r-1 {
			fmt.Print("#")
		} else {
			fmt.Print(".")
		}
	}
	fmt.Print("\n")
}

func parseAddx(instruction string) int {
	if instruction == "noop" {
		return 0
	}
	change := strings.Split(instruction, " ")[1]
	return helpers.StringToInt(change)
}
