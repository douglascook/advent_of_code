package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Println("Day 3 - Reorganising rucksacks!")

	prioritiesMap := buildPrioritiesMap()
	elves := readLines("input.txt")
	var score int

	fmt.Println("Finding items in both rucksacks, per elf")
	score = 0
	for _, elf := range elves {
		rucksacks := splitRucksacks(elf)
		score = score + getItemInAll(rucksacks, prioritiesMap)
	}
	fmt.Println("Sum of priorities of items =", score)

	fmt.Println("Finding items carried by every elf in each group")
	score = 0
	// Each group is three elves
	for i := 0; i < len(elves); i += 3 {
		group := elves[i : i+3]
		score = score + getItemInAll(group, prioritiesMap)
	}
	fmt.Println("Sum of priorities of items =", score)
}

func buildPrioritiesMap() map[rune]int {
	prioritiesMap := make(map[rune]int)

	// Priorities for a-z start at 1 -> ascii - 96
	for charCode := 97; charCode < 123; charCode++ {
		prioritiesMap[rune(charCode)] = charCode - 96
	}
	// Priorities for A-Z start at 27 -> ascii - 38
	for charCode := 65; charCode < 91; charCode++ {
		prioritiesMap[rune(charCode)] = charCode - 38
	}

	return prioritiesMap
}

func splitRucksacks(rucksacks string) []string {
	rucksackSize := len(rucksacks) / 2
	return []string{rucksacks[:rucksackSize], rucksacks[rucksackSize:]}
}

func getItemInAll(rucksacks []string, prioritiesMap map[rune]int) int {
	var inAll bool
	for char, priority := range prioritiesMap {
		inAll = true
		for _, sack := range rucksacks {
			if !strings.ContainsRune(sack, char) {
				inAll = false
				break
			}
		}
		if inAll {
			return priority
		}
	}
	panic("Expected one item to be found in both rucksacks :(")
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
