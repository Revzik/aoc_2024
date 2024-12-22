package day6

import (
	"errors"
	"fmt"

	f "github.com/revzik/aoc_2024/common/files"
	s "github.com/revzik/aoc_2024/common/structures"
)

func RunTask() {
	path := "day6/input"
	lines := f.ReadLines(path)

	board := s.CreateBoard(lines)
	originalPlane := s.CopyMatrix(board.Plane)

	fmt.Printf("Total squares walked: %d\n", followPath(board))

	resetBoard(board, originalPlane)
	fmt.Printf("Possible loops: %d\n", countPossibleLoops(board))
}

// Guard data structure
type Guard struct {
	Dir rune
	X   int
	Y   int
}

func (g Guard) Move(direction s.Vector) Guard {
	return Guard{Dir: g.Dir, X: g.X + direction.X, Y: g.Y + direction.Y}
}

func (g *Guard) Turn() {
	switch g.Dir {
	case '^':
		g.Dir = '>'
	case '>':
		g.Dir = 'v'
	case 'v':
		g.Dir = '<'
	case '<':
		g.Dir = '^'
	default:
		panic("Unknown guard character")
	}
}

// Board related
func resetBoard(board *s.Board, originalPlane [][]rune) {
	for i := range originalPlane {
		copy(board.Plane[i], originalPlane[i])
	}
}

func isWithinBounds(board *s.Board, g Guard) bool {
	if g.X >= board.MinX() && g.X <= board.MaxX() && g.Y >= board.MinY() && g.Y <= board.MaxY() {
		return true
	}
	return false
}

func getDirection(g Guard) s.Vector {
	switch g.Dir {
	case '>':
		return s.Vector{X: 1, Y: 0}
	case 'v':
		return s.Vector{X: 0, Y: 1}
	case '<':
		return s.Vector{X: -1, Y: 0}
	case '^':
		return s.Vector{X: 0, Y: -1}
	default:
		panic("Guard character is not correct")
	}
}

func findGuard(board *s.Board) Guard {
	for y, line := range board.Plane {
		for x, char := range line {
			if char == '^' || char == '>' || char == 'v' || char == '<' {
				return Guard{Dir: char, X: x, Y: y}
			}
		}
	}

	panic("Guard not found on the board!")
}

// First part
func followPath(board *s.Board) int {
	guard := findGuard(board)
	followGuardPath(board, guard)
	return countVisitedPoints(board)
}

func countVisitedPoints(board *s.Board) int {
	total := 0

	for _, line := range board.Plane {
		for _, point := range line {
			if point == 'X' {
				total++
			}
		}
	}

	return total
}

// Second part
func countPossibleLoops(board *s.Board) int {
	total := 0

	guard := findGuard(board)
	originalPlane := s.CopyMatrix(board.Plane)

	for y, line := range board.Plane {
		for x := range line {
			resetBoard(board, originalPlane)
			if isObstructed(board, s.Vector{X: x, Y: y}) {
				continue
			}

			board.Plane[y][x] = 'O'

			err := followGuardPath(board, guard)

			if err != nil {
				total++
			}
		}
	}

	return total
}

func isObstructed(b *s.Board, p s.Vector) bool {
	return b.Plane[p.Y][p.X] == '#' || b.Plane[p.Y][p.X] == '^' || b.Plane[p.Y][p.X] == '>' || b.Plane[p.Y][p.X] == 'v' || b.Plane[p.Y][p.X] == '<'
}

func isInLoop(guard Guard, visitedTurningPoints map[Guard]bool) bool {
	if _, ok := visitedTurningPoints[guard]; ok {
		return true
	}
	return false
}

// Common part
func followGuardPath(board *s.Board, guard Guard) error {
	direction := getDirection(guard)
	nextPosition := guard.Move(direction)

	visitedTurningPoints := make(map[Guard]bool)

	for isWithinBounds(board, nextPosition) {
		for board.Plane[nextPosition.Y][nextPosition.X] == '#' || board.Plane[nextPosition.Y][nextPosition.X] == 'O' {
			if isInLoop(guard, visitedTurningPoints) {
				return errors.New("loop detected")
			}

			visitedTurningPoints[guard] = true
			guard.Turn()
			direction = getDirection(guard)
			nextPosition = guard.Move(direction)
		}

		board.Plane[guard.Y][guard.X] = 'X'
		board.Plane[nextPosition.Y][nextPosition.X] = guard.Dir
		guard = nextPosition

		direction = getDirection(guard)
		nextPosition = guard.Move(direction)
	}

	board.Plane[guard.Y][guard.X] = 'X'

	return nil
}
