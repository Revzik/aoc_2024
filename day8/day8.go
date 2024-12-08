package day8

import (
	"fmt"
	"unicode"

	"github.com/revzik/aoc_2024/common/files"
)

type Point struct {
	x int
	y int
}

type Board struct {
	plane [][]rune
	minX  int
	maxX  int
	minY  int
	maxY  int
}

func RunTask() {
	path := "day8/input.txt"
	matrix := files.ReadMatrix(path)
	board := createBoard(matrix)

	fmt.Printf("Number of weak antinodes: %d\n", countWeakAntinodes(board))
	fmt.Printf("Number of resonant antinodes: %d\n", countResonantAntinodes(board))
}

func createBoard(matrix [][]rune) Board {
	return Board{matrix, 0, len(matrix[0]) - 1, 0, len(matrix) - 1}
}

func countWeakAntinodes(board Board) int {
	total := 0
	antinodes := make(map[Point]bool)

	for y := 0; y < board.maxY; y++ {
		for x := 0; x < board.maxX; x++ {
			if unicode.IsLetter(board.plane[y][x]) || unicode.IsDigit(board.plane[y][x]) {
				markWeakAntinodes(board, antinodes, Point{x, y})
			}
		}
	}

	validateBoundaries(antinodes, board)

	for _, exists := range antinodes {
		if exists {
			total++
		}
	}

	return total
}

func markWeakAntinodes(board Board, antinodes map[Point]bool, p Point) {
	searchedRune := board.plane[p.y][p.x]

	// check only to the bottom, top ones are dealt with
	for i := p.y; i <= board.maxY; i++ {
		for j := 0; j <= board.maxX; j++ {
			if board.plane[i][j] == searchedRune && !(p.x == j && p.y == i) {
				xDist := p.x - j
				yDist := p.y - i

				point1 := Point{p.x + xDist, p.y + yDist}
				point2 := Point{j - xDist, i - yDist}

				antinodes[point1] = true
				antinodes[point2] = true
			}
		}
	}
}

func countResonantAntinodes(board Board) int {
	antinodes := make(map[Point]bool)

	for y := 0; y <= board.maxY; y++ {
		for x := 0; x <= board.maxX; x++ {
			if unicode.IsLetter(board.plane[y][x]) || unicode.IsDigit(board.plane[y][x]) {
				markResonantAntinodes(board, antinodes, Point{x, y})
			}
		}
	}

	return len(antinodes)
}

func markResonantAntinodes(board Board, antinodes map[Point]bool, p Point) {
	searchedRune := board.plane[p.y][p.x]

	// check only to the bottom, top ones are dealt with
	for i := p.y; i <= board.maxY; i++ {
		for j := 0; j <= board.maxX; j++ {
			if board.plane[i][j] == searchedRune && !(p.x == j && p.y == i) {
				markResonanceInLine(board, antinodes, p, Point{j, i})
			}
		}
	}
}

func markResonanceInLine(board Board, antinodes map[Point]bool, p1 Point, p2 Point) {
	xDist := p1.x - p2.x
	yDist := p1.y - p2.y

	x := p1.x
	y := p1.y
	for x >= board.minX && x <= board.maxX && y >= board.minY && y <= board.maxY {
		antinodes[Point{x, y}] = true
		x += xDist
		y += yDist
	}

	x = p1.x
	y = p1.y
	for x >= board.minX && x <= board.maxX && y >= board.minY && y <= board.maxY {
		antinodes[Point{x, y}] = true
		x -= xDist
		y -= yDist
	}
}

func validateBoundaries(antinodes map[Point]bool, board Board) {
	for point := range antinodes {
		if point.x < board.minX || point.x > board.maxX || point.y < board.minY || point.y > board.maxY {
			antinodes[point] = false
		}
	}
}
