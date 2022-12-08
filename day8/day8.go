package day8

import (
	"fmt"
	"math"
	"strconv"

	"github.com/adgai19/aoc2022/utils"
)

func Day8a(input string) int {
	grid := make([][]int, 0)
	for _, line := range utils.GetLines(input) {
		row := make([]int, 0)
		for _, char := range []rune(line) {
			int, _ := strconv.Atoi(string(char))
			row = append(row, int)
		}
		grid = append(grid, row)
	}
	visible := len(grid)*4 - 4
	for i := 1; i < len(grid)-1; i++ {
		for j := 1; j < len(grid[i])-1; j++ {
			if isVisible(grid, i, j, grid[i][j]) {
				visible += 1
			}
		}
	}

	return visible
}

func Day8b(input string) int {
	grid := make([][]int, 0)
	for _, line := range utils.GetLines(input) {
		row := make([]int, 0)
		for _, char := range []rune(line) {
			int, _ := strconv.Atoi(string(char))
			row = append(row, int)
		}
		grid = append(grid, row)
	}
	max := math.MinInt
	dirArray := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	for i := 1; i < len(grid)-1; i++ {
		for j := 1; j < len(grid[i])-1; j++ {
			count := 1
			for k := range dirArray {
				count *= score(grid, i, j, grid[i][j], dirArray[k], 1)
			}
			if count > max {
				max = count
			}
		}
	}
	fmt.Println(max)

	return max
}

func isVisible(grid [][]int, i, j, cur int) bool {
	dirArray := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	for k := range dirArray {
		count := check(grid, i, j, cur, dirArray[k])
		if count == 1 {
			return true
		}
	}
	return false
}

func check(graph [][]int, row int, col int, currVal int, dir []int) int {
	length := len(graph)
	newRow := row + dir[0]
	newCol := col + dir[1]
	// reached the end
	if newRow >= length || newRow < 0 || newCol >= length || newCol < 0 {
		return 1
	}
	// found a bigger number
	if graph[newRow][newCol] >= currVal {
		return 0
	}
	return check(graph, newRow, newCol, currVal, dir)
}

func score(graph [][]int, row int, col int, currVal int, dir []int, currScore int) int {
	length := len(graph)
	newRow := row + dir[0]
	newCol := col + dir[1]
	if newRow >= length || newRow < 0 || newCol >= length || newCol < 0 {
		return 0
	}
	// found a bigger number
	if graph[newRow][newCol] >= currVal {
		return currScore
	}
	return currScore + score(graph, newRow, newCol, currVal, dir, 1)
}
