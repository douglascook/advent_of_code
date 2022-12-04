package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Section struct {
	start int
	end   int
}

func (s Section) contains(otherSection Section) bool {
	return s.start <= otherSection.start && s.end >= otherSection.end
}

func (s Section) overlaps(otherSection Section) bool {
	disjoint := s.end < otherSection.start || s.start > otherSection.end
	return !disjoint
}

func main() {
	fmt.Println("Day 4 - Camp Cleanup!")

	sections := readLines("input.txt")
	findOverlaps(sections)
}

func findOverlaps(sections []string) {
	encloses := 0
	overlaps := 0
	for _, s := range sections {
		first, second, _ := strings.Cut(s, ",")
		s1 := parseSection(first)
		s2 := parseSection(second)

		if s1.contains(s2) || s2.contains(s1) {
			encloses++
		}
		if s1.overlaps(s2) {
			overlaps++
		}
	}
	fmt.Println("Number of sections enclosing another is", encloses)
	fmt.Println("Number of sections with an overlap is", overlaps)
}

func parseSection(section string) Section {
	start, end, _ := strings.Cut(section, "-")
	sectionStart, err := strconv.Atoi(start)
	if err != nil {
		log.Fatal(err)
	}
	sectionEnd, err := strconv.Atoi(end)
	if err != nil {
		log.Fatal(err)
	}
	return Section{sectionStart, sectionEnd}
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
