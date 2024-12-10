package day10

import (
	"fmt"

	"github.com/revzik/aoc_2024/common/files"
	"github.com/revzik/aoc_2024/common/parsers"
	"github.com/revzik/aoc_2024/common/types"
)

func RunTask() {
	path := "day10/input"
	lines := files.ReadLines(path)

	matrix := types.CreateMatrix(lines)
	intMatrix := parsers.RuneToIntMatrix(matrix)
	topographyMap := Map{Plane: intMatrix}

	fmt.Printf("Trailheads sum: %d\n", countTrails(topographyMap, false))
	fmt.Printf("Trailheads rating sum: %d\n", countTrails(topographyMap, true))
}

type Map struct {
	Plane [][]int
}

func (m Map) Get(p types.Point) int {
	return m.Plane[p.Y][p.X]
}

func (m Map) MinX() int {
	return 0
}

func (m Map) MaxX() int {
	return len(m.Plane[0]) - 1
}

func (m Map) MinY() int {
	return 0
}

func (m Map) MaxY() int {
	return len(m.Plane) - 1
}

func countTrails(topographyMap Map, countRating bool) int {
	total := 0
	for i, row := range topographyMap.Plane {
		for j, height := range row {
			if height == 0 {
				total += countTrailsForPoint(types.Point{X: j, Y: i}, topographyMap, countRating)
			}
		}
	}
	return total
}

func countTrailsForPoint(point types.Point, topographyMap Map, countRating bool) int {
	// BFS inspired algorithm
	// neighbour is only valid if value is higher by one
	queue := []types.Point{point}
	visited := map[types.Point]bool{point: true}
	totalCount := 0

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		visited[current] = true

		height := topographyMap.Get(current)
		neighbours := getNeighbours(current, topographyMap)

		for _, neighbour := range neighbours {

			neighbourHeight := topographyMap.Get(neighbour)
			_, wasVisited := visited[neighbour]

			if neighbourHeight != height+1 || wasVisited {
				continue
			}

			if neighbourHeight == 9 {
				// don't know why this works, made it by mistake
				// task failed succesfully I guess
				if !countRating {
					visited[neighbour] = true
				}
				totalCount++
			} else {
				queue = append(queue, neighbour)
			}
		}
	}

	return totalCount
}

func getNeighbours(point types.Point, topographyMap Map) []types.Point {
	neighbours := make([]types.Point, 0)

	if point.X+1 <= topographyMap.MaxX() {
		neighbours = append(neighbours, types.Point{X: point.X + 1, Y: point.Y})
	}
	if point.Y+1 <= topographyMap.MaxY() {
		neighbours = append(neighbours, types.Point{X: point.X, Y: point.Y + 1})
	}
	if point.X-1 >= topographyMap.MinX() {
		neighbours = append(neighbours, types.Point{X: point.X - 1, Y: point.Y})
	}
	if point.Y-1 >= topographyMap.MinY() {
		neighbours = append(neighbours, types.Point{X: point.X, Y: point.Y - 1})
	}

	return neighbours
}
