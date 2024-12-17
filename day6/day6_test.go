package day6

import (
	"testing"

	s "github.com/revzik/aoc_2024/common/structures"
	"github.com/stretchr/testify/assert"
)

func TestFindGuard(t *testing.T) {
	board1 := s.CreateBoard([]string{
		".....",
		".^...",
		".....",
		".....",
		".....",
	})
	board2 := s.CreateBoard([]string{
		".....",
		".....",
		".....",
		".....",
		"...v.",
	})
	board3 := s.CreateBoard([]string{
		".......",
		".......",
		"..>....",
		".......",
	})

	assert.Equal(t, s.Vector{X: 1, Y: 1}, findGuard(board1))
	assert.Equal(t, s.Vector{X: 3, Y: 4}, findGuard(board2))
	assert.Equal(t, s.Vector{X: 2, Y: 2}, findGuard(board3))
}

func TestCountVisited(t *testing.T) {
	board1 := s.CreateBoard([]string{
		".#X..",
		".^XX#",
		"..XX.",
		"##XX.",
		"..##.",
	})
	board2 := s.CreateBoard([]string{
		".X...",
		".XXX.",
		".X.X.",
		"XXXX.",
		".X...",
	})

	assert.Equal(t, 7, countVisitedPoints(board1))
	assert.Equal(t, 11, countVisitedPoints(board2))
}

func TestGuardWithinBounds(t *testing.T) {
	board := s.CreateBoard([]string{
		"......",
		"......",
		"......",
		"......",
	})

	assert.True(t, isWithinBounds(board, Guard{Dir: '^', X: 0, Y: 0}))
	assert.True(t, isWithinBounds(board, Guard{Dir: '^', X: 4, Y: 2}))
	assert.False(t, isWithinBounds(board, Guard{Dir: '^', X: -1, Y: 3}))
	assert.False(t, isWithinBounds(board, Guard{Dir: '^', X: 0, Y: 4}))
	assert.False(t, isWithinBounds(board, Guard{Dir: '^', X: 6, Y: 0}))
	assert.False(t, isWithinBounds(board, Guard{Dir: '^', X: 5, Y: -1}))
}

func TestGuardPath(t *testing.T) {
	board := s.CreateBoard([]string{
		".#...",
		"....#",
		".^...",
		"##...",
		"..##.",
	})
	expectedResult := s.CreateBoard([]string{
		".#X..",
		".XXX#",
		".XXX.",
		"##XX.",
		"..##.",
	})

	followGuardPath(board, Guard{Dir: '^', X: 1, Y: 2})

	assert.Equal(t, expectedResult, board)
}

func TestLoopDetection(t *testing.T) {
	board1 := s.CreateBoard([]string{
		".#...",
		"....#",
		".^...",
		"#....",
		"...#.",
	})
	board2 := s.CreateBoard([]string{
		"...#.....",
		".......#.",
		".#.......",
		".....#...",
		"#..^.....",
		"......#..",
		"..#......",
		"....#....",
	})
	board3 := s.CreateBoard([]string{
		"...#.....",
		".......#.",
		".#.......",
		".....#...",
		"#..^.....",
		"......#..",
		".........",
		"....#....",
	})

	err1 := followGuardPath(board1, Guard{Dir: '^', X: 1, Y: 2})
	err2 := followGuardPath(board2, Guard{Dir: '^', X: 3, Y: 4})
	err3 := followGuardPath(board3, Guard{Dir: '^', X: 3, Y: 4})

	assert.NotNil(t, err1)
	assert.NotNil(t, err2)
	assert.Nil(t, err3)
}

func TestCountLoops(t *testing.T) {
	board := s.CreateBoard([]string{
		"....#.....",
		".........#",
		"..........",
		"..#.......",
		".......#..",
		"..........",
		".#..^.....",
		"........#.",
		"#.........",
		"......#...",
	})

	loopCount := countPossibleLoops(board)

	assert.Equal(t, 6, loopCount)
}
