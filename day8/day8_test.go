package day8

import (
	"testing"

	"github.com/revzik/aoc_2024/common/types"
	"github.com/stretchr/testify/assert"
)

func TestMarkingAntinodes(t *testing.T) {
	board := types.CreateBoard([]string{
		".BB..",
		"...A.",
		".A...",
		".....",
		"..A..",
	})
	antinodes := make(map[types.Point]bool, 0)

	// Check first B from the left
	markWeakAntinodes(board, antinodes, types.Point{X: 1, Y: 0})
	assert.True(t, antinodes[types.Point{X: 0, Y: 0}])
	assert.True(t, antinodes[types.Point{X: 3, Y: 0}])

	// Check the first A from the top
	markWeakAntinodes(board, antinodes, types.Point{X: 3, Y: 1})
	assert.True(t, antinodes[types.Point{X: 5, Y: 0}])
	assert.True(t, antinodes[types.Point{X: -1, Y: 3}])
	assert.True(t, antinodes[types.Point{X: 4, Y: -2}])
	assert.True(t, antinodes[types.Point{X: 1, Y: 7}])

	// Check the other A for overlap with B
	markWeakAntinodes(board, antinodes, types.Point{X: 1, Y: 2})
	assert.True(t, antinodes[types.Point{X: 0, Y: 0}])
	assert.True(t, antinodes[types.Point{X: 3, Y: 6}])

	// Make sure no other places were marked
	assert.Equal(t, 7, len(antinodes))
}

func TestCheckBoundaries(t *testing.T) {
	board := types.CreateBoard([]string{
		"...",
		"...",
		"...",
	})
	antinodes := map[types.Point]bool{
		{X: -1, Y: 0}: true,
		{X: 1, Y: 1}:  true,
		{X: 0, Y: 2}:  true,
		{X: 0, Y: 4}:  true,
		{X: -4, Y: 8}: true,
	}

	validateBoundaries(antinodes, board)

	assert.True(t, antinodes[types.Point{X: 1, Y: 1}])
	assert.True(t, antinodes[types.Point{X: 0, Y: 2}])

	assert.False(t, antinodes[types.Point{X: -1, Y: 0}])
	assert.False(t, antinodes[types.Point{X: 0, Y: 4}])
	assert.False(t, antinodes[types.Point{X: -4, Y: 8}])
}

func TestCountAntinodes(t *testing.T) {
	board := types.CreateBoard([]string{
		".#.#.#",
		"......",
		".#BB#.",
		"....A.",
		"..AB..",
		"#.....",
		"...A#.", // here A is also an antinode
	})

	assert.Equal(t, 8, countWeakAntinodes(board))
}

func TestCountResonantAntinodes(t *testing.T) {
	board := types.CreateBoard([]string{
		"##.#.#",
		"......",
		"##BB##",
		"....A.",
		"..AB..",
		"#.....",
		"...A#.",
	})

	assert.Equal(t, 16, countResonantAntinodes(board))
}
