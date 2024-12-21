package graphs

import (
	"container/heap"
	"errors"
	"math"

	s "github.com/revzik/aoc_2024/common/structures"
)

// general graph representation
type Graph struct {
	Vertices map[s.Vector][]Edge
}

type Edge struct {
	Vertex s.Vector
	Weight int
}

func NewGraph() *Graph {
	return &Graph{Vertices: make(map[s.Vector][]Edge)}
}

func (g *Graph) GetEdges(n s.Vector) []Edge {
	val, ok := g.Vertices[n]
	if ok {
		return val
	}
	return []Edge{}
}

func (g *Graph) AddNode(n s.Vector) {
	g.Vertices[n] = make([]Edge, 0)
}

func (g *Graph) AddEdge(src, dst s.Vector, weight int) {
	e := Edge{Vertex: dst, Weight: weight}
	val, ok := g.Vertices[src]
	if ok {
		val = append(val, e)
		g.Vertices[src] = val
	} else {
		g.Vertices[src] = []Edge{e}
	}
}

// Dijkstra internal structure
type Node struct {
	Vertex   s.Vector
	Distance int
	Path     []s.Vector
}

// Minheap for nodes
type NodeHeap []*Node

func (h NodeHeap) Len() int {
	return len(h)
}

func (h NodeHeap) Less(i, j int) bool {
	return h[i].Distance < h[j].Distance
}

func (h NodeHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *NodeHeap) Push(x interface{}) {
	*h = append(*h, x.(*Node))
}

func (h *NodeHeap) Pop() interface{} {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

func NewHeap() *NodeHeap {
	var nodeHeap NodeHeap
	heap.Init(&nodeHeap)
	return &nodeHeap
}

// actual algorithm
func ReindeerDijkstra(g *Graph, src, dst s.Vector) (*Node, error) {

	if _, ok := g.Vertices[src]; !ok {
		panic("source node not found in graph")
	}
	if _, ok := g.Vertices[dst]; !ok {
		panic("destination node not found in graph")
	}

	paths := make(map[s.Vector]*Node)
	visited := make(map[s.Vector]bool)
	for vertex := range g.Vertices {
		node := Node{
			Vertex:   vertex,
			Distance: math.MaxInt32,
			Path:     make([]s.Vector, 0),
		}
		paths[vertex] = &node
		visited[vertex] = false
	}

	srcNode := Node{
		Vertex:   src,
		Distance: 0,
		Path:     make([]s.Vector, 0),
	}
	unvisitedNodes := NewHeap()
	unvisitedNodes.Push(&srcNode)

	for unvisitedNodes.Len() > 0 {
		node := unvisitedNodes.Pop().(*Node)

		if visited[node.Vertex] {
			continue
		}
		visited[node.Vertex] = true

		if node.Vertex == dst {
			return node, nil
		}

		for _, edge := range g.GetEdges(node.Vertex) {
			if visited[edge.Vertex] {
				continue
			}

			// TODO: add chech if previous node and next one are on the same line
			// if not, add 1000 to distance
			newDistance := node.Distance + edge.Weight

			if newDistance < paths[edge.Vertex].Distance {
				paths[edge.Vertex].Distance = newDistance
				paths[edge.Vertex].Path = append(paths[edge.Vertex].Path, node.Vertex)
			}

			unvisitedNodes.Push(paths[edge.Vertex])
		}
	}

	return nil, errors.New("path not found")
}
