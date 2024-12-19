package day18

import (
	"fmt"
	"strings"

	f "github.com/revzik/aoc_2024/common/files"
	p "github.com/revzik/aoc_2024/common/parsers"
	g "github.com/revzik/aoc_2024/common/structures/graphs"
)

const SIZE = 70

func RunTask() {
	path := "day18/input"
	lines := f.ReadLines(path)

	start := Vertex{X: 0, Y: 0}
	end := Vertex{X: SIZE, Y: SIZE}

	fmt.Printf("Shortest path after 1024 bytes blocked: %d\n", shortestPath1024(lines, start, end))
	blockingByteIndex := findWhenExitBlocked(lines, start, end)
	fmt.Printf("Exit blocked by byte with coordinates: %s\n", lines[blockingByteIndex])
}

func shortestPath1024(lines []string, start Vertex, end Vertex) int {
	graph := createGraph(lines, 1024)
	path, err := g.Dijkstra(graph, start, end)
	if err != nil {
		panic("path not found for 1024 bytes")
	}
	return path.Distance
}

func findWhenExitBlocked(lines []string, start, end Vertex) int {
	low := 0
	high := len(lines)
	current := (high + low) / 2

	for high-low > 1 {
		graph := createGraph(lines, current)
		_, err := g.Dijkstra(graph, start, end)

		if err == nil {
			low = current
		} else {
			high = current
		}

		current = (high + low) / 2
	}

	return current
}

func createGraph(lines []string, bytes int) *g.Graph[Vertex] {

	blockedFields := make(map[Vertex]bool, len(lines))
	for i := 0; i < bytes; i++ {
		splitLine := strings.Split(lines[i], ",")
		v := Vertex{
			X: p.StringToInt(splitLine[0]),
			Y: p.StringToInt(splitLine[1]),
		}
		blockedFields[v] = true
	}

	graph := g.NewGraph[Vertex]()

	directions := []Vertex{
		{X: 1, Y: 0},
		{X: 0, Y: 1},
		{X: -1, Y: 0},
		{X: 0, Y: -1},
	}
	for x := 0; x <= SIZE; x++ {
		for y := 0; y <= SIZE; y++ {

			src := Vertex{X: x, Y: y}
			if _, ok := blockedFields[src]; ok {
				continue
			}

			for _, dir := range directions {
				dst := src.Add(dir)
				if fieldAvailable(dst, blockedFields, SIZE) {
					graph.AddEdge(src, dst, 1)
				}
			}
		}
	}

	return graph
}

func fieldAvailable(v Vertex, blockedFields map[Vertex]bool, size int) bool {
	if v.X >= 0 && v.X <= size && v.Y >= 0 && v.Y <= size {
		return !blockedFields[v]
	}
	return false
}

type Vertex struct {
	X int
	Y int
}

func (v1 Vertex) Add(v2 Vertex) Vertex {
	return Vertex{X: v1.X + v2.X, Y: v1.Y + v2.Y}
}
