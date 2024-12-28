package day16

import (
	"testing"

	"github.com/revzik/aoc_2024/common/structures"
	"github.com/revzik/aoc_2024/day16/graphcreator"
	"github.com/revzik/aoc_2024/day16/pathfinder"
	"github.com/stretchr/testify/assert"
)

func TestCreateGraph(t *testing.T) {
	maze := []string{
		"#####",
		"#.#E#",
		"#...#",
		"#.#.#",
		"#...#",
		"#S#.#",
		"#####",
	}
	board := structures.CreateBoard(maze)

	startVector, endVector := replaceStartEnd(board)
	graph := graphcreator.CreateGraph(board)
	startNode := createStartNode(startVector, graph, board)
	endNode := createEndNode(endVector, graph, board)

	debugPrint(graph, board)

	path, err := pathfinder.FindAllPaths(graph, startNode, endNode)

	assert.Nil(t, err)
	assert.Equal(t, 3006, path.Distance)
	assert.Equal(t, 10, countAllSpots(path, board))
}
