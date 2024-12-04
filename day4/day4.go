package day4

import (
	"fmt"

	"github.com/revzik/aoc_2024/common/files"
)

func RunTask() {
	lines := files.ReadLines("day4/input.txt")

	fmt.Printf("XMAS occurrences: %d\n", countXMAS(lines))
	fmt.Printf("X-MAS occurrences: %d\n", countXShapedMAS(lines))
}

func countXMAS(lines []string) int {
	matrix := linesToMatrix(lines)
	total := 0

	// We'll find X and then search nearby runes for matches
	for i, row := range matrix {
		for j := 0; j < len(row); j++ {
			total += searchXMAS(matrix, 'X', j, i, 0)
		}
	}

	return total
}

func searchXMAS(matrix [][]rune, char rune, x int, y int, dir int) int {
	if matrix[y][x] != char {
		return 0
	}

	if matrix[y][x] != 'X' && matrix[y][x] != 'M' && matrix[y][x] != 'A' && matrix[y][x] != 'S' {
		return 0
	}

	if char == 'S' {
		return 1
	}

	var searchChar rune
	if char == 'X' {
		searchChar = 'M'
	} else if char == 'M' {
		searchChar = 'A'
	} else if char == 'A' {
		searchChar = 'S'
	}

	found := 0
	maxX := len(matrix[0])
	maxY := len(matrix)

	if (dir == 0 || dir == 1) && x-1 >= 0 && y-1 >= 0 {
		found += searchXMAS(matrix, searchChar, x-1, y-1, 1)
	}
	if (dir == 0 || dir == 2) && y-1 >= 0 {
		found += searchXMAS(matrix, searchChar, x, y-1, 2)
	}
	if (dir == 0 || dir == 3) && x+1 < maxX && y-1 >= 0 {
		found += searchXMAS(matrix, searchChar, x+1, y-1, 3)
	}
	if (dir == 0 || dir == 4) && x-1 >= 0 {
		found += searchXMAS(matrix, searchChar, x-1, y, 4)
	}
	if (dir == 0 || dir == 5) && x+1 < maxX {
		found += searchXMAS(matrix, searchChar, x+1, y, 5)
	}
	if (dir == 0 || dir == 6) && x-1 >= 0 && y+1 < maxY {
		found += searchXMAS(matrix, searchChar, x-1, y+1, 6)
	}
	if (dir == 0 || dir == 7) && y+1 < maxY {
		found += searchXMAS(matrix, searchChar, x, y+1, 7)
	}
	if (dir == 0 || dir == 8) && x+1 < maxX && y+1 < maxY {
		found += searchXMAS(matrix, searchChar, x+1, y+1, 8)
	}

	return found
}

func countXShapedMAS(lines []string) int {
	matrix := linesToMatrix(lines)
	total := 0

	for i := 1; i < len(matrix)-1; i++ {
		for j := 1; j < len(matrix[i])-1; j++ {
			if searchXShapedMAS(matrix, j, i) {
				total++
			}
		}
	}

	return total
}

func searchXShapedMAS(matrix [][]rune, x int, y int) bool {
	if matrix[y][x] != 'A' {
		return false
	}

	if matrix[y-1][x-1] == 'M' && matrix[y-1][x+1] == 'M' && matrix[y+1][x-1] == 'S' && matrix[y+1][x+1] == 'S' {
		return true
	}
	if matrix[y-1][x-1] == 'M' && matrix[y-1][x+1] == 'S' && matrix[y+1][x-1] == 'M' && matrix[y+1][x+1] == 'S' {
		return true
	}
	if matrix[y-1][x-1] == 'S' && matrix[y-1][x+1] == 'S' && matrix[y+1][x-1] == 'M' && matrix[y+1][x+1] == 'M' {
		return true
	}
	if matrix[y-1][x-1] == 'S' && matrix[y-1][x+1] == 'M' && matrix[y+1][x-1] == 'S' && matrix[y+1][x+1] == 'M' {
		return true
	}

	return false
}

func linesToMatrix(lines []string) [][]rune {
	// assume the input is sanitized
	x := len(lines)

	matrix := make([][]rune, x)
	for i, line := range lines {
		matrix[i] = []rune(line)
	}

	return matrix
}
