package day16

import (
	"testing"

	"github.com/revzik/aoc_2024/common/structures"
	"github.com/revzik/aoc_2024/day16/graphs"
	"github.com/stretchr/testify/assert"
)

func TestMaze(t *testing.T) {
	lines := []string{
		"###############",
		"#.......#....E#",
		"#.#.###.#.###.#",
		"#.....#.#...#.#",
		"#.###.#####.#.#",
		"#.#.#.......#.#",
		"#.#.#####.###.#",
		"#...........#.#",
		"###.#.#####.#.#",
		"#...#.....#.#.#",
		"#.#.#.###.#.#.#",
		"#.....#...#.#.#",
		"#.###.#.#.#.#.#",
		"#S..#.....#...#",
		"###############",
	}
	board := structures.CreateBoard(lines)
	maze, src, dst := createMaze(board)

	end, err := graphs.ReindeerDijkstra(maze, src, dst, board)

	assert.Nil(t, err)
	assert.Equal(t, 7036, end.Distance)
	assert.Equal(t, 45, len(end.Path))
}
