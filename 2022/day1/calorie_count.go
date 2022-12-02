package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	fmt.Println("Day 1 - Counting calories!")

	lines := readLines("input.txt")

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
