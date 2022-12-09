package day9

import (
	"adventofcode/helpers"
	"fmt"
	"log"
	"strings"
)

const (
	up    = "U"
	down  = "D"
	left  = "L"
	right = "R"
)

type point struct {
	x int
	y int
}

func (p point) toString() string {
	return fmt.Sprintf("%d,%d", p.x, p.y)
}

func (p point) touching(other point) bool {
	return p.x >= other.x-1 && p.x <= other.x+1 && p.y >= other.y-1 && p.y <= other.y+1
}

func Day9(filepath string) {
	fmt.Println("Day 9 - Rope Bridge Physics Modelling!")

	motions := helpers.ReadLines(filepath)
	ropeSimulation(motions, 2)
	ropeSimulation(motions, 10)
}

func ropeSimulation(motions []string, ropeLength int) {
	fmt.Println("Simulating rope of length", ropeLength)

	knots := make([]point, 0)
	for k := 0; k < 10; k++ {
		knots = append(knots, point{0, 0})
	}
	tailVisited := map[string]int{"0,0": 0}

	for _, m := range motions {
		mot := strings.Split(m, " ")
		direction := mot[0]
		magnitude := helpers.StringToInt(mot[1])

		for i := 0; i < magnitude; i++ {
			// Update head
			knots[0] = updateHead(knots[0], direction)

			// Update middle knots
			for k := 1; k < ropeLength-1; k++ {
				if !knots[k].touching(knots[k-1]) {
					knots[k] = updateTail(knots[k-1], knots[k], nil)
				}
			}
			// Update tail and keep track of visited
			if !knots[ropeLength-1].touching(knots[ropeLength-2]) {
				knots[ropeLength-1] = updateTail(knots[ropeLength-2], knots[ropeLength-1], &tailVisited)
			}
		}
	}
	fmt.Println("Tail visited", len(tailVisited), "positions:")
}

func updateHead(head point, direction string) point {
	switch direction {
	case right:
		head.x++
	case left:
		head.x--
	case up:
		head.y++
	case down:
		head.y--
	default:
		log.Fatal("Cannot process unexpected direction", direction)
	}
	return head
}

func updateTail(head point, tail point, visited *map[string]int) point {
	// No need to change x if head and tail are on same row
	if tail.x < head.x {
		tail.x++
	} else if tail.x > head.x {
		tail.x--
	}
	// No need to change y if head and tail are on same column
	if tail.y < head.y {
		tail.y++
	} else if tail.y > head.y {
		tail.y--
	}

	if visited != nil {
		value, exists := (*visited)[tail.toString()]
		if !exists {
			value = 0
		}
		(*visited)[tail.toString()] = value + 1
	}
	return tail
}
