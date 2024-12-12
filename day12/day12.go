package day12

import (
	"fmt"
	"sort"

	"github.com/revzik/aoc_2024/common/files"
)

func RunTask() {
	path := "day12/input"
	lines := files.ReadLines(path)
	field := createField(lines)

	regularPrice, discountedPrice := calculatePrices(field)
	fmt.Printf("Regular fence price: %d\n", regularPrice)
	fmt.Printf("Discounted fence price: %d\n", discountedPrice)
}

// structures to make things a bit more readable
type Region struct {
	X int
	Y int
}

type NeighbourRegion struct {
	Reg  Region
	Side rune
}

type Field struct {
	Plants  [][]rune
	Visited [][]bool
}

func (f Field) MinX() int {
	return 0
}

func (f Field) MinY() int {
	return 0
}

func (f Field) MaxX() int {
	return len(f.Plants[0]) - 1
}

func (f Field) MaxY() int {
	return len(f.Plants) - 1
}

func (f Field) GetPlant(r Region) rune {
	return f.Plants[r.Y][r.X]
}

func (f Field) IsVisited(r Region) bool {
	return f.Visited[r.Y][r.X]
}

func (f Field) Visit(r Region) {
	f.Visited[r.Y][r.X] = true
}

// parsing input
func createField(lines []string) Field {
	plants := make([][]rune, len(lines))
	visited := make([][]bool, len(lines))

	for i, line := range lines {
		plants[i] = []rune(line)
		visited[i] = make([]bool, len(line))
		for j := range visited[i] {
			visited[i][j] = false
		}
	}

	return Field{Plants: plants, Visited: visited}
}

// main functions for calculating both prices
func calculatePrices(field Field) (int, int) {
	regularTotal, discountedTotal := 0, 0

	for y, row := range field.Visited {
		for x, visited := range row {
			if visited {
				continue
			}

			regular, discounted := calculatePlotPrices(field, Region{X: x, Y: y})
			regularTotal += regular
			discountedTotal += discounted
		}
	}

	return regularTotal, discountedTotal
}

func calculatePlotPrices(field Field, region Region) (int, int) {
	plot := findPlot(region, field)

	return calculateRegularPrice(plot), calculateDiscoutedPrice(plot)
}

// finds individual plot starting from selected region
// returns all of the plot regions together with their egdes (walls)
// if a region has has edges to the left and right it'll be marked as:
// Region: {'R': true, 'D': false, 'L': true, 'U': false}
func findPlot(region Region, field Field) map[Region]map[rune]bool {
	plot := make(map[Region]map[rune]bool, 0)
	queue := []Region{region}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		walls := map[rune]bool{
			'R': true,
			'D': true,
			'L': true,
			'U': true,
		}

		neighbours := getNeighbours(current, field)
		for _, neighbour := range neighbours {
			if field.GetPlant(current) != field.GetPlant(neighbour.Reg) {
				continue
			}

			walls[neighbour.Side] = false

			if field.IsVisited(neighbour.Reg) {
				continue
			}

			queue = append(queue, neighbour.Reg)
			field.Visit(neighbour.Reg)
		}

		plot[current] = walls
	}

	return plot
}

func getNeighbours(region Region, field Field) []NeighbourRegion {
	neighbours := make([]NeighbourRegion, 0)

	if region.X+1 <= field.MaxX() {
		neighbours = append(neighbours, NeighbourRegion{Reg: Region{X: region.X + 1, Y: region.Y}, Side: 'R'})
	}
	if region.Y+1 <= field.MaxY() {
		neighbours = append(neighbours, NeighbourRegion{Reg: Region{X: region.X, Y: region.Y + 1}, Side: 'D'})
	}
	if region.X-1 >= field.MinX() {
		neighbours = append(neighbours, NeighbourRegion{Reg: Region{X: region.X - 1, Y: region.Y}, Side: 'L'})
	}
	if region.Y-1 >= field.MinY() {
		neighbours = append(neighbours, NeighbourRegion{Reg: Region{X: region.X, Y: region.Y - 1}, Side: 'U'})
	}

	return neighbours
}

// part one price calculation
// count all regions as area, count all walls as perimeter
func calculateRegularPrice(plot map[Region]map[rune]bool) int {
	area, perimeter := 0, 0

	for _, walls := range plot {
		area++
		for _, exists := range walls {
			if exists {
				perimeter++
			}
		}
	}

	return area * perimeter
}

// part two price calculation
// consider a shape:
//
//	A.AA.
//	AAAAA
//
// idea is to:
//  1. group plot regions that have a certain wall ('U' for example)
//  2. compute a map of walls that are on the same line:
//     0: [0, 2, 3]
//     1: [1, 4]
//  3. count how many consecutive numbers are bigger than more than 1
func calculateDiscoutedPrice(plot map[Region]map[rune]bool) int {
	wallPositions := []rune{'R', 'D', 'L', 'U'}
	totalWalls := 0

	for _, position := range wallPositions {
		regions := getRegionsWithWall(plot, position)
		wallsByPosition := getWallsByPosition(regions, position)
		totalWalls += countContinuousWalls(wallsByPosition)
	}

	return len(plot) * totalWalls
}

func getRegionsWithWall(plot map[Region]map[rune]bool, side rune) []Region {
	regions := make([]Region, 0)
	for region, walls := range plot {
		if walls[side] {
			regions = append(regions, region)
		}
	}
	return regions
}

// for horizontal walls, we're interested in chains of X positions
// for vertical, we're interested in Y positions
func getWallsByPosition(regions []Region, wall rune) map[int][]int {
	wallsByPosition := make(map[int][]int, 0)

	for _, region := range regions {
		if wall == 'R' || wall == 'L' {
			putOrAppend(wallsByPosition, region.X, region.Y)
		} else {
			putOrAppend(wallsByPosition, region.Y, region.X)
		}
	}

	for _, lengths := range wallsByPosition {
		sort.Ints(lengths)
	}

	return wallsByPosition
}

func putOrAppend(m map[int][]int, k int, v int) {
	if val, ok := m[k]; ok {
		m[k] = append(val, v)
	} else {
		m[k] = []int{v}
	}
}

func countContinuousWalls(wallsByPosition map[int][]int) int {
	allContinuousWalls := 0

	for _, lengths := range wallsByPosition {
		previousWall := lengths[0]
		continuousWalls := 1

		for i := 1; i < len(lengths); i++ {
			if previousWall+1 != lengths[i] {
				continuousWalls++
			}
			previousWall = lengths[i]
		}

		allContinuousWalls += continuousWalls
	}

	return allContinuousWalls
}
