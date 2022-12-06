package day6

import (
	"fmt"
	"log"
	"os"
)

// Day6 Puzzle
func Day6(filepath string) {
	fmt.Println("Day 6 - Tuning communication device thing")

	datastream, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}

	var marker int
	marker = FindUniqueCharsChunk(string(datastream), 4)
	fmt.Println("Found start of packet marker in datastream ending at character", marker)

	marker = FindUniqueCharsChunk(string(datastream), 14)
	fmt.Println("Found start of message marker in datastream ending at character", marker)
}

// FindUniqueCharsChunk returns the index of the end of the first chunk of given
// size containing unique characters
func FindUniqueCharsChunk(datastream string, chunkSize int) int {
	for i := chunkSize; i <= len(datastream); i++ {
		chunk := datastream[i-chunkSize : i]

		found := make(map[rune]bool)
		for _, char := range chunk {
			// If same char is encountered again the chunk contains non-unique chars
			_, exists := found[char]
			if exists {
				break
			}
			found[char] = true
		}
		// If all the characters are unique then this is desired chunk
		if len(found) == chunkSize {
			return i
		}
	}
	log.Fatal("Did not find any marker")
	return 0
}
