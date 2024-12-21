package day16

import (
	"fmt"

	f "github.com/revzik/aoc_2024/common/files"
	s "github.com/revzik/aoc_2024/common/structures"
	g "github.com/revzik/aoc_2024/day16/graphs"
)

func RunTask() {
	path := "day16/input"
	lines := f.ReadLines(path)

	maze, src, dst := createMaze(lines)
	end, err := g.ReindeerDijkstra(maze, src, dst)
	if err != nil {
		panic("no possible path")
	}

	fmt.Printf("Shortest path through maze: %d\n", end.Distance)
}

// each intersection, corner or dead end
func createMaze(lines []string) (*g.Graph, s.Vector, s.Vector) {
	board := s.CreateBoard(lines)
	maze := g.NewGraph()
	var start s.Vector
	var end s.Vector

	for y, row := range board.Plane {
		for x, item := range row {
			if item == '.' || item == 'S' || item == 'E' {
				addVertex(maze, board, x, y)
			}
			if item == 'S' {
				start = s.Vector{X: x, Y: y}
			}
			if item == 'E' {
				end = s.Vector{X: x, Y: y}
			}
		}
	}

	return maze, start, end
}

func addVertex(maze *g.Graph, board s.Board, x, y int) {
	directions := []s.Vector{
		{X: 1, Y: 0},
		{X: 0, Y: 1},
		{X: -1, Y: 0},
		{X: 0, Y: -1},
	}
	src := s.Vector{X: x, Y: y}
	maze.AddNode(src)

	for _, direction := range directions {
		dst := src.Add(direction)
		if board.Get(dst) == '.' || board.Get(dst) == 'S' || board.Get(dst) == 'E' {
			maze.AddEdge(src, dst, 1)
		}
	}
}
