package graphcreator

import (
	"testing"

	s "github.com/revzik/aoc_2024/common/structures"
	"github.com/stretchr/testify/assert"
)

func TestIntToNode(t *testing.T) {
	board := &s.Board{Plane: make([][]rune, 10)}
	for i := range board.Plane {
		board.Plane[i] = make([]rune, 10)
	}

	n1 := &Node{
		Vertex: s.Vector{X: 1, Y: 1},
		Side:   1,
	}
	n2 := &Node{
		Vertex: s.Vector{X: 5, Y: 7},
		Side:   3,
	}
	n3 := &Node{
		Vertex: s.Vector{X: 9, Y: 0},
		Side:   0,
	}

	assert.Equal(t, n1, IntToNode(45, board))
	assert.Equal(t, n2, IntToNode(303, board))
	assert.Equal(t, n3, IntToNode(36, board))

	assert.Equal(t, 45, n1.ToInt(board))
	assert.Equal(t, 303, n2.ToInt(board))
	assert.Equal(t, 36, n3.ToInt(board))
}

func TestCheckDeadEnd(t *testing.T) {
	maze := []string{
		"#######",
		"#.#...#",
		"#...#.#",
		"#.##..#",
		"#.#.#.#",
		"#.....#",
		"#######",
	}
	board := s.CreateBoard(maze)

	v1 := s.Vector{X: 1, Y: 1}
	v2 := s.Vector{X: 3, Y: 4}
	v3 := s.Vector{X: 4, Y: 3}
	v4 := s.Vector{X: 4, Y: 1}
	v5 := s.Vector{X: 1, Y: 5}

	v1Side := DOWN
	v2Side := DOWN
	v3Side := RIGHT

	assert.Equal(t, v1Side, checkDeadEnd(v1, board))
	assert.Equal(t, v2Side, checkDeadEnd(v2, board))
	assert.Equal(t, v3Side, checkDeadEnd(v3, board))

	assert.Equal(t, -1, checkDeadEnd(v4, board))
	assert.Equal(t, -1, checkDeadEnd(v5, board))
}

func TestCheckCorner(t *testing.T) {
	maze := []string{
		"#######",
		"#.#...#",
		"#...#.#",
		"#.##..#",
		"#.#.#.#",
		"#.....#",
		"#######",
	}
	board := s.CreateBoard(maze)

	v1 := s.Vector{X: 1, Y: 2}
	v2 := s.Vector{X: 5, Y: 1}
	v3 := s.Vector{X: 3, Y: 5}
	v4 := s.Vector{X: 1, Y: 3}
	v5 := s.Vector{X: 2, Y: 2}

	corners1 := []int{UP, RIGHT, DOWN}
	corners2 := []int{DOWN, LEFT}
	corners3 := []int{UP, LEFT, RIGHT}

	assert.ElementsMatch(t, corners1, checkCorner(v1, board))
	assert.ElementsMatch(t, corners2, checkCorner(v2, board))
	assert.ElementsMatch(t, corners3, checkCorner(v3, board))

	assert.Equal(t, 0, len(checkCorner(v4, board)))
	assert.Equal(t, 0, len(checkCorner(v5, board)))

}

func TestFindNodes(t *testing.T) {
	maze := []string{
		"#####",
		"#..##",
		"##.##",
		"##..#",
		"##.##",
		"#####",
	}
	board := s.CreateBoard(maze)

	expectedNodes := map[s.Vector][]int{
		{X: 1, Y: 1}: {RIGHT},
		{X: 2, Y: 1}: {LEFT, DOWN},
		{X: 2, Y: 3}: {RIGHT, DOWN, UP},
		{X: 3, Y: 3}: {LEFT},
		{X: 2, Y: 4}: {UP},
	}

	nodes := findNodes(board)

	assert.Equal(t, len(expectedNodes), len(nodes))
	for key := range nodes {
		assert.ElementsMatch(t, expectedNodes[key], nodes[key])
	}
}

func TestConnectNodes(t *testing.T) {
	maze := []string{
		"#####",
		"#..##",
		"##.##",
		"##..#",
		"##.##",
		"#####",
	}
	board := s.CreateBoard(maze)
	vectorMap := map[s.Vector][]int{
		{X: 1, Y: 1}: {RIGHT},
		{X: 2, Y: 1}: {LEFT, DOWN},
		{X: 2, Y: 3}: {RIGHT, DOWN, UP},
		{X: 3, Y: 3}: {LEFT},
		{X: 2, Y: 4}: {UP},
	}

	graph := connectNodes(vectorMap, board)

	assert.NotNil(t, graph)
}
