package day17

import (
	"fmt"
	"log"
	"os"
)

const (
	LEFT       = '<'
	RIGHT      = '>'
	CAVE_WIDTH = 7
	ROCK       = "#"
)

// Day17 puzzle
func Day17(filepath string) {
	fmt.Println("Day 17 - Pyrocastic Tetris!")

	gasJets, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(gasJets))

	rocksToDrop := 2022
	gasJetIndex := 0
	cave := make([][]string, 0)

	var block RockBlock
	for r := 1; r <= rocksToDrop; r++ {
		// fmt.Println("\nBlock number", r)
		// Block (defined by its bottom left coord) starts two units away from left
		// side and 3 units above the highest rock
		block = makeBlock(r, point{2, len(cave) + 3})
		dropBlock(block, &cave, gasJets, &gasJetIndex)
		// printCave(cave)
	}
	printCave(cave)
	fmt.Println("Height of the cave when block", rocksToDrop+1, "appears is", len(cave))
}

func dropBlock(block RockBlock, cave *[][]string, gasJets []byte, gasJetIndex *int) {
	stopped := false
	for !stopped {
		// fmt.Println("Move number", *gasJetIndex, ": moving", string(gasJets[*gasJetIndex]))
		switch gasJets[*gasJetIndex] {
		case LEFT:
			block = tryMoveLeft(*cave, block)
		case RIGHT:
			block = tryMoveRight(*cave, block)
		default:
			panic("No move!")
		}

		// Move down
		block = block.moveDown()
		for _, b := range block.downBoundary() {
			// fmt.Println("Checking Down boundary", b.x, b.y)
			// No other blocks at this level so nothing to check
			if b.y >= len(*cave) {
				continue
			} else if b.y < 0 || (*cave)[b.y][b.x] == ROCK {
				// Block is stopped, add rocks to the cave
				block = block.moveUp()
				stopped = true
				addToRocks(cave, block)
				break
			}
		}

		*gasJetIndex++
		// Jets restart at beginning again once they've all fired
		if *gasJetIndex == len(gasJets)-1 {
			*gasJetIndex = 0
		}
	}
}

func tryMoveLeft(cave [][]string, block RockBlock) RockBlock {
	block = block.moveLeft()
	for _, b := range block.leftBoundary() {
		// Hit the wall -> move can't happen so move back
		if b.x < 0 {
			block = block.moveRight()
			break
			// No rocks at this level yet -> nothing to bump into
		} else if b.y >= len(cave) {
			continue
			// Hit another rock -> move can't happen so move back
		} else if cave[b.y][b.x] == ROCK {
			block = block.moveRight()
			break
		}
	}
	return block
}

func tryMoveRight(cave [][]string, block RockBlock) RockBlock {
	block = block.moveRight()
	for _, b := range block.rightBoundary() {
		if b.x >= CAVE_WIDTH {
			block = block.moveLeft()
			break
		} else if b.y >= len(cave) {
			continue
		} else if cave[b.y][b.x] == ROCK {
			block = block.moveLeft()
			break
		}
	}
	return block
}

func addToRocks(cave *[][]string, b RockBlock) {
	for _, r := range b.allCoords() {
		// Need to do this from lowest block first
		if len(*cave) == 0 || r.y >= len(*cave) {
			*cave = append(*cave, []string{".", ".", ".", ".", ".", ".", "."})
		}
		(*cave)[r.y][r.x] = ROCK
	}
}

func makeBlock(rockType int, position point) RockBlock {
	var block RockBlock
	switch rockType % 5 {
	case 1:
		block = Rock1{position}
	case 2:
		block = Rock2{position}
	case 3:
		block = Rock3{position}
	case 4:
		block = Rock4{position}
	case 0:
		block = Rock5{position}
	}
	return block
}

func printCave(cave [][]string) {
	for i := len(cave) - 1; i >= 0; i-- {
		fmt.Println(fmt.Sprintf("% 4d", i), cave[i])
	}
}
