package day15

import (
	"adventofcode/helpers"
	"fmt"
	"log"
	"regexp"
)

type point struct {
	x int
	y int
}

// Calculate manhattan distance between p1 and p2
func (p1 point) distanceTo(p2 point) int {
	deltaX := p1.x - p2.x
	deltaY := p1.y - p2.y

	score := 0
	if deltaX > 0 {
		score += deltaX
	} else {
		score -= deltaX
	}
	if deltaY > 0 {
		score += deltaY
	} else {
		score -= deltaY
	}
	return score
}

// Each sensor has a single closest beacon. The reach of a sensor is the distance
// to the closest beacon.
type sensor struct {
	location      point
	closestBeacon point
	reach         int
}

func (s sensor) findPointsCoveredAtRow(row int) []int {
	distanceFromRow := s.location.y - row
	if distanceFromRow < 0 {
		distanceFromRow = -distanceFromRow
	}

	overshoot := s.reach - distanceFromRow
	// Row is outside reach of the beacon
	if overshoot < 0 {
		return make([]int, 0)
	}
	return []int{s.location.x - overshoot, s.location.x + overshoot}
}

// Day15 puzzle
func Day15(filepath string) {
	fmt.Println("Day 15 - Beacon Search!")

	lines := helpers.ReadLines(filepath)
	sensors := parseSensors(lines)

	targetRow := 2000000
	// targetRow := 10
	fmt.Println("Part 1 - counting points on row", targetRow, "covered by sensors")
	countPointsCoveredOnRow(sensors, targetRow)

	maxCoord := 4000000
	// maxCoord := 20
	fmt.Println("Part 2 - finding the beacon!")
	notCovered := findBeacon(sensors, maxCoord)
	fmt.Println("Sensors cover:", notCovered)
}

func countPointsCoveredOnRow(sensors []sensor, targetRow int) {
	covered := make(map[int]bool)
	for _, s := range sensors {
		interval := s.findPointsCoveredAtRow(targetRow)
		if len(interval) > 0 {
			for i := interval[0]; i < interval[1]; i++ {
				covered[i] = true
			}
		}
	}
	fmt.Println("Beacons cover", len(covered), "points on target row", targetRow)
}

func findBeacon(sensors []sensor, maxCoord int) [][]int {
	for y := 0; y < maxCoord; y++ {
		fullyCovered := false

		// fmt.Println("\nChecking row", y)

		covered := make([][]int, 0)
		for _, s := range sensors {
			// fmt.Println("Checking sensor", s.location)
			sensorCovers := s.findPointsCoveredAtRow(y)

			// Sensor doesn't reach current row
			if len(sensorCovers) == 0 {
				// fmt.Println("Covered is still", covered)
				continue
			}
			start := sensorCovers[0]
			if start < 0 {
				start = 0
			}
			end := sensorCovers[1]
			if end > maxCoord {
				end = maxCoord
			}
			// fmt.Println("New section covered", start, end)

			if len(covered) == 0 {
				covered = append(covered, []int{start, end})
				// fmt.Println("Covered is now", covered)
			} else {
				covered = mergeSections(covered, start, end)
				// fmt.Println("Covered is now", covered)
			}
			if covered[0][0] == 0 && covered[0][1] == maxCoord {
				fullyCovered = true
				break
			}
		}
		// fmt.Println("Covered by sensors", covered)
		if !fullyCovered {
			fmt.Println("Row", y, "is not fully covered!")
			return covered
		}
	}
	return make([][]int, 0)
}

func mergeSections(sections [][]int, start int, end int) [][]int {
	newSections := make([][]int, 0)

	for _, s := range sections {
		// Current section range is contained in existing, nothing changes
		if start >= s[0] && end <= s[1] {
			return sections
		}

		// Current section contains existing, skip existing so it's not included twice
		if start < s[0] && end > s[1] {
			continue
		}

		// Partial overlap, extend section to the right
		if start < s[0] && end >= s[0] && end <= s[1] {
			end = s[1]
			// Partial overlap, extend section to the left
		} else if start >= s[0] && start <= s[1] && end > s[1] {
			start = s[0]
			// This section is disjoint from new section, needs to be returned
		} else {
			newSections = append(newSections, s)
		}
	}

	// Add the (potentially extended) new section
	newSections = append(newSections, []int{start, end})
	return newSections
}

func parseSensors(lines []string) []sensor {
	pattern, err := regexp.Compile(
		`Sensor at x=(-?\d+), y=(-?\d+): closest beacon is at x=(-?\d+), y=(-?\d+)`,
	)
	if err != nil {
		log.Fatal(err)
	}

	sensors := make([]sensor, 0)
	for _, l := range lines {
		sensors = append(sensors, parseSensor(*pattern, l))
	}
	return sensors
}

func parseSensor(pattern regexp.Regexp, text string) sensor {
	match := pattern.FindStringSubmatch(text)
	location := point{helpers.StringToInt(match[1]), helpers.StringToInt(match[2])}
	beacon := point{helpers.StringToInt(match[3]), helpers.StringToInt(match[4])}
	return sensor{location, beacon, location.distanceTo(beacon)}
}

func plotCave(sensors [][]point) [][]string {
	// First decide dimensions of cave. Not sure how you're meant to know when
	// to stop because the cave can go on forever,
	return make([][]string, 0)
}
