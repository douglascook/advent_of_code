package day1

import (
	"adventofcode/helpers"
	"fmt"
	"log"
	"sort"
	"strconv"
)

// Day1 Puzzle
func Day1(filepath string) {
	fmt.Println("Day 1 - Counting calories!")
	lines := helpers.ReadLines(filepath)

	var elfCalories []int
	elf := 0
	for i := 0; i < len(lines); i++ {
		line := lines[i]

		/* readLines removes newlines, so blank line is empty */
		if line == "" {
			elfCalories = append(elfCalories, elf)
			elf = 0
		} else {
			food, err := strconv.Atoi(line)
			if err != nil {
				log.Fatal(err)
			}
			elf = elf + food
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(elfCalories)))
	total := 0
	for i, elf := range elfCalories[:3] {
		fmt.Println("Elf carrying", i+1, "most has", elf, "calories")
		total = total + elf
	}
	fmt.Println("Together they are carrying", total)
}
