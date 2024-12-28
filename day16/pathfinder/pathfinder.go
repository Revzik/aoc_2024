package pathfinder

import (
	"container/heap"
	"errors"
	"math"

	g "github.com/revzik/aoc_2024/common/structures/graphs"
)

type Path struct {
	Node     int
	Distance int
	Previous []*Path
}

type PathHeap []*Path

func (h PathHeap) Len() int {
	return len(h)
}

func (h PathHeap) Less(i, j int) bool {
	return h[i].Distance < h[j].Distance
}

func (h PathHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *PathHeap) Push(x interface{}) {
	*h = append(*h, x.(*Path))
}

func (h *PathHeap) Pop() interface{} {
	old := *h

	size := len(old)
	x := old[size-1]
	*h = old[:size-1]

	return x
}

// slightly modified dijkstra
func FindAllPaths(graph *g.Graph[int], src, dst int) (*Path, error) {

	if _, ok := graph.Nodes[src]; !ok {
		panic("source node not found in graph")
	}
	if _, ok := graph.Nodes[dst]; !ok {
		panic("destination node not found in graph")
	}

	paths := make(map[int]*Path, len(graph.Nodes))
	visited := make(map[int]bool, len(graph.Nodes))
	for node := range graph.Nodes {
		paths[node] = &Path{
			Node:     node,
			Distance: math.MaxInt32,
			Previous: make([]*Path, 0)}
		visited[node] = false
	}
	paths[src].Distance = 0

	unvisitedNodes := &PathHeap{}
	heap.Push(unvisitedNodes, paths[src])

	for unvisitedNodes.Len() > 0 {
		current := heap.Pop(unvisitedNodes).(*Path)

		if visited[current.Node] {
			continue
		}
		visited[current.Node] = true

		for _, edge := range graph.GetEdges(current.Node) {
			newDistance := current.Distance + edge.Weight

			if newDistance <= paths[edge.Node].Distance {
				paths[edge.Node].Distance = newDistance
				paths[edge.Node].Previous = append(paths[edge.Node].Previous, current)
			}

			heap.Push(unvisitedNodes, paths[edge.Node])
		}
	}

	if val, ok := paths[dst]; ok {
		return val, nil
	} else {
		return nil, errors.New("path not found")
	}
}
