package day8

import (
	"adventofcode/helpers"
	"fmt"
)

// Day8 Puzzle
func Day8(filepath string) {
	fmt.Println("Day 8 - Tree top treehouse scouting!")

	trees := parseTrees(filepath)

	visible := findVisibleTrees(trees)
	fmt.Println("Found", len(visible), "visible trees")
	// fmt.Println(visible)

	score := findHighestScenicScore(trees)
	fmt.Println("Highest scenic score of any tree =", score)
}

func findVisibleTrees(trees [][]int) map[string]bool {
	// These aren't *really* height and width now but makes everything else more concise
	height := len(trees) - 1
	width := len(trees[0]) - 1
	visible := make(map[string]bool)

	// All trees at the edges are visible
	for i := 0; i <= height; i++ {
		visible[fmt.Sprintf("%d,0", i)] = true
		visible[fmt.Sprintf("%d,%d", i, height)] = true
	}
	for j := 0; j <= width; j++ {
		visible[fmt.Sprintf("0,%d", j)] = true
		visible[fmt.Sprintf("%d,%d", width, j)] = true
	}

	// Find trees visible from left or right
	for i, row := range trees {
		maxHeightFromLeft := row[0]
		maxHeightFromRight := row[width]

		for j, treeHeight := range row {
			oppositeHeight := row[width-j]

			// Tallest tree from left
			if treeHeight > maxHeightFromLeft {
				visible[fmt.Sprintf("%d,%d", i, j)] = true
				maxHeightFromLeft = treeHeight
			}
			// Tallest tree from right
			if oppositeHeight > maxHeightFromRight {
				visible[fmt.Sprintf("%d,%d", i, width-j)] = true
				maxHeightFromRight = oppositeHeight
			}
		}
	}
	// Find trees visible from top or bottom
	for j := 0; j <= width; j++ {
		maxHeightFromTop := trees[0][j]
		maxHeightFromBottom := trees[height][j]

		for i := 0; i <= height; i++ {
			treeHeight := trees[i][j]
			oppositeHeight := trees[height-i][j]

			if treeHeight > maxHeightFromTop {
				visible[fmt.Sprintf("%d,%d", i, j)] = true
				maxHeightFromTop = treeHeight
			}
			if oppositeHeight > maxHeightFromBottom {
				visible[fmt.Sprintf("%d,%d", height-i, j)] = true
				maxHeightFromBottom = oppositeHeight
			}
		}
	}
	return visible
}

func findHighestScenicScore(trees [][]int) int {
	highest := 0
	for i, row := range trees {
		for j := range row {
			score := calculateScenicScore(trees, i, j)
			if score > highest {
				highest = score
			}
		}
	}
	return highest
}

func calculateScenicScore(trees [][]int, x int, y int) int {
	height := len(trees) - 1
	width := len(trees[0]) - 1
	thisHeight := trees[x][y]

	visibleAbove := 0
	for i := x - 1; i >= 0; i-- {
		visibleAbove++
		if trees[i][y] >= thisHeight {
			break
		}
	}
	visibleBelow := 0
	for i := x + 1; i <= height; i++ {
		visibleBelow++
		if trees[i][y] >= thisHeight {
			break
		}
	}
	visibleLeft := 0
	for j := y - 1; j >= 0; j-- {
		visibleLeft++
		if trees[x][j] >= thisHeight {
			break
		}
	}
	visibleRight := 0
	for j := y + 1; j <= width; j++ {
		visibleRight++
		if trees[x][j] >= thisHeight {
			break
		}
	}
	return visibleAbove * visibleLeft * visibleRight * visibleBelow
}

func parseTrees(filepath string) [][]int {
	trees := make([][]int, 0)
	for _, l := range helpers.ReadLines(filepath) {
		row := make([]int, 0)
		for _, t := range l {
			row = append(row, helpers.RuneToInt(t))
		}
		trees = append(trees, row)
	}
	return trees
}
