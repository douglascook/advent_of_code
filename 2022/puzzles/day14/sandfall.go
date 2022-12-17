package day14

import (
	"adventofcode/helpers"
	"fmt"
	"math"
	"strings"
)

const (
	ROCK = "#"
	AIR  = "."
	SAND = "o"
)

type point struct {
	x int
	y int
}

// Cave is a 2D structure consisting of sand and air
type Cave struct {
	_cave [][]string
	minX  int
	maxX  int
}

func Day14(filepath string) {
	fmt.Println("Day 14 - Sand Fall!")
	paths, minX, maxX, maxY := parsePaths(helpers.ReadLines(filepath))

	fmt.Println("Part 1 - Checking how many grains of sand it takes until overflow.")
	cave := plotCave(paths, minX, maxX, maxY)
	overflowing := false
	blocked := false
	var i int
	for i = 0; !overflowing; i++ {
		overflowing, blocked = cave.dropSand()
		// cave.print(i, false)
	}
	fmt.Println("Took", i-1, "grains of sand until cave was overflowing")

	fmt.Println("Part 2 - Checking how many grains of sand until inflow blocked")
	// Make cave wider to cater for grains falling off right over rock paths
	cave = plotCave(paths, minX, maxX*2, maxY)
	addFloor(&cave, maxX*2, maxY)
	// cave.print(0, true)

	overflowing = false
	blocked = false
	for i = 0; !blocked; i++ {
		overflowing, blocked = cave.dropSand()
		// cave.print(i, true)
	}
	// Final grain blocking entrance is counted
	fmt.Println("Took", i, "grains of sand until cave was blocked at input")
}

func parsePaths(lines []string) ([][]point, int, int, int) {
	minX := 100000
	maxX := 0
	minY := 100000
	maxY := 0

	paths := make([][]point, 0)
	for _, l := range lines {
		path := make([]point, 0)

		for _, p := range strings.Split(l, " -> ") {
			coords := strings.Split(p, ",")
			x := helpers.StringToInt(coords[0])
			y := helpers.StringToInt(coords[1])
			path = append(path, point{x, y})

			if x < minX {
				minX = x
			} else if x > maxX {
				maxX = x
			}
			if y < minY {
				minY = y
			} else if y > maxY {
				maxY = y
			}
		}
		paths = append(paths, path)
	}
	return paths, minX, maxX, maxY
}

func plotCave(paths [][]point, minX int, maxX int, maxY int) Cave {
	// Create empty cave big enough to hold all paths
	cave := createCave(maxX, maxY)
	cave.minX = minX
	cave.maxX = maxX

	for _, p := range paths {
		cave.plotRockPath(p)
	}
	return cave
}

func createCave(x, y int) Cave {
	fmt.Println("Cave has size", x, "x", y)
	cave := make([][]string, 0)
	for j := 0; j <= y; j++ {
		row := make([]string, 0)
		for i := 0; i <= x; i++ {
			row = append(row, AIR)
		}
		cave = append(cave, row)
	}
	return Cave{cave, 100000, 0}
}

func addFloor(cave *Cave, maxX int, maxY int) {
	fmt.Println("Adding floor to cave")
	(*cave)._cave = append(cave._cave, make([]string, maxX+1))
	(*cave)._cave = append(cave._cave, make([]string, maxX+1))
	for i := 0; i <= maxX; i++ {
		(*cave)._cave[maxY+1][i] = AIR
		(*cave)._cave[maxY+2][i] = ROCK
	}
}

func (c Cave) plotRockPath(path []point) {
	for i := 1; i < len(path); i++ {
		p1 := path[i-1]
		p2 := path[i]

		startY := int(math.Min(float64(p1.y), float64(p2.y)))
		endY := int(math.Max(float64(p1.y), float64(p2.y)))
		startX := int(math.Min(float64(p1.x), float64(p2.x)))
		endX := int(math.Max(float64(p1.x), float64(p2.x)))

		for x := startX; x <= endX; x++ {
			for y := startY; y <= endY; y++ {
				c._cave[y][x] = ROCK
			}
		}
	}
}

func (c Cave) print(i int, bigCave bool) {
	fmt.Println()
	fmt.Println("Grain", i)
	for _, l := range c._cave {
		if bigCave {
			fmt.Println(l[c.minX-10 : c.maxX/2+10])
		} else {
			fmt.Println(l[c.minX-1 : c.maxX+1])
		}
	}
}

func (c Cave) dropSand() (bool, bool) {
	x := 500
	y := 0
	stopped := false
	for !stopped {
		// If sand has gone past bounds of the cave then it is overflowing
		if y >= len(c._cave)-1 {
			return true, false
			// Must be in cave with floor. Grain has gone off to left or right doesn't
			// matter where so continue with next grain.
		} else if x < 0 || x >= c.maxX {
			stopped = true
			// Fall straight down
		} else if c._cave[y+1][x] == AIR {
			y++
			// Diagonal to left
		} else if c._cave[y+1][x-1] == AIR {
			y++
			x--
			// Diagonal to right
		} else if c._cave[y+1][x+1] == AIR {
			y++
			x++
			// Sand can go no further
		} else {
			c._cave[y][x] = SAND
			stopped = true
			// No more sand can come into the cave because entrance is blocked
			if x == 500 && y == 0 {
				c.print(10000, true)
				return false, true
			}
		}
	}
	return false, false
}
