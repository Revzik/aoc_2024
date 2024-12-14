package day14

import (
	"fmt"
	"strings"

	"github.com/revzik/aoc_2024/common/files"
	"github.com/revzik/aoc_2024/common/parsers"
)

func RunTask() {
	path := "day14/input"
	lines := files.ReadLines(path)

	originalHall := createHall(lines, 101, 103)
	hall := calculatePositionsAfter(originalHall, 100)
	safetyFactor := calculateSafetyFactor(hall)

	fmt.Printf("Safety factor: %d\n", safetyFactor)
	printTree(originalHall)
	fmt.Printf("Seconds after tree is formed: %d\n", searchForTree(originalHall))
}

type Hall struct {
	Robots []Robot
	Width  int
	Height int
}

type Robot struct {
	Position Vector
	Velocity Vector
}

type Vector struct {
	X int
	Y int
}

func createHall(lines []string, width, height int) Hall {
	robots := make([]Robot, len(lines))

	for i, line := range lines {
		splitLine := strings.Split(line, " ")
		positionString := strings.Split(splitLine[0][2:], ",")
		velocityString := strings.Split(splitLine[1][2:], ",")

		position := Vector{
			X: parsers.StringToInt(positionString[0]),
			Y: parsers.StringToInt(positionString[1]),
		}
		velocity := Vector{
			X: parsers.StringToInt(velocityString[0]),
			Y: parsers.StringToInt(velocityString[1]),
		}

		robots[i] = Robot{Position: position, Velocity: velocity}
	}

	return Hall{Robots: robots, Width: width, Height: height}
}

// first part
func calculateSafetyFactor(hall Hall) int {
	quadrants := []int{0, 0, 0, 0}
	midVertical := (hall.Width - 1) / 2
	midHorizontal := (hall.Height - 1) / 2

	for _, robot := range hall.Robots {
		position := robot.Position
		if position.X < midVertical && position.Y < midHorizontal {
			quadrants[0]++
		} else if position.X > midVertical && position.Y < midHorizontal {
			quadrants[1]++
		} else if position.X < midVertical && position.Y > midHorizontal {
			quadrants[2]++
		} else if position.X > midVertical && position.Y > midHorizontal {
			quadrants[3]++
		}
	}

	return quadrants[0] * quadrants[1] * quadrants[2] * quadrants[3]
}

func calculatePositionsAfter(hall Hall, seconds int) Hall {
	newRobots := make([]Robot, len(hall.Robots))
	for i, robot := range hall.Robots {
		newRobot := Robot{Position: Vector{}, Velocity: robot.Velocity}
		newRobot.Position.X = robot.Position.X + robot.Velocity.X*seconds
		newRobot.Position.Y = robot.Position.Y + robot.Velocity.Y*seconds
		newRobots[i] = normalizePosition(newRobot, hall)
	}
	hall.Robots = newRobots
	return hall
}

func normalizePosition(robot Robot, hall Hall) Robot {
	x := robot.Position.X
	y := robot.Position.Y

	if x >= hall.Width {
		overlaps := x / hall.Width
		x -= hall.Width * overlaps
	}
	if y >= hall.Height {
		overlaps := y / hall.Height
		y -= hall.Height * overlaps
	}
	if x < 0 {
		overlaps := -x / hall.Width
		x += hall.Width * overlaps
		// could use ceil to calculate overlaps but this works
		if x < 0 {
			x += hall.Width
		}
	}
	if y < 0 {
		overlaps := -y / hall.Height
		y += hall.Height * overlaps
		if y < 0 {
			y += hall.Height
		}
	}

	robot.Position = Vector{X: x, Y: y}

	return robot
}

// second part
//
// semi-automatic approach
// the idea (not mine) is to use statistics to find clustered robots and validate the image
//
// calculate variance of X and Y coordinates of robots
// high values mean a uniform distribution
// low values mean clustered distribution
// if values are low, print out the hall and check if it looks like a chrismas tree
//
// Variance for X hits a low after certain amount of seconds and repeats after 101
// Variance for Y hits a low after certain amount of seconds and repeats after 103
// When they hit a common low, that's where a potential cristmas tree is formed
func searchForTree(hall Hall) int {
	elapsedSeconds := 0
	potentialSeconds := make([]int, 0)
	varianceThreshold := 500.0

	// left for testing purposes
	for len(potentialSeconds) < 1 {
		elapsedSeconds++
		processSecond(hall)
		varX, varY := calculateVariance(hall.Robots)
		if varX < varianceThreshold && varY < varianceThreshold {
			potentialSeconds = append(potentialSeconds, elapsedSeconds)
			fmt.Printf("Variances X: %.3f, Y: %.3f, second: %d\n", varX, varY, elapsedSeconds)
			printTree(hall)
		}
	}

	return elapsedSeconds
}

func processSecond(hall Hall) {
	robots := hall.Robots
	for i, robot := range robots {
		robot.Position.X += robot.Velocity.X
		robot.Position.Y += robot.Velocity.Y
		robots[i] = normalizePosition(robot, hall)
	}
	hall.Robots = robots
}

func calculateVariance(robots []Robot) (float64, float64) {
	sumX, sumY := 0, 0
	for _, robot := range robots {
		sumX += robot.Position.X
		sumY += robot.Position.Y
	}

	meanX := float64(sumX) / float64(len(robots))
	meanY := float64(sumY) / float64(len(robots))

	diffX, diffY := 0.0, 0.0
	squareSumX, squareSumY := 0.0, 0.0
	for _, robot := range robots {
		diffX = float64(robot.Position.X) - meanX
		diffY = float64(robot.Position.Y) - meanY
		squareSumX += diffX * diffX
		squareSumY += diffY * diffY
	}

	varianceX := squareSumX / float64(len(robots))
	varianceY := squareSumY / float64(len(robots))

	return varianceX, varianceY
}

func printTree(hall Hall) {
	matrix := make([][]rune, hall.Height)
	for i := range matrix {
		matrix[i] = make([]rune, hall.Width)
		for j := range matrix[i] {
			matrix[i][j] = ' '
		}
	}

	for _, robot := range hall.Robots {
		matrix[robot.Position.Y][robot.Position.X] = '#'
	}

	for _, row := range matrix {
		fmt.Println(string(row))
	}
}
