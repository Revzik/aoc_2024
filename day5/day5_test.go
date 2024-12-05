package day5

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseMap(t *testing.T) {
	orderRules := [][]int{
		{3, 5},
		{1, 3},
		{2, 3},
		{2, 4},
		{1, 2},
		{4, 5},
		{3, 4},
		{1, 4},
		{2, 5},
		{1, 5},
	}
	expectedRuleMap := map[int][]int{
		1: {3, 2, 4, 5},
		2: {3, 4, 5},
		3: {5, 4},
		4: {5},
	}

	assert.Equal(t, expectedRuleMap, getRuleMap(orderRules))
}

func TestCheckCorrectlyOrdered(t *testing.T) {
	orderRules := map[int][]int{
		1: {2, 3, 4},
		2: {3, 4},
		3: {4},
	}

	assert.True(t, isCorrectlyOrdered([]int{1, 2, 3, 4}, orderRules))
	assert.False(t, isCorrectlyOrdered([]int{4, 1, 2, 3}, orderRules))
}

func TestSortPages(t *testing.T) {
	orderRules := map[int][]int{
		1: {2, 3, 4},
		2: {3, 4},
		3: {4},
	}
	incorrectPages := []int{4, 1, 2, 3}
	expectedPages := []int{1, 2, 3, 4}

	assert.Equal(t, expectedPages, sortPages(incorrectPages, orderRules))
}
