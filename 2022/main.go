package main

import (
	"adventofcode/puzzles/day1"
	"adventofcode/puzzles/day2"
	"adventofcode/puzzles/day3"
	"adventofcode/puzzles/day4"
	"adventofcode/puzzles/day5"
	"adventofcode/puzzles/day6"
	"adventofcode/puzzles/day7"
	"adventofcode/puzzles/day8"
	"fmt"
	"os"
)

func main() {
	day := os.Args[1]
	inputFile := os.Args[2]

	fmt.Println("Running Day", day, "with input", inputFile)

	days := map[string]func(string){
		"1": day1.Day1,
		"2": day2.Day2,
		"3": day3.Day3,
		"4": day4.Day4,
		"5": day5.Day5,
		"6": day6.Day6,
		"7": day7.Day7,
		"8": day8.Day8,
	}
	filepath := "puzzles/day" + day + "/" + inputFile
	days[day](filepath)
}
