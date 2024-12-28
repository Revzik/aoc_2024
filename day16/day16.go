package day16

import (
	"fmt"

	f "github.com/revzik/aoc_2024/common/files"
	s "github.com/revzik/aoc_2024/common/structures"
	g "github.com/revzik/aoc_2024/common/structures/graphs"
	gc "github.com/revzik/aoc_2024/day16/graphcreator"
	pf "github.com/revzik/aoc_2024/day16/pathfinder"
)

func RunTask() {
	inputFile := "day16/input"
	lines := f.ReadLines(inputFile)
	board := s.CreateBoard(lines)

	startVector, endVector := replaceStartEnd(board)
	graph := gc.CreateGraph(board)
	startNode := createStartNode(startVector, graph, board)
	endNode := createEndNode(endVector, graph, board)

	path, _ := pf.FindAllPaths(graph, startNode, endNode)

	fmt.Printf("Path lenght: %d\n", path.Distance)
	fmt.Printf("Best spots: %d\n", countAllSpots(path, board))
}

func replaceStartEnd(board *s.Board) (s.Vector, s.Vector) {
	var start s.Vector
	var end s.Vector

	for y, row := range board.Plane {
		for x, item := range row {
			if item == 'S' {
				start = s.Vector{X: x, Y: y}
				board.Set(start, '.')
			} else if item == 'E' {
				end = s.Vector{X: x, Y: y}
				board.Set(end, '.')
			}
		}
	}

	return start, end
}

func createStartNode(v s.Vector, graph *g.Graph[int], board *s.Board) int {
	start := &gc.Node{
		Vertex: v,
		Side:   gc.RIGHT,
	}
	startCode := start.ToInt(board)

	// if it was already created, there is nothing to do
	if _, ok := graph.Nodes[startCode]; ok {
		return startCode
	}

	for nCode := range graph.Nodes {
		n := gc.IntToNode(nCode, board)
		if v != n.Vertex {
			continue
		}

		weight := 0
		if n.Side != gc.RIGHT {
			weight = 1000
		}

		graph.AddEdge(startCode, nCode, weight)
	}

	return startCode
}

func createEndNode(v s.Vector, graph *g.Graph[int], board *s.Board) int {
	end := &gc.Node{
		Vertex: v,
		Side:   0, // doesn't matter
	}
	endCode := end.ToInt(board)

	for nCode := range graph.Nodes {
		n := gc.IntToNode(nCode, board)
		if v == n.Vertex {
			graph.AddEdge(nCode, endCode, 0)
		}
	}

	graph.AddNode(endCode)

	return endCode
}

func countAllSpots(path *pf.Path, board *s.Board) int {
	toProcess := []*pf.Path{path}
	vertices := make(map[s.Vector]bool)
	visited := make(map[int]bool)

	for len(toProcess) > 0 {
		p := toProcess[0]
		toProcess = toProcess[1:]

		n := gc.IntToNode(p.Node, board)
		vertices[n.Vertex] = true
		visited[p.Node] = true

		// there is a loop in the path - might be worth inspecting
		for _, prev := range p.Previous {
			if _, ok := visited[prev.Node]; !ok {
				toProcess = append(toProcess, prev)
			}
		}
	}

	return len(vertices)
}
