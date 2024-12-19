package graphs

import (
	"errors"
	"math"
	"slices"
)

// general graph representation
type Graph[N comparable] struct {
	Nodes map[N][]Edge[N]
}

type Edge[N comparable] struct {
	Node   N
	Weight int
}

func NewGraph[N comparable]() *Graph[N] {
	return &Graph[N]{Nodes: make(map[N][]Edge[N])}
}

func (g *Graph[N]) GetEdges(n N) []Edge[N] {
	val, ok := g.Nodes[n]
	if ok {
		return val
	}
	return []Edge[N]{}
}

func (g *Graph[N]) AddNode(n N) {
	g.Nodes[n] = make([]Edge[N], 0)
}

func (g *Graph[N]) AddEdge(src, dst N, weight int) {
	e := Edge[N]{Node: dst, Weight: weight}
	val, ok := g.Nodes[src]
	if ok {
		val = append(val, e)
		g.Nodes[src] = val
	} else {
		g.Nodes[src] = []Edge[N]{e}
	}
}

// dijkstra's algorithm implementation
func Dijkstra[N comparable](g *Graph[N], src, dst N) (*Path[N], error) {

	if _, ok := g.Nodes[src]; !ok {
		panic("source node not found in graph")
	}
	if _, ok := g.Nodes[dst]; !ok {
		panic("destination node not found in graph")
	}

	paths := make(map[N]*Path[N], len(g.Nodes))
	visited := make(map[N]bool, len(g.Nodes))
	for node := range g.Nodes {
		paths[node] = &Path[N]{Distance: math.MaxInt32, Nodes: make([]N, 0)}
		visited[node] = false
	}
	paths[src].Distance = 0

	// could replace with minHeap or priority queue
	unvisitedNodes := []N{src}

	for len(unvisitedNodes) > 0 {
		// yes, this really needs a replacement
		slices.SortFunc(unvisitedNodes, func(a, b N) int {
			return paths[a].Distance - paths[b].Distance
		})

		node := unvisitedNodes[0]
		unvisitedNodes = unvisitedNodes[1:]

		if visited[node] {
			continue
		}
		visited[node] = true

		if node == dst {
			return paths[dst], nil
		}

		for _, edge := range g.GetEdges(node) {
			if visited[edge.Node] {
				continue
			}

			newDistance := paths[node].Distance + edge.Weight

			if newDistance < paths[edge.Node].Distance {
				paths[edge.Node].Distance = newDistance
				paths[edge.Node].Nodes = append(paths[edge.Node].Nodes, node)
			}

			unvisitedNodes = append(unvisitedNodes, edge.Node)
		}
	}

	return nil, errors.New("path not found")
}

type Path[N comparable] struct {
	Distance int
	Nodes    []N
}
