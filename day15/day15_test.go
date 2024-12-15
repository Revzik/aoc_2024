package day15

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindRobot(t *testing.T) {
	warehouse := [][]rune{
		[]rune("########"),
		[]rune("#...#.O#"),
		[]rune("#.###..#"),
		[]rune("#...O..#"),
		[]rune("#...O.O#"),
		[]rune("#.@....#"),
		[]rune("#...O..#"),
		[]rune("########"),
	}

	robot := findRobot(warehouse)

	assert.Equal(t, Vector{X: 2, Y: 5}, robot)
}

func TestMoveRobotSmallWarehouse(t *testing.T) {
	warehouse := [][]rune{
		[]rune("########"),
		[]rune("#......#"),
		[]rune("#..@.#.#"),
		[]rune("#...O..#"),
		[]rune("#...O..#"),
		[]rune("#...O..#"),
		[]rune("#......#"),
		[]rune("########"),
	}
	robot := Vector{X: 3, Y: 2}

	robot = moveRobotSmallWarehouse(warehouse, '>', robot)
	assert.Equal(t, Vector{X: 4, Y: 2}, robot)

	robot = moveRobotSmallWarehouse(warehouse, '>', robot)
	assert.Equal(t, Vector{X: 4, Y: 2}, robot)

	robot = moveRobotSmallWarehouse(warehouse, 'v', robot)
	assert.Equal(t, Vector{X: 4, Y: 3}, robot)

	robot = moveRobotSmallWarehouse(warehouse, 'v', robot)
	assert.Equal(t, Vector{X: 4, Y: 3}, robot)

	robot = moveRobotSmallWarehouse(warehouse, '<', robot)
	assert.Equal(t, Vector{X: 3, Y: 3}, robot)

	finalWarehouse := [][]rune{
		[]rune("########"),
		[]rune("#......#"),
		[]rune("#....#.#"),
		[]rune("#..@...#"),
		[]rune("#...O..#"),
		[]rune("#...O..#"),
		[]rune("#...O..#"),
		[]rune("########"),
	}
	assert.Equal(t, finalWarehouse, warehouse)
}

func TestGetWideWarehouse(t *testing.T) {
	warehouse := [][]rune{
		[]rune("########"),
		[]rune("#...#.O#"),
		[]rune("#.###..#"),
		[]rune("#...O..#"),
		[]rune("#...O.O#"),
		[]rune("#.@....#"),
		[]rune("#...O..#"),
		[]rune("########"),
	}
	expectedWide := [][]rune{
		[]rune("################"),
		[]rune("##......##..[]##"),
		[]rune("##..######....##"),
		[]rune("##......[]....##"),
		[]rune("##......[]..[]##"),
		[]rune("##..@.........##"),
		[]rune("##......[]....##"),
		[]rune("################"),
	}

	wideWarehouse := getWideWarehouse(warehouse)

	assert.Equal(t, expectedWide, wideWarehouse)
}
