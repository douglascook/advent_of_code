package day7

import (
	"adventofcode/helpers"
	"fmt"
	"strings"
)

type dir struct {
	name        string
	path        string
	directories map[string]*dir
	files       map[string]int
}

// Day7 Puzzle
func Day7(filepath string) {
	fmt.Println("Day 7 - Filesystem no space left!")
	commands := helpers.ReadLines(filepath)

	root := parseFilesystem(commands)
	sizes := make(map[string]int)

	calculateDirSizes(&root, sizes)
	fmt.Println("Directory sizes are", sizes)

	total := getSumOfDirsWithSizeLessThan(sizes, 100000)
	fmt.Println("Sum of dir sizes < 100000 =", total)

	dirToDelete := getDirToDelete(sizes)
	fmt.Println("Size of directory to delete to free up space =", dirToDelete)
}

func parseFilesystem(commands []string) dir {
	root := newDir("/", nil)

	var d *dir = &root
	var name string

	// Skip first command, we are already in root dir
	for _, c := range commands[1:] {
		parts := strings.Split(c, " ")

		// Moving directory - the directory must already have been listed so just
		// update current d
		if parts[0] == "$" && parts[1] == "cd" {
			name = parts[2]
			fmt.Println("Current directory is", d.path)
			next := d.directories[name]
			fmt.Println("Moving into", next.name)
			d = next

			// No information from this command, just skip it
		} else if parts[0] == "$" && parts[1] == "ls" {
			fmt.Println("Listing current directory")
			continue

			// Found a new directory, add to current dir's directories
		} else if parts[0] == "dir" {
			name = parts[1]
			fmt.Println("Found new directory", name, "under", d.path)
			nested := newDir(name, d)
			d.directories[name] = &nested
			fmt.Println("Updated directories under", d.path, "are :", d.directories)

			// Found a new file, add to current dir's files
		} else {
			name = parts[1]
			fmt.Println("Found file", name)
			d.files[name] = helpers.StringToInt(parts[0])
		}
	}
	return root
}

func calculateDirSizes(d *dir, sizes map[string]int) {
	size := 0
	for _, s := range d.files {
		size += s
	}

	for dirName, pointer := range d.directories {
		if dirName != ".." {
			calculateDirSizes(pointer, sizes)
			size += sizes[pointer.path]
		}
	}
	sizes[d.path] = size
}

func getSumOfDirsWithSizeLessThan(sizes map[string]int, maxSize int) int {
	total := 0
	for d, s := range sizes {
		if s <= maxSize {
			fmt.Println(d, "has size", s, "LESS than", maxSize)
			total += s
		}
	}
	return total
}

func getDirToDelete(sizes map[string]int) int {
	// 70m in total
	// 30m required so <= 40m used
	// delete dirs >= current used - 40m
	requiredToFree := sizes["/"] - 40000000
	fmt.Println("Need to delete at least", requiredToFree)

	var smallest int = 70000000
	for _, s := range sizes {
		if s >= requiredToFree && s < smallest {
			smallest = s
		}
	}
	return smallest
}

func newDir(name string, parent *dir) dir {
	// Directory names are not unique, so need to keep track of full path as well
	var path string
	if parent == nil {
		path = "/"
	} else {
		path = parent.path + name + "/"
	}
	return dir{
		name,
		path,
		// Instantiate directories with parent
		map[string]*dir{"..": parent},
		// Files empty to begin with
		make(map[string]int),
	}
}
