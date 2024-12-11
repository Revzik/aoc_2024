package day11

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProcessSingleBlink(t *testing.T) {
	stones := map[int]int{
		0:   1,
		1:   1,
		10:  1,
		100: 1,
	}
	expectedStones := map[int]int{
		1:      2,
		2024:   1,
		0:      1,
		202400: 1,
	}

	resultStones := processBlink(stones)

	assert.Equal(t, expectedStones, resultStones)
}

func TestPartOne(t *testing.T) {
	stones := []int{0, 1, 2, 3, 4, 2024, 17, 64531, 987639407}
	expectedResult := 251176

	result := countStones(stones, 25)

	assert.Equal(t, expectedResult, result)
}
