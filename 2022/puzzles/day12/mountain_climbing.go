package day12

import (
	"adventofcode/helpers"
	"fmt"
)

type node struct {
	x int
	y int
}

type heightMap struct {
	_map  [][]rune
	start node
	end   node
}

// Return neighbours whose height is at most 1 greater than current height
func (m heightMap) getValidSteps(n node) []node {
	currentHeight := m._map[n.x][n.y]
	steps := make([]node, 0)

	// Up neighbour
	if n.x > 0 && m._map[n.x-1][n.y] <= currentHeight+1 {
		steps = append(steps, node{n.x - 1, n.y})
	}
	// Down neighbour
	if n.x < len(m._map)-1 && m._map[n.x+1][n.y] <= currentHeight+1 {
		steps = append(steps, node{n.x + 1, n.y})
	}
	// Left neighbour
	if n.y > 0 && m._map[n.x][n.y-1] <= currentHeight+1 {
		steps = append(steps, node{n.x, n.y - 1})
	}
	// Right neighbour
	if n.y < len(m._map[0])-1 && m._map[n.x][n.y+1] <= currentHeight+1 {
		steps = append(steps, node{n.x, n.y + 1})
	}
	return steps
}

// Manhattan distance from end node
func (m heightMap) distanceFromEnd(n node) int {
	distance := 0
	if m.end.x < n.x {
		distance += (n.x - m.end.x)
	} else {
		distance += (m.end.x - n.x)
	}
	if m.end.y < n.y {
		distance += (n.y - m.end.y)
	} else {
		distance += (m.end.y - n.y)
	}
	return distance
}

func (m heightMap) getAlternativeStarts() []node {
	lowPoints := make([]node, 0)
	for i, row := range m._map {
		for j, height := range row {
			n := node{i, j}
			if n == m.start {
				continue
			} else if height == 'a' {
				lowPoints = append(lowPoints, node{i, j})
			}
		}
	}
	return lowPoints
}

// Day12 puzzle
func Day12(filepath string) {
	fmt.Println("Day 12 - Mountain Climbing Algorithm!")

	_map := readMap(filepath)
	distance := findShortestPath(_map, _map.start)
	fmt.Println("Shortest path from start to end has length", distance)

	shortest := 1000000
	for _, start := range _map.getAlternativeStarts() {
		fmt.Println("Starting at", start)
		distance = findShortestPath(_map, start)
		if distance < shortest {
			shortest = distance
		}
	}
	fmt.Println("Shortest path from any low point to end has length", shortest)
}

// findShortestPath from start to end of map using A* algorithm
func findShortestPath(_map heightMap, start node) int {
	// maps node to score from start to the node, using currently shortest path
	scoreToNode := make(map[node]int)
	// nodes to search over, mapped to predicted score if path goes through the node
	frontier := make(map[node]int)

	// Initialise search with node to start at
	frontier[start] = _map.distanceFromEnd(start)
	scoreToNode[start] = 0

	for len(frontier) > 0 {
		current := getNodeWithLowestScore(frontier)
		if current == _map.end {
			return frontier[current]
		}
		scoreToCurrent := scoreToNode[current]

		delete(frontier, current)
		for _, next := range _map.getValidSteps(current) {
			// If new shortest path to next goes through current then update score
			scoreToNext, visited := scoreToNode[next]
			if !visited || (scoreToCurrent+1 < scoreToNext) {
				// Cost from current to next is always 1
				scoreToNode[next] = scoreToCurrent + 1
				// Add node to search
				frontier[next] = scoreToNode[next] + _map.distanceFromEnd(next)
			}
		}
	}
	fmt.Println("Could not find a path from start to end!")
	return 1000000
}

func getNodeWithLowestScore(scores map[node]int) node {
	var best node
	lowest := 1000000
	for n := range scores {
		if scores[n] < lowest {
			lowest = scores[n]
			best = n
		}
	}
	return best
}

func readMap(filepath string) heightMap {
	_map := make([][]rune, 0)
	var start, end node

	for i, line := range helpers.ReadLines(filepath) {
		row := make([]rune, 0)
		for j, l := range line {
			// Start has height a
			if l == 'S' {
				start = node{i, j}
				row = append(row, 'a')
				// End has height z
			} else if l == 'E' {
				end = node{i, j}
				row = append(row, 'z')
			} else {
				row = append(row, l)
			}
		}
		_map = append(_map, row)
	}
	return heightMap{_map, start, end}
}
