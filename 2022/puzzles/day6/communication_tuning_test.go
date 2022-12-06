package day6

import "testing"

// Start of PACKET markers are indicated by chunk of 4 unique characters
func TestFindStartOfPacketMarker(t *testing.T) {
	testExpectedMarkerFound(t, "mjqjpqmgbljsphdztnvjfqwrcgsmlb", 4, 7)
	testExpectedMarkerFound(t, "bvwbjplbgvbhsrlpgdmjqwftvncz", 4, 5)
	testExpectedMarkerFound(t, "nppdvjthqldpwncqszvftbrmjlhg", 4, 6)
	testExpectedMarkerFound(t, "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 4, 10)
	testExpectedMarkerFound(t, "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 4, 11)
}

// Start of MESSAGE markers are indicated by chunk of 14 unique characters
func TestFindStartOfMessageMarker(t *testing.T) {
	testExpectedMarkerFound(t, "mjqjpqmgbljsphdztnvjfqwrcgsmlb", 14, 19)
	testExpectedMarkerFound(t, "bvwbjplbgvbhsrlpgdmjqwftvncz", 14, 23)
	testExpectedMarkerFound(t, "nppdvjthqldpwncqszvftbrmjlhg", 14, 23)
	testExpectedMarkerFound(t, "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 14, 29)
	testExpectedMarkerFound(t, "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 14, 26)
}

func testExpectedMarkerFound(t *testing.T, datastream string, chunkSize int, expected int) {
	marker := FindUniqueCharsChunk(datastream, chunkSize)
	if !(marker == expected) {
		t.Fatal("FindMarker", datastream, "=", marker, ", expected", expected)
	}
}
