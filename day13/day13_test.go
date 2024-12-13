package day13

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParsingInput(t *testing.T) {
	input := []string{
		"Button A: X+94, Y+34",
		"Button B: X+22, Y+67",
		"Prize: X=8400, Y=5400",
		"",
		"Button A: X+26, Y+66",
		"Button B: X+67, Y+21",
		"Prize: X=12748, Y=12176",
	}
	expectedArcades := []Arcade{
		{A: Vector{X: 94, Y: 34}, B: Vector{X: 22, Y: 67}, Origin: Vector{X: 0, Y: 0}, Target: Vector{X: 8400, Y: 5400}},
		{A: Vector{X: 26, Y: 66}, B: Vector{X: 67, Y: 21}, Origin: Vector{X: 0, Y: 0}, Target: Vector{X: 12748, Y: 12176}},
	}

	arcades := parseInput(input)

	assert.Equal(t, expectedArcades, arcades)
}

func TestCalculateTokens(t *testing.T) {
	arcade1 := Arcade{
		Origin: Vector{X: 0, Y: 0},
		Target: Vector{X: 10, Y: 10},
		A:      Vector{X: 4, Y: 4},
		B:      Vector{X: 2, Y: 2},
	}
	arcade2 := Arcade{
		Origin: Vector{X: 0, Y: 0},
		Target: Vector{X: 10, Y: 14},
		A:      Vector{X: 4, Y: 4},
		B:      Vector{X: 1, Y: 3},
	}
	arcade3 := Arcade{
		Origin: Vector{X: 0, Y: 0},
		Target: Vector{X: 10000000012748, Y: 10000000012176},
		A:      Vector{X: 26, Y: 66},
		B:      Vector{X: 67, Y: 21},
	}
	arcade4 := Arcade{
		Origin: Vector{X: 0, Y: 0},
		Target: Vector{X: 20, Y: 10},
		A:      Vector{X: 4, Y: 4},
		B:      Vector{X: 2, Y: 2},
	}

	tokens1, err1 := calculateTokens(arcade1)
	tokens2, err2 := calculateTokens(arcade2)
	tokens3, err3 := calculateTokens(arcade3)
	_, err4 := calculateTokens(arcade4)

	assert.Nil(t, err1)
	assert.Equal(t, int64(5), tokens1)

	assert.Nil(t, err2)
	assert.Equal(t, int64(8), tokens2)

	assert.Nil(t, err3)
	assert.Equal(t, int64(459236326669), tokens3)

	assert.NotNil(t, err4)
}

func TestIsReachableByVector(t *testing.T) {
	origins := []Vector{
		{X: 0, Y: 0},
		{X: 0, Y: 0},
		{X: 10, Y: 10},
		{X: 25, Y: 40},
	}
	targets := []Vector{
		{X: 1, Y: 2},
		{X: -2, Y: -4},
		{X: 10, Y: 20},
		{X: 15, Y: 20},
	}
	vectors := []Vector{
		{X: 1, Y: 2},
		{X: -1, Y: -2},
		{X: 0, Y: 2},
		{X: -5, Y: -10},
	}

	assert.True(t, canReachByVector(origins[0], targets[0], vectors[0]))
	assert.True(t, canReachByVector(origins[1], targets[1], vectors[1]))
	assert.True(t, canReachByVector(origins[2], targets[2], vectors[2]))
	assert.True(t, canReachByVector(origins[3], targets[3], vectors[3]))
}

func TestIsNotReachableByVector(t *testing.T) {
	origins := []Vector{
		{X: 0, Y: 0},
		{X: 0, Y: 0},
		{X: 10, Y: 10},
		{X: 25, Y: 40},
	}
	targets := []Vector{
		{X: 1, Y: 2},
		{X: -2, Y: -4},
		{X: 10, Y: 20},
		{X: 15, Y: 20},
	}
	vectors := []Vector{
		{X: 0, Y: 0},
		{X: -2, Y: -2},
		{X: -1, Y: 2},
		{X: 5, Y: -10},
	}

	assert.False(t, canReachByVector(origins[0], targets[0], vectors[0]))
	assert.False(t, canReachByVector(origins[1], targets[1], vectors[1]))
	assert.False(t, canReachByVector(origins[2], targets[2], vectors[2]))
	assert.False(t, canReachByVector(origins[3], targets[3], vectors[3]))
}
