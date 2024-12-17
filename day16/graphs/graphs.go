package graphs

import (
	"errors"
	"math"
	"slices"
)

// general graph representation
type Graph struct {
	Nodes map[Node][]Edge
}

type Edge struct {
	Node   Node
	Weight int
}

type Node struct {
	X int
	Y int
}

type Path struct {
	Distance int
	Nodes    []Node
}

func NewGraph() *Graph {
	return &Graph{Nodes: make(map[Node][]Edge)}
}

func (g *Graph) GetEdges(n Node) []Edge {
	val, ok := g.Nodes[n]
	if ok {
		return val
	}
	return []Edge{}
}

func (g *Graph) AddNode(n Node) {
	g.Nodes[n] = make([]Edge, 0)
}

func (g *Graph) AddEdge(src, dst Node, weight int) {
	e := Edge{Node: dst, Weight: weight}
	val, ok := g.Nodes[src]
	if ok {
		val = append(val, e)
		g.Nodes[src] = val
	} else {
		g.Nodes[src] = []Edge{e}
	}
}

func ReindeerDijkstra(g *Graph, src, dst Node) (*Path, error) {

	if _, ok := g.Nodes[src]; !ok {
		panic("source node not found in graph")
	}
	if _, ok := g.Nodes[dst]; !ok {
		panic("destination node not found in graph")
	}

	paths := make(map[Node]*Path)
	visited := make(map[Node]bool)
	for node := range g.Nodes {
		paths[node] = &Path{Distance: math.MaxInt32, Nodes: make([]Node, 0)}
		visited[node] = false
	}

	// could replace with minHeap or priority queue
	unvisitedNodes := []Node{src}

	for len(unvisitedNodes) > 0 {
		// yes, this really needs a replacement
		slices.SortFunc(unvisitedNodes, func(a, b Node) int {
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
