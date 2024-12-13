package day13

import (
	"errors"
	"fmt"
	"math"
	"strings"

	"github.com/revzik/aoc_2024/common/files"
	"github.com/revzik/aoc_2024/common/parsers"
)

func RunTask() {
	path := "day13/input"
	lines := files.ReadLines(path)
	arcades := parseInput(lines)
	fixedArcades := fixArcades(arcades)

	fmt.Printf("Minimal amount of tokens, pre fix: %d\n", calculateAllTokens(arcades))
	fmt.Printf("Minimal amount of tokens, post fix: %d\n", calculateAllTokens(fixedArcades))
}

type Vector struct {
	X int64
	Y int64
}

func (v Vector) Opposite() Vector {
	return Vector{X: -v.X, Y: -v.Y}
}

type Arcade struct {
	Origin Vector
	Target Vector
	A      Vector
	B      Vector
}

func parseInput(lines []string) []Arcade {
	arcades := make([]Arcade, 0)

	arcade := Arcade{Origin: Vector{X: 0, Y: 0}}
	for _, line := range lines {
		if line == "" {
			arcades = append(arcades, arcade)
			arcade = Arcade{Origin: Vector{X: 0, Y: 0}}
			continue
		}

		splitLine := strings.Split(line, ": ")
		coords := strings.Split(splitLine[1], ", ")

		x := parsers.StringToInt64(coords[0][2:])
		y := parsers.StringToInt64(coords[1][2:])
		position := Vector{X: x, Y: y}

		if splitLine[0] == "Button A" {
			arcade.A = position
		} else if splitLine[0] == "Button B" {
			arcade.B = position
		} else {
			arcade.Target = position
		}
	}

	emptyArcade := Arcade{Origin: Vector{X: 0, Y: 0}}
	if arcade != emptyArcade {
		arcades = append(arcades, arcade)
	}

	return arcades
}

func fixArcades(arcades []Arcade) []Arcade {
	newArcades := make([]Arcade, len(arcades))

	for i, arcade := range arcades {
		newArcade := arcade
		newArcade.Target.X += 10000000000000
		newArcade.Target.Y += 10000000000000
		newArcades[i] = newArcade
	}

	return newArcades
}

func calculateAllTokens(arcades []Arcade) int64 {
	total := int64(0)

	for _, arcade := range arcades {
		tokens, err := calculateTokens(arcade)
		if err == nil {
			total += tokens
		}
	}

	return total
}

// given there are only 2 vectors that can lead us to the target
// there is really only one combination of them, unless they have the same direction
//
// in fact we can designate 2 straight lines, one with vector A from origin and one with vector B from target
// then we calculate where they cross
// if in the first quarter, then it might be valid
// if it's past the target, it's invalid
//
// if midpoint is reachable from origin by a whole amount of A vectors
// and if target is reachable from midpoint by a whole amount of B vectors
// the combination is correct and amount of tokens can be calculated
func calculateTokens(arcade Arcade) (int64, error) {

	if canReachByVector(arcade.Origin, arcade.Target, arcade.B) {
		return arcade.Target.X / arcade.B.X, nil
	}
	if canReachByVector(arcade.Origin, arcade.Target, arcade.A) {
		return arcade.Target.X / arcade.A.X * 3, nil
	}

	position, err := getValidMidpoint(arcade.Origin, arcade.Target, arcade.A, arcade.B)
	if err == nil {
		if canReachByVector(arcade.Origin, position, arcade.A) && canReachByVector(arcade.Target, position, arcade.B.Opposite()) {
			aTokens := position.X / arcade.A.X * 3
			bTokens := (arcade.Target.X - position.X) / arcade.B.X
			return aTokens + bTokens, nil
		}
	}

	return 0, errors.New("unreachable target")
}

func getValidMidpoint(origin, target, fromOrigin, fromTarget Vector) (Vector, error) {
	// after simplification of:
	//     y1 - y2   orY - (orY - vY)   vy
	// a = ------- = ---------------- = --
	//     x1 - x2   orX - (orX - vX)   vx
	a1 := float64(fromOrigin.Y) / float64(fromOrigin.X)
	b1 := float64(origin.Y) - a1*float64(origin.X)

	a2 := float64(fromTarget.Y) / float64(fromTarget.X)
	b2 := float64(target.Y) - a2*float64(target.X)

	if a1 < 0 || a2 < 0 {
		return Vector{}, errors.New("a coefficient cannot be negative")
	}

	x := (b2 - b1) / (a1 - a2)
	y := a1*x + b1

	if x > float64(target.X) || y > float64(target.Y) {
		return Vector{}, errors.New("midpoint beyond target")
	}

	return Vector{X: int64(math.Round(x)), Y: int64(math.Round(y))}, nil
}

func canReachByVector(origin, target, v Vector) bool {
	adjustedTarget := Vector{X: target.X - origin.X, Y: target.Y - origin.Y}

	if v.X == 0 && v.Y == 0 {
		return false
	}
	if v.X == 0 && adjustedTarget.Y%v.Y == 0 {
		return true
	}
	if v.Y == 0 && adjustedTarget.X%v.X == 0 {
		return true
	}

	if adjustedTarget.X%v.X == 0 && adjustedTarget.Y%v.Y == 0 {
		amountX := adjustedTarget.X / v.X
		amountY := adjustedTarget.Y / v.Y
		return amountX == amountY
	}
	return false
}
