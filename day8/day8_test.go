package day8

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarkingAntinodes(t *testing.T) {
	board := createBoard([][]rune{
		{'.', 'B', 'B', '.', '.'},
		{'.', '.', '.', 'A', '.'},
		{'.', 'A', '.', '.', '.'},
		{'.', '.', '.', '.', '.'},
		{'.', '.', 'A', '.', '.'},
	})
	antinodes := make(map[Point]bool, 0)

	// Check first B from the left
	markWeakAntinodes(board, antinodes, Point{1, 0})
	assert.True(t, antinodes[Point{0, 0}])
	assert.True(t, antinodes[Point{3, 0}])

	// Check the first A from the top
	markWeakAntinodes(board, antinodes, Point{3, 1})
	assert.True(t, antinodes[Point{5, 0}])
	assert.True(t, antinodes[Point{-1, 3}])
	assert.True(t, antinodes[Point{4, -2}])
	assert.True(t, antinodes[Point{1, 7}])

	// Check the other A for overlap with B
	markWeakAntinodes(board, antinodes, Point{1, 2})
	assert.True(t, antinodes[Point{0, 0}])
	assert.True(t, antinodes[Point{3, 6}])

	// Make sure no other places were marked
	assert.Equal(t, 7, len(antinodes))
}

func TestCheckBoundaries(t *testing.T) {
	board := createBoard([][]rune{
		{'.', '.', '.'},
		{'.', '.', '.'},
		{'.', '.', '.'},
	})
	antinodes := map[Point]bool{
		{-1, 0}: true,
		{1, 1}:  true,
		{0, 2}:  true,
		{0, 4}:  true,
		{-4, 8}: true,
	}

	validateBoundaries(antinodes, board)

	assert.True(t, antinodes[Point{1, 1}])
	assert.True(t, antinodes[Point{0, 2}])

	assert.False(t, antinodes[Point{-1, 0}])
	assert.False(t, antinodes[Point{0, 4}])
	assert.False(t, antinodes[Point{-4, 8}])
}

func TestCountAntinodes(t *testing.T) {
	board := createBoard([][]rune{
		{'.', '#', '.', '#', '.', '#'},
		{'.', '.', '.', '.', '.', '.'},
		{'.', '#', 'B', 'B', '#', '.'},
		{'.', '.', '.', '.', 'A', '.'},
		{'.', '.', 'A', 'B', '.', '.'},
		{'#', '.', '.', '.', '.', '.'},
		{'.', '.', '.', 'A', '#', '.'}, // here A is also an antinode
	})

	assert.Equal(t, 8, countWeakAntinodes(board))
}

func TestCountResonantAntinodes(t *testing.T) {
	board := createBoard([][]rune{
		{'#', '#', '.', '#', '.', '#'},
		{'.', '.', '.', '.', '.', '.'},
		{'#', '#', 'B', 'B', '#', '#'},
		{'.', '.', '.', '.', 'A', '.'},
		{'.', '.', 'A', 'B', '.', '.'},
		{'#', '.', '.', '.', '.', '.'},
		{'.', '.', '.', 'A', '#', '.'},
	})

	assert.Equal(t, 16, countResonantAntinodes(board))
}
