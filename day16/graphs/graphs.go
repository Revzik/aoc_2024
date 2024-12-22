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

	size := len(old)
	x := old[size-1]
	*h = old[:size-1]

	return x
}

// actual algorithm
func ReindeerDijkstra(g *Graph, src, dst s.Vector) (*Node, error) {

	if _, ok := g.Vertices[src]; !ok {
		panic("source node not found in graph")
	}
	if _, ok := g.Vertices[dst]; !ok {
		panic("destination node not found in graph")
	}

	nodes := make(map[s.Vector]*Node)
	visited := make(map[s.Vector]bool)
	for vertex := range g.Vertices {
		node := Node{
			Vertex:   vertex,
			Distance: math.MaxInt32,
			Path:     make([]s.Vector, 0),
		}
		nodes[vertex] = &node
		visited[vertex] = false
	}

	srcNode := Node{
		Vertex:   src,
		Distance: 0,
		Path:     []s.Vector{src.Add(s.Vector{X: -1, Y: 0})},
	}
	unvisitedNodes := &NodeHeap{}
	heap.Push(unvisitedNodes, &srcNode)

	for unvisitedNodes.Len() > 0 {
		node := heap.Pop(unvisitedNodes).(*Node)

		if visited[node.Vertex] {
			continue
		}
		visited[node.Vertex] = true

		for _, edge := range g.GetEdges(node.Vertex) {
			if visited[edge.Vertex] {
				continue
			}

			newDistance := node.Distance + edge.Weight
			if !areInLine(edge.Vertex, node.Path[len(node.Path)-1]) {
				newDistance += 1000
			}

			// could this be changed to <= to find all shortest paths?
			if newDistance < nodes[edge.Vertex].Distance {
				nodes[edge.Vertex].Distance = newDistance
				nodes[edge.Vertex].Path = append(node.Path, node.Vertex)
			}

			heap.Push(unvisitedNodes, nodes[edge.Vertex])
		}
	}

	if end, ok := nodes[dst]; ok {
		return end, nil
	} else {
		return nil, errors.New("path not found")
	}
}

func areInLine(v1, v2 s.Vector) bool {
	if v1.X == v2.X {
		return true
	}
	if v1.Y == v2.Y {
		return true
	}
	return false
}
