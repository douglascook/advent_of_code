package day5

import (
	"adventofcode/helpers"
	"fmt"
	"log"
	"strings"
)

type moveCrates struct {
	quantity int
	from     int
	to       int
}

type stack struct {
	crates []string
}

func (s *stack) push(crate string) {
	s.crates = append(s.crates, crate)
}

func (s *stack) pop() string {
	if len(s.crates) == 0 {
		log.Fatal("Cannot pop from empty stack")
	}
	crates := s.crates[len(s.crates)-1]
	s.crates = s.crates[:len(s.crates)-1]
	return crates
}

func (s *stack) takeTopCrates(n int) []string {
	crates := s.crates[len(s.crates)-n:]
	// fmt.Println(crates)
	s.crates = s.crates[:len(s.crates)-n]
	// fmt.Println(s.crates)
	return crates
}

func (s stack) topItem() string {
	return s.crates[len(s.crates)-1]
}

// Day5 Puzzle
func Day5(filepath string) {
	fmt.Println("Day 5 - Crane and Crates!")

	var stacks []stack
	var moves []moveCrates
	// TODO a visualisation of the crates moving would be fun...
	fmt.Println("Crane model CrateMover 9000 - moving crates one at a time")
	stacks, moves = parseInput(filepath)
	for _, m := range moves {
		moveCratesOneByOne(stacks, m)
	}

	fmt.Println("Top crates are now:")
	for _, s := range stacks {
		fmt.Print(s.topItem())
	}
	fmt.Println()

	fmt.Println("Crane model CrateMover 9001 - moving crates whole group at once")
	// FIXME really can't work out how to copy the stacks so that nothing is mutated
	// so just read it all again :|
	stacks, moves = parseInput(filepath)
	for _, m := range moves {
		moveCratesAllAtOnce(stacks, m)
	}

	fmt.Println("Top crates are now:")
	for _, s := range stacks {
		fmt.Print(s.topItem())
	}
	fmt.Println()
}

// CrateMover 9000 moves one crate at a time so pop and push
func moveCratesOneByOne(stacks []stack, move moveCrates) {
	for i := 0; i < move.quantity; i++ {
		item := stacks[move.from].pop()
		stacks[move.to].push(item)
	}
}

// CrateMover 9001 moves crates whole group at once so take slice and append
func moveCratesAllAtOnce(stacks []stack, move moveCrates) {
	crates := stacks[move.from].takeTopCrates(move.quantity)
	for _, crate := range crates {
		stacks[move.to].push(crate)
	}
}

func parseInput(filepath string) ([]stack, []moveCrates) {
	var crates []string
	var moves []moveCrates

	cratesDone := false
	for _, l := range helpers.ReadLines(filepath) {
		if l == "" {
			cratesDone = true
		} else if !cratesDone {
			crates = append(crates, l)
		} else {
			moves = append(moves, parseMove(l))
		}
	}
	stacks := parseCrates(crates)
	return stacks, moves
}

func parseMove(moveDescription string) moveCrates {
	parts := strings.Split(moveDescription, " ")
	return moveCrates{
		helpers.StringToInt(parts[1]),
		// Use zero indexing for stack number
		helpers.StringToInt(parts[3]) - 1,
		helpers.StringToInt(parts[5]) - 1,
	}
}

func parseCrates(lines []string) []stack {
	// This assumes there are no empty stacks to begin with - fine for both inputs
	stackCount := strings.Count(lines[len(lines)-2], "]")
	fmt.Println("Found", stackCount, "stacks")

	stacks := []stack{}

	for i := 0; i < stackCount; i++ {
		stacks = append(stacks, newStack())
	}
	// Iterate over crates from bottom row up so we can push onto the stacks
	for i := len(lines) - 2; i >= 0; i-- {
		level := lines[i]
		for stackNo := 0; stackNo < stackCount; stackNo++ {
			// Crates diagram is padded with spaces, so we will never overshoot
			char := string(level[(stackNo*4)+1])
			// Non-space character means there is a crate for this stack at this level
			if char != " " {
				stacks[stackNo].push(char)
			}
		}
	}
	return stacks
}

func newStack() stack {
	empty := []string{}
	return stack{empty}
}
