package day16

import (
	"testing"

	"github.com/revzik/aoc_2024/day16/graphs"
	"github.com/stretchr/testify/assert"
)

func TestPartOne(t *testing.T) {
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
	maze, src, dst := createMaze(lines)

	end, err := graphs.ReindeerDijkstra(maze, src, dst)

	assert.Nil(t, err)
	assert.Equal(t, 7036, end.Distance)
}
