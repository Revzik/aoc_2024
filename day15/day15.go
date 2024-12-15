package day15

import (
	"fmt"

	"github.com/revzik/aoc_2024/common/files"
)

func RunTask() {
	path := "day15/input"
	lines := files.ReadLines(path)
	warehouse, moves := parseInput(lines)
	wideWarehouse := getWideWarehouse(warehouse)

	processInstructions(warehouse, moves, false)
	fmt.Printf("GPS positions in small warehouse: %d\n", calculateGpsSum(warehouse))

	processInstructions(wideWarehouse, moves, true)
	fmt.Printf("GPS positions in wide warehouse: %d\n", calculateGpsSum(warehouse))
}

type Vector struct {
	X int
	Y int
}

func (v1 Vector) Add(v2 Vector) Vector {
	return Vector{X: v1.X + v2.X, Y: v1.Y + v2.Y}
}

func getObstacle(m [][]rune, v Vector) rune {
	return m[v.Y][v.X]
}

func setObstacle(m [][]rune, v Vector, o rune) {
	m[v.Y][v.X] = o
}

func parseInput(lines []string) ([][]rune, []rune) {
	warehouse := make([][]rune, 0)
	moves := make([]rune, 0)

	warehouseDone := false

	for _, line := range lines {
		if line == "" {
			warehouseDone = true
			continue
		}
		if warehouseDone {
			moves = append(moves, []rune(line)...)
		} else {
			warehouse = append(warehouse, []rune(line))
		}
	}

	return warehouse, moves
}

func getWideWarehouse(warehouse [][]rune) [][]rune {
	wideWarehouse := make([][]rune, len(warehouse))

	for i, row := range warehouse {
		wideRow := make([]rune, len(row)*2)
		for j, char := range row {
			if char == '#' {
				wideRow[j*2] = '#'
				wideRow[j*2+1] = '#'
			} else if char == '.' {
				wideRow[j*2] = '.'
				wideRow[j*2+1] = '.'
			} else if char == 'O' {
				wideRow[j*2] = '['
				wideRow[j*2+1] = ']'
			} else if char == '@' {
				wideRow[j*2] = '@'
				wideRow[j*2+1] = '.'
			}
		}
		wideWarehouse[i] = wideRow
	}

	return wideWarehouse
}

// common functions
func processInstructions(warehouse [][]rune, moves []rune, wideWarehouse bool) {
	robot := findRobot(warehouse)

	for _, move := range moves {
		if wideWarehouse {
			robot = moveRobotSmallWarehouse(warehouse, move, robot)
		} else {
			robot = moveRobotWideWarehouse(warehouse, move, robot)
		}
	}
}

func findRobot(warehouse [][]rune) Vector {
	for y, row := range warehouse {
		for x, char := range row {
			if char == '@' {
				return Vector{X: x, Y: y}
			}
		}
	}
	panic("no robot found")
}

func getDirection(directionChar rune) Vector {
	if directionChar == '>' {
		return Vector{X: 1, Y: 0}
	}
	if directionChar == 'v' {
		return Vector{X: 0, Y: 1}
	}
	if directionChar == '<' {
		return Vector{X: -1, Y: 0}
	}
	if directionChar == '^' {
		return Vector{X: 0, Y: -1}
	}
	panic("unsupported direction character")
}

func calculateGpsSum(warehouse [][]rune) int {
	total := 0

	for y, row := range warehouse {
		for x, obstacle := range row {
			if obstacle == 'O' {
				total += x + 100*y
			}
		}
	}

	return total
}

// part one
func moveRobotSmallWarehouse(warehouse [][]rune, move rune, robot Vector) Vector {

	direction := getDirection(move)
	position := robot
	boxes := make([]Vector, 0)
	isBlocked := false

	for {
		position = position.Add(direction)
		obstacle := getObstacle(warehouse, position)

		if obstacle == '.' {
			break
		}

		if obstacle == '#' {
			isBlocked = true
			break
		}

		if getObstacle(warehouse, position) == 'O' {
			boxes = append(boxes, position)
		}
	}

	if isBlocked {
		return robot
	}

	var newRobot Vector
	if len(boxes) == 0 {
		newRobot = position
		setObstacle(warehouse, newRobot, '@')
	} else {
		newRobot = boxes[0]
		setObstacle(warehouse, position, 'O')
		setObstacle(warehouse, newRobot, '@')
	}
	setObstacle(warehouse, robot, '.')

	return newRobot
}

// part two
func moveRobotWideWarehouse(warehouse [][]rune, move rune, robot Vector) Vector {
	panic("unimplemented")
}
