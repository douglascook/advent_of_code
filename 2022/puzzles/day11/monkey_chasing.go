package day11

import (
	"adventofcode/helpers"
	"fmt"
	"sort"
	"strings"
)

type monkey struct {
	index          int
	items          []int
	operation      func(int) int
	divisor        int
	trueTarget     int
	falseTarget    int
	itemsInspected int
}

// Test first item in the list and throw it to one of the target monkeys
func (m monkey) throwItem(monkeys *map[int]monkey, worryNormaliser int) monkey {
	m.itemsInspected++

	i := m.items[0]
	m.items = m.items[1:]

	// Monkey inspects item and worry level is modified...
	i = m.operation(i)
	// Not super worried - didn't break it so worry drops down to a third
	if worryNormaliser == 0 {
		i = i / 3
		// Super worried - need to manage worry differently
	} else {
		i = i % worryNormaliser
	}

	// Throw item based on divisibility test
	var target monkey
	if i%m.divisor == 0 {
		target = (*monkeys)[m.trueTarget]
	} else {
		target = (*monkeys)[m.falseTarget]
	}
	target.items = append(target.items, i)
	(*monkeys)[target.index] = target

	return m
}

// Day11 Puzzle
func Day11(filepath string) {
	fmt.Println("Day 11 - Monkey Business!")
	lines := helpers.ReadLines(filepath)

	fmt.Println("Watching monkeys for 20 rounds with normal worry levels")
	monkeys := parseMonkeys(lines)
	monkeyBusiness := observeMonkeys(monkeys, 20, false)
	fmt.Println("Monkey business score =", monkeyBusiness)

	fmt.Println("Watching monkeys for 10000 rounds with heightened worry levels")
	monkeys = parseMonkeys(lines)
	monkeyBusiness = observeMonkeys(monkeys, 10000, true)
	fmt.Println("Monkey business score =", monkeyBusiness)
}

func observeMonkeys(monkeys map[int]monkey, numberRounds int, extraWorried bool) int {
	worryNormaliser := 0
	if extraWorried {
		// Need to keep worry levels down, normalise worry by product of all
		// divisors to maintain test outcome with lower worry levels.
		worryNormaliser = 1
		for _, m := range monkeys {
			worryNormaliser *= m.divisor
		}
	}

	for r := 0; r < numberRounds; r++ {
		for m := 0; m < len(monkeys); m++ {
			// fmt.Println("MONKEY", m)
			// fmt.Println("ITEMS TO THROW", monkeys[m].items)
			for range monkeys[m].items {
				monkeys[m] = monkeys[m].throwItem(&monkeys, worryNormaliser)
			}
		}
	}
	return calculateMonkeyBusiness(monkeys)
}

func calculateMonkeyBusiness(monkeys map[int]monkey) int {
	inspectedCounts := make([]int, 0)
	for _, m := range monkeys {
		inspectedCounts = append(inspectedCounts, m.itemsInspected)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(inspectedCounts)))

	// Monkey business is product of two busiest monkeys
	return inspectedCounts[0] * inspectedCounts[1]
}

func parseMonkeys(lines []string) map[int]monkey {
	monkeys := make(map[int]monkey)
	for i := 0; i <= len(lines)/7; i++ {
		monkeys[i] = parseMonkey(lines[i*7:(i+1)*7], i)
	}
	return monkeys
}

func parseMonkey(lines []string, index int) monkey {
	items := make([]int, 0)
	for _, i := range strings.Split(lines[1][18:], ", ") {
		items = append(items, helpers.StringToInt(i))
	}
	operation := parseOperation(lines[2])

	divisor := helpers.StringToInt(lines[3][21:len(lines[3])])
	trueTarget := helpers.StringToInt(lines[4][29:len(lines[4])])
	falseTarget := helpers.StringToInt(lines[5][30:len(lines[5])])

	return monkey{
		index,
		items,
		operation,
		divisor,
		trueTarget,
		falseTarget,
		0,
	}
}

func parseOperation(line string) func(int) int {
	operator := line[23]
	modifier := line[25:len(line)]

	var operation func(int) int
	switch operator {
	case '+':
		operation = add(modifier)
	case '*':
		operation = multiply(modifier)
	}
	return operation
}

func add(modifier string) func(int) int {
	if modifier == "old" {
		return func(i int) int {
			return i + i
		}
	}
	modifierInt := helpers.StringToInt(modifier)
	return func(i int) int {
		return i + modifierInt
	}
}

func multiply(modifier string) func(int) int {
	if modifier == "old" {
		return func(i int) int {
			return i * i
		}
	}
	modifierInt := helpers.StringToInt(modifier)
	return func(i int) int {
		return i * modifierInt
	}
}
