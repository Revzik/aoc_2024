package day8

import (
	"fmt"
	"unicode"

	f "github.com/revzik/aoc_2024/common/files"
	s "github.com/revzik/aoc_2024/common/structures"
)

func RunTask() {
	path := "day8/input"
	matrix := f.ReadLines(path)
	board := s.CreateBoard(matrix)

	fmt.Printf("Number of weak antinodes: %d\n", countWeakAntinodes(board))
	fmt.Printf("Number of resonant antinodes: %d\n", countResonantAntinodes(board))
}

func countWeakAntinodes(board *s.Board) int {
	total := 0
	antinodes := make(map[s.Vector]bool)

	for y := 0; y < board.MaxY(); y++ {
		for x := 0; x < board.MaxX(); x++ {
			if unicode.IsLetter(board.Plane[y][x]) || unicode.IsDigit(board.Plane[y][x]) {
				markWeakAntinodes(board, antinodes, s.Vector{X: x, Y: y})
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

func markWeakAntinodes(board *s.Board, antinodes map[s.Vector]bool, p s.Vector) {
	searchedRune := board.Plane[p.Y][p.X]

	// check only to the bottom, top ones are dealt with
	for i := p.Y; i <= board.MaxY(); i++ {
		for j := 0; j <= board.MaxX(); j++ {
			if board.Plane[i][j] == searchedRune && !(p.X == j && p.Y == i) {
				xDist := p.X - j
				yDist := p.Y - i

				point1 := s.Vector{X: p.X + xDist, Y: p.Y + yDist}
				point2 := s.Vector{X: j - xDist, Y: i - yDist}

				antinodes[point1] = true
				antinodes[point2] = true
			}
		}
	}
}

func countResonantAntinodes(board *s.Board) int {
	antinodes := make(map[s.Vector]bool)

	for y := 0; y <= board.MaxY(); y++ {
		for x := 0; x <= board.MaxX(); x++ {
			if unicode.IsLetter(board.Plane[y][x]) || unicode.IsDigit(board.Plane[y][x]) {
				markResonantAntinodes(board, antinodes, s.Vector{X: x, Y: y})
			}
		}
	}

	return len(antinodes)
}

func markResonantAntinodes(board *s.Board, antinodes map[s.Vector]bool, p s.Vector) {
	searchedRune := board.Plane[p.Y][p.X]

	// check only to the bottom, top ones are dealt with
	for i := p.Y; i <= board.MaxY(); i++ {
		for j := 0; j <= board.MaxX(); j++ {
			if board.Plane[i][j] == searchedRune && !(p.X == j && p.Y == i) {
				markResonanceInLine(board, antinodes, p, s.Vector{X: j, Y: i})
			}
		}
	}
}

func markResonanceInLine(board *s.Board, antinodes map[s.Vector]bool, p1 s.Vector, p2 s.Vector) {
	xDist := p1.X - p2.X
	yDist := p1.Y - p2.Y

	x := p1.X
	y := p1.Y
	for x >= board.MinX() && x <= board.MaxX() && y >= board.MinY() && y <= board.MaxY() {
		antinodes[s.Vector{X: x, Y: y}] = true
		x += xDist
		y += yDist
	}

	x = p1.X
	y = p1.Y
	for x >= board.MinX() && x <= board.MaxX() && y >= board.MinY() && y <= board.MaxY() {
		antinodes[s.Vector{X: x, Y: y}] = true
		x -= xDist
		y -= yDist
	}
}

func validateBoundaries(antinodes map[s.Vector]bool, board *s.Board) {
	for point := range antinodes {
		if point.X < board.MinX() || point.X > board.MaxX() || point.Y < board.MinY() || point.Y > board.MaxY() {
			antinodes[point] = false
		}
	}
}
