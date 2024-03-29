package main

import (
	"adventofcode/puzzles/day1"
	"adventofcode/puzzles/day10"
	"adventofcode/puzzles/day11"
	"adventofcode/puzzles/day12"
	"adventofcode/puzzles/day14"
	"adventofcode/puzzles/day15"
	"adventofcode/puzzles/day17"
	"adventofcode/puzzles/day2"
	"adventofcode/puzzles/day3"
	"adventofcode/puzzles/day4"
	"adventofcode/puzzles/day5"
	"adventofcode/puzzles/day6"
	"adventofcode/puzzles/day7"
	"adventofcode/puzzles/day8"
	"adventofcode/puzzles/day9"
	"fmt"
	"os"
)

func main() {
	day := os.Args[1]
	inputFile := os.Args[2]

	fmt.Println("Running Day", day, "with input", inputFile)

	days := map[string]func(string){
		"1":  day1.Day1,
		"2":  day2.Day2,
		"3":  day3.Day3,
		"4":  day4.Day4,
		"5":  day5.Day5,
		"6":  day6.Day6,
		"7":  day7.Day7,
		"8":  day8.Day8,
		"9":  day9.Day9,
		"10": day10.Day10,
		"11": day11.Day11,
		"12": day12.Day12,
		"14": day14.Day14,
		"15": day15.Day15,
		"17": day17.Day17,
	}
	filepath := "puzzles/day" + day + "/" + inputFile
	days[day](filepath)
}
