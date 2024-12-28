package graphcreator

import (
	"github.com/revzik/aoc_2024/common/intmath"
	s "github.com/revzik/aoc_2024/common/structures"
	"github.com/revzik/aoc_2024/common/structures/graphs"
)

const RIGHT = 0
const DOWN = 1
const LEFT = 2
const UP = 3

type Node struct {
	Vertex s.Vector
	Side   int
}

func (n *Node) ToInt(b *s.Board) int {
	return (n.Vertex.X+n.Vertex.Y*b.Width())*4 + n.Side
}

func IntToNode(n int, b *s.Board) *Node {
	side := n % 4
	n = (n - side) / 4

	x := n % b.Width()
	y := n / b.Width()

	return &Node{
		Vertex: s.Vector{X: x, Y: y},
		Side:   side,
	}
}

func CreateGraph(board *s.Board) *graphs.Graph[int] {
	allNodes := findNodes(board)
	return connectNodes(allNodes, board)
}

func checkCorner(position s.Vector, board *s.Board) []int {
	dirVectors := []s.Vector{
		{X: 1, Y: 0},
		{X: 0, Y: 1},
		{X: -1, Y: 0},
		{X: 0, Y: -1},
	}
	dirs := []int{RIGHT, DOWN, LEFT, UP}

	cornerMap := make(map[int]bool, 0)

	for i := 0; i < 4; i++ {
		if board.Get(position.Add(dirVectors[0])) == '.' &&
			board.Get(position.Add(dirVectors[1])) == '.' {

			cornerMap[dirs[0]] = true
			cornerMap[dirs[1]] = true
		}

		dirVectors = append(dirVectors, dirVectors[0])
		dirVectors = dirVectors[1:]
		dirs = append(dirs, dirs[0])
		dirs = dirs[1:]
	}

	corners := make([]int, 0)
	for side := range cornerMap {
		corners = append(corners, side)
	}

	return corners
}

func checkDeadEnd(position s.Vector, board *s.Board) int {
	dirVectors := []s.Vector{
		{X: 1, Y: 0},
		{X: 0, Y: 1},
		{X: -1, Y: 0},
		{X: 0, Y: -1},
	}
	dirs := []int{RIGHT, DOWN, LEFT, UP}

	for _, dir := range dirs {
		if board.Get(position.Add(dirVectors[0])) == '.' &&
			board.Get(position.Add(dirVectors[1])) == '#' &&
			board.Get(position.Add(dirVectors[2])) == '#' &&
			board.Get(position.Add(dirVectors[3])) == '#' {

			return dir
		}

		dirVectors = append(dirVectors, dirVectors[0])
		dirVectors = dirVectors[1:]
	}

	return -1
}

func findNodes(board *s.Board) map[s.Vector][]int {
	nodes := make(map[s.Vector][]int)

	for y, row := range board.Plane {
		for x, item := range row {
			if item != '.' {
				continue
			}

			position := s.Vector{X: x, Y: y}

			endSide := checkDeadEnd(position, board)
			if endSide >= 0 {
				nodes[position] = []int{endSide}
			}

			nodeSides := checkCorner(position, board)
			if len(nodeSides) > 0 {
				nodes[position] = nodeSides
			}
		}
	}

	return nodes
}

func connectNodes(vectors map[s.Vector][]int, board *s.Board) *graphs.Graph[int] {
	graph := graphs.NewGraph[int]()
	dirVectors := []s.Vector{
		{X: 1, Y: 0},
		{X: 0, Y: 1},
		{X: -1, Y: 0},
		{X: 0, Y: -1},
	}

	for pos, sides := range vectors {
		nodes := make([]*Node, len(sides))

		for i, side := range sides {
			nodes[i] = &Node{
				Vertex: pos,
				Side:   side,
			}
		}

		// connection in a corner or an intersection
		if len(nodes) > 1 {
			for i, srcNode := range nodes {
				for j, dstNode := range nodes {
					if i == j {
						continue
					}

					weight := 1000
					if intmath.Abs(srcNode.Side-dstNode.Side) == 2 {
						weight = 0
					}

					graph.AddEdge(srcNode.ToInt(board), dstNode.ToInt(board), weight)
				}
			}
		}

		// connections in line
		for _, srcNode := range nodes {
			dir := dirVectors[srcNode.Side]

			nextVector := pos.Add(dir)
			dstSide := srcNode.Side + 2
			if dstSide > 3 {
				dstSide -= 4
			}

			prevNode := srcNode.ToInt(board)

			for board.WithinBounds(nextVector) {
				next := &Node{
					Vertex: nextVector,
					Side:   dstSide,
				}
				nextNode := next.ToInt(board)

				graph.AddEdge(prevNode, nextNode, 1)

				if _, ok := vectors[nextVector]; ok {
					break
				}

				prevNode = nextNode
				nextVector = nextVector.Add(dir)
			}
		}
	}

	return graph
}
