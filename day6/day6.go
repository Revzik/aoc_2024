package day6

import (
	"errors"
	"fmt"

	"github.com/revzik/aoc_2024/common/files"
	"github.com/revzik/aoc_2024/common/types"
)

func RunTask() {
	path := "day6/input.txt"
	lines := files.ReadLines(path)

	board := types.CreateBoard(lines)
	originalPlane := types.CopyMatrix(board.Plane)

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

func (g Guard) Move(direction types.Point) Guard {
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
func resetBoard(board types.Board, originalPlane [][]rune) {
	for i := range originalPlane {
		copy(board.Plane[i], originalPlane[i])
	}
}

func isWithinBounds(board types.Board, g Guard) bool {
	if g.X >= board.MinX && g.X <= board.MaxX && g.Y >= board.MinY && g.Y <= board.MaxY {
		return true
	}
	return false
}

func getDirection(g Guard) types.Point {
	switch g.Dir {
	case '>':
		return types.Point{X: 1, Y: 0}
	case 'v':
		return types.Point{X: 0, Y: 1}
	case '<':
		return types.Point{X: -1, Y: 0}
	case '^':
		return types.Point{X: 0, Y: -1}
	default:
		panic("Guard character is not correct")
	}
}

func findGuard(board types.Board) Guard {
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
func followPath(board types.Board) int {
	guard := findGuard(board)
	followGuardPath(board, guard)
	return countVisitedPoints(board)
}

func countVisitedPoints(board types.Board) int {
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
func countPossibleLoops(board types.Board) int {
	total := 0

	guard := findGuard(board)
	originalPlane := types.CopyMatrix(board.Plane)

	for y, line := range board.Plane {
		for x := range line {
			resetBoard(board, originalPlane)
			if isObstructed(board, types.Point{X: x, Y: y}) {
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

func isObstructed(b types.Board, p types.Point) bool {
	return b.Plane[p.Y][p.X] == '#' || b.Plane[p.Y][p.X] == '^' || b.Plane[p.Y][p.X] == '>' || b.Plane[p.Y][p.X] == 'v' || b.Plane[p.Y][p.X] == '<'
}

func isInLoop(guard Guard, visitedTurningPoints map[Guard]bool) bool {
	if _, ok := visitedTurningPoints[guard]; ok {
		return true
	}
	return false
}

// Common part
func followGuardPath(board types.Board, guard Guard) error {
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
