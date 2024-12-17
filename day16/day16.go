package day16

import (
	f "github.com/revzik/aoc_2024/common/files"
	g "github.com/revzik/aoc_2024/common/structures/graphs"
)

func RunTask() {
	path := "day16/input"
	lines := f.ReadLines(path)

}

type Node struct {
	X    int
	Y    int
	Side rune
}

// each intersection, corner or dead end
func createMaze(lines []string) *g.Graph {
	return nil
}
