package day8

import (
	"fmt"
	"testing"
)

func TestVisibleFromLeft(t *testing.T) {
	trees := [][]int{
		{9, 9, 9},
		{1, 2, 9},
		{9, 9, 9},
	}
	coords := []string{"1,1"}
	checkVisible(trees, coords, t)
}

func TestVisibleFromLeft2(t *testing.T) {
	trees := [][]int{
		{9, 9, 9, 9},
		{1, 2, 3, 9},
		{1, 2, 3, 9},
		{9, 9, 9, 9},
	}
	coords := []string{"1,0", "1,1", "1,2", "2,0", "2,1", "2,2"}
	checkVisible(trees, coords, t)
}

func TestVisibleFromLeft3(t *testing.T) {
	trees := [][]int{
		{9, 9, 9, 9},
		{1, 3, 1, 9},
		{2, 1, 4, 9},
		{9, 9, 9, 9},
	}
	coords := []string{"1,1", "2,3"}
	checkVisible(trees, coords, t)
	checkNotVisible(trees, []string{"1,2", "2,1"}, t)
}

func TestVisibleFromRight(t *testing.T) {
	trees := [][]int{
		{9, 9, 9},
		{9, 2, 1},
		{9, 9, 9},
	}
	coords := []string{"1,1"}
	checkVisible(trees, coords, t)
}

func TestVisibleFromRight2(t *testing.T) {
	trees := [][]int{
		{9, 9, 9, 9},
		{9, 3, 2, 1},
		{9, 3, 2, 1},
		{9, 9, 9, 9},
	}
	coords := []string{"2,3", "2,2", "2,1", "1,3", "1,2", "1,1"}
	checkVisible(trees, coords, t)
}

func TestVisibleFromRight3(t *testing.T) {
	trees := [][]int{
		{9, 9, 9, 9},
		{9, 4, 1, 2},
		{9, 1, 3, 1},
		{9, 9, 9, 9},
	}
	checkVisible(trees, []string{"2,3", "2,2", "1,3", "1,1"}, t)
	checkNotVisible(trees, []string{"2,1", "1,2"}, t)
}

func TestVisibleFromTop(t *testing.T) {
	trees := [][]int{
		{9, 1, 9},
		{9, 2, 9},
		{9, 9, 9},
	}
	coords := []string{"1,1"}
	checkVisible(trees, coords, t)
}

func TestVisibleFromTop2(t *testing.T) {
	trees := [][]int{
		{9, 1, 1, 9},
		{9, 2, 2, 9},
		{9, 3, 3, 9},
		{9, 9, 9, 9},
	}
	coords := []string{"0,1", "1,1", "2,1", "0,2", "1,2", "2,2"}
	checkVisible(trees, coords, t)
}

func TestVisibleFromTop3(t *testing.T) {
	trees := [][]int{
		{9, 1, 2, 9},
		{9, 3, 1, 9},
		{9, 1, 4, 9},
		{9, 9, 9, 9},
	}
	coords := []string{"0,1", "1,1", "0,2", "2,2"}
	checkVisible(trees, coords, t)
	checkNotVisible(trees, []string{"2,1", "1,2"}, t)
}

func TestVisibleFromBottom1(t *testing.T) {
	trees := [][]int{
		{9, 9, 9},
		{9, 2, 9},
		{9, 1, 9},
	}
	coords := []string{"1,1"}
	checkVisible(trees, coords, t)
}

func TestVisibleFromBottom2(t *testing.T) {
	trees := [][]int{
		{9, 9, 9, 9},
		{9, 3, 3, 9},
		{9, 2, 2, 9},
		{9, 1, 1, 9},
	}
	checkVisible(trees, []string{"3,1", "2,1", "1,1", "3,2", "2,2", "1,2"}, t)
}

func TestVisibleFromBottom3(t *testing.T) {
	trees := [][]int{
		{9, 9, 9, 9},
		{9, 1, 4, 9},
		{9, 3, 1, 9},
		{9, 1, 2, 9},
	}
	checkVisible(trees, []string{"3,1", "2,1", "3,2", "1,2"}, t)
	checkNotVisible(trees, []string{"1,1", "2,2"}, t)
}

func TestNotVisible(t *testing.T) {
	trees := [][]int{
		{1, 1, 1, 1, 1},
		{1, 5, 5, 5, 1},
		{1, 5, 2, 5, 1},
		{1, 5, 5, 5, 1},
		{1, 1, 1, 1, 1},
	}
	coords := []string{"2,2"}
	checkNotVisible(trees, coords, t)
}

func checkVisible(trees [][]int, coords []string, t *testing.T) {
	visible := findVisibleTrees(trees)
	for _, c := range coords {
		if !visible[c] {
			fmt.Println(visible)
			t.Fatal("Couldn't see", c)
		}
	}
}

func checkNotVisible(trees [][]int, coords []string, t *testing.T) {
	visible := findVisibleTrees(trees)
	for _, c := range coords {
		if visible[c] {
			fmt.Println(visible)
			t.Fatal("Shouldn't be able to see", c)
		}
	}
}
