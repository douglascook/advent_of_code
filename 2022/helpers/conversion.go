package helpers

import (
	"log"
	"strconv"
)

// StringToInt converts given string to an int, failing on any error
func StringToInt(s string) int {
	integer, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return integer
}

// RuneToInt converts given rune to an int, failing on any error
func RuneToInt(r rune) int {
	integer, err := strconv.Atoi(string(r))
	if err != nil {
		log.Fatal(err)
	}
	return integer
}
